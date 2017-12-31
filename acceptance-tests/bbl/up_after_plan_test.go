package acceptance_test

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"time"

	yaml "gopkg.in/yaml.v2"

	acceptance "github.com/cloudfoundry/bosh-bootloader/acceptance-tests"
	"github.com/cloudfoundry/bosh-bootloader/acceptance-tests/actors"
	"github.com/cloudfoundry/bosh-bootloader/storage"
	"github.com/cloudfoundry/bosh-bootloader/testhelpers"
	proxy "github.com/cloudfoundry/socks5-proxy"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/onsi/gomega/gexec"
)

var _ = Describe("up after plan", func() {
	var (
		bbl actors.BBL

		stateDir   string
		jumpboxURL string
		iaas       string
	)

	BeforeEach(func() {
		acceptance.SkipUnless("up-after-plan")

		configuration, err := acceptance.LoadConfig()
		Expect(err).NotTo(HaveOccurred())

		iaas = configuration.IAAS
		stateDir = configuration.StateFileDir

		bbl = actors.NewBBL(stateDir, pathToBBL, configuration, "up-after-plan-env")
	})

	AfterEach(func() {
		acceptance.SkipUnless("up-after-plan")

		deleteDirectorPath := filepath.Join(stateDir, "delete-director.sh")
		deleteJumpboxPath := filepath.Join(stateDir, "delete-jumpbox.sh")
		noOpScript := []byte("#!/bin/bash\n")

		err := ioutil.WriteFile(deleteDirectorPath, noOpScript, storage.ScriptMode)
		Expect(err).NotTo(HaveOccurred())

		err = ioutil.WriteFile(deleteJumpboxPath, noOpScript, storage.ScriptMode)
		Expect(err).NotTo(HaveOccurred())

		session := bbl.Down()
		Eventually(session, 10*time.Minute).Should(gexec.Exit())
	})

	It("preserves files modified after plan", func() {
		tempDir, err := ioutil.TempDir("", "")
		Expect(err).NotTo(HaveOccurred())

		createEnvOutputPath := filepath.Join(tempDir, "create-env-output")

		By("running bbl plan", func() {
			var (
				cert string
				key  string
			)

			if iaas == "azure" {
				cert = testhelpers.PFX_BASE64
				key = testhelpers.PFX_PASSWORD
			} else {
				cert = testhelpers.BBL_CERT
				key = testhelpers.BBL_KEY
			}

			certPath, err := testhelpers.WriteContentsToTempFile(cert)
			Expect(err).NotTo(HaveOccurred())

			keyPath, err := testhelpers.WriteContentsToTempFile(key)
			Expect(err).NotTo(HaveOccurred())

			session := bbl.Plan("--name", bbl.PredefinedEnvID(), "--lb-type", "cf", "--lb-cert", certPath, "--lb-key", keyPath)
			Eventually(session, 40*time.Minute).Should(gexec.Exit(0))
		})

		By("starting an SSH server as a double for the jumpbox", func() {
			httpServer := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
				rw.WriteHeader(http.StatusOK)
			}))
			httpServerHostPort := strings.Split(httpServer.URL, "http://")[1]

			err := exec.Command("bosh",
				"int", filepath.Join(stateDir, "jumpbox-deployment", "jumpbox.yml"),
				"--vars-store", filepath.Join(stateDir, "vars", "jumpbox-vars-store.yml"),
			).Run()
			Expect(err).NotTo(HaveOccurred())

			vars, err := ioutil.ReadFile(filepath.Join(stateDir, "vars", "jumpbox-vars-store.yml"))
			Expect(err).NotTo(HaveOccurred())

			key := getJumpboxPrivateKey(string(vars))
			jumpboxURL = proxy.StartTestSSHServer(httpServerHostPort, key)
		})

		By("modifying the plan", func() {
			createDirectorPath := filepath.Join(stateDir, "create-director.sh")
			newCreateDirector := []byte(fmt.Sprintf("#!/bin/bash\necho 'director' >> %s\n", createEnvOutputPath))

			err := ioutil.WriteFile(createDirectorPath, newCreateDirector, storage.ScriptMode)
			Expect(err).NotTo(HaveOccurred())

			createJumpboxPath := filepath.Join(stateDir, "create-jumpbox.sh")
			newCreateJumpbox := []byte(fmt.Sprintf("#!/bin/bash\necho 'jumpbox' >> %s\n", createEnvOutputPath))

			err = ioutil.WriteFile(createJumpboxPath, newCreateJumpbox, storage.ScriptMode)
			Expect(err).NotTo(HaveOccurred())

			terraformTemplatePath := filepath.Join(stateDir, "terraform", "template.tf")
			newTerraformTemplate := []byte(fmt.Sprintf(`
output "jumpbox_url" { value = "%s" }

output "internal_security_group" { value = "some-security-group" }
output	"cf_router_lb_name" { value = "some-router-lb-name" }
output	"cf_router_lb_internal_security_group" { value = "some-router-internal-security-group" }
output	"cf_ssh_lb_name" { value = "some-ssh-lb-name" }
output	"cf_ssh_lb_internal_security_group" { value = "some-ssh-internal-security-group" }
output	"cf_tcp_lb_name" { value = "some-tcp-lb-name" }
output	"cf_tcp_lb_internal_security_group" { value = "some-tcp-internal-security-group" }
output "internal_az_subnet_cidr_mapping" {
  value = "${
    map(
		"us-east-1a", "10.0.16.0/20",
		"us-east-1c", "10.0.48.0/20",
		"us-east-1b", "10.0.32.0/20"
    )
  }"
}

output "internal_az_subnet_id_mapping" {
  value = "${
    map(
	    "us-east-1c", "some-internal-subnet-ids-3",
		"us-east-1a", "some-internal-subnet-ids-1",
		"us-east-1b", "some-internal-subnet-ids-2"
    )
  }"
}
`, jumpboxURL))

			err = ioutil.WriteFile(terraformTemplatePath, newTerraformTemplate, storage.StateMode)
			Expect(err).NotTo(HaveOccurred())
		})

		By("running bbl up", func() {
			time.Sleep(5 * time.Second)

			session := bbl.Up()
			Eventually(session, 40*time.Minute).Should(gexec.Exit())
			// Don't check the exit code of up because upload cloud config fails.
			// We don't yet have a way to inject different behavior for that step.
		})

		By("verifying that artifacts are created in state dir", func() {
			checkExists := func(dir string, filenames []string) {
				for _, f := range filenames {
					_, err := os.Stat(filepath.Join(dir, f))
					Expect(err).NotTo(HaveOccurred())
				}
			}

			checkExists(stateDir, []string{"bbl-state.json"})
			checkExists(stateDir, []string{"create-jumpbox.sh"})
			checkExists(stateDir, []string{"create-director.sh"})
			checkExists(stateDir, []string{"delete-jumpbox.sh"})
			checkExists(stateDir, []string{"delete-director.sh"})
			checkExists(filepath.Join(stateDir, "cloud-config"), []string{
				"cloud-config.yml",
				"ops.yml",
			})
			checkExists(filepath.Join(stateDir, "bosh-deployment"), []string{
				"bosh.yml",
				filepath.Join(iaas, "cpi.yml"),
				"credhub.yml",
				"jumpbox-user.yml",
				"uaa.yml",
			})
			checkExists(filepath.Join(stateDir, "jumpbox-deployment"), []string{
				filepath.Join(iaas, "cpi.yml"),
				"jumpbox.yml",
			})
			checkExists(filepath.Join(stateDir, "terraform"), []string{
				"template.tf",
			})
			checkExists(filepath.Join(stateDir, "vars"), []string{
				//  These don't exist because we overwrite the create-env scripts
				// "bosh-state.json",
				// "jumpbox-state.json",
				// "director-variables.yml",
				// "jumpbox-variables.yml",
				"director-vars-file.yml",
				"jumpbox-vars-file.yml",
				"terraform.tfstate",
				"user-ops-file.yml",
			})
		})

		By("verifying that vm extensions were added to the cloud config", func() {
			if iaas == "azure" {
				return
			}

			var cloudConfig struct {
				VMExtensions []struct {
					Name            string                 `yaml:"name"`
					CloudProperties map[string]interface{} `yaml:"cloud_properties"`
				} `yaml:"vm_extensions"`
			}
			output := bbl.CloudConfig()
			err := yaml.Unmarshal([]byte(output), &cloudConfig)
			Expect(err).NotTo(HaveOccurred())

			var names []string
			for _, extension := range cloudConfig.VMExtensions {
				names = append(names, extension.Name)
			}

			Expect(names).To(ContainElement("cf-router-network-properties"))
			Expect(names).To(ContainElement("diego-ssh-proxy-network-properties"))
			Expect(names).To(ContainElement("cf-tcp-router-network-properties"))
		})

		By("verifying the bbl lbs output", func() {
			if iaas == "azure" {
				return
			}

			stdout := bbl.Lbs()
			Expect(stdout).To(MatchRegexp("CF Router LB:.*"))
			Expect(stdout).To(MatchRegexp("CF SSH Proxy LB:.*"))
			Expect(stdout).To(MatchRegexp("CF TCP Router LB:.*"))

			if iaas == "gcp" {
				Expect(stdout).To(MatchRegexp("CF WebSocket LB:.*"))
				Expect(stdout).To(MatchRegexp("CF Credhub LB:.*"))
			}
		})

		By("verifying that modified scripts were run", func() {
			createEnvOutput, err := ioutil.ReadFile(createEnvOutputPath)
			Expect(err).NotTo(HaveOccurred())
			Expect(string(createEnvOutput)).To(Equal("jumpbox\ndirector\n"))
		})

		By("destroying the director and the jumpbox", func() {
			session := bbl.Down()
			Eventually(session, 10*time.Minute).Should(gexec.Exit(0))
		})

		By("verifying that artifacts are removed from state dir", func() {
			f, err := os.Open(stateDir)
			Expect(err).NotTo(HaveOccurred())

			filenames, err := f.Readdirnames(0)
			Expect(err).NotTo(HaveOccurred())
			Expect(filenames).To(BeEmpty())
		})
	})
})

func getJumpboxPrivateKey(v string) string {
	var vars struct {
		JumpboxSSH struct {
			PrivateKey string `yaml:"private_key"`
		} `yaml:"jumpbox_ssh"`
	}

	err := yaml.Unmarshal([]byte(v), &vars)
	Expect(err).NotTo(HaveOccurred())

	return vars.JumpboxSSH.PrivateKey
}

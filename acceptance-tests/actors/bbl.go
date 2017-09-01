package actors

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"os/exec"
	"strings"
	"time"

	acceptance "github.com/cloudfoundry/bosh-bootloader/acceptance-tests"

	. "github.com/onsi/gomega"
	"github.com/onsi/gomega/gexec"
)

type BBL struct {
	stateDirectory string
	pathToBBL      string
	configuration  acceptance.Config
	envID          string
}

type IAAS int

func NewBBL(stateDirectory string, pathToBBL string, configuration acceptance.Config, envIDSuffix string) BBL {
	envIDPrefix := os.Getenv("BBL_TEST_ENV_ID_PREFIX")
	if envIDPrefix == "" {
		envIDPrefix = "bbl-test"
	}

	return BBL{
		stateDirectory: stateDirectory,
		pathToBBL:      pathToBBL,
		configuration:  configuration,
		envID:          fmt.Sprintf("%s-%s", envIDPrefix, envIDSuffix),
	}
}

func (b BBL) PredefinedEnvID() string {
	return b.envID
}

func (b BBL) Up(additionalArgs ...string) *gexec.Session {
	args := []string{
		"--state-dir", b.stateDirectory,
		"--debug",
		"up",
	}

	args = append(args, additionalArgs...)

	return b.execute(args, os.Stdout, os.Stderr)
}

func (b BBL) Rotate() *gexec.Session {
	return b.execute([]string{
		"--state-dir", b.stateDirectory,
		"rotate",
	}, os.Stdout, os.Stderr)
}

func (b BBL) Destroy() *gexec.Session {
	return b.execute([]string{
		"--state-dir", b.stateDirectory,
		"destroy",
		"--no-confirm",
	}, os.Stdout, os.Stderr)
}

func (b BBL) Down() *gexec.Session {
	return b.execute([]string{
		"--state-dir", b.stateDirectory,
		"down",
		"--no-confirm",
	}, os.Stdout, os.Stderr)
}

func (b BBL) CreateLB(loadBalancerType string, cert string, key string, chain string) *gexec.Session {
	args := []string{
		"--state-dir", b.stateDirectory,
		"--debug",
		"create-lbs",
		"--type", loadBalancerType,
	}

	if loadBalancerType == "cf" || b.configuration.IAAS == "aws" {
		args = append(args,
			"--cert", cert,
			"--key", key,
			"--chain", chain,
		)
	}

	return b.execute(args, os.Stdout, os.Stderr)
}

func (b BBL) UpdateLB(certPath, keyPath, chainPath string) *gexec.Session {
	args := []string{
		"--state-dir", b.stateDirectory,
		"--debug",
		"update-lbs",
		"--cert", certPath,
		"--key", keyPath,
	}

	if chainPath != "" {
		args = append(args, "--chain", chainPath)
	}

	return b.execute(args, os.Stdout, os.Stderr)
}

func (b BBL) LBs() *gexec.Session {
	args := []string{
		"--state-dir", b.stateDirectory,
		"lbs",
	}

	return b.execute(args, os.Stdout, os.Stderr)
}

func (b BBL) DeleteLBs() *gexec.Session {
	args := []string{
		"--state-dir", b.stateDirectory,
		"delete-lbs",
	}

	return b.execute(args, os.Stdout, os.Stderr)
}

func (b BBL) DirectorUsername() string {
	return b.fetchValue("director-username")
}

func (b BBL) DirectorPassword() string {
	return b.fetchValue("director-password")
}

func (b BBL) DirectorAddress() string {
	return b.fetchValue("director-address")
}

func (b BBL) DirectorCACert() string {
	return b.fetchValue("director-ca-cert")
}

func (b BBL) JumpboxAddress() string {
	return b.fetchValue("jumpbox-address")
}

func (b BBL) SSHKey() string {
	return b.fetchValue("ssh-key")
}

func (b BBL) EnvID() string {
	return b.fetchValue("env-id")
}

func (b BBL) BOSHDeploymentVars() string {
	return b.fetchValue("bosh-deployment-vars")
}

func (b BBL) CloudConfig() string {
	return b.fetchValue("cloud-config")
}

func (b BBL) PrintEnv() string {
	return b.fetchValue("print-env")
}

func (b BBL) SaveDirectorCA() string {
	stdout := bytes.NewBuffer([]byte{})
	session := b.execute([]string{
		"--state-dir", b.stateDirectory,
		"director-ca-cert",
	}, stdout, os.Stderr)
	Eventually(session, 10*time.Minute).Should(gexec.Exit(0))

	file, err := ioutil.TempFile("", "")
	defer file.Close()
	Expect(err).NotTo(HaveOccurred())

	file.Write(stdout.Bytes())

	return file.Name()
}

func (b BBL) fetchValue(value string) string {
	args := []string{
		"--state-dir", b.stateDirectory,
		value,
	}

	stdout := bytes.NewBuffer([]byte{})
	stderr := bytes.NewBuffer([]byte{})
	b.execute(args, stdout, stderr).Wait(30 * time.Second)

	return strings.TrimSpace(string(stdout.Bytes()))
}

func (b BBL) execute(args []string, stdout io.Writer, stderr io.Writer) *gexec.Session {
	cmd := exec.Command(b.pathToBBL, args...)
	session, err := gexec.Start(cmd, stdout, stderr)
	Expect(err).NotTo(HaveOccurred())

	return session
}

func LBURL(config acceptance.Config, bbl BBL, state acceptance.State) (string, error) {
	lbs := bbl.fetchValue("lbs")
	var url string
	if config.IAAS == "aws" {
		cutLBsPrefix := strings.Split(lbs, "[")[1]
		url = strings.Split(cutLBsPrefix, "]")[0]
	} else {
		url = strings.Split(lbs, " ")[2]
	}

	return fmt.Sprintf("https://%s", url), nil
}

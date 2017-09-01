package azure

const (
	BaseOps = `
azs:
- name: z1
- name: z2
- name: z3

vm_types:
- name: default
  cloud_properties:
    instance_type: Standard_D1_v2
- name: large
  cloud_properties:
    instance_type: Standard_D3_v2

disk_types:
- name: default
  disk_size: 3000
- name: large
  disk_size: 50_000

networks:
- name: vip
  type: vip

compilation:
  workers: 5
  reuse_compilation_vms: true
  az: z1
  vm_type: default
  network: default
`
)

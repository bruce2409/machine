package provision

import (
	"github.com/docker/machine/libmachine/drivers"
)

func init() {
	Register("AlmaLinux", &RegisteredProvisioner{
		New: NewAlmaLinuxProvisioner,
	})
}

func NewAlmaLinuxProvisioner(d drivers.Driver) Provisioner {
	return &AlmaLinuxProvisioner{
		NewRedHatProvisioner("almalinux", d),
	}
}

type AlmaLinuxProvisioner struct {
	*RedHatProvisioner
}

func (provisioner *AlmaLinuxProvisioner) String() string {
	return "almalinux"
}

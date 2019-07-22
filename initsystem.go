package initsystem

import (
	"fmt"
	"os/exec"
)

// InitSystem is the interface that describe behaviors of an init system
type InitSystem interface {
	Start(service string) error
	Stop(service string) error
	Enable(service string) error
	Disable(service string) error
	IsActive(service string) (bool, error)
	IsEnabled(service string) (bool, error)
	EnableAndStartService(service string) error
	DisableAndStopService(service string) error
}

// GetInitSystem returns an InitSystem for the current system, or error
// if we cannot detect a supported init system.
func GetInitSystem() (InitSystem, error) {
	_, err := exec.LookPath("systemctl")
	if err == nil {
		return &SystemdInitSystem{}, nil
	}

	return nil, fmt.Errorf("no supported init system detected")
}

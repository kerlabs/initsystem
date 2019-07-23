package initsystem

import (
	"fmt"
	"os/exec"
	"strings"
)

type SystemdInitSystem struct{}

func (s SystemdInitSystem) reloadSystemd() error {
	if err := exec.Command("systemctl", "daemon-reload").Run(); err != nil {
		return fmt.Errorf("failed to reload systemd: %v", err)
	}
	return nil
}

// Start a service
func (s SystemdInitSystem) Start(service string) error {
	// Before we try to start any service, make sure that systemd is ready
	if err := s.reloadSystemd(); err != nil {
		return err
	}
	args := []string{"start", service}
	return exec.Command("systemctl", args...).Run()
}

// Stop a service
func (s SystemdInitSystem) Stop(service string) error {
	// Before we try to start any service, make sure that systemd is ready
	if err := s.reloadSystemd(); err != nil {
		return err
	}
	args := []string{"stop", service}
	return exec.Command("systemctl", args...).Run()
}

func (s SystemdInitSystem) Restart(service string) error {
	// Before we try to start any service, make sure that systemd is ready
	if err := s.reloadSystemd(); err != nil {
		return err
	}
	args := []string{"restart", service}
	return exec.Command("systemctl", args...).Run()
}

// Enable a service
func (s SystemdInitSystem) Enable(service string) error {
	// Before we try to enable any service, make sure that systemd is ready
	if err := s.reloadSystemd(); err != nil {
		return err
	}
	args := []string{"enable", service}
	return exec.Command("systemctl", args...).Run()
}

// Disable a service
func (s SystemdInitSystem) Disable(service string) error {
	// Before we try to disable any service, make sure that systemd is ready
	if err := s.reloadSystemd(); err != nil {
		return err
	}
	args := []string{"disable", service}
	return exec.Command("systemctl", args...).Run()
}

// IsActive checks if the systemd unit is active
func (s SystemdInitSystem) IsActive(service string) (bool, error) {
	args := []string{"is-active", service}
	outBytes, err := exec.Command("systemctl", args...).Output()
	if err != nil {
		return false, fmt.Errorf("failed to run command: %s", err)
	}
	out := strings.TrimSpace(string(outBytes))
	if out == "active" || out == "activating" {
		return true, nil
	}
	return false, nil
}

// IsEnabled checks if the systemd unit is enabled
func (s SystemdInitSystem) IsEnabled(service string) (bool, error) {
	args := []string{"is-enabled", service}
	outBytes, err := exec.Command("systemctl", args...).Output()
	if err != nil {
		return false, fmt.Errorf("failed to run command: %s", err)
	}
	out := strings.TrimSpace(string(outBytes))
	if out == "enabled" {
		return true, nil
	}
	return false, nil
}

func (s SystemdInitSystem) IsExists(service string) (bool, error) {
	args := []string{"status", service}
	outBytes, err := exec.Command("systemctl", args...).Output()
	if err != nil {
		return false, fmt.Errorf("failed to run command: %s", err)
	}
	out := strings.TrimSpace(string(outBytes))
	if strings.Contains(out, "could not be found") || strings.Contains(out, "Loaded: not-found") {
		return false, nil
	}
	return true, nil
}

package runc

import (
	"encoding/json"
	"fmt"
	"os/exec"
	"time"
)

type Runc struct {
	command string
}

func New() *Runc {
	return &Runc{
		command: "runc",
	}
}

// stdout,stderr cannot be extracted from execution
// or the command will hang indefinitely
func (r *Runc) Create(id, bundlePath, pidPath string) error {
	return exec.Command(r.command, "create", "--bundle", bundlePath, "--pid-file", pidPath, id).Run()
}

func (r *Runc) Start(id string) error {
	return exec.Command(r.command, "start", id).Run()
}

func (r *Runc) HasContainer(id string) (bool, error) {
	output, err := exec.Command(r.command, "list", "--format", "json").Output()
	if err != nil {
		return false, fmt.Errorf("failed to query vms: %s", err)
	}

	var vms []struct {
		ID string `json:"id"`
	}

	err = json.Unmarshal(output, &vms)
	if err != nil {
		return false, fmt.Errorf("failed to query vms: %s", err)
	}

	for _, vm := range vms {
		if vm.ID == id {
			return true, nil
		}
	}

	return false, nil
}

func (r *Runc) DeleteContainer(id string) {
	r.stopContainer(id)
	exec.Command(r.command, "delete", id).Run()
}

func (r *Runc) stopContainer(id string) {
	exec.Command(r.command, "kill", "--all", id).Run()

	timeout := time.After(10 * time.Second)

	for {
		select {
		case <-timeout:
			exec.Command(r.command, "kill", "--all", id, "KILL").Run()
			return
		default:
			status, _ := r.containerStatus(id)
			if status == "stopped" {
				return
			}
		}
	}
}

func (r *Runc) containerStatus(id string) (string, error) {
	output, err := exec.Command(r.command, "state", id).Output()
	if err != nil {
		return "", fmt.Errorf("failed to query container %s: %s", id, err)
	}

	var vm struct {
		Status string `json:"status"`
	}

	err = json.Unmarshal(output, &vm)
	if err != nil {
		return "", fmt.Errorf("failed to parse status for container %s: %s", id, err)
	}

	return vm.Status, nil
}


package runc

import (
	"encoding/json"
	"fmt"
	"os/exec"
)

type Runc struct {
	command string
}

func New() *Runc {
	return &Runc{
		command: "runc",
	}
}

func (r *Runc) Run(id, bundlePath string) error {
	return exec.Command(r.command, "run", "--bundle", bundlePath, "--detach", id).Run()
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
	exec.Command(r.command, "kill", "--all", id, "KILL").Run()
	exec.Command(r.command, "delete", id).Run()
}

package runc

import (
	"encoding/json"
	"fmt"
	"github.com/aemengo/bosh-runc-cpi/utils"
	"os/exec"
	"regexp"
	"strings"
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

func (r *Runc) StopProcesses(id string) error {
	utils.Do(3, time.Second, func() error {
		return exec.Command(r.command, "exec", id, "monit", "stop", "all").Run()
	})

	var (
		timeout = time.After(5 * time.Minute)
		ticker  = time.NewTicker(2 * time.Second)
		output  []byte
	)

	for {
		select {
		case <-timeout:
			return fmt.Errorf("timed out waiting for processes to stop on vm: %s: %s", id, output)
		case <-ticker.C:
			var err error
			output, err = exec.Command(r.command, "exec", id, "monit", "summary").Output()
			if err != nil {
				return fmt.Errorf("failed to query monit status of %s", id)
			}

			if allProcesses(output, "not monitored") {
				return nil
			}
		}
	}
}

func (r *Runc) StartProcesses(id string) error {
	utils.Do(3, 5*time.Second, func() error {
		return exec.Command(r.command, "exec", id, "monit", "start", "all").Run()
	})

	var (
		timeout = time.After(5 * time.Minute)
		ticker  = time.NewTicker(2 * time.Second)
		output  []byte
	)

	for {
		select {
		case <-timeout:
			return fmt.Errorf("timed out waiting for processes to start on vm: %s: %s", id, output)
		case <-ticker.C:
			var err error
			output, err = exec.Command(r.command, "exec", id, "monit", "summary").Output()
			if err != nil {
				return fmt.Errorf("failed to query monit status of %s", id)
			}

			if allProcesses(output, "running") {
				return nil
			}
		}
	}
}

func (r *Runc) Checkpoint(id, imagePath, workPath, parentPath string) error {
	// sometimes the bosh-agent process is 'busy'
	// so we try 'till we get a point in time
	// when it doesn't interfere with checkpoint
	return utils.Do(11, time.Second, func() error {
		return utils.RunCommand(
			r.command,
			"checkpoint",
			"--tcp-established",
			"--image-path", imagePath,
			"--work-path", workPath,
			"--parent-path", parentPath,
			id,
		)
	})
}

func (r *Runc) Restore(id, bundlePath, imagePath, workPath, pidPath string) error {
	return utils.RunCommand(
		r.command,
		"restore",
		"-d",
		"--image-path", imagePath,
		"--work-path", workPath,
		"--pid-file", pidPath,
		"--bundle", bundlePath,
		id,
	)
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

func allProcesses(output []byte, status string) bool {
	for _, line := range strings.Split(string(output), "\n") {
		r := regexp.MustCompile(`^(Process|System)`)
		s := regexp.MustCompile(`\s` + status + `$`)

		if r.MatchString(line) {
			if !s.MatchString(line) {
				return false
			}
		}
	}

	return true
}

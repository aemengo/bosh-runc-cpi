package utils

import (
	"os"
	"os/exec"
	"fmt"
)

func Exists(path string) bool  {
	_, err := os.Stat(path)
	if os.IsNotExist(err) {
		return false
	}

	return true
}

func RunCommand(path string, args ...string) error {
	output, err := exec.Command(path, args...).CombinedOutput()
	if err != nil {
		return fmt.Errorf("failed to execute: %s %v: %s: %s", path, args, err, output)
	}

	return nil
}
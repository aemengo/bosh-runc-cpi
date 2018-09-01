package utils

import (
	"os"
	"os/exec"
	"fmt"
	"time"
)

func Do(attempts int, delay time.Duration, task func() error) (err error) {
	ticker := time.NewTicker(delay)

	for {
		select {
		case <-ticker.C:
			err = task()
			if err == nil {
				return
			}

			attempts--
			if attempts == 0 {
				return
			}
		}
	}
}

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

func MkdirAll(dirs ...string) error {
	for _, dir := range dirs {
		if err := os.MkdirAll(dir, os.ModePerm); err != nil {
			return fmt.Errorf("failed to make dir at %s: %s", dir, err)
		}
	}
	return nil
}

func RemoveAll(dirs ...string) error {
	for _, dir := range dirs {
		if err := os.RemoveAll(dir); err != nil {
			return fmt.Errorf("failed to make dir at %s: %s", dir, err)
		}
	}

	return nil
}
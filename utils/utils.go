package utils

import (
	"encoding/json"
	"fmt"
	"os"
	"os/exec"
	"time"
)

func Do(attempts int, delay time.Duration, task func() error) (err error) {
	var (
		currentAttempt = 0
		maxAttempts    = attempts
		ticker         = time.NewTicker(delay)
	)

	for {
		select {
		case <-ticker.C:
			currentAttempt++
			attempts--

			err = task()
			if err == nil {
				return
			}

			err = fmt.Errorf("[%d/%d] %s", currentAttempt, maxAttempts, err)

			if attempts == 0 {
				return
			}
		}
	}
}

func Exists(path string) bool {
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
			return fmt.Errorf("failed to remove dir at %s: %s", dir, err)
		}
	}

	return nil
}

func EncodeFile(path string, obj interface{}) error {
	f, err := os.Create(path)
	if err != nil {
		return fmt.Errorf("failed to write at %s: %s", path, err)
	}
	defer f.Close()

	return json.NewEncoder(f).Encode(obj)
}

func DecodeFile(path string, obj interface{}) error {
	f, err := os.Open(path)
	if err != nil {
		return fmt.Errorf("failed to read at %s: %s", path, err)
	}
	defer f.Close()

	return json.NewDecoder(f).Decode(obj)
}

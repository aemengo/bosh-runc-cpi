package service

import (
	"path/filepath"
	"os"
	"io/ioutil"
)

func saveDiskState(vmPath string, diskID string) {
	diskStateDir := filepath.Join(vmPath, "disks")
	diskStatePath := filepath.Join(diskStateDir, diskID)

	os.MkdirAll(diskStateDir, os.ModePerm)
	ioutil.WriteFile(diskStatePath, []byte(""), 0666)
}

func removeDiskState(vmPath string, diskID string) {
	diskStatePath := filepath.Join(vmPath, "disks", diskID)

	os.RemoveAll(diskStatePath)
}

func vmsWithPersistentDisk(vmsPath string, diskID string) []string {
	files, _ := filepath.Glob(filepath.Join(vmsPath, "*", "disks", diskID))

	var vms []string

	for _, file := range files {
		diskDir := filepath.Dir(file)
		vmPath := filepath.Dir(diskDir)
		vmID := filepath.Base(vmPath)
		vms = append(vms, vmID)
	}

	return vms
}
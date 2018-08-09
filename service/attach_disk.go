package service

import (
	"context"
	"fmt"
	"github.com/aemengo/bosh-containerd-cpi/pb"
	"io/ioutil"
	"path/filepath"
)

func (s *Service) AttachDisk(ctx context.Context, req *pb.DisksOpts) (*pb.Void, error) {
	var (
		persistentDiskDir = "/persistent-disk"
		diskPath          = filepath.Join(s.config.DiskDir, req.DiskID)
		vmPath            = filepath.Join(s.config.VMDir, req.VmID)
		pidPath           = filepath.Join(vmPath, "pid")
		specPath          = filepath.Join(vmPath, "config.json")
		agentSettingsPath = filepath.Join(vmPath, "warden-cpi-agent-env.json")
	)

	spec, err := ioutil.ReadFile(specPath)
	if err != nil {
		return nil, fmt.Errorf("failed to read spec file: %s", err)
	}

	spec = attachBindMount(spec, diskPath, persistentDiskDir)

	agentSettings, err := ioutil.ReadFile(agentSettingsPath)
	if err != nil {
		return nil, fmt.Errorf("failed to read agent settings file: %s", err)
	}

	agentSettings = attachPersistentDisk(agentSettings, req.DiskID, persistentDiskDir)

	s.runc.DeleteContainer(req.VmID)

	if err := ioutil.WriteFile(specPath, spec, 0666); err != nil {
		return nil, fmt.Errorf("failed to write spec file: %s", err)
	}

	if err := ioutil.WriteFile(agentSettingsPath, agentSettings, 0666); err != nil {
		return nil, fmt.Errorf("failed to write agent settings file: %s", err)
	}

	err = s.startContainer(ctx, req.VmID, vmPath, pidPath, agentSettings)
	if err != nil {
		return nil, err
	}

	saveDiskState(vmPath, req.DiskID)

	return &pb.Void{}, nil
}

package service

import (
	"github.com/aemengo/bosh-containerd-cpi/pb"
	"context"
	"path/filepath"
	"io/ioutil"
	"fmt"
)

func (s *Service) DetachDisk(ctx context.Context, req *pb.DisksOpts) (*pb.Void, error) {
	var (
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

	spec = detachBindMount(spec, diskPath)

	agentSettings, err := ioutil.ReadFile(agentSettingsPath)
	if err != nil {
		return nil, fmt.Errorf("failed to read agent settings file: %s", err)
	}

	agentSettings = detachPersistentDisk(agentSettings)

	s.runc.DeleteContainer(req.VmID)

	if err := ioutil.WriteFile(specPath, spec, 0666); err != nil {
		return nil, fmt.Errorf("failed to write spec file: %s", err)
	}

	if err := ioutil.WriteFile(agentSettingsPath, agentSettings, 0666); err != nil {
		return nil, fmt.Errorf("failed to write agent settings file: %s", err)
	}

	err = s.runc.Create(req.VmID, vmPath, pidPath)
	if err != nil {
		return nil, fmt.Errorf("failed to create container: %s", err)
	}

	ip, mask, gatewayIP, err := extractNetValues(agentSettings)
	if err != nil {
		return nil, fmt.Errorf("failed to extract network values: %s", err)
	}

	err = s.configureNetworking(vmPath, pidPath, ip, mask, gatewayIP)
	if err != nil {
		return nil, fmt.Errorf("failed to configure networker: %s", err)
	}

	err = s.runc.Start(req.VmID)
	if err != nil {
		return nil, fmt.Errorf("failed to start container: %s", err)
	}

	return &pb.Void{}, nil
}
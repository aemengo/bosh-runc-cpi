package service

import (
	"context"
	"fmt"
	"github.com/aemengo/bosh-runc-cpi/pb"
	"io/ioutil"
	"path/filepath"
)

func (s *Service) RestoreVM(ctx context.Context, req *pb.VMFilterOpts) (*pb.Void, error) {
	vmIDs, err := s.vmIDs(req)
	if err != nil {
		return nil, err
	}

	for _, vmID := range vmIDs {
		var (
			vmPath            = filepath.Join(s.config.VMDir, vmID)
			imagePath         = filepath.Join(vmPath, "criu", "image")
			workPath          = filepath.Join(vmPath, "criu", "work")
			agentSettingsPath = filepath.Join(vmPath, "warden-cpi-agent-env.json")
			pidPath           = filepath.Join(vmPath, "pid")
		)

		err = s.runc.Restore(vmID, vmPath, imagePath, workPath, pidPath)
		if err != nil {
			return nil, err
		}

		agentSettings, err := ioutil.ReadFile(agentSettingsPath)
		if err != nil {
			return nil, fmt.Errorf("failed to read agent settings file: %s", err)
		}

		ip, mask, gatewayIP, err := extractNetValues(agentSettings)
		if err != nil {
			return nil, fmt.Errorf("failed to extract network values: %s", err)
		}

		err = s.configureNetworking(vmPath, pidPath, ip, mask, gatewayIP)
		if err != nil {
			return nil, fmt.Errorf("failed to configure networker: %s", err)
		}

		err = s.runc.StartProcesses(vmID)
		if err != nil {
			return nil, err
		}
	}

	return &pb.Void{}, nil
}

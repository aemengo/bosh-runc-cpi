package service

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/aemengo/bosh-containerd-cpi/pb"
	"io/ioutil"
	"path/filepath"
)

func (s *Service) AttachDisk(ctx context.Context, req *pb.AttachDisksOpts) (*pb.Void, error) {
	var (
		persistentDiskDir = "/persistent-disk"
		diskPath          = filepath.Join(s.config.DiskDir, req.DiskID)
		vmPath            = filepath.Join(s.config.VMDir, req.VmID)
		pidPath           = filepath.Join(vmPath, "pid")
		specPath          = filepath.Join(vmPath, "config.json")
		agentSettingsPath = filepath.Join(vmPath, "warden-cpi-agent-env.json")
	)
	contents, err := ioutil.ReadFile(specPath)
	if err != nil {
		return nil, fmt.Errorf("failed to read spec file: %s", err)
	}

	var spec map[string]interface{}
	err = json.Unmarshal(contents, &spec)
	if err != nil {
		return nil, fmt.Errorf("failed to parse spec file: %s", err)
	}

	if mounts, ok := spec["mounts"].([]interface{}); ok {
		spec["mounts"] = append(mounts, map[string]interface{}{
			"source":      diskPath,
			"destination": persistentDiskDir,
			"type":        "bind",
			"options": []string{
				"rw",
				"rbind",
				"rprivate",
			},
		})
	}

	agentContents, err := ioutil.ReadFile(agentSettingsPath)
	if err != nil {
		return nil, fmt.Errorf("failed to read agent settings file: %s", err)
	}

	var agentSettings map[string]interface{}
	err = json.Unmarshal(agentContents, &agentSettings)
	if err != nil {
		return nil, fmt.Errorf("failed to parse spec file: %s", err)
	}

	agentSettings["disks"] = map[string]interface{}{
		"system":    "",
		"ephemeral": nil,
		"persistent": map[string]interface{}{
			req.DiskID: map[string]string{
				"path": persistentDiskDir,
			},
		},
	}

	s.runc.DeleteContainer(req.VmID)

	newSpec, _ := json.Marshal(spec)
	newAgentSettings, _ := json.Marshal(agentSettings)

	if err := ioutil.WriteFile(specPath, newSpec, 0666); err != nil {
		return nil, fmt.Errorf("failed to write spec file: %s", err)
	}

	if err := ioutil.WriteFile(agentSettingsPath, newAgentSettings, 0666); err != nil {
		return nil, fmt.Errorf("failed to write agent settings file: %s", err)
	}

	err = s.runc.Create(req.VmID, vmPath, pidPath)
	if err != nil {
		return nil, fmt.Errorf("failed to create container: %s", err)
	}

	ip, mask, gatewayIP, err := s.extractNetValues(agentContents)
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

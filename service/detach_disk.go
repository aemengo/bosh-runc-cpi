package service

import (
	"context"
	"fmt"
	"github.com/aemengo/bosh-runc-cpi/pb"
	"io/ioutil"
	"path/filepath"
	"github.com/opencontainers/runtime-spec/specs-go"
	"github.com/aemengo/bosh-runc-cpi/utils"
	"github.com/aemengo/bosh-runc-cpi/runc"
)

func (s *Service) DetachDisk(ctx context.Context, req *pb.DisksOpts) (*pb.Void, error) {
	var diskPath = filepath.Join(s.config.DiskDir, req.DiskID)

	// The cpi ignores calls of method type 'set_disk_metadata' and 'set_vm_metadata'
	// so the VM id passed by the CPI is unreliable for 'detach_disk'
	// we must use an alternate means to tracking vm to disk associations
	for _, vmID := range vmsWithPersistentDisk(s.config.VMDir, req.DiskID) {

		var (
			vmPath            = filepath.Join(s.config.VMDir, vmID)
			pidPath           = filepath.Join(vmPath, "pid")
			specPath          = filepath.Join(vmPath, "config.json")
			agentSettingsPath = filepath.Join(vmPath, "warden-cpi-agent-env.json")
		)

		spec := &specs.Spec{}
		err := utils.DecodeFile(specPath, spec)
		if err != nil {
			return nil, fmt.Errorf("failed to read spec file: %s", err)
		}

		runc.Apply(spec, runc.WithoutMount(diskPath))

		agentSettings, err := ioutil.ReadFile(agentSettingsPath)
		if err != nil {
			return nil, fmt.Errorf("failed to read agent settings file: %s", err)
		}

		agentSettings = detachPersistentDisk(agentSettings)

		s.runc.DeleteContainer(vmID)

		if err := utils.EncodeFile(specPath, spec); err != nil {
			return nil, fmt.Errorf("failed to write spec file: %s", err)
		}

		if err := ioutil.WriteFile(agentSettingsPath, agentSettings, 0666); err != nil {
			return nil, fmt.Errorf("failed to write agent settings file: %s", err)
		}

		err = s.startContainer(ctx, vmID, vmPath, pidPath, agentSettings)
		if err != nil {
			return nil, err
		}

		removeDiskState(vmPath, req.DiskID)
	}

	return &pb.Void{}, nil
}

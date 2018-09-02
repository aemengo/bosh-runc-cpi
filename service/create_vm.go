package service

import (
	"context"
	"fmt"
	"github.com/aemengo/bosh-runc-cpi/pb"
	"github.com/aemengo/bosh-runc-cpi/runc"
	"github.com/aemengo/bosh-runc-cpi/utils"
	"github.com/opencontainers/runtime-spec/specs-go"
	"github.com/satori/go.uuid"
	"io/ioutil"
	"path/filepath"
)

func (s *Service) CreateVM(ctx context.Context, req *pb.CreateVMOpts) (*pb.TextParcel, error) {
	var (
		id                = uuid.NewV4().String()
		persistentDiskDir = "/persistent-disk"
		vmPath            = filepath.Join(s.config.VMDir, id)
		rootFsPath        = filepath.Join(vmPath, "rootfs")
		workDirPath       = filepath.Join(vmPath, "workdir")
		upperDirPath      = filepath.Join(vmPath, "upperdir")
		specPath          = filepath.Join(vmPath, "config.json")
		pidPath           = filepath.Join(vmPath, "pid")
		agentSettingsPath = filepath.Join(vmPath, "warden-cpi-agent-env.json")
		stemcellPath      = filepath.Join(s.config.StemcellDir, req.StemcellID)
		agentSettings     = attachVMID(req.AgentSettings, id)
		spec              = runc.DefaultSpec()
	)

	err := utils.MkdirAll(rootFsPath, workDirPath, upperDirPath)
	if err != nil {
		return nil, fmt.Errorf("failed to create vm directory: %s", err)
	}

	if req.DiskID != "" {
		agentSettings = attachPersistentDisk(agentSettings, req.DiskID, persistentDiskDir)
	}

	err = ioutil.WriteFile(agentSettingsPath, agentSettings, 0666)
	if err != nil {
		return nil, fmt.Errorf("failed to write agent settings: %s", err)
	}

	// Destination "/" mount must be prepended because it
	// loads the entire filesystem for subsequent mounts
	runc.Apply(spec, runc.PrependMounts([]specs.Mount{
		{
			Destination: "/",
			Source:      "overlay",
			Type:        "overlay",
			Options: []string{
				"suid",
				"upperdir=" + upperDirPath,
				"lowerdir=" + stemcellPath,
				"workdir=" + workDirPath,
			},
		},
		{
			Destination: "/var/vcap/bosh/warden-cpi-agent-env.json",
			Source:      agentSettingsPath,
			Type:        "bind",
			Options: []string{
				"mode=666",
				"bind",
			},
		},
	}))

	if req.DiskID != "" {
		runc.Apply(spec, runc.AppendMounts([]specs.Mount{{
			Destination: persistentDiskDir,
			Source:      filepath.Join(s.config.DiskDir, req.DiskID),
			Type:        "bind",
			Options: []string{
				"bind",
				"rw",
			},
		}}))
	}

	err = utils.EncodeFile(specPath, spec)
	if err != nil {
		return nil, fmt.Errorf("failed to write container spec: %s", err)
	}

	err = s.startContainer(ctx, id, vmPath, pidPath, agentSettings, true)
	if err != nil {
		return nil, err
	}

	if req.DiskID != "" {
		saveDiskState(vmPath, req.DiskID)
	}

	return &pb.TextParcel{Value: id}, nil
}

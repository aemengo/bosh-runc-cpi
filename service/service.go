package service

import (
	cfg "github.com/aemengo/bosh-runc-cpi/config"
	nt "github.com/aemengo/bosh-runc-cpi/network"
	rc "github.com/aemengo/bosh-runc-cpi/runc"
	"log"
	"github.com/aemengo/bosh-runc-cpi/pb"
	"fmt"
	"context"
	"path/filepath"
)

//go:generate protoc -I ../pb --go_out=plugins=grpc:../pb ../pb/messages.proto

type Service struct {
	config  cfg.Config
	logger  *log.Logger
	runc    *rc.Runc
	network *nt.Network
}

func New(config cfg.Config, runc *rc.Runc, network *nt.Network, logger *log.Logger) *Service {
	return &Service{
		config:  config,
		runc:    runc,
		network: network,
		logger:  logger,
	}
}

func (s *Service) startContainer(ctx context.Context, id string, vmPath string, pidPath string, agentSettings []byte, options ...bool) error {
	var deleteOnError bool

	if len(options) > 0 {
		deleteOnError = options[0]
	}

	err := s.runc.Create(id, vmPath, pidPath)
	if err != nil {
		return fmt.Errorf("failed to create container: %s", err)
	}

	ip, mask, gatewayIP, err := extractNetValues(agentSettings)
	if err != nil {
		if deleteOnError { s.DeleteVM(ctx, &pb.TextParcel{Value: id}) }
		return fmt.Errorf("failed to extract network values: %s", err)
	}

	err = s.configureNetworking(vmPath, pidPath, ip, mask, gatewayIP)
	if err != nil {
		if deleteOnError { s.DeleteVM(ctx, &pb.TextParcel{Value: id}) }
		return fmt.Errorf("failed to configure networker: %s", err)
	}

	err = s.runc.Start(id)
	if err != nil {
		if deleteOnError { s.DeleteVM(ctx, &pb.TextParcel{Value: id}) }
		return fmt.Errorf("failed to start container: %s", err)
	}

	return nil
}

func (s *Service) vmIDs(req *pb.VMFilterOpts) ([]string, error) {
	if req.VmID != "" {
		return []string{req.VmID}, nil
	}

	if req.All == true {
		return s.vms(), nil
	}

	return nil, fmt.Errorf("must specify a vmID or all to the 'Checkpoint VM' rpc call")
}

func (s *Service) vms() []string {
	files, _ := filepath.Glob(filepath.Join(s.config.VMDir, "*"))

	var vms []string

	for _, file := range files {
		vmID := filepath.Base(file)
		vms = append(vms, vmID)
	}

	return vms
}

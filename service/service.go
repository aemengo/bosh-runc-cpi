package service

import (
	cfg "github.com/aemengo/bosh-containerd-cpi/config"
	nt "github.com/aemengo/bosh-containerd-cpi/network"
	rc "github.com/aemengo/bosh-containerd-cpi/runc"
	"log"
)

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

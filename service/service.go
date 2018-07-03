package service

import (
	cfg "github.com/aemengo/bosh-containerd-cpi/config"
	rc "github.com/aemengo/bosh-containerd-cpi/runc"
	"log"
)

type Service struct {
	config cfg.Config
	logger *log.Logger
	runc   *rc.Runc
}

func New(config cfg.Config, runc *rc.Runc, logger *log.Logger) *Service {
	return &Service{
		config: config,
		runc:   runc,
		logger: logger,
	}
}

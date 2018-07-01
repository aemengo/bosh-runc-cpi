package service

import (
	cfg "github.com/aemengo/bosh-containerd-cpi/config"
	"log"
)

type Service struct {
	config cfg.Config
	logger *log.Logger
}

func New(config cfg.Config, logger *log.Logger) *Service {
	return &Service{
		config: config,
		logger: logger,
	}
}
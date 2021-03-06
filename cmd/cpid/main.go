package main

import (
	"net"
	"log"
	"os"
	"google.golang.org/grpc"
	cfg "github.com/aemengo/bosh-runc-cpi/config"
	rc "github.com/aemengo/bosh-runc-cpi/runc"
	nt "github.com/aemengo/bosh-runc-cpi/network"
	"github.com/aemengo/bosh-runc-cpi/pb"
	"github.com/aemengo/bosh-runc-cpi/service"
	"os/signal"
	"syscall"
	"path/filepath"
)

var logger *log.Logger

func main() {
	logger = log.New(os.Stdout, "[CPID] ", log.LstdFlags)

	if len(os.Args) != 2 {
		logger.Fatalf("[USAGE] %s <config-path>", os.Args[0])
	}

	config, err := cfg.NewServerConfig(os.Args[1])
	expectNoError(err)
	expectNoError(os.MkdirAll(config.VMDir, os.ModePerm))
	expectNoError(os.MkdirAll(config.StemcellDir, os.ModePerm))
	expectNoError(os.MkdirAll(config.DiskDir, os.ModePerm))
	expectNoError(os.MkdirAll(config.TempDir, os.ModePerm))
	expectNoError(prepareUnixSocket(config))

	network, err := nt.New(config.CIDR)
	expectNoError(err)

	lis, err := net.Listen(config.NetworkType, config.NetworkAddress)
	expectNoError(err)

	s := grpc.NewServer()
	pb.RegisterCPIDServer(s, service.New(config, rc.New(), network, logger))

	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGTERM, syscall.SIGKILL)
	go killServerWhenStopped(sigs, s, logger)

	logger.Println("Initializing bosh-runc-cpid...")
	err = s.Serve(lis)
	expectNoError(err)
}

func prepareUnixSocket(config cfg.ServerConfig) error {
	if config.NetworkType != "unix" {
		return nil
	}

	if err := os.RemoveAll(config.NetworkAddress); err != nil {
		return err
	}

	dir := filepath.Dir(config.NetworkAddress)

	return os.MkdirAll(dir, os.ModePerm)
}

func killServerWhenStopped(sigs chan os.Signal, server *grpc.Server, logger *log.Logger) {
	<-sigs
	logger.Println("Shutting down bosh-runc-cpid...")
	server.Stop()
}

func expectNoError(err error) {
	if err != nil {
		logger.Fatalf("failed to initialize: %s\n", err)
	}
}

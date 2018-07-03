package main

import (
	"net"
	"log"
	"os"
	"google.golang.org/grpc"
	cfg "github.com/aemengo/bosh-containerd-cpi/config"
	rc "github.com/aemengo/bosh-containerd-cpi/runc"
	"github.com/aemengo/bosh-containerd-cpi/pb"
	"github.com/aemengo/bosh-containerd-cpi/service"
)

var logger *log.Logger

//go:generate protoc -I ../../pb --go_out=plugins=grpc:../../pb ../../pb/messages.proto

func main() {
	logger = log.New(os.Stdout, "[CPID] ", log.LstdFlags)

	if len(os.Args) != 2 {
		logger.Fatalf("[USAGE] %s <config-path>", os.Args[0])
	}

	config, err := cfg.New(os.Args[1])
	expectNoError(err)
	expectNoError(os.MkdirAll(config.VMDir, os.ModePerm))
	expectNoError(os.MkdirAll(config.StemcellDir, os.ModePerm))
	expectNoError(os.MkdirAll(config.DiskDir, os.ModePerm))

	lis, err := net.Listen("tcp", ":" + config.ServerPort)
	expectNoError(err)

	s := grpc.NewServer()
	runc := rc.New()

	pb.RegisterCPIDServer(s, service.New(config, runc, logger))

	err = s.Serve(lis)
	expectNoError(err)
}

func expectNoError(err error) {
	if err != nil {
		logger.Fatalf("failed to initialize: %s\n", err)
	}
}

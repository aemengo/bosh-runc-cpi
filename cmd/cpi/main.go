package main

import (
	"encoding/json"
	"github.com/aemengo/bosh-containerd-cpi/bosh"
	cfg "github.com/aemengo/bosh-containerd-cpi/config"
	"log"
	"os"
	"context"
	cmd "github.com/aemengo/bosh-containerd-cpi/command"
	"google.golang.org/grpc"
	"github.com/aemengo/bosh-containerd-cpi/pb"
)

func main() {
	if len(os.Args) != 2 {
		log.Fatalf("[USAGE] - | %s <config-path>", os.Args[0])
	}

	config, err := cfg.New(os.Args[1])
	expectNoError(err)

	var args struct {
		Method    string        `json:"method"`
		Arguments []interface{} `json:"arguments"`
		Context   struct {
					  DirectorUUID string `json:"director_uuid"`
				  } `json:"context"`
	}

	expectNoError(json.NewDecoder(os.Stdin).Decode(&args))

	conn, err := grpc.Dial(config.ServerAddr(), grpc.WithInsecure())
	expectNoError(err)
	defer conn.Close()

	ctx := context.Background()
	cpidClient := pb.NewCPIDClient(conn)

	command := cmd.New(ctx, cpidClient, args.Method, args.Arguments, config)
	response := command.Run()
	json.NewEncoder(os.Stdout).Encode(&response)

	if response.Error != nil {
		os.Exit(1)
	} else {
		os.Exit(0)
	}
}

func expectNoError(err error) {
	if err == nil {
		return
	}

	response := bosh.CPIError(
		"failed to initialize",
		err,
		"cpi has been misconfigured",
	)

	json.NewEncoder(os.Stdout).Encode(&response)
	os.Exit(1)
}
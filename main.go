package main

import (
	"encoding/json"
	"github.com/aemengo/bosh-containerd-cpi/bosh"
	cfg "github.com/aemengo/bosh-containerd-cpi/config"
	cmd "github.com/aemengo/bosh-containerd-cpi/command"
	"log"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		log.Println("[USAGE] - | %v <config-path>", os.Args[0])
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

	err = json.NewDecoder(os.Stdin).Decode(&args)
	expectNoError(err)

	command, err := cmd.New(args.Method, args.Arguments, config)
	expectNoError(err)

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

	var response = bosh.Response{
		Error: &bosh.Error{
			Type:      "Bosh::Clouds::ConfigurationError",
			Message:   err.Error(),
			OkToRetry: false,
		},
		Log: "bosh-containerd-cpi has been misconfigured",
	}

	json.NewEncoder(os.Stdout).Encode(&response)
	os.Exit(1)
}

package config

import (
	"errors"
	"fmt"
	"github.com/aemengo/bosh-containerd-cpi/utils"
	"gopkg.in/yaml.v2"
	"os"
)

type Config struct {
	StemcellDir string `yaml:"stemcell_dir"`
	DiskDir     string `yaml:"disk_dir"`
}

func New(configPath string) (Config, error) {
	if !utils.Exists(configPath) {
		return Config{}, errors.New("config does not exist at path " + configPath)
	}

	f, err := os.Open(configPath)
	if err != nil {
		return Config{}, err
	}
	defer f.Close()

	var config Config
	err = yaml.NewDecoder(f).Decode(&config)
	if err != nil {
		return Config{}, fmt.Errorf("failed to decode yaml config at path %s: %s", configPath, err)
	}

	return config, nil
}

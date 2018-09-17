package config

import (
	"errors"
	"fmt"
	"github.com/aemengo/bosh-runc-cpi/utils"
	"gopkg.in/yaml.v2"
	"os"
)

type Config struct {
	StemcellDir    string `yaml:"stemcell_dir"`
	DiskDir        string `yaml:"disk_dir"`
	VMDir          string `yaml:"vm_dir"`
	NetworkType    string `yaml:"network_type"`
	NetworkAddress string `yaml:"address"`
	Agent          Agent  `yaml:"agent"`
}

type Agent struct {
	Mbus      string    `yaml:"mbus"`
	NTP       []string  `yaml:"ntp"`
	Blobstore Blobstore `yaml:"blobstore"`
}

type Blobstore struct {
	Provider string                 `yaml:"provider"`
	Options  map[string]interface{} `yaml:"options"`
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

func (c *Config) ServerAddr() string {
	switch c.NetworkType {
	case "tcp":
		return c.NetworkAddress
	case "unix":
		return "unix://" + c.NetworkAddress
	default:
		panic(fmt.Sprintf("Invalid network type submitted: received '%s', accepting (unix/tcp)", c.NetworkType))
	}
}

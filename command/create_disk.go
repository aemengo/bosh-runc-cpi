package command

import (
	cfg "github.com/aemengo/bosh-containerd-cpi/config"
	"errors"
	"github.com/aemengo/bosh-containerd-cpi/bosh"
	"strconv"
	"path/filepath"
	"github.com/satori/go.uuid"
	"os"
	"io/ioutil"
	"github.com/aemengo/bosh-containerd-cpi/utils"
)

type createDisk struct {
	config cfg.Config
	size int
}

func NewCreateDisk(arguments []interface{}, config cfg.Config) (*createDisk, error) {
	if len(arguments) == 0 {
		return nil, errors.New("invalid disk size passed to create_disk command")
	}

	size, ok := arguments[0].(int)
	if !ok {
		return nil, errors.New("invalid disk size passed to create_disk command")
	}

	return &createDisk{
		config: config,
		size: size,
	}, nil
}

func (c *createDisk) Run() bosh.Response {
	id := uuid.NewV4().String()
	diskPath := filepath.Join(c.config.DiskDir, id)

	err := os.MkdirAll(c.config.DiskDir, os.ModePerm)
	if err != nil {
		return bosh.CPIError("disk directory could not be created", err)
	}

	err = ioutil.WriteFile(diskPath, []byte{}, 0600)
	if err != nil {
		return bosh.CPIError("disk directory could not be created", err)
	}

	sizeStr := strconv.Itoa(c.size) + "MB"

	err = utils.RunCommand("truncate", "-s", sizeStr, diskPath)
	if err != nil {
		return bosh.CPIError("failed to resize disk", err)
	}

	err = utils.RunCommand("/sbin/mkfs", "-t", "ext4", "-F", diskPath)
	if err != nil {
		return bosh.CPIError("failed to build disk filesystem", err)
	}

	return bosh.Response{Result: id}
}
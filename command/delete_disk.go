package command

import (
	cfg "github.com/aemengo/bosh-containerd-cpi/config"
	"errors"
	"github.com/aemengo/bosh-containerd-cpi/bosh"
	"os"
	"path/filepath"
)

type deleteDisk struct {
	config cfg.Config
	diskPath string
}

func NewDeleteDisk(arguments []interface{}, config cfg.Config) (*deleteDisk, error) {
	if len(arguments) == 0 {
		return nil, errors.New("invalid disk path passed to delete_disk command")
	}

	path, ok := arguments[0].(string)
	if !ok {
		return nil, errors.New("invalid disk path passed to delete_disk command")
	}

	return &deleteDisk{
		config: config,
		diskPath: path,
	}, nil
}

func (c *deleteDisk) Run() bosh.Response  {
	diskPath := filepath.Join(c.config.DiskDir, c.diskPath)
	err := os.RemoveAll(diskPath)

	if err != nil {
		return bosh.CPIError("failed to delete disk " + c.diskPath, err)
	} else {
		return bosh.Response{}
	}
}
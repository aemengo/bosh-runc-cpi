package command

import (
	cfg "github.com/aemengo/bosh-containerd-cpi/config"
	"errors"
	"github.com/aemengo/bosh-containerd-cpi/bosh"
	"path/filepath"
	"github.com/aemengo/bosh-containerd-cpi/utils"
)

type hasDisk struct {
	config cfg.Config
	diskPath string
}

func NewHasDisk(arguments []interface{}, config cfg.Config) (*hasDisk, error) {
	if len(arguments) == 0 {
		return nil, errors.New("invalid disk path passed to has_disk command")
	}

	path, ok := arguments[0].(string)
	if !ok {
		return nil, errors.New("invalid disk path passed to has_disk command")
	}

	return &hasDisk{
		config: config,
		diskPath: path,
	}, nil
}

func (c *hasDisk) Run() bosh.Response {
	diskPath := filepath.Join(c.config.DiskDir, c.diskPath)
	exists := utils.Exists(diskPath)
	return bosh.Response{Result: exists}
}
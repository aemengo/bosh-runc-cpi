package command

import (
	"github.com/aemengo/bosh-containerd-cpi/bosh"
	cfg "github.com/aemengo/bosh-containerd-cpi/config"
	"os"
	"errors"
	"path/filepath"
)

type deleteStemcell struct {
	config cfg.Config
	stemcellPath string
}

func NewDeleteStemcell(arguments []interface{}, config cfg.Config) (*deleteStemcell, error) {
	if len(arguments) == 0 {
		return nil, errors.New("invalid stemcell path passed to delete_stemcell command")
	}

	path, ok := arguments[0].(string)
	if !ok {
		return nil, errors.New("invalid stemcell path passed to delete_stemcell command")
	}

	return &deleteStemcell{
		config: config,
		stemcellPath: path,
	}, nil
}

func (c *deleteStemcell) Run() bosh.Response {
	stemcellPath := filepath.Join(c.config.StemcellDir, c.stemcellPath)
	err := os.RemoveAll(stemcellPath)

	if err != nil {
		return bosh.CPIError("failed to delete stemcell " + c.stemcellPath, err)
	} else {
		return bosh.Response{}
	}
}
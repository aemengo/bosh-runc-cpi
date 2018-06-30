package command

import (
	"github.com/aemengo/bosh-containerd-cpi/bosh"
	cfg "github.com/aemengo/bosh-containerd-cpi/config"
	"github.com/satori/go.uuid"
	"os"
	"errors"
	"path/filepath"
	"io"
)

type createStemcell struct {
	config cfg.Config
	stemcellSrcPath string
}

func NewCreateStemcell(arguments []interface{}, config cfg.Config) (*createStemcell, error) {
	if len(arguments) == 0 {
		return nil, errors.New("invalid stemcell path passed to create_stemcell command")
	}

	path, ok := arguments[0].(string)
	if !ok {
		return nil, errors.New("invalid stemcell path passed to create_stemcell command")
	}

	return &createStemcell{
		config: config,
		stemcellSrcPath: path,
	}, nil
}

func (c *createStemcell) Run() bosh.Response {
	r, err := os.Open(c.stemcellSrcPath)
	if err != nil {
		return bosh.CPIError("failed to read stemcell to path " + c.stemcellSrcPath, err)
	}

	err = os.MkdirAll(c.config.StemcellDir, os.ModePerm)
	if err != nil {
		return bosh.CPIError("stemcell directory could not be created", err)
	}

	id := uuid.NewV4().String()
	destPath := filepath.Join(c.config.StemcellDir, id)

	f, err := os.Create(destPath)
	if err != nil {
		return bosh.CPIError("failed to create stemcell at path " + destPath, err)
	}

	_, err = io.Copy(f, r)
	if err != nil {
		return bosh.CPIError("failed to create stemcell", err)
	}

	return bosh.Response{Result: id}
}
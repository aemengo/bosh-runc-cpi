package command

import (
	"github.com/aemengo/bosh-containerd-cpi/bosh"
	cfg "github.com/aemengo/bosh-containerd-cpi/config"
	"errors"
)

type Command interface {
	Run() bosh.Response
}

func New(method string, arguments []interface{}, config cfg.Config) (Command, error) {
	switch method {
	case "create_stemcell":
		return NewCreateStemcell(arguments, config)
	case "info":
		return NewInfo()
	default:
		return nil, errors.New("")
	}
}

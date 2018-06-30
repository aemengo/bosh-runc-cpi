package command

import "github.com/aemengo/bosh-containerd-cpi/bosh"

type setVMMetadata struct {}

func NewSetVMMetadata() (*setVMMetadata, error) {
	return &setVMMetadata{}, nil
}

func (c *setVMMetadata) Run() bosh.Response {
	return bosh.Response{}
}

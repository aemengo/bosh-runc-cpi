package command

import "github.com/aemengo/bosh-containerd-cpi/bosh"

type setVMMetadata struct {}

func NewSetVMMetadata() *setVMMetadata {
	return &setVMMetadata{}
}

func (c *setVMMetadata) Run() bosh.Response {
	return bosh.Response{}
}
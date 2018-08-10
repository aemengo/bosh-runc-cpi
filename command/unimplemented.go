package command

import (
	"github.com/aemengo/bosh-runc-cpi/bosh"
)

type unimplemented struct {
	method string
}

func NewUnimplemented(method string) *unimplemented {
	return &unimplemented{
		method: method,
	}
}

func (c *unimplemented) Run() bosh.Response  {
	return bosh.UnimplementedError(c.method)
}
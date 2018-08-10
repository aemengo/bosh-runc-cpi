package command

import "github.com/aemengo/bosh-runc-cpi/bosh"

type noOP struct {}

func NewNoOP() *noOP {
	return &noOP{}
}

func (c *noOP) Run() bosh.Response {
	return bosh.Response{}
}
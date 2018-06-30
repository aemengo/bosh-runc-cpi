package command

import (
	"github.com/aemengo/bosh-containerd-cpi/bosh"
	"fmt"
)

type unimplemented struct {
	method string
}

func NewUnimplemented(method string) (*unimplemented, error) {
	return &unimplemented{
		method: method,
	}, nil
}

func (c *unimplemented) Run() bosh.Response  {
	return bosh.Response{
		Error: &bosh.Error{
			Type: "Bosh::Clouds::NotImplemented",
			Message: fmt.Sprintf("'%s' is not yet supported. Please call implemented method", c.method),
		},
	}
}

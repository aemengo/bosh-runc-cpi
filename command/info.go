package command

import "github.com/aemengo/bosh-containerd-cpi/bosh"

type info struct {}

func NewInfo() (*info, error) {
	return &info{}, nil
}

func (c *info) Run() bosh.Response {
	return bosh.Response{
		Result: map[string][]string{
			"stemcell_formats": []string{"warden-tar", "general-tar"},
		},
	}
}
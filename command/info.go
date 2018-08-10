package command

import "github.com/aemengo/bosh-runc-cpi/bosh"

type info struct {}

func NewInfo() *info {
	return &info{}
}

func (c *info) Run() bosh.Response {
	return bosh.Response{
		Result: map[string][]string{
			"stemcell_formats": []string{"warden-tar", "general-tar"},
		},
	}
}
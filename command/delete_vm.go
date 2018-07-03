package command

import (
	cfg "github.com/aemengo/bosh-containerd-cpi/config"
	"github.com/aemengo/bosh-containerd-cpi/bosh"
	"github.com/aemengo/bosh-containerd-cpi/pb"
	"context"
	"errors"
)

type deleteVM struct {
	pb.CPIDClient

	ctx context.Context
	arguments []interface{}
	config cfg.Config
	logPrefix string
}

func NewDeleteVM(ctx context.Context, cpidClient pb.CPIDClient, arguments []interface{}, config cfg.Config) *deleteVM {
	return &deleteVM{
		CPIDClient: cpidClient,
		ctx: ctx,
		arguments: arguments,
		config: config,
		logPrefix: "delete_vm",
	}
}

func (c *deleteVM) Run() bosh.Response {
	if len(c.arguments) == 0 {
		return bosh.CPIError(c.logPrefix, errors.New("invalid vm id submitted"))
	}

	vmID, ok := c.arguments[0].(string)
	if !ok {
		return bosh.CPIError(c.logPrefix, errors.New("invalid vm id submitted"))
	}

	c.DeleteVM(c.ctx, &pb.IDParcel{Value: vmID})
	return bosh.Response{}
}
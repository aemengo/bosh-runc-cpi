package command

import (
	cfg "github.com/aemengo/bosh-containerd-cpi/config"
	"errors"
	"github.com/aemengo/bosh-containerd-cpi/bosh"
	"github.com/aemengo/bosh-containerd-cpi/pb"
	"context"
)

type deleteDisk struct {
	pb.CPIDClient

	ctx context.Context
	config cfg.Config
	arguments []interface{}
	logPrefix string
}

func NewDeleteDisk(ctx context.Context, cpidClient pb.CPIDClient, arguments []interface{}, config cfg.Config) *deleteDisk {
	return &deleteDisk{
		CPIDClient: cpidClient,
		ctx: ctx,
		arguments: arguments,
		config: config,
		logPrefix: "delete_disk",
	}
}

func (c *deleteDisk) Run() bosh.Response {
	if len(c.arguments) == 0 {
		return bosh.CPIError(c.logPrefix, errors.New("invalid disk id submitted"))
	}

	path, ok := c.arguments[0].(string)
	if !ok {
		return bosh.CPIError(c.logPrefix, errors.New("invalid disk id submitted"))
	}

	_, err := c.DeleteDisk(c.ctx, &pb.IDParcel{Value: path})
	if err != nil {
		return bosh.CloudError(c.logPrefix, err)
	}

	return bosh.Response{}
}
package command

import (
	cfg "github.com/aemengo/bosh-containerd-cpi/config"
	"errors"
	"github.com/aemengo/bosh-containerd-cpi/bosh"
	"github.com/aemengo/bosh-containerd-cpi/pb"
	"context"
)

type attachDisk struct {
	pb.CPIDClient

	ctx context.Context
	config cfg.Config
	arguments []interface{}
	logPrefix string
}

func NewAttachDisk(ctx context.Context, cpidClient pb.CPIDClient, arguments []interface{}, config cfg.Config) *attachDisk {
	return &attachDisk{
		CPIDClient: cpidClient,
		ctx: ctx,
		arguments: arguments,
		config: config,
		logPrefix: "attach_disk",
	}
}

func (c *attachDisk) Run() bosh.Response {
	if len(c.arguments) != 2 {
		return bosh.CPIError(c.logPrefix, errors.New("invalid arguments submitted"))
	}

	vmID, ok := c.arguments[0].(string)
	if !ok {
		return bosh.CPIError(c.logPrefix, errors.New("invalid vm id submitted"))
	}

	diskID, ok := c.arguments[1].(string)
	if !ok {
		return bosh.CPIError(c.logPrefix, errors.New("invalid disk id submitted"))
	}

	_, err := c.AttachDisk(c.ctx, &pb.AttachDisksOpts{VmID: vmID, DiskID: diskID})
	if err != nil {
		return bosh.CloudError(c.logPrefix, err)
	}

	return bosh.Response{}
}
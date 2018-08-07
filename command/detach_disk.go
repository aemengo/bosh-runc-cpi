package command

import (
	"context"
	"errors"
	"github.com/aemengo/bosh-containerd-cpi/bosh"
	cfg "github.com/aemengo/bosh-containerd-cpi/config"
	"github.com/aemengo/bosh-containerd-cpi/pb"
)

type detachDisk struct {
	pb.CPIDClient

	ctx       context.Context
	config    cfg.Config
	arguments []interface{}
	logPrefix string
}

func NewDetachDisk(ctx context.Context, cpidClient pb.CPIDClient, arguments []interface{}, config cfg.Config) *detachDisk {
	return &detachDisk{
		CPIDClient: cpidClient,
		ctx:        ctx,
		arguments:  arguments,
		config:     config,
		logPrefix:  "detach_disk",
	}
}

func (c *detachDisk) Run() bosh.Response {
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

	_, err := c.DetachDisk(c.ctx, &pb.DisksOpts{VmID: vmID, DiskID: diskID})
	if err != nil {
		return bosh.CloudError(c.logPrefix, err)
	}

	return bosh.Response{}
}

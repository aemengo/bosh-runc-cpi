package command

import (
	cfg "github.com/aemengo/bosh-runc-cpi/config"
	"github.com/aemengo/bosh-runc-cpi/bosh"
	"github.com/aemengo/bosh-runc-cpi/pb"
	"context"
	"errors"
)

type createDisk struct {
    pb.CPIDClient

	ctx context.Context
    arguments []interface{}
    config cfg.Config
	logPrefix string
}

func NewCreateDisk(ctx context.Context, cpidClient pb.CPIDClient, arguments []interface{}, config cfg.Config) *createDisk {
	return &createDisk{
		CPIDClient: cpidClient,
		ctx: ctx,
		arguments: arguments,
		config: config,
		logPrefix: "create_disk",
	}
}

func (c *createDisk) Run() bosh.Response {
	if len(c.arguments) == 0 {
		return bosh.CPIError(c.logPrefix, errors.New("invalid disk size submitted"))
	}

	size, ok := c.parseInt32(c.arguments[0])
	if !ok {
		return bosh.CPIError(c.logPrefix, errors.New("invalid disk size submitted"))
	}

	id, err := c.CreateDisk(c.ctx, &pb.ValueParcel{Value: size})
	if err != nil {
		return bosh.CloudError(c.logPrefix, err)
	}

	return bosh.Response{Result: id.Value}
}

func (*createDisk) parseInt32(arg interface{}) (int32, bool) {
	vali32, ok := arg.(int32)
	if ok {
		return vali32, true
	}

	valf64, ok := arg.(float64)
	if ok {
		return int32(valf64), true
	}

	return 0, false
}
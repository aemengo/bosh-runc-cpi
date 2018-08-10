package command

import (
	"errors"
	"github.com/aemengo/bosh-runc-cpi/bosh"
	cfg "github.com/aemengo/bosh-runc-cpi/config"
	"github.com/aemengo/bosh-runc-cpi/pb"
	"context"
)

type hasDisk struct {
	pb.CPIDClient

	ctx        context.Context
	arguments  []interface{}
	config     cfg.Config
	logPrefix string
}

func NewHasDisk(ctx context.Context, cpidClient pb.CPIDClient, arguments []interface{}, config cfg.Config) *hasDisk {
	return &hasDisk{
		CPIDClient: cpidClient,
		ctx: ctx,
		arguments: arguments,
		config:   config,
		logPrefix: "has_disk",
	}
}

func (c *hasDisk) Run() bosh.Response {
	if len(c.arguments) == 0 {
		return bosh.CPIError(c.logPrefix, errors.New("invalid disk id submitted"))
	}

	path, ok := c.arguments[0].(string)
	if !ok {
		return bosh.CPIError(c.logPrefix, errors.New("invalid disk id submitted"))
	}

	exists, err := c.CPIDClient.HasDisk(c.ctx, &pb.IDParcel{Value: path})
	if err != nil {
		return bosh.CloudError(c.logPrefix, err)
	}

	return bosh.Response{Result: exists.Value}
}
package command

import (
	"errors"
	"github.com/aemengo/bosh-runc-cpi/bosh"
	cfg "github.com/aemengo/bosh-runc-cpi/config"
	"github.com/aemengo/bosh-runc-cpi/pb"
	"context"
)

type hasVM struct {
	pb.CPIDClient

	ctx        context.Context
	arguments  []interface{}
	config     cfg.Config
	logPrefix string
}

func NewHasVM(ctx context.Context, cpidClient pb.CPIDClient, arguments []interface{}, config cfg.Config) *hasVM {
	return &hasVM{
		CPIDClient: cpidClient,
		ctx: ctx,
		arguments: arguments,
		config:   config,
		logPrefix: "has_vm",
	}
}

func (c *hasVM) Run() bosh.Response {
	if len(c.arguments) == 0 {
		return bosh.CPIError(c.logPrefix, errors.New("invalid vm id submitted"))
	}

	id, ok := c.arguments[0].(string)
	if !ok {
		return bosh.CPIError(c.logPrefix, errors.New("invalid vm id submitted"))
	}

	success, err := c.CPIDClient.HasVM(c.ctx, &pb.IDParcel{Value: id})
	if err != nil {
		return bosh.CloudError(c.logPrefix, err)
	}

	return bosh.Response{Result: success.Value}
}
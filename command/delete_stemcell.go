package command

import (
	"github.com/aemengo/bosh-runc-cpi/bosh"
	cfg "github.com/aemengo/bosh-runc-cpi/config"
	"errors"
	"context"
	"github.com/aemengo/bosh-runc-cpi/pb"
)

type deleteStemcell struct {
	pb.CPIDClient

	ctx context.Context
	config cfg.Config
	arguments []interface{}
	logPrefix string

}

func NewDeleteStemcell(ctx context.Context, cpidClient pb.CPIDClient, arguments []interface{}, config cfg.Config) *deleteStemcell {
	return &deleteStemcell{
		CPIDClient: cpidClient,
		ctx: ctx,
		arguments: arguments,
		config: config,
		logPrefix: "delete_stemcell",
	}
}

func (c *deleteStemcell) Run() bosh.Response {
	if len(c.arguments) == 0 {
		return bosh.CPIError(c.logPrefix, errors.New("invalid stemcell id submitted"))
	}

	path, ok := c.arguments[0].(string)
	if !ok {
		return bosh.CPIError(c.logPrefix, errors.New("invalid stemcell id submitted"))
	}

	_, err := c.DeleteStemcell(c.ctx, &pb.IDParcel{Value: path})
	if err != nil {
		return bosh.CloudError(c.logPrefix, err)
	}

	return bosh.Response{}
}
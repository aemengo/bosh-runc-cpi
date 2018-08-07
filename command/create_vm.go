package command

import (
	"context"
	"errors"
	"github.com/aemengo/bosh-containerd-cpi/bosh"
	cfg "github.com/aemengo/bosh-containerd-cpi/config"
	"github.com/aemengo/bosh-containerd-cpi/pb"
)

type createVM struct {
	pb.CPIDClient

	ctx       context.Context
	arguments []interface{}
	config    cfg.Config
	logPrefix string
}

func NewCreateVM(ctx context.Context, cpidClient pb.CPIDClient, arguments []interface{}, config cfg.Config) *createVM {
	return &createVM{
		CPIDClient: cpidClient,
		ctx:        ctx,
		arguments:  arguments,
		config:     config,
		logPrefix:  "create_vm",
	}
}

func (c *createVM) Run() bosh.Response {
	if len(c.arguments) < 2 {
		return bosh.CPIError(c.logPrefix, errors.New("invalid stemcell id submitted"))
	}

	stemcellID, ok := c.arguments[1].(string)
	if !ok {
		return bosh.CPIError(c.logPrefix, errors.New("invalid stemcell id submitted"))
	}

	var diskID string
	if values, ok := c.arguments[4].([]interface{}); ok {
		if len(values) > 0 {
			if str, ok := values[0].(string); ok {
				diskID = str
			}
		}
	}

	agentSettings := bosh.ConvertAgentSettings(c.arguments, c.config)

	id, err := c.CreateVM(c.ctx, &pb.CreateVMOpts{
		StemcellID: stemcellID,
		AgentSettings: agentSettings,
		DiskID: diskID,
	})
	if err != nil {
		return bosh.CloudError(c.logPrefix, err)
	}

	return bosh.Response{Result: id.Value}
}

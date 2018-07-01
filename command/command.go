package command

import (
	"github.com/aemengo/bosh-containerd-cpi/bosh"
	cfg "github.com/aemengo/bosh-containerd-cpi/config"
	"context"
	"github.com/aemengo/bosh-containerd-cpi/pb"
)

type Command interface {
	Run() bosh.Response
}

func New(ctx context.Context, cpidClient pb.CPIDClient, method string, arguments []interface{}, config cfg.Config) Command {
	switch method {
	case "create_stemcell":
		return NewCreateStemcell(ctx, cpidClient, arguments, config)
	case "info":
		return NewInfo()
	case "delete_stemcell":
		return NewDeleteStemcell(ctx, cpidClient, arguments, config)
	case "create_disk":
		return NewCreateDisk(ctx, cpidClient, arguments, config)
	case "delete_disk":
		return NewDeleteDisk(ctx, cpidClient, arguments, config)
	case "has_disk":
		return NewHasDisk(ctx, cpidClient, arguments, config)
	case "set_vm_metadata":
		return NewSetVMMetadata()
	default:
		return NewUnimplemented(method)
	}
}

package command

import (
	"github.com/aemengo/bosh-containerd-cpi/bosh"
	cfg "github.com/aemengo/bosh-containerd-cpi/config"
	"os"
	"errors"
	"io"
	"github.com/aemengo/bosh-containerd-cpi/pb"
	"context"
	"fmt"
)

type createStemcell struct {
	pb.CPIDClient

	ctx context.Context
	arguments []interface{}
	config cfg.Config
	logPrefix string
}

func NewCreateStemcell(ctx context.Context, cpidClient pb.CPIDClient, arguments []interface{}, config cfg.Config) *createStemcell {
	return &createStemcell{
		CPIDClient: cpidClient,
		ctx: ctx,
		arguments: arguments,
		config: config,
		logPrefix: "create_stemcell",
	}
}

func (c *createStemcell) Run() bosh.Response {
	if len(c.arguments) == 0 {
		return bosh.CPIError(c.logPrefix, errors.New("invalid stemcell id submitted"))
	}

	path, ok := c.arguments[0].(string)
	if !ok {
		return bosh.CPIError(c.logPrefix, errors.New("invalid stemcell id submitted"))
	}

	f, err := os.Open(path)
	if err != nil {
		return bosh.CPIError(c.logPrefix, fmt.Errorf("failed to read stemcell to path: %s: %s" + path, err))
	}
	defer f.Close()

	stream, err := c.CreateStemcell(c.ctx)
	if err != nil {
		return bosh.CloudError(c.logPrefix, err)
	}

	for {
		chunk := make([]byte, 64 * 1024)
		n, err := f.Read(chunk)
		if err == io.EOF {
			break
		} else if err != nil {
			return bosh.CloudError(c.logPrefix, err)
		}

		if n < len(chunk) {
			chunk = chunk[:n]
		}

		err = stream.Send(&pb.DataParcel{Value: chunk})
		if err != nil {
			return bosh.CloudError(c.logPrefix, err)
		}
	}

	id, err := stream.CloseAndRecv()
	if err != nil {
		return bosh.CloudError(c.logPrefix, err)
	}

	return bosh.Response{Result: id.Value}
}
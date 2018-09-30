package client

import (
	"context"
	"fmt"
	"github.com/aemengo/bosh-runc-cpi/pb"
	"google.golang.org/grpc"
)

func Ping(ctx context.Context, target string) error {
	client, err := newClient(target)
	if err != nil {
		return err
	}

	result, err := client.Ping(ctx, &pb.Void{})
	if err != nil {
		return err
	}

	if result.Value != "pong" {
		return fmt.Errorf("server returned unexpected value: %s", result.Value)
	}

	return nil
}

func newClient(target string) (pb.CPIDClient, error) {
	conn, err := grpc.Dial(target, grpc.WithInsecure())
	if err != nil {
		return nil, err
	}

	return pb.NewCPIDClient(conn), nil
}


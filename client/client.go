package client

import (
	"context"
	"fmt"
	"github.com/aemengo/bosh-runc-cpi/pb"
	"google.golang.org/grpc"
	"io"
	"os"
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

func Prune(ctx context.Context, target string) error {
	client, err := newClient(target)
	if err != nil {
		return err
	}

	_, err = client.Prune(ctx, &pb.Void{})
	return err
}

func StreamOut(ctx context.Context, target, src, dest string) error {
	client, err := newClient(target)
	if err != nil {
		return err
	}

	f, err := os.Create(dest)
	if err != nil {
		return err
	}

	stream, err := client.StreamOut(ctx, &pb.TextParcel{Value: src})
	if err != nil {
		return err
	}

	for {
		data, err := stream.Recv()
		if err == io.EOF {
			break
		} else if err != nil {
			return err
		}

		f.Write(data.Value)
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


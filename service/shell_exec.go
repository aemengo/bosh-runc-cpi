package service

import (
	"context"
	"github.com/aemengo/bosh-runc-cpi/pb"
	"os/exec"
	"fmt"
	"errors"
)

func (s *Service) ShellExec(ctx context.Context, req *pb.TextParcel) (*pb.TextParcel, error) {
	if !s.config.Debug {
		return nil, errors.New("unsupported action")
	}

	output, err := exec.Command("sh", "-c", req.Value).CombinedOutput()
	if err != nil {
		return nil, fmt.Errorf("failed to execute: %q: %s: %s", req.Value, err, output)
	}

	return &pb.TextParcel{Value: string(output)}, nil
}

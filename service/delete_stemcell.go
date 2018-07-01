package service

import (
	"context"
	"github.com/aemengo/bosh-containerd-cpi/pb"
	"os"
	"path/filepath"
)

func (s *Service) DeleteStemcell(ctx context.Context, req *pb.IDParcel) (*pb.Void, error) {
	stemcellPath := filepath.Join(s.config.StemcellDir, req.Value)

	err := os.RemoveAll(stemcellPath)

	if err != nil {
		return nil, err
	} else {
		return &pb.Void{}, nil
	}
}

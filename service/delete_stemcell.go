package service

import (
	"context"
	"github.com/aemengo/bosh-runc-cpi/pb"
	"os"
	"path/filepath"
)

func (s *Service) DeleteStemcell(ctx context.Context, req *pb.TextParcel) (*pb.Void, error) {
	stemcellPath := filepath.Join(s.config.StemcellDir, req.Value)
	os.RemoveAll(stemcellPath)
	return &pb.Void{}, nil
}

package service

import (
	"context"
	"github.com/aemengo/bosh-containerd-cpi/pb"
	"os"
	"path/filepath"
)

func (s *Service) DeleteStemcell(ctx context.Context, req *pb.IDParcel) (*pb.Void, error) {
	var (
		stemcellPath    = filepath.Join(s.config.StemcellDir, req.Value)
		stemcellTarPath = stemcellPath + ".tgz"
	)

	os.RemoveAll(stemcellPath)
	os.RemoveAll(stemcellTarPath)
	return &pb.Void{}, nil
}

package service

import (
	"github.com/aemengo/bosh-containerd-cpi/pb"
	"os"
	"github.com/satori/go.uuid"
	"path/filepath"
	"fmt"
	"io"
)

func (s *Service) CreateStemcell(stream pb.CPID_CreateStemcellServer) error  {
	id := uuid.NewV4().String()
	destPath := filepath.Join(s.config.StemcellDir, id)

	f, err := os.Create(destPath)
	if err != nil {
		return fmt.Errorf("failed to create stemcell at path: %s: %s " + destPath, err)
	}
	defer f.Close()

	for {
		data, err := stream.Recv()
		if err == io.EOF {
			return stream.SendAndClose(&pb.IDParcel{Value: id})
		} else if err != nil {
			return err
		}

		f.Write(data.Value)
	}
}
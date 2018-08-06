package service

import (
	"fmt"
	"github.com/aemengo/bosh-containerd-cpi/pb"
	"github.com/satori/go.uuid"
	"io"
	"os"
	"path/filepath"
	"github.com/aemengo/bosh-containerd-cpi/utils"
)

func (s *Service) CreateStemcell(stream pb.CPID_CreateStemcellServer) error {
	var (
		id       = uuid.NewV4().String()
		destPath = filepath.Join(s.config.StemcellDir, id)
		tarPath  = destPath + ".tgz"
	)

	f, err := os.Create(tarPath)
	if err != nil {
		return fmt.Errorf("failed to create stemcell at path: %s: %s "+destPath, err)
	}
	defer f.Close()
	defer os.RemoveAll(tarPath)

	for {
		data, err := stream.Recv()
		if err == io.EOF {
			break
		} else if err != nil {
			return err
		}

		f.Write(data.Value)
	}

	err = os.MkdirAll(destPath, os.ModePerm)
	if err != nil {
		return fmt.Errorf("failed to make stemcell path: %s", err)
	}

	err = utils.RunCommand("tar", "-xf", tarPath, "-C", destPath)
	if err != nil {
		return fmt.Errorf("failed to unpack stemcell: %s", err)
	}

	return stream.SendAndClose(&pb.IDParcel{Value: id})
}

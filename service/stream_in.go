package service

import (
	"errors"
	"fmt"
	"github.com/aemengo/bosh-runc-cpi/pb"
	"github.com/aemengo/bosh-runc-cpi/utils"
	"github.com/satori/go.uuid"
	"google.golang.org/grpc/metadata"
	"io"
	"os"
	"path/filepath"
)

func (s *Service) StreamIn(stream pb.CPID_StreamInServer) error {
	md, ok := metadata.FromIncomingContext(stream.Context())
	if !ok {
		return errors.New("StreamIn request must specify 'destination_dir' in the metadata")
	}

	dest := md["destination_dir"]

	if dest == nil || len(dest) == 0 {
		return errors.New("StreamIn request must specify 'destination_dir' in the metadata")
	}

	var (
		id       = uuid.NewV4().String()
		tmpPath  = filepath.Join(s.config.TempDir, id)
		tarPath  = tmpPath + ".tgz"
		destPath = dest[0]
	)

	f, err := os.Create(tarPath)
	if err != nil {
		return fmt.Errorf("failed to create archive at path: %s: %s ", tarPath, err)
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

	err = utils.RunCommand("tar", "-xf", tarPath, "-C", destPath)
	if err != nil {
		return fmt.Errorf("failed to unpack stemcell: %s", err)
	}

	return stream.SendAndClose(&pb.Void{})
}

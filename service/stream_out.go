package service

import (
	"fmt"
	"github.com/aemengo/bosh-runc-cpi/pb"
	"io"
	"os"
	"os/exec"
)

func (s *Service) StreamOut(path *pb.TextParcel, stream pb.CPID_StreamOutServer) error {
	tarPath := path.Value+".tgz"
	err := exec.Command("tar", "-czf", tarPath, path.Value).Run()
	if err != nil {
		return fmt.Errorf("cannot create archive at %q from %q: %s", tarPath, path.Value, err)
	}
	defer os.RemoveAll(tarPath)

	f, err := os.Open(tarPath)
	if err != nil {
		return err
	}
	defer f.Close()

	for {
		chunk := make([]byte, 64 * 1024)
		n, err := f.Read(chunk)
		if err == io.EOF {
			break
		} else if err != nil {
			return err
		}

		if n < len(chunk) {
			chunk = chunk[:n]
		}

		err = stream.Send(&pb.DataParcel{Value: chunk})
		if err != nil {
			return err
		}
	}

	return nil
}

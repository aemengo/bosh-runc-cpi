package service

import (
	"fmt"
	"github.com/aemengo/bosh-runc-cpi/pb"
	"github.com/satori/go.uuid"
	"io"
	"os"
	"os/exec"
	"path/filepath"
)

func (s *Service) StreamOut(path *pb.TextParcel, stream pb.CPID_StreamOutServer) error {
	var (
		id       = uuid.NewV4().String()
		destPath = filepath.Join(s.config.TempDir, id)
		tarPath  = destPath + ".tgz"
	)

	err := exec.Command("tar", buildCompressArguments(tarPath, path.Value)...).Run()
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
		chunk := make([]byte, 64*1024)
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

func buildCompressArguments(dest, src string) []string {
	doesPigzExist := func() bool {
		_, err := exec.LookPath("pigz")
		return err == nil
	}

	args := []string{
		"-cf", dest,
	}

	if doesPigzExist() {
		args = append(args, "--use-compress-program=pigz")
	} else {
		args = append(args, "--use-compress-program=gzip")
	}

	args = append(args, src)
	return args
}

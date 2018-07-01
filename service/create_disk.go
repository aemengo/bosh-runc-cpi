package service

import (
	"context"
	"github.com/aemengo/bosh-containerd-cpi/pb"
	"github.com/satori/go.uuid"
	"path/filepath"
	"io/ioutil"
	"strconv"
	"github.com/aemengo/bosh-containerd-cpi/utils"
	"fmt"
)

func (s *Service) CreateDisk(ctx context.Context, req *pb.ValueParcel) (*pb.IDParcel, error) {
	id := uuid.NewV4().String()
	diskPath := filepath.Join(s.config.DiskDir, id)

	err := ioutil.WriteFile(diskPath, []byte{}, 0600)
	if err != nil {
		return nil, fmt.Errorf("disk directory could not be created: %s", err)
	}

	sizeStr := strconv.FormatInt(int64(req.Value), 10) + "MB"

	err = utils.RunCommand("truncate", "-s", sizeStr, diskPath)
	if err != nil {
		return nil, fmt.Errorf("failed to resize disk: %s", err)
	}

	err = utils.RunCommand("/sbin/mkfs", "-t", "ext4", "-F", diskPath)
	if err != nil {
		return nil, fmt.Errorf("failed to build disk filesystem: %s", err)
	}

	return &pb.IDParcel{Value: id}, nil
}

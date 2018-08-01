package service

import (
	"context"
	"github.com/aemengo/bosh-containerd-cpi/pb"
	"github.com/aemengo/bosh-containerd-cpi/utils"
	"github.com/vishvananda/netlink"
	"io/ioutil"
	"os"
	"path/filepath"
)

func (s *Service) DeleteVM(ctx context.Context, req *pb.IDParcel) (*pb.Void, error) {
	var (
		vmPath      = filepath.Join(s.config.VMDir, req.Value)
		rootFsPath  = filepath.Join(vmPath, "rootfs")
		ethNamePath = filepath.Join(vmPath, "eth-name")
	)

	deleteVirtualInterfaceIfExists(ethNamePath)
	s.runc.DeleteContainer(req.Value)
	utils.RunCommand("umount", rootFsPath)
	os.RemoveAll(vmPath)
	return &pb.Void{}, nil
}

func deleteVirtualInterfaceIfExists(path string) {
	contents, err := ioutil.ReadFile(path)
	if err != nil || string(contents) == "" {
		return
	}

	link, err := netlink.LinkByName(string(contents))
	if err != nil {
		return
	}

	netlink.LinkDel(link)
}

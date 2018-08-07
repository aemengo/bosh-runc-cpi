package service

import (
	"context"
	"github.com/aemengo/bosh-containerd-cpi/pb"
	"github.com/vishvananda/netlink"
	"io/ioutil"
	"os"
	"path/filepath"
)

func (s *Service) DeleteVM(ctx context.Context, req *pb.IDParcel) (*pb.Void, error) {
	var (
		vmPath      = filepath.Join(s.config.VMDir, req.Value)
		ethNamePath = filepath.Join(vmPath, "eth-name")
	)

	s.runc.DeleteContainer(req.Value)
	deleteVirtualInterfaceIfExists(ethNamePath)
	os.RemoveAll(vmPath)
	return &pb.Void{}, nil
}

// deleting a non-existent virtual interface will panic
// so we have to check first
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

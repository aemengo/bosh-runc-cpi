package service

import (
	"io/ioutil"
	"strconv"
	"path/filepath"
	"github.com/vishvananda/netlink"
)

func (s *Service) configureNetworking(vmPath string, pidPath string, ip string, mask string, gatewayIP string) error {
	contents, err := ioutil.ReadFile(pidPath)
	if err != nil {
		return err
	}

	pid, err := strconv.Atoi(string(contents))
	if err != nil {
		return err
	}

	vEthPair := s.network.CreateVirtualEthernetPair(pid)

	err = ioutil.WriteFile(
		ethNamePath(vmPath),
		[]byte(vEthPair.Name),
		0666,
	)
	if err != nil {
		return err
	}

	err = s.network.InstallVirtualEthernetPair(vEthPair, pid)
	if err != nil {
		return err
	}

	addr, err := netlink.ParseAddr(ip + "/16")
	if err != nil {
		return err
	}

	return s.network.ConfigurePeerInterface(
		vEthPair.PeerName,
		pid,
		addr,
		gatewayIP,
	)
}

// deleting a non-existent virtual interface will panic
// so we have to check first
func (s *Service) tearDownNetworking(vmPath string) error {
	path := ethNamePath(vmPath)

	contents, err := ioutil.ReadFile(path)
	if err != nil || string(contents) == "" {
		return nil
	}

	link, err := netlink.LinkByName(string(contents))
	if err != nil {
		return nil
	}

	return netlink.LinkDel(link)
}

func ethNamePath(vmPath string) string  {
	return filepath.Join(vmPath, "eth-name")
}
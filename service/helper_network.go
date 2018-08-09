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
		filepath.Join(vmPath, "eth-name"),
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

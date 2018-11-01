package service

import (
	"github.com/vishvananda/netlink"
	"io/ioutil"
	"regexp"
	"strconv"
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

	err = s.network.InstallVirtualEthernetPair(vEthPair, pid)
	if err != nil {
		tearDownNetworking(vEthPair.Name)
		return err
	}

	regex := regexp.MustCompile(`^\d+\.\d+\.\d+\.\d+(\/\d+)$`)
	matches := regex.FindStringSubmatch(s.config.CIDR)

	addr, err := netlink.ParseAddr(ip + matches[1])
	if err != nil {
		tearDownNetworking(vEthPair.Name)
		return err
	}

	err = s.network.ConfigurePeerInterface(
		vEthPair.PeerName,
		pid,
		addr,
		gatewayIP,
	)

	if err != nil {
		tearDownNetworking(vEthPair.Name)
		return err
	}

	return nil
}

// deleting a non-existent virtual interface will panic
// so we have to check first
func tearDownNetworking(ethName string) error {
	link, err := netlink.LinkByName(ethName)
	if err != nil {
		return nil
	}

	return netlink.LinkDel(link)
}

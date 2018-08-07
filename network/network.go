package network

import (
	"fmt"
	"github.com/vishvananda/netlink"
	"github.com/vishvananda/netns"
	"github.com/genuinetools/netns/bridge"
	"net"
	"runtime"
)

type Network struct {
	bridge *net.Interface
}

func New() (*Network, error) {
	br, err := bridge.Init(bridge.Opt{Name: "netns0", IPAddr: "10.0.0.1/16"})
	if err != nil {
		return nil, err
	}

	return &Network{
		bridge: br,
	}, nil
}

func (n *Network) CreateVirtualEthernetPair(pid int) *netlink.Veth {
	la := netlink.NewLinkAttrs()
	la.Name = fmt.Sprintf("netnsv0-%d", pid)
	la.MasterIndex = n.bridge.Index

	return &netlink.Veth{
		LinkAttrs: la,
		PeerName:  fmt.Sprintf("ethc-%d", pid),
	}
}

func (n *Network) InstallVirtualEthernetPair(vEthPair *netlink.Veth, namespacePid int) error {
	if err := netlink.LinkAdd(vEthPair); err != nil {
		return fmt.Errorf("create veth pair named [ %#v ] failed: %v", vEthPair, err)
	}

	peer, err := netlink.LinkByName(vEthPair.PeerName)
	if err != nil {
		return fmt.Errorf("getting peer interface %s failed: %v", vEthPair.PeerName, err)
	}

	if err := netlink.LinkSetNsPid(peer, namespacePid); err != nil {
		return fmt.Errorf("adding peer interface to network namespace of pid %d failed: %v", namespacePid, err)
	}

	if err := netlink.LinkSetUp(vEthPair); err != nil {
		return fmt.Errorf("bringing local veth pair [ %#v ] up failed: %v", vEthPair, err)
	}

	return nil
}

func (n *Network) ConfigurePeerInterface(name string, namespacePid int, addr *netlink.Addr, gatewayIP string) error {
	runtime.LockOSThread()
	defer runtime.UnlockOSThread()

	origins, err := netns.Get()
	if err != nil {
		return fmt.Errorf("getting current network namespace failed: %v", err)
	}
	defer origins.Close()

	newns, err := netns.GetFromPid(namespacePid)
	if err != nil {
		return fmt.Errorf("getting network namespace for pid %d failed: %v", namespacePid, err)
	}
	defer newns.Close()

	if err := netns.Set(newns); err != nil {
		return fmt.Errorf("entering network namespace failed: %v", err)
	}
	defer netns.Set(origins)

	iface, err := netlink.LinkByName(name)
	if err != nil {
		return fmt.Errorf("getting link %s failed: %v", name, err)
	}

	if err := netlink.LinkSetDown(iface); err != nil {
		return fmt.Errorf("bringing interface [ %#v ] down failed: %v", iface, err)
	}

	if err := netlink.LinkSetName(iface, "eth0"); err != nil {
		return fmt.Errorf("renaming interface %s to eth0 failed: %v", name, err)
	}

	if err := netlink.AddrAdd(iface, addr); err != nil {
		return fmt.Errorf("setting %s interface ip to %s failed: %v", name, addr.String(), err)
	}

	if err := netlink.LinkSetUp(iface); err != nil {
		return fmt.Errorf("bringing interface [ %#v ] up failed: %v", iface, err)
	}

	gw := net.ParseIP(gatewayIP)
	err = netlink.RouteAdd(&netlink.Route{
		Scope:     netlink.SCOPE_UNIVERSE,
		LinkIndex: iface.Attrs().Index,
		Gw:        gw,
	})

	if err != nil {
		return fmt.Errorf("adding route %s to interface %s failed: %v", gw.String(), name, err)
	}

	return nil
}
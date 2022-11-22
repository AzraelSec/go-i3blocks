package main

import (
	"errors"
	"net"
	"strings"
)

const LOCAL_IP_ADDR = "192.168."

type Network struct {
	Interface net.Interface
	Address   net.Addr
}

func FindNetwork(name string) (*Network, error) {
	return identifyNetwork(func(n *Network) bool {
		return n.Interface.Name == name
	})
}

func CurrentNetwork() (*Network, error) {
	return identifyNetwork(func(n *Network) bool {
		return internalIPAddr(n.Address)
	})
}

func identifyNetwork(p func(*Network) bool) (*Network, error) {
	intfs := AvailableInterfaces()
	if len(intfs) == 0 {
		return nil, errors.New("no viable interface found")
	}

	for i := 0; i < len(intfs); i++ {
		intf := intfs[i]
		addrs, err := intf.Addrs()
		if err != nil {
			continue
		}

		for _, addr := range addrs {
			network := &Network{
				Interface: intf,
				Address:   addr,
			}

			if p(network) {
				return network, nil
			}
		}
	}

	return nil, errors.New("no active local network found")
}

func internalIPAddr(addr net.Addr) bool {
	addrStr := addr.String()
	return strings.HasPrefix(addrStr, LOCAL_IP_ADDR)
}

func AvailableInterfaces() []net.Interface {
	res := make([]net.Interface, 0)

	intfs, err := net.Interfaces()
	if err != nil {
		return res
	}

	for _, intf := range intfs {
		if intf.Flags&net.FlagUp != 0 && intf.Flags&net.FlagLoopback == 0 {
			res = append(res, intf)
		}
	}

	return res
}

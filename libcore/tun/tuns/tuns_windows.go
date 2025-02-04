package tuns

import (
	"errors"
	"libcore/tun"
	"os"

	"gvisor.dev/gvisor/pkg/tcpip"
)

func NewGvisor(dev int32, mtu int32, handler tun.Handler, nicId tcpip.NICID, pcap bool, pcapFile *os.File, snapLen uint32, ipv6Mode int32) (tun.Tun, error) {
	return nil, errors.New("platform depended")
}

func NewSystem(dev int32, mtu int32, handler tun.Handler, ipv6Mode int32, errorHandler func(err string)) (tun.Tun, error) {
	return nil, errors.New("platform depended")
}

func NewTun2Socket(fd int32, handler tun.Handler) (tun.Tun, error) {
	return nil, errors.New("platform depended")
}

//go:build !darwin && !freebsd && !linux && !windows

package tfo

import (
	"context"
	"net"
)

func SetTFOListener(fd uintptr) error {
	return ErrPlatformUnsupported
}

func SetTFODialer(fd uintptr) error {
	return ErrPlatformUnsupported
}

func (d *Dialer) dialTFOContext(ctx context.Context, network, address string, b []byte) (*net.TCPConn, error) {
	return nil, ErrPlatformUnsupported
}

func dialTCPAddr(network string, laddr, raddr *net.TCPAddr, b []byte) (*net.TCPConn, error) {
	return nil, ErrPlatformUnsupported
}

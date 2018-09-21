package udp

import (
	"context"
	"net"
)

type Puncher struct {
}

func New() *Puncher {
	return &Puncher{}
}

func (p *Puncher) Connect(ctx context.Context, ip net.IP) (addr *net.UDPAddr, err error) {
	return
}

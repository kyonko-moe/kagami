package model

import (
	"net"
	"time"
)

type UDPServer struct {
	AddressBook  map[string]*UDPUser
	LocalAddress *net.UDPAddr
	Connection   *net.UDPConn
}

type UDPUser struct {
	Name         string
	UDPAddr      *net.UDPAddr
	RegisterTime *time.Time
}

package model

import (
	"net"
)

type ReqUDP struct {
	Type string
	Data interface{}
}

type RspUDP struct {
	Ok   bool
	Desc string
	Data interface{}
}

type ArgRegister struct {
	Name string
}

type RetRegister struct {
	UDPAddr *net.UDPAddr
}

type ArgListNeighbor struct {
	Name string
}

type RetListNeighbor struct {
	Neighbors []*UDPUser
}

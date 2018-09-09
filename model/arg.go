package model

import (
	"net"
)

// ArgRegister is.
type ArgRegister struct {
	SourcIP *net.IPAddr
	Name    string
}

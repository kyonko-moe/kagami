package model

import (
	"net"
)

type nodeType string

const (
	BranchNode nodeType = "branch"
	LeafNode   nodeType = "leaf"
)

type Node struct {
	Name string      `json:"name"`
	ID   int64       `json:"id"`
	Addr *net.IPAddr `json:"-"`
	Type nodeType    `json:"type"`
}

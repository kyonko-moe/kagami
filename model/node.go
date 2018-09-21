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
	Name string   `json:"name"`
	ID   uint64   `json:"id"`
	IP   net.IP   `json:"-"`
	Type nodeType `json:"type"`
}

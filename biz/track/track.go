package track

import (
	"context"
	"net"

	"github.com/kyonko-moe/kagami/biz/naming"
	"github.com/kyonko-moe/kagami/biz/track/register"
	"github.com/kyonko-moe/kagami/model"
)

type Tracker struct {
	reg   *register.Default
	namer *naming.Default
}

func New() *Tracker {
	return &Tracker{
		reg:   register.NewDefault(),
		namer: naming.NewDefault(0x1, make([]int64, 0), 0),
	}
}

func (t *Tracker) Register(c context.Context, ipAddr *net.IPAddr, name string) (n *model.Node, err error) {
	n = &model.Node{
		Name: name,
		Addr: ipAddr,
		Type: model.LeafNode,
	}
	if n.ID, err = t.namer.Node(n); err != nil {
		return
	}
	if err = t.reg.SetNode(n); err != nil {
		return
	}
	return
}

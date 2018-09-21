package track

import (
	"context"
	"net"

	"github.com/kyonko-moe/kagami/biz/naming"
	"github.com/kyonko-moe/kagami/biz/register"
	"github.com/kyonko-moe/kagami/model"

	"github.com/pkg/errors"
)

type Tracker struct {
	reg   *register.Default
	namer *naming.Default
}

func New() *Tracker {
	return &Tracker{
		reg:   register.NewDefault(),
		namer: naming.NewDefault(0x0, make([]uint64, 0), 0),
	}
}

func (t *Tracker) Register(c context.Context, ip net.IP, name string) (n *model.Node, err error) {
	n = &model.Node{
		Name: name,
		IP:   ip,
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

func (t *Tracker) Locate(c context.Context, name string) (n *model.Node, err error) {
	n = t.reg.NodeByName(name)
	if n == nil {
		err = errors.Errorf("Unknown node : %s", name)
		return
	}
	return
}

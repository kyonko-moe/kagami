package pick

import (
	"github.com/kyonko-moe/kagami/model"
)

type Default struct{}

func (d *Default) PickOne(myself *model.Node, ns []*model.Node) (n *model.Node) {
	if len(ns) > 0 {
		return ns[0]
	}
	return
}

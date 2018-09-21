package register

import (
	"sync"

	"github.com/kyonko-moe/kagami/model"

	"github.com/pkg/errors"
)

type Default struct {
	m map[string]*model.Node

	sync.RWMutex
}

func NewDefault() *Default {
	return &Default{
		m: make(map[string]*model.Node),
	}
}

func (d *Default) NodeByName(name string) (n *model.Node) {
	d.RLock()
	defer d.RUnlock()

	n, _ = d.m[name]
	return
}

func (d *Default) NodeByID(id int64) (n *model.Node) {
	return
}

func (d *Default) SetNode(n *model.Node) (err error) {
	if n == nil {
		return
	}
	d.Lock()
	defer d.Unlock()

	if _, ok := d.m[n.Name]; ok {
		err = errors.Errorf("duplicated name of node : %s", n.Name)
		return
	}
	d.m[n.Name] = n
	return
}

func (d *Default) Clear() {
	d.Lock()
	defer d.Unlock()

	d.m = make(map[string]*model.Node)
}

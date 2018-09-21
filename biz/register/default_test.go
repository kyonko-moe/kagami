package register

import (
	"net"
	"testing"

	"github.com/kyonko-moe/kagami/model"

	. "github.com/smartystreets/goconvey/convey"
)

func TestDefault(t *testing.T) {
	Convey("Init default", t, func(ctx C) {
		reg := NewDefault()
		ctx.Convey("test set node", func(ctx C) {
			node := &model.Node{
				Name: "ut",
				IP:   net.ParseIP("127.0.0.1"),
				Type: model.LeafNode,
			}
			err := reg.SetNode(node)
			So(err, ShouldBeNil)
			ctx.Convey("get node", func(ctx C) {
				// node2 := reg.NodeByID(node.ID)
				// So(node2, ShouldResemble, node)

				node2 := reg.NodeByName(node.Name)
				So(node2, ShouldResemble, node)
			})
		})
	})
}

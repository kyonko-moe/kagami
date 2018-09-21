package naming

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestDefault(t *testing.T) {
	Convey("Init default", t, func(ctx C) {
		var (
			region         uint64 = 0x0
			maxNodeID      uint64 = 0
			assignedRegion        = make([]uint64, 0)
		)
		namer := NewDefault(region, assignedRegion, maxNodeID)
		ctx.Convey("test name node", func(ctx C) {
			id, err := namer.Node(nil)
			So(err, ShouldBeNil)
			So(id, ShouldEqual, region|maxNodeID+1)

			id, err = namer.Node(nil)
			So(err, ShouldBeNil)
			So(id, ShouldEqual, region|maxNodeID+2)
		})

		ctx.Convey("test name region", func(ctx C) {
			r, err := namer.Region(nil)
			So(err, ShouldBeNil)
			So(r, ShouldEqual, uint64(0x8000000000000000))

			r, err = namer.Region(nil)
			So(err, ShouldBeNil)
			So(r, ShouldEqual, uint64(0xc000000000000000))

			r, err = namer.Region(nil)
			So(err, ShouldBeNil)
			So(r, ShouldEqual, uint64(0xe000000000000000))
		})
	})
}

package naming

import (
	"math"
	"sync"

	"github.com/kyonko-moe/kagami/model"

	"github.com/pkg/errors"
)

// Default is a default node ID namer.
// Node ID assign from minimum id in the current region space.
// BranchNode region assgin from highest but not assgined bit.
type Default struct {
	region         uint64
	assignedRegion []uint64
	maxNodeID      uint64
	sync.Mutex
}

func NewDefault(region uint64, assignedRegion []uint64, maxNodeID uint64) *Default {
	return &Default{
		region:         region,
		assignedRegion: assignedRegion,
		maxNodeID:      maxNodeID | region,
	}
}

func (d *Default) Node(n *model.Node) (id uint64, err error) {
	d.Lock()
	defer d.Unlock()

	regionMask := d.regionMask()
	if math.MaxUint64^regionMask <= d.maxNodeID+1 {
		err = errors.Errorf("no more leaf id can be assigned , maxNodeID : %d %x , regionMask : %x", d.maxNodeID, d.maxNodeID, regionMask)
		return
	}
	d.maxNodeID++
	id = d.maxNodeID
	return
}

func (d *Default) Region(n *model.Node) (region uint64, err error) {
	d.Lock()
	defer d.Unlock()

	var (
		availRegionBit uint = 0
		regionMask          = d.regionMask()
	)
	for ; (regionMask|d.maxNodeID)<<availRegionBit > 0; availRegionBit++ {
	}
	if availRegionBit >= 64 {
		err = errors.Errorf("no more region can be assigned , maxNodeID : %d %x , regionMask : %x", d.maxNodeID, d.maxNodeID, regionMask)
		return
	}
	region = d.maxNodeID + 0x1<<(63-availRegionBit) | regionMask

	d.assignedRegion = append(d.assignedRegion, region)
	return
}

func (d *Default) regionMask() (mask uint64) {
	var (
		region = d.region
	)
	mask = d.region
	for i := uint(1); region > 0; region = region << 1 {
		mask |= 0x1 << i
		i++
	}
	for _, ar := range d.assignedRegion {
		mask |= ar
	}
	return
}

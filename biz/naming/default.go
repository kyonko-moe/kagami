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
	region         int64
	assignedRegion []int64
	maxNodeID      int64
	sync.Mutex
}

func NewDefault(region int64, assignedRegion []int64, maxNodeID int64) *Default {
	return &Default{
		region:         region,
		assignedRegion: assignedRegion,
		maxNodeID:      maxNodeID,
	}
}

func (d *Default) Node(n *model.Node) (id int64, err error) {
	d.Lock()
	defer d.Unlock()

	var (
		regionMask int64 = 0x0
	)
	for _, ar := range d.assignedRegion {
		regionMask |= ar
	}
	if math.MaxInt64^regionMask <= d.maxNodeID {
		err = errors.Errorf("no more leaf id can be assigned , maxNodeID : %d %x , regionMask : %x", d.maxNodeID, d.maxNodeID, regionMask)
		return
	}
	d.maxNodeID++
	id = d.maxNodeID
	return
}

func (d *Default) Region(n *model.Node) (region int64, err error) {
	d.Lock()
	defer d.Unlock()

	var (
		bitIndex   uint  = 1
		regionMask int64 = 0x0
	)
	for _, ar := range d.assignedRegion {
		regionMask |= ar
	}
	for ; bitIndex <= 63; bitIndex++ {
		if regionMask<<bitIndex == 0 {
			break
		}
	}
	if math.MaxInt64<<bitIndex <= d.maxNodeID {
		err = errors.Errorf("no more region can be assigned , maxNodeID : %d %x , regionMask : %x", d.maxNodeID, d.maxNodeID, regionMask)
		return
	}
	region = regionMask + 0x1<<(63-bitIndex)
	d.assignedRegion = append(d.assignedRegion, region)
	return
}

// SPDX-License-Identifier: AGPL-3.0-only
package index

import (
	"encoding/binary"
	"fmt"
	"hash/fnv"
)

// FractalIdx is a 2-level consistent-hash ring.
// Level-0 shards by slot / 4096 → node.
// Level-1 inside each node shards by addr hash → partition.
type FractalIdx struct {
	shifts uint8 // 12 → 4096 slots per shard
}

func New() *FractalIdx { return &FractalIdx{shifts: 12} }

func (f *FractalIdx) Shard(slot uint64) uint32 {
	return uint32(slot >> f.shifts)
}

func (f *FractalIdx) Partition(addr string, shards uint32) uint32 {
	h := fnv.New32a()
	h.Write([]byte(addr))
	return h.Sum32() % shards
}

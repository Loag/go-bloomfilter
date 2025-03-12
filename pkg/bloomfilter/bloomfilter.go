package bloomfilter

import "github.com/Loag/FNV-1a/pkg/fnv1a_fast"

type BloomFilter struct {
	size  uint32
	items []bool
}

func NewBloomFilter(size uint32) *BloomFilter {
	items := make([]bool, size)
	return &BloomFilter{size: size, items: items}
}

func (h *BloomFilter) Add(item []byte) {
	val := fnv1a_fast.FNV1a32(item)
	index := val % h.size
	h.items[index] = true
}

func (h *BloomFilter) Contains(item []byte) bool {
	val := fnv1a_fast.FNV1a32(item)
	index := val % h.size
	return h.items[index] == true
}

func (h *BloomFilter) Remove(item []byte) {
	val := fnv1a_fast.FNV1a32(item)
	index := val % h.size
	h.items[index] = false
}

func (h *BloomFilter) RemoveAll() {
	for i := range h.items {
		h.items[i] = false
	}
}

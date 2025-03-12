package bloomfilter

import "testing"

func TestBloomFilter(t *testing.T) {
	bloomFilter := NewBloomFilter(500)
	bloomFilter.Add([]byte("hello"))
	bloomFilter.Add([]byte("world"))
	bloomFilter.Add([]byte("foo"))
	bloomFilter.Add([]byte("bar"))
	bloomFilter.Add([]byte("baz"))

	if !bloomFilter.Contains([]byte("hello")) {
		t.Errorf("hello should be in the bloomfilter")
	}

}

func TestHashSetRemove(t *testing.T) {
	bloomFilter := NewBloomFilter(500)
	bloomFilter.Add([]byte("hello"))
	bloomFilter.Add([]byte("world"))
	bloomFilter.Add([]byte("foo"))
	bloomFilter.Add([]byte("bar"))

	bloomFilter.Remove([]byte("world"))

	if bloomFilter.Contains([]byte("world")) {
		t.Errorf("world should be removed from the hashset")
	}
}

func TestHashSetRemoveAll(t *testing.T) {
	bloomFilter := NewBloomFilter(500)
	bloomFilter.Add([]byte("hello"))
	bloomFilter.Add([]byte("world"))
	bloomFilter.Add([]byte("foo"))
	bloomFilter.Add([]byte("bar"))

	bloomFilter.RemoveAll()

	if bloomFilter.Contains([]byte("hello")) {
		t.Errorf("hello should be removed from the bloomfilter")
	}
}

func BenchmarkBloomFilterAdd(b *testing.B) {
	bloomFilter := NewBloomFilter(500)
	for i := 0; i < b.N; i++ {
		bloomFilter.Add([]byte("hello"))
	}
}

func BenchmarkBloomFilterContains(b *testing.B) {
	bloomFilter := NewBloomFilter(500)
	bloomFilter.Add([]byte("hello"))
	for i := 0; i < b.N; i++ {
		bloomFilter.Contains([]byte("hello"))
	}
}

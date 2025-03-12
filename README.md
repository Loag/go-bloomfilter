# go-bloomfilter

The intent behind this package is a low overhead, as fast as possible bloom filter.

This projects original use case is for windowed udp packet deduplication.

## Usage

``` go
	bloomFilter := NewBloomFilter(500)

	bloomFilter.Add([]byte("hello"))
	bloomFilter.Contains([]byte("hello"))
```
package main

type BloomFilter struct {
	bitset   []bool
	hashFunc hash.Hash64
}

func NewBloomFilter(n uint32) *BloomFilter {
	return &BloomFilter{
		bitset:   make([]bool, n),
		hashFunc: fnv.New64(),
		}
}

func (bf *BloomFilter) Add(item []byte) {
	bf.hashFunc.Reset()
	bf.hashFunc.Write(item)
	hashVal := bf.hashFunc.Sum(nil)
	index1 := binary.BigEndian.Uint32(hashVal[:4]) % uint32(len(bf.bitset))
	index2 := binary.BigEndian.Uint32(hashVal[4:]) % uint32(len(bf.bitset))
	bf.bitset[index1] = true
	bf.bitset[index2] = true
}

func (bf *BloomFilter) Contains(item []byte) bool {
	bf.hashFunc.Reset()
	bf.hashFunc.Write(item)
	hashVal := bf.hashFunc.Sum(nil)
	index1 := binary.BigEndian.Uint32(hashVal[:4]) % uint32(len(bf.bitset))
	index2 := binary.BigEndian.Uint32(hashVal[4:]) % uint32(len(bf.bitset))
	return bf.bitset[index1] && bf.bitset[index2]
}

package filter

// Bloom Filter 是一个基于概率的数据结构
// 只能告诉我们一个元素绝对不在集合内或可能在集合内
// so we can disable a lot of unnecessary request,such as keys the redis doesn't have at all

import (
	"hash"
	"hash/fnv"
	"math"
)

// p	false-positive
// m	bit size
// n	expected element size
// k	expected hash function size

// return the optimal bit size by expected element number and expected false-positive
func optimalBitSize(n uint, p float64) uint {
	return uint(math.Ceil(-1 * float64(n) * math.Log(p) / math.Pow(math.Log(2), 2)))
}

// return the optimal hash function size by bit size and expected element size
func optimalHashFuncSize(n uint, m uint) uint {
	return uint(math.Ceil(math.Log(2) * float64(m) / float64(n)))
}

type BloomFilter struct {
	k    uint   // the hash function number calculated by optimalHashFuncSize()
	m    uint   // the bit map length calculated by optimalBitSize()
	bits []bool // the bit map
}

// NewBloomFilter return a *BloomFilter by expected element size and expected false-positive
func NewBloomFilter(n uint, p float64) (bf *BloomFilter) {
	bf.m = optimalBitSize(n, p)
	bf.k = optimalHashFuncSize(n, uint(len(bf.bits)))
	bf.bits = make([]bool, bf.m)
	return
}

func (bf *BloomFilter) hash(b []byte) (uint32, uint32) {
	bf.hashFunc.Reset()
	_, _ = bf.hashFunc.Write(b)
	res := bf.hashFunc.Sum64()
	h1 := uint32(res&(1<<32) - 1)
	h2 := uint32(res >> 32)
	return h1, h2
}

func (bf *BloomFilter) Add(b []byte) {
	h1, h2 := bf.hash(b)
	for i := 0; i < int(bf.k); i++ {
		tmp := (h1 + uint32(i)*h2) % uint32(bf.m)
		bf.bits[tmp] = true
	}
	bf.n++
}

// Check	true可能存在, false绝对不存在
func (bf *BloomFilter) Check(b []byte) bool {
	h1, h2 := bf.hash(b)
	res := true
	for i := 0; i < int(bf.k); i++ {
		ind := (h1 + uint32(i)*h2) % uint32(bf.m)
		res = res && bf.bits[ind]
	}
	return res
}

// FalsePositiveRate return the false-positive
func (bf *BloomFilter) FalsePositiveRate() float64 {
	return math.Pow(1-math.Exp(-float64(bf.k*bf.n)/
		float64(bf.m)), float64(bf.k))
}

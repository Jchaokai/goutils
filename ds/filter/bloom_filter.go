package filter

// Bloom Filter 是一个基于概率的数据结构
// 只能告诉我们一个元素绝对不在集合内或可能在集合内
// so we can disable a lot of unnecessary request,such as keys the redis doesn't have at all

import (
	"github.com/Jchaokai/goutils/ds/set"
	"hash/fnv"
	"math"
	"sync"
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
	k            uint        // the hash function number calculated by optimalHashFuncSize()
	m            uint        // the bit map size calculated by optimalBitSize()
	bits         *set.BitSet // bit map
	sync.RWMutex             // ensure thread safe
}

// NewBloomFilter return a *BloomFilter by expected element size and expected false-positive
//
// Example:
//		bf := NewBloomFilter(1000000,0.0001)
//
func NewBloomFilter(n uint, p float64) *BloomFilter {
	bf := &BloomFilter{}
	bf.m = optimalBitSize(n, p)
	bf.k = optimalHashFuncSize(n, bf.m)
	bf.bits = set.NewBitSet()
	if bf.k < 4 { // hash function minimum size is 4
		bf.k = 4
	}
	return bf
}

func (bf *BloomFilter) hash(b []byte) uint64 {
	hash64 := fnv.New64()
	_, _ = hash64.Write(b)
	return hash64.Sum64() % uint64(bf.m)
}

func (bf *BloomFilter) Add(b []byte) {
	bf.Lock()
	defer bf.Unlock()
	for i := 0; i < int(bf.k); i++ {
		hashed := bf.hash(append(b, byte(i)))
		bf.bits.Add(int(hashed))
	}
}

// Check	true可能存在, false绝对不存在
func (bf *BloomFilter) Check(b []byte) bool {
	bf.RLock()
	defer bf.RUnlock()
	for i := 0; i < int(bf.k); i++ {
		hashed := bf.hash(append(b, byte(i)))
		if !bf.bits.Contains(int(hashed)) {
			return false
		}
	}
	return true
}

// FalsePositiveRate return the false-positive
func (bf *BloomFilter) FalsePositiveRate() float64 {
	//return math.Pow((1 - math.Exp(-float64(bf.k*bf.elementNum)/
	//	float64(bf.m))), float64(bf.k))
	return 0
}

package filter

import (
	"math/rand"
	"strconv"
	"testing"
)

func TestBloomFilter_Add(t *testing.T) {
	bloomFilter := NewBloomFilter(100000, 0.01)
	for i := 0; i < 100; i++ {
		bloomFilter.Add([]byte("image" + strconv.Itoa(i)))
	}
}

func TestBloomFilter_Check(t *testing.T) {
	bloomFilter := NewBloomFilter(100000, 0.01)
	for i := 0; i < 100; i++ {
		bloomFilter.Add([]byte("image" + strconv.Itoa(i)))
	}

	for i := 0; i < 100; i++ {
		randID := rand.Intn(200)
		if !bloomFilter.Check([]byte("image" + strconv.Itoa(randID))) {
			if randID < 100 {
				t.Error(randID)
			}
		}
	}
}

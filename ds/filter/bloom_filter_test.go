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

//  goos: windows
//  goarch: amd64
//  pkg: github.com/Jchaokai/goutils/ds/filter
//  cpu: Intel(R) Core(TM) i5-9400F CPU @ 2.90GHz
//  BenchmarkBloomFilter_Check
//  BenchmarkBloomFilter_Check-6   	 5702342	       243.2 ns/op
//  PASS

func BenchmarkBloomFilter_Check(b *testing.B) {
	bloomFilter := NewBloomFilter(uint(b.N), 0.02)
	for i := 0; i < b.N; i++ {
		bloomFilter.Add([]byte("image" + strconv.Itoa(i)))
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		rand.Intn(b.N)
		bloomFilter.Check([]byte("image" + strconv.Itoa(i)))
	}
}

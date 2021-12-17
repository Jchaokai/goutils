package set

import (
	"math"
	"math/bits"
)

const (
	IntSize = 32 << (^uint(0) >> 63) // 32 or 64

	shift = 6    // get the index in slice	n/64 => n>>6
	mask  = 0x3f // get the position in 64 bits	n%64 => n&63 => n&0x3f
)

type (
	// BitSet not thread safe
	BitSet struct {
		data []uint64
	}
)

// NewBitSet negative element will not included in bitset
func NewBitSet(n ...int) *BitSet {
	if len(n) == 0 {
		return new(BitSet)
	}
	max := n[0]
	for _, v := range n {
		if v > max {
			max = v
		}
	}
	if max < 0 {
		return new(BitSet)
	}

	s := &BitSet{make([]uint64, max>>shift+1)}
	for _, v := range n {
		if v >= 0 {
			s.data[v>>6] |= 1 << uint(v&mask)
		}
	}
	return s
}

func (s *BitSet) realloc(n int) bool {
	if c := cap(s.data); c < n {
		s.data = make([]uint64, n, newCap(n, c))
		return true
	}
	// set 0 if shrinking
	d := s.data
	for i := len(d) - 1; i >= n; i-- {
		d[i] = 0
	}
	s.data = d[:n]
	// todo reallocate if cap(s.data) too large than len(s.data)
	return false
}

func newCap(newCap int, oldCap int) int {
	i := pow2(oldCap)
	if newCap > i {
		return newCap
	}
	return i
}

// return 1 2 4 8 ...... math.MaxInt
func pow2(cap int) int {
	if cap <= 0 {
		return 1
	}
	if k := 64 - bits.LeadingZeros64(uint64(cap)); k < IntSize-1 {
		return 1 << uint(k)
	}
	return math.MaxInt
}

func (s *BitSet) resize(n int) {
	d := s.data
	if s.realloc(n) {
		copy(s.data, d)
	}
}

// remove all trailing words == 0
func (s *BitSet) trim() {
	n := len(s.data) - 1
	for n >= 0 && s.data[n] == 0 {
		n--
	}
	s.data = s.data[:n+1]
}

// Add negative element will not be added
func (s *BitSet) Add(n int) {
	if n < 0 {
		return
	}

	i := n >> shift
	if i >= len(s.data) {
		s.resize(i + 1)
	}
	s.data[i] |= 1 << uint(n&mask)
}

func (s *BitSet) Remove(n int) {
	if n < 0 {
		return
	}
	i := n >> shift
	if i >= len(s.data) {
		return
	}

	s.data[i] &^= 1 << uint(n&mask)
	s.trim()
}

func (s *BitSet) Contains(n int) bool {
	if n < 0 {
		return false
	}
	i := n >> shift
	if i >= len(s.data) {
		return false
	}
	return s.data[i]&(1<<uint(n&mask)) != 0
}

// Max if bitset is empty return -1
func (s *BitSet) Max() int {
	if s.Empty() {
		return -1
	}
	i := len(s.data) - 1
	return i<<shift + 63 - bits.LeadingZeros64(s.data[i])
}

func (s *BitSet) Size() int {
	n := 0
	for i, lens := 0, len(s.data); i < lens; i++ {
		if w := s.data[i]; w != 0 {
			n += bits.OnesCount64(w)
		}
	}
	return n
}

func (s *BitSet) Empty() bool {
	return len(s.data) == 0
}

// Next return the next element > m, or -1 if no such element
func (s *BitSet) Next(m int) int {
	if s.Empty() {
		return -1
	}
	if m < 0 {
		// 最小值是1
		if s.data[0]&1 != 0 {
			return 0
		}
		m = 0
	}

	i := m >> shift
	if i > len(s.data) {
		return -1
	}

	t := 1 + uint(m&mask)
	w := s.data[i] >> t << t // 将m(包括m)右边的比特位全部归0
	// 是否需要跨行寻找结果
	for i < len(s.data)-1 && w == 0 {
		i++
		w = s.data[i]
	}
	if w == 0 {
		return -1
	}
	return i<<shift + bits.TrailingZeros64(w)
}

// Prev return the prev element < m, or -1 if no such element
func (s *BitSet) Prev(m int) int {
	if s.Empty() || m <= 0 {
		return -1
	}
	if max := s.Max(); max < m {
		return max
	}
	i := m >> shift
	t := 64 - uint(m&mask)
	w := s.data[i] << t >> t // 将m(包括m)左边的比特位全部归0
	// 跨行查找
	for i > 0 && w == 0 {
		i--
		w = s.data[i]
	}
	if w == 0 {
		return -1
	}
	return i<<shift + 63 - bits.LeadingZeros64(w)
}

package bit_set

//  the size of bit_set normally is fixed

type BitSet struct {
	data []int64
	size int
}

func (b *BitSet) Size() int {
	return b.size
}

package set

import (
	"fmt"
	"testing"
)

func TestNewBitSet(t *testing.T) {
	bitSet := NewBitSet(1, 200, 333, 4333, 53333)
	fmt.Println(bitSet.Size())
}

func TestBitSet_Size(t *testing.T) {
	bitSet := NewBitSet(1, 22, 333)
	if bitSet.Size() != 3 {
		t.Error()
	}
	t.Log(bitSet.Size())

	bitSet2 := NewBitSet()
	if bitSet2.Size() != 0 {
		t.Error()
	}
	t.Log(bitSet2.Size())
}

func TestBitSet_Empty(t *testing.T) {
	bitSet := NewBitSet(1, 23, 333)
	if bitSet.Empty() {
		t.Error()
	} else {
		t.Log(bitSet.Size())
	}
}

func TestBitSet_Max(t *testing.T) {
	bitSet := NewBitSet(1, 22, 333, 4444)
	if bitSet.Max() != 4444 {
		t.Error()
	}
	t.Log(bitSet.Max())
}

func TestBitSet_Contains(t *testing.T) {
	bitSet := NewBitSet(1, 22, 333, 321, 21, 11)
	if !bitSet.Contains(22) {
		t.Error()
	}
	if !bitSet.Contains(11) {
		t.Error()
	}
	if !bitSet.Contains(321) {
		t.Error()
	}
	if !bitSet.Contains(21) {
		t.Error()
	}
	if !bitSet.Contains(333) {
		t.Error()
	}
	if !bitSet.Contains(1) {
		t.Error()
	}
}

func TestBitSet_Remove(t *testing.T) {
	bitSet := NewBitSet(1, 22, 333, 321, 21, 11)
	bitSet.Remove(22)
	if bitSet.Contains(22) {
		t.Error()
	}
	bitSet.Remove(11)
	if bitSet.Contains(11) {
		t.Error()
	}
	bitSet.Remove(321)
	if bitSet.Contains(321) {
		t.Error()
	}
	bitSet.Remove(21)
	if bitSet.Contains(21) {
		t.Error()
	}
	bitSet.Remove(333)
	if bitSet.Contains(333) {
		t.Error()
	}
	bitSet.Remove(1)
	if bitSet.Contains(1) {
		t.Error()
	}
	t.Log(bitSet.Size())
}

func TestBitSet_Add(t *testing.T) {
	bitSet := NewBitSet(1, 22, 333, 321, 4444)
	bitSet.Add(111)
	if !bitSet.Contains(111) {
		t.Error()
	}
	bitSet.Add(-111)
	if bitSet.Contains(-111) {
		t.Error()
	}
	bitSet.Remove(333)
	bitSet.Add(333)
	if !bitSet.Contains(333) {
		t.Error()
	}
	t.Log(bitSet.Size())
}

func TestBitSet_Next(t *testing.T) {
	bitSet := NewBitSet(1, 22, 333, 321, 300, 301)
	if bitSet.Next(1) != 22 {
		t.Error()
	}
	if bitSet.Next(22) != 300 {
		t.Error()
	}
	if bitSet.Next(300) != 301 {
		t.Error()
	}
	t.Log(bitSet.Next(333))
	if bitSet.Next(-111) != 1 {
		t.Error()
	} else {
		t.Log(bitSet.Next(-111))
	}
}

func TestBitSet_Prev(t *testing.T) {
	bitSet := NewBitSet(1, 22, 333, 321, 300, 301)
	if bitSet.Prev(1) != -1 {
		t.Error()
	}
	if bitSet.Prev(22) != 1 {
		t.Error()
	}
	if bitSet.Prev(301) != 300 {
		t.Error()
	}
	if bitSet.Prev(333) != 321 {
		t.Error()
	}
}

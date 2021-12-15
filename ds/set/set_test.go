package set

import (
	"testing"
)

func TestSafeSet_Add(t *testing.T) {
	s := NewSet(1, 2, 3)
	s.Add("dsadas")
	if s.Size() != 4 {
		t.Error()
	}
}

func TestSafeSet_Contains(t *testing.T) {
	s := NewSet("a", "b", "c")
	s.Clear()
	if s.Size() != 0 {
		t.Error()
	}
}

// todo

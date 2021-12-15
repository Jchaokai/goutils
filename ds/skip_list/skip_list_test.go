package skip_list

import (
	"math/rand"
	"testing"
	"time"
)

func TestSkipListForDebug(t *testing.T) {
	skiplist := NewSkipList()
	skiplist.Add(1)
	skiplist.Add(2)
	skiplist.Add(3)
	if skiplist.Search(0) {
		t.Fatal("no '0' but search success")
	}
	skiplist.Add(4)
	if !skiplist.Search(1) {
		t.Fatal("do not search '1' but exists")
	}
	if skiplist.Delete(0) {
		t.Fatal("no '0' to delete")
	}
	if !skiplist.Delete(1) {
		t.Fatal("delete '1' not success")
	}
	if skiplist.Delete(1) {
		t.Fatal("no '1' exists to delete,but success")
	}
}

func randInt(min, max int) int {
	rand.Seed(time.Now().UnixNano())
	return min + rand.Intn(max-min)
}

func TestSkipList(t *testing.T) {
	const numOfTask = 1000
	contrastSkiplist := NewSkipList()
	skiplist := NewSkipList()
	for i := 0; i < numOfTask; i++ {
		num := randInt(0, numOfTask>>3)
		switch randInt(0, 3) {
		case 0:
			skiplist.Add(num)
			contrastSkiplist.Add(num)
		case 1:
			if skiplist.Search(num) != contrastSkiplist.Search(num) {
				t.Fail()
			}
		case 2:
			if skiplist.Delete(num) != contrastSkiplist.Delete(num) {
				t.Fail()
			}
		}
	}
}

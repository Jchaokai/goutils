package distribute

import (
	"sort"
	"testing"
)

func checkNum(num, expected int, t *testing.T) {
	if num != expected {
		t.Errorf("got %d, expected %d", num, expected)
	}
}

func TestNew(t *testing.T) {
	x := NewConsistentHash()
	if x == nil {
		t.Errorf("expected obj")
	}
	checkNum(x.ReplicasNum, 100, t)
}

func TestConsHash_Add(t *testing.T) {
	x := NewConsistentHash()
	x.Add("abcdefg")
	checkNum(len(x.circle), 100, t)
	checkNum(len(x.sortedHashes), 100, t)

	if !sort.SliceIsSorted(x.sortedHashes, func(i, j int) bool {
		return x.sortedHashes[i] < x.sortedHashes[j]
	}) {
		t.Log("sortedHashed un sorted")
	}

	x.Add("qwer")
	checkNum(len(x.circle), 200, t)
	checkNum(len(x.sortedHashes), 200, t)

	if !sort.SliceIsSorted(x.sortedHashes, func(i, j int) bool {
		return x.sortedHashes[i] < x.sortedHashes[j]
	}) {
		t.Errorf("sortedHashes un sorted")
	}
}

func TestConsHash_Remove(t *testing.T) {
	x := NewConsistentHash()
	x.Add("qwer")
	t.Logf("len(sortedHashes): %d, cap(sortedHashes): %d", len(x.sortedHashes), cap(x.sortedHashes))
	x.Remove("qwer")
	checkNum(len(x.circle), 0, t)
	checkNum(len(x.sortedHashes), 0, t)
	t.Logf("len(sortedHashes): %d, cap(sortedHashes): %d", len(x.sortedHashes), cap(x.sortedHashes))
}

func TestConsHash_RemoveNoExisting(t *testing.T) {
	x := NewConsistentHash()
	x.Add("qwer")
	x.Remove("qwertyuiop")
	checkNum(len(x.circle), 100, t)
	checkNum(len(x.sortedHashes), 100, t)
}

// todo another test

//package main
//
//import (
//"fmt"
//"github.com/Jchaokai/goutils/distribute"
//"strconv"
//)
//
//func main() {
//	// test consistent hash in package 'distribute'
//	c := distribute.NewConsistentHash()
//	// add 10 physical node
//	for i := 0; i < 10; i++ {
//		c.Add("node"+strconv.Itoa(i))
//	}
//	fmt.Println("the node size: ",c.Size())
//	nodeSize := make(map[string]uint64)
//	// now we have 100000 images, which node will storage some of them
//	for i := 0; i < 100000; i++ {
//		get,_ := c.Get("image" + strconv.Itoa(i))
//		nodeSize[get]++
//	}
//	// print the size of image storage in these 10 node
//	for name, size := range nodeSize {
//		fmt.Println(name," => ",size)
//	}
//	fmt.Println("===============after add node10 node11 =================")
//	// after add a node
//	c.Add("node10")
//	c.Add("node11")
//	nodeSizeAdded := make(map[string]uint64)
//	for i := 0; i < 100000; i++ {
//		get,_ := c.Get("image" + strconv.Itoa(i))
//		nodeSizeAdded[get]++
//	}
//	for i := 0; i < 10; i++ {
//		remap := nodeSize["node"+strconv.Itoa(i)] - nodeSizeAdded["node"+strconv.Itoa(i)]
//		fmt.Println("node"+strconv.Itoa(i), " remap size: ",remap," current new size: ", nodeSizeAdded["node"+strconv.Itoa(i)])
//	}
//	fmt.Println("new node10 => ", nodeSizeAdded["node10"])
//	fmt.Println("new node11 => ", nodeSizeAdded["node11"])
//
//	fmt.Println("=============after remove node 11 ======================")
//	c.Remove("node11")
//	nodeSizeRemoved := make(map[string]uint64)
//	for i := 0; i < 100000; i++ {
//		get, _ := c.Get("image" + strconv.Itoa(i))
//		nodeSizeRemoved[get]++
//	}
//	for i := 0; i <= 10; i++ {
//		remap := nodeSizeRemoved["node"+strconv.Itoa(i)] - nodeSizeAdded["node"+strconv.Itoa(i)]
//		fmt.Println("node"+strconv.Itoa(i)," remap size: ",remap," current new size: ",nodeSizeRemoved["node"+strconv.Itoa(i)])
//	}
//}

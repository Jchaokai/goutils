package distribute

// consistent hash implemented in a Ring-based way
// hash function used is "hash/fnv"

import (
	"errors"
	"hash/fnv"
	"sort"
	"strconv"
	"sync"
)

var (
	ErrEmptyCircle = errors.New("the consistent hash circle is empty")
)

type ConsHash struct {
	circle       map[uint32]string // consistent hash ring, key: hash of physical node, value: physical node info
	ReplicasNum  int               // the num of virtual nodes associated with each physical node
	sortedHashes []uint32          // the hash of physical node is sorted
	members      map[string]bool   // the info of physical node
	count        uint64
	sync.RWMutex
}

// NewConsistentHash return new Consistent hash
// the default of ReplicasNum is 100, you can set it before Add() called
func NewConsistentHash() *ConsHash {
	return &ConsHash{
		circle:       make(map[uint32]string),
		ReplicasNum:  100,
		sortedHashes: make([]uint32, 0),
		members:      make(map[string]bool),
	}
}

// using fnv hash function as default
func (ch *ConsHash) hashKey(key string) uint32 {
	hash32 := fnv.New32a()
	_, _ = hash32.Write([]byte(key))
	return hash32.Sum32()
}

// wrap the physical node information of index
func (ch *ConsHash) elementKey(ele string, idx int) string {
	return strconv.Itoa(idx) + ele
}

// Add a new physical node
func (ch *ConsHash) Add(ele string) {
	ch.Lock()
	defer ch.Unlock()
	for i := 0; i < ch.ReplicasNum; i++ {
		ch.circle[ch.hashKey(ch.elementKey(ele, i))] = ele
	}
	ch.sortHashes()
	ch.count++
}

// Remove a physical node
func (ch *ConsHash) Remove(ele string) {
	ch.Lock()
	defer ch.Unlock()
	for i := 0; i < ch.ReplicasNum; i++ {
		delete(ch.circle, ch.hashKey(ch.elementKey(ele, i)))
	}
	delete(ch.members, ele)
	ch.sortHashes()
	ch.count--
}

//
func (ch *ConsHash) sortHashes() {
	hashes := ch.sortedHashes[:0]
	// todo reallocate if cap(ch.sortedHashes) too large
	//if cap(ch.sortedHashes)/(ch.ReplicasNum*4) > len(ch.circle) {
	//	hashes = nil
	//}

	for k := range ch.circle {
		hashes = append(hashes, k)
	}
	sort.Slice(hashes, func(i, j int) bool {
		return hashes[i] < hashes[j]
	})
	ch.sortedHashes = hashes
}

// Get return a first node > input hashes in the circle
func (ch *ConsHash) Get(name string) (string, error) {
	ch.RLock()
	defer ch.RUnlock()
	if len(ch.circle) == 0 {
		return "", ErrEmptyCircle
	}
	key := ch.hashKey(name)
	i := ch.search(key)
	return ch.circle[ch.sortedHashes[i]], nil
}

func (ch *ConsHash) GetN(name string, n int) ([]string, error) {
	ch.RLock()
	defer ch.RUnlock()

	if len(ch.circle) == 0 {
		return nil, ErrEmptyCircle
	}

	if n <= 0 {
		return nil, errors.New("invalid input")
	}

	if ch.count < uint64(n) {
		n = int(ch.count)
	}
	key := ch.hashKey(name)
	i := ch.search(key)
	start := i
	res := make([]string, 0, n)
	elem := ch.circle[ch.sortedHashes[i]]

	res = append(res, elem)

	for i = start + 1; i != start; i++ {
		if i > len(ch.sortedHashes) {
			i = 0
		}
		elem = ch.circle[ch.sortedHashes[i]]
		// todo append same elem ?
		// res = append(res, elem)
		if len(res) == n {
			break
		}
	}
	return res, nil
}

// search return the index of first node which hash > input hash
func (ch *ConsHash) search(key uint32) (i int) {
	i = sort.Search(len(ch.sortedHashes), func(i int) bool {
		return ch.sortedHashes[i] > key
	})
	if i >= len(ch.sortedHashes) {
		i = 0 // no element > key,we should set zero because it's a ring
	}
	return
}

func (ch *ConsHash) Members() []string {
	ch.RLock()
	defer ch.RUnlock()
	var res []string
	for k := range ch.members {
		res = append(res, k)
	}
	return res
}

func (ch *ConsHash) Size() uint64 {
	return ch.count
}

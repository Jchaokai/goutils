package skip_list

// https://leetcode-cn.com/problems/design-skiplist/solution/zui-jian-dan-de-golangtiao-biao-shi-xian-by-jayust/

import (
	"math"
	"math/rand"
)

type (
	Node struct {
		val        int
		next, down *Node
	}

	SkipList struct {
		level int // 0: null
		Head  *Node
	}
)

func NewNode(val int, r, d *Node) *Node {
	return &Node{val, r, d}
}

func NewSkipList() *SkipList {
	return &SkipList{1, NewNode(math.MinInt, nil, nil)}
}

func (s *SkipList) Contains(target int) bool {
	cur := s.Head
	for cur != nil {
		for cur.next != nil && cur.next.val < target {
			cur = cur.next
		}
		if cur.next != nil && cur.next.val == target {
			return true
		}
		cur = cur.down
	}
	return false
}

func (s *SkipList) Add(val int) {
	// 新建的val 需要在最上面几层添加
	nlevel := 1
	for nlevel <= s.level && rand.Int()&1 == 0 {
		nlevel++
	}
	// 是否新加level
	if nlevel > s.level {
		s.level = nlevel
		s.Head = NewNode(val, nil, s.Head)
	}
	cur := s.Head
	var last *Node // 保存上一层新加的Node,为了与下一层要新加的Node的连接 last.down = cur.next
	for l := s.level; l >= 1; l-- {
		for cur.next != nil && cur.next.val < val {
			cur = cur.next
		}
		if l <= nlevel { //此时cur.next为空, 或 > val，新加节点
			cur.next = NewNode(val, cur.next, nil)
			if last != nil {
				last.down = cur.next
			}
			last = cur.next
		}
		cur = cur.down
	}
}

func (s *SkipList) Delete(val int) bool {
	cur := s.Head
	seen := false
	for cur != nil {
		for cur.next != nil && cur.next.val < val {
			cur = cur.next
		}
		if cur.next != nil && cur.next.val == val {
			seen = true
			cur.next = cur.next.next
		}
		cur = cur.down
	}
	return seen
}

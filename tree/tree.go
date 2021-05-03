package tree

import (
	"fmt"
	"goutils/queue"
)

type Tree struct {
	Val   interface{}
	Left  *Tree
	Right *Tree
	//节点数
	Num int
}

// 通过带有nil的 slice 创建 二叉树
// s := []interface{}{1,2,3,4,5,nil,6}
func NewBTree(slice []interface{}) *Tree {
	current := &Tree{slice[0], nil, nil, len(slice)}
	root := current
	for i := 1; i < len(slice); i += 2 {
		current.Left = &Tree{slice[i], nil, nil, 0}
		current.Right = &Tree{slice[i+1], nil, nil, 0}
		current = current.Right
	}
	return root
}

// 前序遍历
func NLR(root *Tree) (res []interface{}) {
	res = append(res, root.Val)
	if root.Left != nil {
		res = append(res, NLR(root.Left)...)
	}
	if root.Right != nil {
		res = append(res, NLR(root.Right)...)
	}
	return
}

// 中序遍历
func LNR(root *Tree) (res []interface{}) {
	if root.Left != nil {
		res = append(res, LNR(root.Left)...)
	}
	res = append(res, root.Val)
	if root.Right != nil {
		res = append(res, LNR(root.Right)...)
	}
	return
}

// 后续遍历
func LRN(root *Tree) (res []interface{}) {
	if root.Left != nil {
		res = append(res, LRN(root.Left)...)
	}
	if root.Right != nil {
		res = append(res, LRN(root.Right)...)
	}
	res = append(res, root.Val)
	return
}

// 层次遍历

// bfs
func bfs(root *Tree) {
	s := make([]*Tree, 0)
	if root != nil {
		s = append(s, root)
	}
	for len(s) != 0 {
		// pop
		pop := s[0]
		s = s[1:]

		// TODO do something with node

		if pop.Left != nil {
			s = append(s, pop.Left)
		}
		if pop.Right != nil {
			s = append(s, pop.Right)
		}
	}
}

// bfs遍历,包括nil
func bfsWithNil(root *Tree) {
	s := make([]*Tree, 0)
	if root != nil {
		s = append(s, root)
	}
	for len(s) != 0 {
		// pop
		pop := s[0]
		s = s[1:]

		if pop != nil {
			// TODO do something with node
		} else {
			// TODO do something with nil node
			continue
		}

		if pop.Left == nil && pop.Right == nil {
			continue
		}
		s = append(s, pop.Left)
		s = append(s, pop.Right)
	}

}

func dfs(currentNode *Tree) {
	if currentNode == nil {
		return
	}

	// TODO do something with "current node"

	if currentNode.Left != nil {
		dfs(currentNode.Left)
	}
	if currentNode.Right != nil {
		dfs(currentNode.Right)
	}
}

// dfs遍历,包括nil节点
func dfsWithNil(currentNode *Tree) {
	if currentNode == nil {
		// TODO do something with  "nil node"
		return
	}
	// TODO do something with "currentNode"
	// "currentNode"有一个非nil的子节点，再递归处理
	if currentNode.Left != nil || currentNode.Right != nil {
		dfsWithNil(currentNode.Left)
		dfsWithNil(currentNode.Right)
	}
}
func (t *Tree) String() string {
	res := "Tree : \n"
	//bfs + queue 遍历二叉树
	q := queue.NewQueue(t.Num)
	if t != nil {
		q.Add(t)
		res += fmt.Sprintf(" %v ", t.Val)
	}
	for !q.IsNil() {
		poptree := q.Pop().(*Tree)
		if poptree.Val != nil {
			res += "\n"
		}
		if poptree.Left != nil {
			q.Add(poptree.Left)
			if poptree.Left.Val != nil {
				res += fmt.Sprintf(" %v ", poptree.Left.Val)
			} else {
				res += " nil "
			}
		}
		if poptree.Right != nil {
			q.Add(poptree.Right)
			if poptree.Right.Val != nil {
				res += fmt.Sprintf(" %v ", poptree.Right.Val)
			} else {
				res += " nil "
			}
		}
	}
	return res
}

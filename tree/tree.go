package tree

import (
	"fmt"
	"goutils/queue"
)

type Tree struct {
	Val interface{}
	Left *Tree
	Right *Tree
	//节点数
	Num int
}

// 通过带有nil的 slice 创建 二叉树
// s := []interface{}{1,2,3,4,5,nil,6}
func NewBTree(slice []interface{}) *Tree{
	current := &Tree{slice[0],nil,nil,len(slice)}
	root := current
	for i:= 1 ; i < len(slice) ; i +=2 {
			current.Left =  &Tree{slice[i],nil,nil,0}
			current.Right = &Tree{slice[i+1],nil,nil,0}
			current = current.Right
	}
	return root
}


func (t *Tree) String() string{
	res := "Tree : \n"
	//bfs + queue 遍历二叉树
	q := queue.NewQueue(t.Num)
	if t != nil {
		q.Add(t)
		res += fmt.Sprintf(" %v ",t.Val)
	}
	for !q.IsNil() {
		poptree := q.Pop().(*Tree)
		if poptree.Val != nil {
			res += "\n"
		}
		if poptree.Left != nil {
			q.Add(poptree.Left)
			if poptree.Left.Val != nil {
				res +=fmt.Sprintf(" %v ",poptree.Left.Val)
			}else{
				res += " nil "
			}
		}
		if poptree.Right != nil {
			q.Add(poptree.Right)
			if poptree.Right.Val != nil {
				res += fmt.Sprintf(" %v ",poptree.Right.Val)
			}else{
				res += " nil "
			}
		}
	}
	return res
}




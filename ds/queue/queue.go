package queue

import "fmt"

//定长不可变队列
type Queue struct {
	slice []interface{}
	len   int
}

//可变队列

// len == 0 is ok
func NewQueue(len int) *Queue {
	return &Queue{
		slice: make([]interface{}, 0),
		len:   len,
	}
}
func (q *Queue) IsNil() bool {
	if len(q.slice) == 0 {
		return true
	}
	return false
}

func (q *Queue) IsFull() bool {
	if len(q.slice) == q.len {
		return true
	}
	return false
}

func (q *Queue) Add(v interface{}) {
	if q.IsFull() {
		panic("queue if full")
	}
	q.slice = append(q.slice, v)
}

func (q *Queue) Pop() interface{} {
	if q.IsNil() {
		return nil
	}
	res := q.slice[0]
	q.slice = q.slice[1:len(q.slice)]
	return res
}

func (q *Queue) String() string {
	s := "Queue : "
	for _, v := range q.slice {
		s += fmt.Sprintf("|%v", v)
	}
	return s + "| end"
}

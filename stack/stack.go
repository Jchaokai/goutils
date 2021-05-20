package stack

type Stack struct {
	i    int //要插入元素的位置，要删除的元素上一个位置
	data [10]interface{}
}

func (s Stack) Push(data interface{}) {
	s.data[s.i] = data
	s.i++
}

func (s Stack) Pop() interface{} {
	s.i--
	return s.data[s.i]
}

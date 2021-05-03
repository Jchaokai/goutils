package queue

/*
	当front=rear时，为空队列，而在循环队列中，队列满时，front也等于rear，将无法判断队满和空的情况。

　　一种办法是设置一个标志变量flag，当front=rear时，通过判断flag是0还是1来确定队列满空情况；

　　另一种方法是，在数组中只剩一个空闲单位时，定义为队列满，如下图所示。（本文程序采用这种办法）
*/

type cqueue struct {
	data  []interface{}
	front int
	rear  int
}

func New(size int) *cqueue {
	return &cqueue{
		data:  make([]interface{}, size),
		front: 0,
		rear:  0,
	}
}

func (q *cqueue) Length() int {
	return (q.rear + len(q.data) - q.front) % len(q.data)
}

func (q *cqueue) Add(date interface{}) bool {
	if (q.rear+1)%len(q.data) == q.front {
		return false
	}
	q.data[q.rear] = date
	q.rear = (q.rear + 1) % len(q.data)
	return true
}

func (q *cqueue) Pull() (data interface{}, ok bool) {
	if q.rear == q.front {
		return nil, false
	}

	e := q.data[q.front]
	q.data[q.front] = nil
	q.front = (q.front + 1) % len(q.data)
	return e, true
}

func (q *cqueue) P() {
	for i := 0; i < len(q.data)-1; i++ {
		if q.data[i] == nil {
			continue
		}
		print(q.data[i].(string) + " ")
	}
	println()
}

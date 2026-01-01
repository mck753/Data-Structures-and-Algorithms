package easy

type Queue struct {
	data              []int
	front, rear, size int
}

func NewQueue(capacity int) *Queue {
	return &Queue{
		data:  make([]int, capacity),
		front: 0,
		rear:  0,
		size:  0,
	}
}

func (q *Queue) Enqueue(val int) {
	if q.size == len(q.data) {
		// 扩容（简化）
		newData := make([]int, len(q.data)*2)
		for i := 0; i < q.size; i++ {
			newData[i] = q.data[(q.front+i)%len(q.data)]
		}
		q.data = newData
		q.front = 0
		q.rear = q.size
	}
	q.data[q.rear] = val
	q.rear = (q.rear + 1) % len(q.data)
	q.size++
}

func (q *Queue) Dequeue() (int, bool) {
	if q.size == 0 {
		return 0, false
	}
	val := q.data[q.front]
	q.front = (q.front + 1) % len(q.data)
	q.size--
	return val, true
}

func (q *Queue) Front() (int, bool) {
	if q.size == 0 {
		return 0, false
	}
	return q.data[q.front], true
}

func (q *Queue) IsEmpty() bool {
	return q.size == 0
}

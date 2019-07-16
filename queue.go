package ga

type Queue struct {
	values []interface{}
}

func NewQueue() *Queue {
	return &Queue{values: []interface{}{}}
}

func (q *Queue) Enqueue(i interface{}) {
	q.values = append(q.values, i)
}

func (q *Queue) Dequeue() interface{} {
	if len(q.values) == 0 {
		return nil
	}
	v := q.values[0]
	q.values = q.values[1:]
	return v
}

func (q Queue) Empty() bool {
	return len(q.values) == 0
}

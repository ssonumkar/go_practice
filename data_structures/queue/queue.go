package queue

//Queue ds
type Queue []string

//IsEmpty used to check queue empty
func (q *Queue) IsEmpty() bool {
	if len(*q) == 0 {
		return true
	}
	return false
}

//Enqueue is method to insert ele from rear pos
func (q *Queue) Enqueue(str string) {
	*q = append(*q, str)
}

//Dequeue is method to delete ele from front pos
func (q *Queue) Dequeue() bool {
	if q.IsEmpty() {
		return false
	}
	*q = (*q)[1:]
	return true
}

//GetRearElement to return rear ele
func (q *Queue) GetRearElement() string {
	return (*q)[len(*q)-1]
}

//GetFrontElement to return ele at front
func (q *Queue) GetFrontElement() string {
	return (*q)[0]
}

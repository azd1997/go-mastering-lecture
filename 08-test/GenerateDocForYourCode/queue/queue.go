package queue

// 先进先出队列
type Queue []int

// 将元素压入栈底
// 		e.g. q.Push(123)
func (q *Queue) Push(v int) {
	*q = append(*q, v)
}

// 将栈顶元素弹出
func (q *Queue) Pop() int {
	head := (*q)[0]
	*q = (*q)[1:]
	return head
}

// 判断队列是否为空
func (q *Queue) IsEmpty() bool {
	return len(*q) == 0
}

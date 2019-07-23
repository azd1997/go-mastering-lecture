package queue

import (
	"fmt"
	"testing"
)

func ExampleQueue_Push() {
	q := Queue{1}
	fmt.Println(q)
	q.Push(2)
	fmt.Println(q)

	// Output:
	// [1]
	// [1 2]
}

func TestQueue_Push(t *testing.T) {
	q := Queue{1}
	fmt.Println(q)
	q.Push(2)
	fmt.Println(q)
}

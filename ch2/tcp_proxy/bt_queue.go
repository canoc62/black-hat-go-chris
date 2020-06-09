package main

import "fmt"

type BinaryTree struct {
	Value       int
	Left, Right *BinaryTree
}

type queue struct {
	data []*BinaryTree
}

func newQueue() *queue {
	q := queue{ data: []*BinaryTree{} }
	return &q
}

func (q *queue) add(bt *BinaryTree) {
	q.data = append(q.data, bt)
}

func(q *queue) dequeue() *BinaryTree {
	if len(q.data) == 0 {
		return nil
	}
	bt := q.data[0]
	q.data = q.data[1:]
	return bt
}

func printQueueVals(q *queue) {
	for _, node := range((*q).data) {
		fmt.Println(node.Value)
	}
}

func main() {
	q1 := newQueue()

	q1.add(&BinaryTree{Value: 1})
	fmt.Println(q1.data[0].Value)
	q1.add(&BinaryTree{Value: 2})
	fmt.Println(q1.data)
	printQueueVals(q1)
	dequeued := q1.dequeue()
	fmt.Println("Dequeued: " + string((*dequeued).Value))
	fmt.Println(q1.data)
}
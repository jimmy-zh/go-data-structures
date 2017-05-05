package queue

import "data-structures/container"

type Queuer interface {
	Enqueue(v interface{})
	Dequeue() interface{}
	container.Container
}

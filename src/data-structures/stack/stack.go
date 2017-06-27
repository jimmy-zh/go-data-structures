package stack

import (
	"data-structures/container"
)

type Stacker interface {
	Push(v interface{})
	Pop() interface{}
	Peek() interface{}
	container.Container
}

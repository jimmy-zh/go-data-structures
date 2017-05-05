package stack

import (
	"../container"
)

type Stacker interface {
	Push(v interface{})
	Pop() interface{}
	Peek() interface{}
	container.Container
}

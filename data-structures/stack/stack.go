package stack

import (
	"github.com/midnight-vivian/go-data-structures/data-structures/container"
)

type Stacker interface {
	Push(v interface{})
	Pop() interface{}
	Peek() interface{}
	container.Container
}

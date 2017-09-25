package container

type Container interface {
	Size() int   //number of items in the container.
	Empty() bool //returns true if this container is empty.
}

package symbol_table

import (
	"algorithms/searching/binary"
	"data-structures/container"
)

type SymbolTabler interface {
	container.Container
	Get(key binary.Comparabler) interface{}
	Put(key binary.Comparabler, value interface{})
	Delete(key binary.Comparabler)
}

type SymbolTableSuper interface {
	SymbolTabler
	Max() binary.Comparabler
	Min() binary.Comparabler
	Contains(key binary.Comparabler) bool
	Select(k int) binary.Comparabler
	Floor(key binary.Comparabler) binary.Comparabler
	Ceiling(key binary.Comparabler) binary.Comparabler
}

package symbol_table

import (
	"algorithms/searching/binary"
	"data-structures/container"
)

type SymbolTabler interface {
	container.Container
	Get(key binary.Comparabler) interface{}        //Returns the value associated with the given key in this symbol table.
	Put(key binary.Comparabler, value interface{}) //Inserts the specified key-value pair into the symbol table, overwriting the old value with the new value if the symbol table already contains the specified key.
	Delete(key binary.Comparabler)                 //Removes the specified key and associated value from this symbol table
}

type SymbolTableSuper interface {
	SymbolTabler
	Max() binary.Comparabler                           //Returns the largest key in this symbol table.
	Min() binary.Comparabler                           //Returns the smallest key in this symbol table.
	Contains(key binary.Comparabler) bool              //Does this symbol table contain the given key?
	Select(k int) binary.Comparabler                   //Return the kth smallest key in this symbol table
	Rank(key binary.Comparabler) int                   //Returns the number of keys in this symbol table strictly less than {@code key}.
	Floor(key binary.Comparabler) binary.Comparabler   //Returns the largest key in this symbol table less than or equal to {@code key}.
	Ceiling(key binary.Comparabler) binary.Comparabler //Returns the smallest key in this symbol table greater than or equal to {@code key}.
}

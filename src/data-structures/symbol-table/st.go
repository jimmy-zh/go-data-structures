package symbol_table

import (
	"data-structures/container"
	"utils"
)

type SymbolTabler interface {
	container.Container
	Get(key utils.Comparabler) interface{}        //Returns the value associated with the given key in this symbol table.
	Put(key utils.Comparabler, value interface{}) //Inserts the specified key-value pair into the symbol table, overwriting the old value with the new value if the symbol table already contains the specified key.
	Delete(key utils.Comparabler)                 //Removes the specified key and associated value from this symbol table
}

type SymbolTableSuper interface {
	SymbolTabler
	Contains(key utils.Comparabler) bool              //Does this symbol table contain the given key?
	Max() utils.Comparabler                           //Returns the largest key in this symbol table.
	Min() utils.Comparabler                           //Returns the smallest key in this symbol table.
	DeleteMax()                                        //Removes the largest key and associated value from this symbol table
	DeleteMin()                                        //Removes the smallest key and associated value from this symbol table
	Floor(key utils.Comparabler) utils.Comparabler   //Returns the largest key in this symbol table less than or equal to {@code key}.
	Ceiling(key utils.Comparabler) utils.Comparabler //Returns the smallest key in this symbol table greater than or equal to {@code key}.
	Rank(key utils.Comparabler) int                   //Returns the number of keys in this symbol table strictly less than {@code key}.
	Select(k int) utils.Comparabler                   //Return the kth smallest key in this symbol table
}

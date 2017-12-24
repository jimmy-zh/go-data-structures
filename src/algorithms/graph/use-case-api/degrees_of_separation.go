package use_case_api

import (
	"fmt"
	dgraph "github.com/midnight-vivian/go-data-structures/src/data-structures/graph"
)

type DegreesOfSeparation struct {
	udsg *dgraph.UdSymbolGraph
	sp   *ShortestPath
}

func NewDegreesOfSeparation(udsg *dgraph.UdSymbolGraph, sVertex string) *DegreesOfSeparation {
	sp := NewShortestPath(udsg.UdGraph(), udsg.Index(sVertex))
	return &DegreesOfSeparation{udsg: udsg, sp: sp}
}

func (dos *DegreesOfSeparation) OutputPath(tVertex string) {
	sll := dos.sp.Path(dos.udsg.Index(tVertex))
	for !sll.Empty() {
		fmt.Println(dos.udsg.Name(sll.Pop().(int)))
	}
}

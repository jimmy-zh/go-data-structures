package graph

import (
	"github.com/midnight-vivian/go-data-structures/src/data-structures/graph"
	"github.com/midnight-vivian/go-data-structures/src/data-structures/stack"
)

func UndirectedDfsRecursion(udg *graph.UdGraph, vertex, pvertex int, processed func(int) bool, process func(int, int)) {
	process(vertex, pvertex)
	for _, v := range udg.Adj(vertex) {
		if !processed(v) {
			UndirectedDfsRecursion(udg, v, vertex, processed, process)
		}
	}
}

func UndirectedDfsIteration(udg *graph.UdGraph, vertex, pvertex int, processed func(int) bool, process func(int, int)) {
	process(vertex, pvertex)
	s := stack.NewStackLinkedList()
	s.Push(vertex)

	for !s.Empty() {
		v := s.Pop().(int)
		for _, sv := range udg.Adj(v) {
			if !processed(sv) {
				process(sv, v)
				s.Push(sv)
			}
		}
	}

}

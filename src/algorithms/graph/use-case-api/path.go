package use_case_api

import (
	dgraph "data-structures/graph"
	agraph "algorithms/graph"
	"data-structures/stack"
)

//Path focuses the path from the source vertex to the target vertex

type Path struct {
	mark   []bool
	record []int
	sID    int
}

func NewPath(udg *dgraph.UndirectedGraph, sID int) *Path {
	path := &Path{
		mark:   make([]bool, udg.GetV()),
		record: make([]int, udg.GetV()),
		sID:    sID,
	}
	agraph.UndirectedDfsRecursion(udg, sID, -1, path.processed, path.process)
	return path
}

func (p *Path) HasPath(v int) bool {
	return p.mark[v]
}

func (p *Path) Path(v int) (*stack.StackLinkedList) {
	s := stack.NewStackLinkedList()
	for i := v; i != p.sID; i = p.record[i] {
		s.Push(i)
	}
	s.Push(p.sID)
	return s
}

func (p *Path) processed(vertex int) bool {
	return p.mark[vertex]
}

func (p *Path) process(vertex, pvertex int) {
	p.mark[vertex] = true
	p.record[vertex] = pvertex
}

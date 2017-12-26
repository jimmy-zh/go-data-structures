package use_case_api

import (
	agraph "github.com/midnight-vivian/go-data-structures/algorithms/graph"
	dgraph "github.com/midnight-vivian/go-data-structures/data-structures/graph"
	"github.com/midnight-vivian/go-data-structures/data-structures/stack"
)

//ShortestPath focuses on the shortest path from the source vertex to the target vertex

type ShortestPath struct {
	mark   []bool
	record []int
	sID    int
}

func NewShortestPath(udg *dgraph.UdGraph, sID int) *ShortestPath {
	sp := &ShortestPath{
		mark:   make([]bool, udg.GetV()),
		record: make([]int, udg.GetV()),
		sID:    sID,
	}

	agraph.UndirectedBfsIteration(udg, sID, -1, sp.Processed, sp.Process)
	return sp
}

func (sp *ShortestPath) Processed(vertex int) bool {
	return sp.mark[vertex]
}

func (sp *ShortestPath) Process(vertex, pvertex int) {
	sp.mark[vertex] = true
	sp.record[vertex] = pvertex
}

func (sp *ShortestPath) Path(vertex int) *stack.StackLinkedList {
	s := stack.NewStackLinkedList()
	for i := vertex; i != sp.sID; i = sp.record[i] {
		s.Push(i)
	}
	s.Push(sp.sID)
	return s
}

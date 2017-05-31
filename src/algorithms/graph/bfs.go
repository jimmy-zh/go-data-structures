package graph

import (
	"data-structures/graph"
	"data-structures/queue"
)

func UndirectedBfsIteration(udg *graph.UdGraph, vertex, pvertex int, processed func(int) bool, process func(int, int)) {
	process(vertex, pvertex)
	q := queue.NewQueueLinkedList()
	q.Enqueue(vertex)
	for !q.Empty() {
		v := q.Dequeue().(int)
		for _, sv := range udg.Adj(v) {
			if !processed(sv) {
				process(sv, v)
				q.Enqueue(sv)
			}
		}
	}
}

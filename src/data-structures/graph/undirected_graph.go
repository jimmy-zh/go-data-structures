package graph

import "bag"

type UndirectedGraph struct {
	adj    []*bag.Bag
	vCount int
	eCount int
}

func NewUndirectedGraph(vCount int) *UndirectedGraph {
	if vCount <= 0 {
		return nil
	}
	adj := make([]*bag.Bag, vCount)
	for k := range adj {
		adj[k] = bag.NewBag()
	}
	return &UndirectedGraph{
		adj:    adj,
		vCount: vCount,
	}
}

func (udg *UndirectedGraph) AddEdge(from, to int) {
	if from >= udg.vCount || to >= udg.vCount {
		return
	}
	udg.adj[from].Add(to)
	udg.adj[to].Add(from)
	udg.eCount++
}

func (udg *UndirectedGraph) GetV() int {
	return udg.vCount
}

func (udg *UndirectedGraph) GetE() int {
	return udg.eCount
}

func (udg *UndirectedGraph) Adj(vID int) []int {
	adj := make([]int, len(udg.adj[vID].Items))
	for k, v := range udg.adj[vID].Items {
		adj[k] = v.(int)
	}
	return adj
}

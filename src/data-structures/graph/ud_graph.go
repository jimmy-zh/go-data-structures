package graph

import "github.com/midnight-vivian/go-data-structures/src/data-structures/bag"

type UdGraph struct {
	adj    []*bag.Bag
	vCount int
	eCount int
}

func NewUdGraph(vCount int) *UdGraph {
	if vCount <= 0 {
		return nil
	}
	adj := make([]*bag.Bag, vCount)
	for k := range adj {
		adj[k] = bag.NewBag()
	}
	return &UdGraph{
		adj:    adj,
		vCount: vCount,
	}
}

func (udg *UdGraph) AddEdge(from, to int) {
	if from >= udg.vCount || to >= udg.vCount {
		return
	}
	udg.adj[from].Add(to)
	udg.adj[to].Add(from)
	udg.eCount++
}

func (udg *UdGraph) GetV() int {
	return udg.vCount
}

func (udg *UdGraph) GetE() int {
	return udg.eCount
}

func (udg *UdGraph) Adj(vID int) []int {
	adj := make([]int, len(udg.adj[vID].Items()))
	for k, v := range udg.adj[vID].Items() {
		adj[k] = v.(int)
	}
	return adj
}

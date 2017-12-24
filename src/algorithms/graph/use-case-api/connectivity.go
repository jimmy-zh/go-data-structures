package use_case_api

import (
	dgraph "github.com/midnight-vivian/go-data-structures/src/data-structures/graph"
	agraph "github.com/midnight-vivian/go-data-structures/src/algorithms/graph"
)

//Connectivity focuses on the connectivity between the source vertex and the target vertex

type Connectivity struct {
	mark  []bool
	count int
}

func NewConnectivity(udg *dgraph.UdGraph, vertex int) *Connectivity {
	connectivity := &Connectivity{
		mark: make([]bool, udg.GetV()),
	}
	agraph.UndirectedDfsRecursion(udg, vertex, -1, connectivity.processed, connectivity.process)
	return connectivity
}

func (c *Connectivity) Connected(vertex int) bool {
	return c.mark[vertex]
}

func (c *Connectivity) Count() int {
	return c.count
}

func (c *Connectivity) processed(vertex int) bool {
	return c.mark[vertex]
}

func (c *Connectivity) process(vertex, pvertex int) {
	c.mark[vertex] = true
	c.count++
}

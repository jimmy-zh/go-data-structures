package use_case_api

import (
	dgraph "github.com/midnight-vivian/go-data-structures/src/data-structures/graph"
)

//Cycle focuses on whether the graph is acylic or not

type Cycle struct {
	mark     []bool
	hasCycle bool
}

func NewCycle(udg *dgraph.UdGraph) *Cycle {
	cycle := &Cycle{
		mark: make([]bool, udg.GetV()),
	}

	for s := 0; s < udg.GetV(); s++ {
		if !cycle.mark[s] {
			cycle.dfs(udg, s, -1)
		}
	}
	return cycle
}

func (c *Cycle) HasCycle() bool {
	return c.hasCycle
}

//algorithms/graph/UndirectedDfsRecursion is not fit,implement cycle's dfs
func (c *Cycle) dfs(udg *dgraph.UdGraph, vertex, pvertex int) {
	c.mark[vertex] = true
	for _, v := range udg.Adj(vertex) {
		if !c.mark[v] {
			c.dfs(udg, v, vertex)
		} else if v != pvertex {
			c.hasCycle = true
		}
	}
}

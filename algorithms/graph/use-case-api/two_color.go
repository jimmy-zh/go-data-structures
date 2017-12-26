package use_case_api

import "github.com/midnight-vivian/go-data-structures/data-structures/graph"

type TwoColor struct {
	mark      []bool
	color     []bool
	bipartite bool
}

//TwoColor focuses on whether the graph is bipartite(two-colorable) or not

func NewTwoColor(udg *graph.UdGraph) *TwoColor {
	tc := &TwoColor{
		mark:      make([]bool, udg.GetV()),
		color:     make([]bool, udg.GetV()),
		bipartite: true,
	}

	for s := 0; s < udg.GetV(); s++ {
		if !tc.mark[s] {
			tc.dfs(udg, s)
		}
	}

	return tc
}

func (tc *TwoColor) Bipartite() bool {
	return tc.bipartite
}

//algorithms/graph/UndirectedDfsRecursion is not fit,implement TwoColor's dfs
func (tc *TwoColor) dfs(udg *graph.UdGraph, vertex int) {
	tc.mark[vertex] = true
	for _, v := range udg.Adj(vertex) {
		if !tc.mark[v] {
			tc.color[v] = !tc.color[vertex]
			tc.dfs(udg, v)
		} else if tc.color[v] == tc.color[vertex] {
			tc.bipartite = false
		}
	}

}

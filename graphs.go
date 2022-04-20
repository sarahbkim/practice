package main

import "fmt"

func findMinHeightTrees(n int, edges [][]int) []int {
	if n <= 2 {
		var ans = make([]int, n)
		for i := 0; i < n; i++ {
			ans[i] = i
		}
		return ans
	}
	var neighbors = make(map[int][]int, n)
	for _, edge := range edges {
		start, end := edge[0], edge[1]
		neighbors[start] = append(neighbors[start], end)
		neighbors[end] = append(neighbors[end], start)
	}
	var leaves []int
	for i := 0; i < n; i++ {
		if len(neighbors[i]) == 1 {
			leaves = append(leaves, i)
		}
	}
	rem := n
	for rem > 2 {
		rem -= len(leaves)
		var newLeaves []int
		for len(leaves) > 0 {
			leaf := leaves[0]
			leaves = leaves[1:]
			// only neighbor for leaf
			neighbor := neighbors[leaf][0]
			neighbors[leaf] = nil
			// remove edge
			for i, v := range neighbors[neighbor] {
				if v == leaf {
					neighbors[neighbor] = append(neighbors[neighbor][:i], neighbors[neighbor][i+1:]...)
					break
				}
			}
			if len(neighbors[neighbor]) == 1 {
				newLeaves = append(newLeaves, neighbor)
			}
		}
		leaves = newLeaves
	}
	return leaves
}

func validTree(n int, edges [][]int) bool {
	if n == 0 {
		return false
	}
	var g = map[int][]int{}
	// 0 -> unvisited, 1 -> visiting, 2 -> processed
	var state = make([]int, n)
	var comps int
	for i := 0; i < len(edges); i++ {
		u, v := edges[i][0], edges[i][1]
		g[u] = append(g[u], v)
		g[v] = append(g[v], u)
	}
	var parent = make([]int, n)
	for i := 0; i < n; i++ {
		if state[i] == 0 {
			comps++
			hasCycle := compDFS(i, g, &state, parent)
			if hasCycle {
				return false
			}
		}
	}
	fmt.Println("comps", comps)
	return comps == 1
}

func compDFS(u int, g map[int][]int, state *[]int, parent []int) bool {
	fmt.Println("visiting ", u, g)
	// visiting
	(*state)[u] = 1
	for _, v := range g[u] {
		if (*state)[v] == 2 {
			if v == u || u != parent[v] {
				return true
			}
		}
		if (*state)[v] == 0 {
			parent[v] = u
			hasCycle := compDFS(v, g, state, parent)
			if hasCycle {
				return true
			}
		}
	}
	// processed
	(*state)[u] = 2
	return false
}

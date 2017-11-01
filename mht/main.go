package main

import "fmt"

func delete(a []int, x int) []int {
	for i := range a {
		if a[i] == x {
			return append(a[:i], a[i+1:]...)
		}
	}
	return a
}

func findMinHeightTrees(n int, edges [][]int) []int {
	if n == 1 {
		return []int{0}
	}

	//adjacency list of all edges
	adj_list := make(map[int][]int, n)
	for i := range edges {
		adj_list[edges[i][0]] = append(adj_list[edges[i][0]], edges[i][1])
		adj_list[edges[i][1]] = append(adj_list[edges[i][1]], edges[i][0])
	}

	//record every node with one edge
	var leaves []int
	for i := range adj_list {
		if len(adj_list[i]) == 1 {
			leaves = append(leaves, i)
		}
	}

	//advance the pointers
	for n > 2 {
		n -= len(leaves)
		var new_nodes []int
		for _, x := range leaves {
			tmp := adj_list[x][0]
			adj_list[x] = adj_list[x][1:]
			adj_list[tmp] = delete(adj_list[tmp], x)
			if len(adj_list[tmp]) == 1 {
				new_nodes = append(new_nodes, tmp)
			}
		}
		leaves = new_nodes
	}

	return leaves
}

func main() {
	n := 4
	edges := [][]int{{1, 0}, {1, 2}, {1, 3}}
	fmt.Println(findMinHeightTrees(n, edges))

	n = 6
	edges = [][]int{{0, 3}, {2, 3}, {1, 3}, {4, 3}, {5, 4}}
	fmt.Println(findMinHeightTrees(n, edges))
}

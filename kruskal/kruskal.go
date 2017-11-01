package kruskal

import (
	"fmt"
	"os"
)

type edge struct {
	v1     int
	v2     int
	weight int
}

//disjoint set operations (useful for cycle finding)
func FindRoot(parents []int, v int) int {
	if parents[v] == 0 {
		return v
	} else {
		return FindRoot(parents, parents[v])
	}
}

func Unify(parents []int, rank []int, v1, v2 int) {
	root1 := FindRoot(parents, v1)
	root2 := FindRoot(parents, v2)

	if rank[root1] < rank[root2] {
		parents[root1] = root2
	} else {
		parents[root2] = root1
	}

	if rank[root1] == rank[root2] {
		rank[root1] += 1
	}
}

// quicksort implementation for edges!
func partition(A []edge, low, high int) int {
	left := low - 1

	for i := low; i < high; i++ {
		if A[i].weight < A[high].weight {
			left++
			A[i], A[left] = A[left], A[i]
		}
	}

	left++
	A[left], A[high] = A[high], A[left]
	return left
}

func quicksort(A []edge, low, high int) {
	if low < high {
		p := partition(A, low, high)

		quicksort(A, low, p-1)
		quicksort(A, p+1, high)
	}
}

func sort(A []edge) {
	quicksort(A, 0, len(A)-1)
}

func Kruskal(testnum int) int {
	var n, e, v1, v2, weight int

	test_files := []string{"test1.txt", "test2.txt"}

	io, ok := os.Open(test_files[testnum])
	if ok != nil {
		fmt.Println("Failed to open file")
		os.Exit(1)
	}
	fmt.Fscan(io, &n)
	fmt.Fscan(io, &e)

	n += 1 //graph begins indexing at 1
	edges := make([]edge, e)
	parents := make([]int, n)
	rank := make([]int, n)

	//fill edges
	for i := range edges {
		fmt.Fscan(io, &v1)
		fmt.Fscan(io, &v2)
		fmt.Fscan(io, &weight)
		edges[i] = edge{v1, v2, weight}
	}

	sort(edges)

	total := 0
	for _, edge := range edges {
		if FindRoot(parents, edge.v1) != FindRoot(parents, edge.v2) {
			total += edge.weight
			Unify(parents, rank, edge.v1, edge.v2)
		}
	}
	return total
}

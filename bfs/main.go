package main

import (
	"fmt"
	"os"
)

func main() {
	//Enter your code here. Read input from STDIN. Print output to STDOUT
	/* The first line contains an integer, , denoting the number of queries. The subsequent lines describe each query in the following format:

	The first line contains two space-separated integers describing the respective values of  (the number of nodes) and  (the number of edges) in the graph.
	Each line  of the  subsequent lines contains two space-separated integers,  and , describing an edge connecting node  to node .
	The last line contains a single integer, , denoting the index of the starting node. */

	//read input file from command line
	if len(os.Args) < 2 {
		fmt.Println("Must include at least one input file.")
		os.Exit(1)
	}

	args := os.Args[1:]

	var q, n, e, v1, v2, s int

	for k, arg := range args {

		fmt.Println("Result from test file ", k)
		//attempt to read input file
		io, ok := os.Open(arg); if ok != nil {
			fmt.Println("Failed to open file")
			os.Exit(1)
		}

		fmt.Fscan(io, &q)
		for i := 0; i < q; i++ {
			fmt.Fscan(io, &n)
			fmt.Fscan(io, &e)
			adj_list := make(map[int][]int, n+1)
			distance := make([]int, n+1)
			for r := range distance {
				distance[r] = -1
			}

			//create adjacency list
			for j := 0; j < e; j++ {
				fmt.Fscan(io, &v1)
				fmt.Fscan(io, &v2)
				adj_list[v1] = append(adj_list[v1], v2)
				adj_list[v2] = append(adj_list[v2], v1)
			}
			fmt.Fscan(io, &s)

			//BFS traversal
			queue := make([]int, 0)
			queue = append(queue, s)

			for len(queue) > 0 {
				curr := queue[0]
				queue = queue[1:]
				distance[s] = 0
				for _, x := range adj_list[curr] {
					if distance[x] == -1 {
						distance[x] = distance[curr] + 6
						queue = append(queue, x)
					}
				}
			}

			for j := 1; j < n+1; j++ {
				if j != s {
					fmt.Print(distance[j], " ")
				}
			}
			fmt.Println("")
		}
	fmt.Println("")
	}
}

package graph

import "fmt"

func (g *Graph) Dfs(start int) {
	stack := []*Vertex{}
	visited := make(map[int]bool)

	//find the starting index and mark it as visited
	startVertex := g.getVertex(start)
	stack = append(stack, startVertex)
	visited[startVertex.key] = true

	//loop until the stack is empty
	for len(stack) > 0 {
		//pop the  top vertex from stack and print it
		current := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		fmt.Printf("%v\n", current.key)

		//iterate through all adjecent vertexes of the current vertex
		for _, val := range current.adjacent {
			if _, ok := visited[val.key]; !ok {
				visited[val.key] = true
				stack = append(stack, val)
			}
		}
	}
}

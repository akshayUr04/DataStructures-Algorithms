package graph

import "fmt"

func (g *Graph) Bfs(start int) {
	visited := make(map[int]bool)
	for _, v := range g.vertices {
		visited[v.key] = false
	}

	queue := make([]*Vertex, 0)

	startVertex := g.getVertex(start)
	visited[start] = true
	queue = append(queue, startVertex)

	for len(queue) != 0 {
		currentVertex := queue[0]
		fmt.Printf("%v\n", currentVertex.key)
		queue = queue[1:]

		for _, adj := range currentVertex.adjacent {
			if !visited[adj.key] {
				visited[adj.key] = true
				queue = append(queue, adj)
			}
		}
	}
}

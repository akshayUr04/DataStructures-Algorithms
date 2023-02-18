package graph

import "fmt"

type Graph struct {
	vertices []*Vertex
}

type Vertex struct {
	key      int
	adjacent []*Vertex
}

func (g *Graph) AddVertex(k int) {
	if contains(g.vertices, k) {
		fmt.Println("Vertices alredy exixts")
	} else {
		g.vertices = append(g.vertices, &Vertex{key: k})
	}

}

func (g *Graph) AddEdge(from, to int) {
	fromVertex := g.getVertex(from)
	toVertex := g.getVertex(to)

	//check errors
	if fromVertex == nil || toVertex == nil {
		fmt.Printf("in valid edge %v to %v \n", from, to)
	} else if contains(fromVertex.adjacent, to) {
		fmt.Printf("existing edge %v to %v \n", from, to)
	} else {
		// add edge
		fromVertex.adjacent = append(fromVertex.adjacent, toVertex)
	}
}

func (g *Graph) getVertex(k int) *Vertex {
	for i, v := range g.vertices {
		if v.key == k {
			return g.vertices[i]
		}
	}
	return nil
}

func (g *Graph) Print() {
	for _, v := range g.vertices {
		fmt.Printf("\nVertices %v : ", v.key)
		for _, v := range v.adjacent {
			fmt.Printf(" %v ", v.key)
		}
	}
	fmt.Println()
}

func contains(s []*Vertex, k int) bool {
	for _, v := range s {
		if v.key == k {
			return true
		}
	}
	return false
}

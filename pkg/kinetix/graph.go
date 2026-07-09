package kinetix

import (
	"fmt"
)

type Node struct {

	Name string
	X int
	Y int
	Edges []*Node

}

type Graph struct {

	Nodes map[string]*Node
	coordsTracker map[string]bool

}

func NewGraph() *Graph{
	return &Graph{
		Nodes: make(map[string]bool)
		coordsTracker: make(map[string]bool)
	}
}

func (g *Graph) AddNode(name string, x, y int) error {

	return nil
}

func (g *Graph) AddEdge(name1, name2 string) error {

	return nil
}
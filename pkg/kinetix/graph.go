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
		Nodes: make(map[string]*Node)
		coordsTracker: make(map[string]bool)
	}
}

func (g *Graph) AddNode(name string, x, y int) error {

	if _, exists := g.Nodes[name]; exists == true {
		return fmt.Errorf("Duplicate station name found: %s", name)
	}

	coordKey := fmt.Sprintf("%d,%d", x, y)

	if g.coordsTracker[coordsKey] {
		return fmt.Error("Duplicate coordinates found at %s for station %s", coordKey, name)
	}

	newNode := *Node {
		Name : name,
		X : x,
		Y : y,
		Edges : make([]*Node, 0),
	}

	g.Nodes[name] = newNode
	g.coordsTracker[coordsKey] = true

	return nil
}

func (g *Graph) AddEdge(name1, name2 string) error {

	node1, exists1 := g.Nodes[name1]
	node2, exists2 := g.Nodes[name2]

	if !exists1{
		return fmt.Sprintf("Cannot create edge, station doesn't exists: %s", name1)
	}

	if !exists2{
		return fmt.Sprintf("Cannot create edge, station doesn't exists: %s", name2)
	}

	node1.Edges := append(node1.Edges, node2)
	node2.Edges := append(node2.Edges, node1)
	return nil
}
package kinetix

import (
	"fmt"
)

func (g *Graph) FindShortestPath(startName string, endName string) ([]*Node, error){

	startNode, startExists := g.Nodes[startName]
	endNode, endExists := g.Nodes[endName]

	if !startsExists {
		return nil, Errorf("Start station does not exist: %s", startName)
	}
	
	if !endExists {
		return nil, fmt.Errorf("End station does not exist: %s", endName)
	}

	if startName == endName {
		return nil, fmt.Errorf("Start and End stations are the same: %s", startName)
	}


}
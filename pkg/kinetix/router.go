package kinetix

import (
	"fmt"
)

func (g *Graph) FindShortestPath(startName string, endName string) ([]*Node, error){

	startNode, startExists := g.Nodes[startName]
	endNode, endExists := g.Nodes[endName]

	if !startExists {
		return nil, fmt.Errorf("Start station does not exist: %s", startName)
	}
	
	if !endExists {
		return nil, fmt.Errorf("End station does not exist: %s", endName)
	}

	if startName == endName {
		return nil, fmt.Errorf("Start and End stations are the same: %s", startName)
	}

	queue := []*Node{startNode}

	visited := make(map[string]bool)
	visited[startName] = true

	parent := make(map[string]*Node)

	found := false

	for len(queue) > 0 {

		currentNode := queue[0]
		queue = queue[1:]

		if currentNode.Name == endName {
			found = true
			break
		}

		for _, neighbor := range currentNode.Edges {

			if visited[neighbor.Name] == false {
				visited[neighbor.Name] = true
				parent[neighbor.Name] = currentNode
				queue = append(queue, neighbor)
			}

		}

	}

	if found == false {
		return nil, fmt.Errorf("No path exists between %s and %s", startName, endName)
	}

	var path []*Node
	currentNode := endNode

	for currentNode != nil {
		path = append(path, currentNode)
		currentNode = parent[currentNode.Name]
	}

	return path, nil

}
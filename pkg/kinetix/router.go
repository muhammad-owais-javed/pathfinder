package kinetix

import (
	"fmt"
)


func (g *Graph) FindOptimalPaths(startName, endName string, numTrains int) ([][]*Node, error) {
	startNode, ok1 := g.Nodes[startName]
	endNode, ok2 := g.Nodes[endName]
	if !ok1 || !ok2 {
		return nil, fmt.Errorf("start or end station does not exist")
	}


	flow := make(map[string]map[string]int)
	for name := range g.Nodes {
		flow[name] = make(map[string]int)
	}

	var bestPaths [][]*Node
	bestCost := -1


	for {
		parent := g.residualBFS(startNode, endNode, flow)
		if parent == nil {
			break 
		}

		curr := endNode
		for curr.Name != startName {
			prev := parent[curr.Name]
			
			flow[prev.Name][curr.Name] += 1
			flow[curr.Name][prev.Name] -= 1
			
			curr = prev
		}

		currentPaths := g.extractPaths(startNode, endNode, flow)

		cost := calculateTotalTurns(currentPaths, numTrains)

		if bestCost == -1 || cost < bestCost {
			bestCost = cost
			bestPaths = currentPaths
		} else {

			break
		}
	}

	if len(bestPaths) == 0 {
		return nil, fmt.Errorf("no valid paths found between %s and %s", startName, endName)
	}

	return bestPaths, nil
}

// residualBFS finds a path while keeping the "Undo" rules of Network Flow.
func (g *Graph) residualBFS(start, end *Node, flow map[string]map[string]int) map[string]*Node {
	queue := []*Node{start}
	parent := make(map[string]*Node)
	visited := make(map[string]bool)
	visited[start.Name] = true

	for len(queue) > 0 {
		curr := queue[0]
		queue = queue[1:]

		if curr.Name == end.Name {
			return parent
		}

		for _, neighbor := range curr.Edges {
			// We can travel to a neighbor IF:
			// 1. We haven't visited it in this BFS run AND
			// 2. The flow from curr -> neighbor is not already 1 (meaning it's not blocked forwards)
			if !visited[neighbor.Name] && flow[curr.Name][neighbor.Name] < 1 {
				
				// VERTEX CAPACITY RULE: 
				// If the neighbor is already part of another path (it has flow going out of it),
				// we are ONLY allowed to visit it if we are walking backwards (undoing) its flow.
				isNodeUsed := false
				for _, outFlow := range flow[neighbor.Name] {
					if outFlow == 1 && neighbor.Name != end.Name && neighbor.Name != start.Name {
						isNodeUsed = true
						break
					}
				}

				if isNodeUsed && flow[curr.Name][neighbor.Name] != -1 {
					continue
				}

				visited[neighbor.Name] = true
				parent[neighbor.Name] = curr
				queue = append(queue, neighbor)
			}
		}
	}
	return nil
}

func (g *Graph) extractPaths(start, end *Node, flow map[string]map[string]int) [][]*Node {
	var paths [][]*Node

	for _, neighbor := range start.Edges {
		if flow[start.Name][neighbor.Name] == 1 {
			path := []*Node{start}
			curr := neighbor

			for curr.Name != end.Name {
				path = append(path, curr)
				for _, nextNode := range curr.Edges {
					if flow[curr.Name][nextNode.Name] == 1 {
						curr = nextNode
						break
					}
				}
			}
			path = append(path, end)
			paths = append(paths, path)
		}
	}
	return paths
}

func calculateTotalTurns(paths [][]*Node, numTrains int) int {
	if len(paths) == 0 {
		return 999999 
	}

	pathCounts := make([]int, len(paths))
	for i := 0; i < numTrains; i++ {
		bestIdx := 0
		bestCost := -1
		for j, p := range paths {
			cost := (len(p) - 1) + pathCounts[j]
			if bestCost == -1 || cost < bestCost {
				bestCost = cost
				bestIdx = j
			}
		}
		pathCounts[bestIdx]++
	}

	maxTurns := 0
	for j, p := range paths {
		if pathCounts[j] > 0 {
			turns := (len(p) - 1) + pathCounts[j] - 1
			if turns > maxTurns {
				maxTurns = turns
			}
		}
	}
	return maxTurns
}

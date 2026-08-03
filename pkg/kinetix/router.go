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

func (g *Graph) residualBFS(start, end *Node, flow map[string]map[string]int) map[string]*Node {
	queue := []*Node{start}
	parent := make(map[string]*Node)
	visited := make(map[string]bool)
	visited[start.Name] = true

	// Helper function to check if a station is already being used by another path.
	// If it is, it returns the name of the station that sent the train there.
	isNodeUsed := func(nodeName string) (bool, string) {
		if nodeName == start.Name || nodeName == end.Name {
			return false, "" // Start and End can hold infinite trains
		}
		for prevName, edges := range flow {
			if edges[nodeName] == 1 {
				return true, prevName
			}
		}
		return false, ""
	}

	for len(queue) > 0 {
		curr := queue[0]
		queue = queue[1:]

		if curr.Name == end.Name {
			return parent
		}

		for _, neighbor := range curr.Edges {
			// Moving backward to the existing stop
			if flow[neighbor.Name][curr.Name] == 1 {
				if !visited[neighbor.Name] {
					visited[neighbor.Name] = true
					parent[neighbor.Name] = curr
					queue = append(queue, neighbor)
				}
				continue
			}

			// Movinf Forward
			if flow[curr.Name][neighbor.Name] == 0 {
				used, prevNode := isNodeUsed(neighbor.Name)
				
				if !used {
					// Normal BFS step
					if !visited[neighbor.Name] {
						visited[neighbor.Name] = true
						parent[neighbor.Name] = curr
						queue = append(queue, neighbor)
					}
				} else {
					// We are allowed to step on the stations, BUT we must immediately go backwards 
					// down its incoming track to force an "Undo"
					if !visited[neighbor.Name] && !visited[prevNode] {
						visited[neighbor.Name] = true
						parent[neighbor.Name] = curr
						
						visited[prevNode] = true
						parent[prevNode] = neighbor
						
						queue = append(queue, g.Nodes[prevNode])
					}
				}
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

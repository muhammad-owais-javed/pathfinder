package kinetix

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func ParseNetworkMap(filePath string) (*Graph, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, fmt.Errorf("Could not open file: %v", err)
	}
	defer file.Close()

	graph := NewGraph()
	scanner := bufio.NewScanner(file)

	currentSection := ""

	for scanner.Scan() {
		line := scanner.Text()

		line = cleanLine(line)
		if line == "" {
			continue
		}

		if line == "stations:" {
			currentSection = "stations"
			continue
		} else if line == "connections:" {
			currentSection = "connections"
			continue
		}

		if currentSection == "stations" {
			err := parseStationLine(graph, line)
			if err != nil {
				return nil, err
			}
		} else if currentSection == "connections" {
			err := parseConnectionLine(graph, line)
			if err != nil {
				return nil, err
			}
		}

		err := scanner.Err()

		if err != nil {
			return nil, fmt.Errorf("Error reading file: %v", err)
		}
	}

	return graph, nil
}

func cleanLine(line string) string {
	idx := strings.Index(line, "#")
	if idx != -1 {
		line = line[:idx]
	}
	return strings.TrimSpace(line)
}

func parseStationLine(graph *Graph, line string) error {
	parts := strings.Split(line, ",")
	if len(parts) != 3 {
		return fmt.Errorf("Invalid station format (expected name,x,y): %s", line)
	}

	name := strings.TrimSpace(parts[0])
	xStr := strings.TrimSpace(parts[1])
	yStr := strings.TrimSpace(parts[2])

	x, errX := strconv.Atoi(xStr)
	y, errY := strconv.Atoi(yStr)

	if errX != nil || errY != nil || x < 0 || y < 0 {
		return fmt.Errorf("Invalid coordinates for station %s: must be positive integers", name)
	}
	err := graph.AddNode(name, x, y)
	if err != nil {
		return err
	}
	return nil
}

func parseConnectionLine(graph *Graph, line string) error {
	parts := strings.Split(line, "-")
	if len(parts) != 2 {
		return fmt.Errorf("Invalid connection format (expected station1-station2): %s", line)
	}

	name1 := strings.TrimSpace(parts[0])
	name2 := strings.TrimSpace(parts[1])

	err := graph.AddEdge(name1, name2)
	if err != nil {
		return err
	}
	return nil
}

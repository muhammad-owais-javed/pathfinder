package main

import (
	"fmt"
	"log"
	"pathfinder/pkg/kinetix"
)

func main() {
	fmt.Println("** Main Function***")

	// graph := kinetix.NewGraph()
	// fmt.Printf("Graph created! Number of stations %d\n", len(graph.Nodes))

	// err := graph.AddNode("waterloo", 3, 1)
	// if err != nil {
	// 	fmt.Println("Error:", err)
	// } else {
	// 	fmt.Println("Station successfully added: waterloo at (3,1)")
	// }

	// graph.AddNode("victoria", 6, 7)
	// fmt.Println("Station successfully added: victoria at (6,7)")

	// err = graph.AddNode("fake_station", 3, 1)
	// if err != nil {
	// 	fmt.Println("EXPECTED ERROR CAUGHT:", err)
	// }

	// err = graph.AddEdge("waterloo", "victoria")
	// if err != nil {
	// 	fmt.Println("Error:", err)
	// } else {
	// 	fmt.Println("Successfully connected waterloo and victoria!")
	// }

	// fmt.Printf("Graph Updated! Number of stations %d\n", len(graph.Nodes))

	// waterlooNode := graph.Nodes["waterloo"]

	// fmt.Printf("I am standing at: %s\n", waterlooNode.Name)
	// fmt.Printf("I have %d connection(s).\n", len(waterlooNode.Edges))

	// for _, neighbor := range waterlooNode.Edges {
	// 	fmt.Printf("-> I can travel to: %s (located at %d,%d)\n", neighbor.Name, neighbor.X, neighbor.Y)
	// }

	graph, err := kinetix.ParseNetworkMap("data/network.map")
	if err != nil {
		log.Fatalf("Failed to parse network map: %v\n", err)
	}
	fmt.Println("Successfully parsed network.map!")
	fmt.Printf("Total Stations Loaded: %d\n\n", len(graph.Nodes))

	fmt.Println("--- Network Connections ---")
	for name, node := range graph.Nodes {
		fmt.Printf("Station [%s] at (%d,%d) connects to:\n", name, node.X, node.Y)
		for _, edge := range node.Edges {
			fmt.Printf("  -> %s\n", edge.Name)
		}
		fmt.Println()
	}

}

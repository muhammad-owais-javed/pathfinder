package main

import (
	"fmt"
	"log"
	"pathfinder/pkg/kinetix"
	"os"
	"strconv"
)

func main() {
	fmt.Println("** Main Function***")

	if len(os.Args) != 5 {
		fmt.Fprintf(os.Stderr, "Error: Incorrect number of command line arguments.\n")
		fmt.Fprintf(os.Stderr, "Usage: go run . [path_to_file] [start_station] [end_station] [number_of_trains]\n")
		os.Exit(1)
	}

	filePath := os.Args[1]
	startStation := os.Args[2]
	endStation := os.Args[3]
	numTrainsStr := os.Args[4]

	numTrains, err := strconv.Atoi(numTrainsStr)
	if err != nil || numTrains <= 0 {
		fmt.Fprintf(os.Stderr, "Error: Number of trains must be a valid positive integer.\n")
		os.Exit(1)
	}

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

	graph, err := kinetix.ParseNetworkMap(filePath)
	//fmt.Println(graph)
	if err != nil {
		log.Fatalf("Failed to parse network map: %v\n", err)
	}
	fmt.Println("Successfully parsed network.map!")
	fmt.Printf("Total Stations Loaded: %d\n\n", len(graph.Nodes))

	// fmt.Println("--- Network Connections ---")
	// for name, node := range graph.Nodes {
	// 	fmt.Printf("Station [%s] at (%d,%d) connects to:\n", name, node.X, node.Y)
	// 	for _, edge := range node.Edges {
	// 		fmt.Printf("  -> %s\n", edge.Name)
	// 	}
	// 	fmt.Println()
	// }

	fmt.Printf("\n--- Calculating Route: %s -> %s ---\n", startStation, endStation)

	paths, err := graph.FindOptimalPaths(startStation, endStation, numTrains)
	//fmt.Printf("%s", path)
	if err != nil {
		log.Fatalf("Routing Error: %v\n", err)
	}

	fmt.Printf("Success! Found %d independent paths:\n", len(paths))
	for pathIdx, path := range paths {

		fmt.Printf("Path %d: ", pathIdx+1)

		for nodeIdx, node := range path {
			if nodeIdx > 0 {
				fmt.Print(" -> ")
			}
			fmt.Print(node.Name)
		}
		fmt.Println() // New line for the next path
	}
	kinetix.Dispatch(paths, numTrains)

}

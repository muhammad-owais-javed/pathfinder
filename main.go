package main

import (
	"fmt"
	"pathfinder/pkg/kinetix"
)

func main(){
	fmt.Println("** Main Function***")

	graph := kinetix.NewGraph()
	fmt.Printf("Graph created! Number of stations %d\n", len(graph.Nodes))

	err := graph.AddNode("waterloo", 3, 1)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Station successfully added: waterloo at (3,1)")
	}

	graph.AddNode("victoria", 6, 7)
	fmt.Println("Station successfully added: victoria at (6,7)")

	err = graph.AddNode("fake_station", 3, 1)
	if err != nil {
		fmt.Println("EXPECTED ERROR CAUGHT:", err)
	}

	err = graph.AddEdge("waterloo", "victoria")
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Successfully connected waterloo and victoria!")
	}

	fmt.Printf("Graph Updated! Number of stations %d\n", len(graph.Nodes))

	waterlooNode := graph.Nodes["waterloo"]

	fmt.Printf("I am standing at: %s\n", waterlooNode.Name)
	fmt.Printf("I have %d connection(s).\n", len(waterlooNode.Edges))

	for _, neighbor := range waterlooNode.Edges {
		fmt.Printf("-> I can travel to: %s (located at %d,%d)\n", neighbor.Name, neighbor.X, neighbor.Y)
	}

}
package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"github.com/muhammad-owais-javed/pathfinder/pkg/kinetix" 
)
type RouteRequest struct {
	Nodes []struct {
		Name string `json:"name"`
		X    int    `json:"x"`
		Y    int    `json:"y"`
	} `json:"nodes"`
	Edges []struct {
		From string `json:"from"`
		To   string `json:"to"`
	} `json:"edges"`
	Start     string `json:"start"`
	End       string `json:"end"`
	NumAgents int    `json:"num_agents"`
}

type RouteResponse struct {
	Turns [][]string `json:"turns"`
	Error string     `json:"error,omitempty"`
}

func main() {
	http.HandleFunc("/api/route", handleRoute )

	fmt.Println("Kinetix Fleet API is running on http://localhost:8080" )
	fmt.Println("Waiting for routing requests...")
	
	log.Fatal(http.ListenAndServe(":8080", nil ))
}

func handleRoute(w http.ResponseWriter, r *http.Request ) {
	if r.Method != http.MethodPost {
		http.Error(w, "Only POST method is allowed", http.StatusMethodNotAllowed )
		return
	}

	var req RouteRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		sendJSONResponse(w, RouteResponse{Error: "Invalid JSON payload"}, http.StatusBadRequest )
		return
	}

	graph := kinetix.NewGraph()
	for _, node := range req.Nodes {
		if err := graph.AddNode(node.Name, node.X, node.Y); err != nil {
			sendJSONResponse(w, RouteResponse{Error: err.Error()}, http.StatusBadRequest )
			return
		}
	}
	for _, edge := range req.Edges {
		if err := graph.AddEdge(edge.From, edge.To); err != nil {
			sendJSONResponse(w, RouteResponse{Error: err.Error()}, http.StatusBadRequest )
			return
		}
	}

	paths, err := graph.FindOptimalPaths(req.Start, req.End, req.NumAgents)
	if err != nil {
		sendJSONResponse(w, RouteResponse{Error: err.Error()}, http.StatusBadRequest )
		return
	}

	turns := kinetix.Dispatch(paths, req.NumAgents)

	sendJSONResponse(w, RouteResponse{Turns: turns}, http.StatusOK )
}

func sendJSONResponse(w http.ResponseWriter, resp RouteResponse, statusCode int ) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(resp)
}

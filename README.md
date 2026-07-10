# Pathfinder (Powered by the Kinetix Engine)

Pathfinder is a high-performance Multi-Agent Path Finding (MAPF) solver written in Go. 

While this project was initially developed to route passenger trains through a fixed-block railway network, the core mathematics have been abstracted into a standalone, agnostic routing engine called **Kinetix**. 

Because Kinetix calculates vertex-disjoint paths for generic "agents" and "nodes", this exact same repository can be used to route autonomous delivery robots, coordinate drone swarms, or manage telecommunication network packets without collisions.

## Architecture

To ensure modularity and real-world applicability, the project is divided into three parts:

1. **`pkg/kinetix` (The Core Engine):** A pure graph-theory library. It uses Breadth-First Search (BFS) combined with augmenting path algorithms to find multiple non-overlapping paths and calculate optimal turn-by-turn movements.
2. **`main.go` (The CLI):** A command-line wrapper built specifically for the train dispatcher objective. It translates "stations" into Kinetix nodes.
3. **`cmd/api-server` (The Bonus API):** A REST API demonstrating how modern robotics companies can use Kinetix to route autonomous fleets via JSON payloads.

---

## Part 1: The Train Dispatcher CLI (Assignment Usage)

The root of this repository contains the CLI tool required for the train routing assignment. It parses a network map, calculates the optimal paths, and outputs the turn-by-turn movements.

### Prerequisites
* Go 1.20 or higher

### Usage
Run the program from the root directory using the following syntax:
```bash
go run . [path_to_map] [start_station] [end_station] [number_of_trains]
```

### Example
```sh
go run . data/network.map waterloo st_pancras 4
```
### Expected Output:
```text
T1-victoria T2-euston
T1-st_pancras T2-st_pancras T3-victoria T4-euston
T3-st_pancras T4-st_pancras
```

## Error Handling
The CLI is built to exit gracefully and print to stderr under the following conditions:
- Invalid number of arguments
- Non-existent start or end stations
- No valid path between stations
- Invalid data formats, duplicate coordinates, or infinite loops in the map file


## Part 2: The Kinetix Engine (Developer Library)
The kinetix package is completely decoupled from the concept of "trains". Developers can import it to solve any vertex-disjoint routing problem.
```Go
import "github.com/muhammad-owais-javed/pathfinder/pkg/kinetix"

// Example: Routing 50 delivery robots through a city grid
graph := kinetix.NewGraph()
graph.AddNode("Warehouse_A")
graph.AddNode("Dropoff_B")
graph.AddEdge("Warehouse_A", "Intersection_1")

engine := kinetix.NewEngine(graph)
movements, err := engine.CalculateOptimalFlow("Warehouse_A", "Dropoff_B", 50)
```


## Part 3: The Fleet API (Bonus Functionality)
To demonstrate the commercial viability of the Kinetix engine, this repository includes a bonus REST API. This API allows external clients (like a drone fleet manager or a game engine) to request routing instructions via HTTP.

**Starting the API**
```Bash
go run ./cmd/api-server/main.go
```
The server will start on http://localhost:8080.

_(Note: API documentation and JSON payload structures can be found in the cmd/api-server directory )._

---
## Algorithmic Approach
To achieve the shortest number of turns, Pathfinder does not simply rely on a single shortest path.
- Graph Construction: Parses the input into an unweighted, undirected graph.
- Path Discovery: Utilizes BFS to find the shortest path, followed by an augmenting path approach (inspired by Edmonds-Karp/Suurballe's) to discover alternative, disjoint paths.
- Flow Distribution: Mathematically distributes the agents (trains) across the discovered paths based on path length to minimize the total number of turns.

## Author
Muhammad Owais Javed


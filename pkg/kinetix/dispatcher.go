package kinetix

import "fmt"

type Train struct {
	ID        int
	Path      []*Node
	PathIndex int
	Position  int
	Finished  bool
}

func Dispatch(paths [][]*Node, numTrains int) {
	pathQueues := make([][]*Train, len(paths))
	var allTrains []*Train

	for i := 1; i <= numTrains; i++ {
		bestPathIdx := 0
		bestCost := -1

		for j, path := range paths {
			cost := (len(path) - 1) + len(pathQueues[j])

			if bestCost == -1 || cost < bestCost {
				bestCost = cost
				bestPathIdx = j
			}
		}

		newTrain := &Train{
			ID:        i,
			Path:      paths[bestPathIdx],
			PathIndex: bestPathIdx,
			Position:  0,
			Finished:  false,
		}

		pathQueues[bestPathIdx] = append(pathQueues[bestPathIdx], newTrain)
		allTrains = append(allTrains, newTrain)
	}

	for {
		moveThisTrain := false
		var turnOutput []string

		pathDispatchTurn := make(map[int]bool)

		for _, t := range allTrains {
			if t.Finished == true {
				continue
			}

			if t.Position > 0 {

				t.Position++

				turnOutput = append(turnOutput, fmt.Sprintf("T%d-%s", t.ID, t.Path[t.Position].Name))

				if t.Position == len(t.Path)-1 {
					t.Finished = true
				}

				moveThisTrain = true
			} else {
				if pathDispatchedThisTurn[t.PathIndex]==false && pathQueues[t.PathIndex][0].ID == t.ID {

					pathQueues[t.PathIndex] = pathQueues[t.PathIndex][1:]

					t.Position = 1
					turnOutput = append(turnOutput, fmt.Sprintf("T%d-%s", t.ID, t.Path[t.Position].Name))

					if t.Position == len(t.Path)-1 {
						t.Finished = true
					}

					pathDispatchedThisTurn[t.PathIndex] = true
					movedThisTurn = true

			}
		}

		if movedThisTurn == false {
			break
		}

		fmt.Println(strings.Join(turnOutput, " "))

	}
}

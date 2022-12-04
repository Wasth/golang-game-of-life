package engine

import (
	"fmt"
)

type Engine struct {
	lifeMap [][]uint8
}

func (e *Engine) InitEmpty(lenByX uint8, lenByY uint8) {
	// creates new empty map

	e.lifeMap = make([][]uint8, lenByY, lenByY)

	for i := range e.lifeMap {

		e.lifeMap[i] = make([]uint8, lenByX, lenByX)

		for j := range e.lifeMap[i] {
			e.lifeMap[i][j] = DEAD
		}

	}

}

func (e *Engine) GetMap() [][]uint8 {
	// returns map
	return e.lifeMap
}

func (e *Engine) SetCellState(x uint8, y uint8, state uint8) error {
	// changes state of cell
	// if cell does not exist or state is invalid, error returns
	if !IsValidState(state) {
		return fmt.Errorf("unknown state was given, available states - %s", STATES)
	}

	if y > uint8(len(e.lifeMap))-1 {
		return fmt.Errorf("y position is out of range")
	}
	if x > uint8(len(e.lifeMap[y]))-1 {
		return fmt.Errorf("x position is out of range")
	}

	e.lifeMap[y][x] = state

	return nil
}

func (e *Engine) PerformGeneration() bool {
	// returns if changes was performed

	var toKill []*uint8
	var toResurrect []*uint8

	for y, row := range e.lifeMap {
		for x, cell := range row { // for each cell

			// get all neighbours values
			neighbours := e.getNeighboursStates(uint8(x), uint8(y))
			aliveCellCount := 0

			// count how many are alive
			for _, neighbourCell := range neighbours {
				if neighbourCell == ALIVE {
					aliveCellCount++
				}
			}

			var addTo *[]*uint8

			// decide to make current cell DEAD or ALIVE
			if cell == DEAD && aliveCellCount == 3 {
				addTo = &toResurrect
			}
			if cell == ALIVE && (aliveCellCount < 2 || aliveCellCount > 3) {
				addTo = &toKill
			}

			if addTo != nil {
				*addTo = append(*addTo, &e.lifeMap[y][x])
			}
		}
	}

	for _, pointer := range toKill {
		*pointer = DEAD
	}
	for _, pointer := range toResurrect {
		*pointer = ALIVE
	}
	return len(toKill) > 0 || len(toResurrect) > 0
}

func (e *Engine) getNeighboursStates(x uint8, y uint8) []uint8 {
	var result []uint8

	/*
		neighbors are in
		y-1:x-1 | y-1:x | y-1:x+1
		y  :x-1 |       | y  :x+1
		y+1:x-1 | y+1:x | y+1:x+1

	*/

	hasLeftNeighbors := x > 0                               // means if there may be neighbour on left side
	hasRightNeighbors := x < (uint8(len(e.lifeMap[y])) - 1) // means if there may be neighbour on right side

	hasUpNeighbors := y > 0                             // means if there may be neighbour above current cell
	hasDownNeighbors := y < (uint8(len(e.lifeMap)) - 1) // means if there may be neighbour below current cell

	if hasUpNeighbors {
		result = append(result, e.lifeMap[y-1][x]) // up

		if hasLeftNeighbors {
			result = append(result, e.lifeMap[y-1][x-1]) // up left
		}
		if hasRightNeighbors {
			result = append(result, e.lifeMap[y-1][x+1]) // up right
		}
	}

	if hasLeftNeighbors {
		result = append(result, e.lifeMap[y][x-1]) // left
	}
	if hasRightNeighbors {
		result = append(result, e.lifeMap[y][x+1]) // right
	}

	if hasDownNeighbors {
		result = append(result, e.lifeMap[y+1][x]) // down

		if hasLeftNeighbors {
			result = append(result, e.lifeMap[y+1][x-1]) // down left
		}
		if hasRightNeighbors {
			result = append(result, e.lifeMap[y+1][x+1]) // down right
		}
	}

	return result
}

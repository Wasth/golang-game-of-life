package main

import (
	"conwaygol/cligui"
	"conwaygol/engine"
)

func main() {
	eng := new(engine.Engine)
	eng.InitEmpty(10, 10)

	// glider
	_ = eng.SetCellState(0, 0, engine.ALIVE)
	_ = eng.SetCellState(2, 0, engine.ALIVE)
	_ = eng.SetCellState(1, 1, engine.ALIVE)
	_ = eng.SetCellState(2, 1, engine.ALIVE)
	_ = eng.SetCellState(1, 2, engine.ALIVE)

	cligui.PrintAllGenerations(*eng, 350)
}

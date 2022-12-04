package cligui

import (
	"conwaygol/engine"
	"fmt"
	"time"
)

var RepresentMap = map[uint8]rune{
	engine.DEAD:  ' ',
	engine.ALIVE: 'X',
}

func PrintMap(eng engine.Engine) {
	for _, row := range eng.GetMap() {
		for _, cell := range row {
			fmt.Printf("%s", string(RepresentMap[cell]))
		}
		fmt.Printf("\n")
	}
	fmt.Println("-----------------------")
}

func PrintAllGenerations(eng engine.Engine, msBetween uint16) {
	for {
		fmt.Printf("\x1bc")
		PrintMap(eng)
		time.Sleep(time.Duration(msBetween) * time.Millisecond)
		if !eng.PerformGeneration() {
			break
		}
	}
}

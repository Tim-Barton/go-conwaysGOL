package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Hello World")

	grid := NewGrid(3, 3)
	grid.Set(0, 0, Grid{1})
	grid.Set(0, 1, Grid{1})
	grid.Set(0, 2, Grid{1})
	
	cont := true
	i := 0

	fmt.Println(grid.Print())
	for (cont) {
		tickStart := time.Now()
		newGrid := grid.Tick()
		tickEnd := time.Now()

		cont = !grid.Same(newGrid)

		fmt.Printf("Tick %d, took %v, while continue %t\n", i, tickEnd.Sub(tickStart), cont)
		i++
		grid = newGrid
		fmt.Println(grid.Print())

		time.Sleep(time.Second)
		
	}
}

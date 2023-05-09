package main

import (
	"fmt"
	"time"

	"github.com/Tim-Barton/go-conwaysGOL/pkg/grid"
)

func main() {

	fmt.Println("Starting LifeGrid Benchmarking")

	var structured grid.LifeGrid = grid.NewStructuredGrid(300, 300)

	structured.Randomize()
	structuredStart := time.Now()
	for i := 0; i < 10000; i++ {
		structured = structured.Tick()
	}
	structuredEnd := time.Now()

	fmt.Printf("Structured took %v\n", structuredEnd.Sub(structuredStart))

	var primitive grid.LifeGrid = grid.NewPrimitiveGrid(300, 300)

	primitive.Randomize()
	primitiveStart := time.Now()
	for i := 0; i < 10000; i++ {
		primitive = primitive.Tick()
	}
	primitiveEnd := time.Now()

	fmt.Printf("Primitive took %v\n", primitiveEnd.Sub(primitiveStart))
}

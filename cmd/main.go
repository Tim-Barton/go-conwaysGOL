package main

import (
	"fmt"
	"image/color"
	"log"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
)

type Game struct {
	alive *ebiten.Image
	dead  *ebiten.Image

	grid LifeGrid
}

func NewGame(grid LifeGrid) Game {
	alive := ebiten.NewImage(300, 300)
	alive.Fill(color.Black)
	dead := ebiten.NewImage(300, 300)
	dead.Fill(color.White)
	return Game{
		alive: alive,
		dead:  dead,
		grid:  grid,
	}
}

func (g *Game) Update() error {
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	//ebitenutil.DebugPrint(screen, "Hello, World!")
	for i := 0; i < g.grid.Rows(); i++ {
		for j := 0; j < g.grid.Cols(); j++ {
			geo := ebiten.GeoM{}
			geo.Translate(float64(300*j), float64(300*i))
			status, _ := g.grid.Get(i, j)
			if status.status == 0 {
				screen.DrawImage(g.dead, &ebiten.DrawImageOptions{GeoM: geo})
			}
		}
	}
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 900, 900
}

func main() {
	fmt.Println("Hello World")

	grid := NewGrid(3, 3)
	grid.Set(0, 0, Grid{1})
	grid.Set(0, 1, Grid{1})
	grid.Set(0, 2, Grid{1})

	cont := true
	i := 0

	fmt.Println(grid.Print())
	for cont {
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

	ebiten.SetWindowSize(900, 900)
	ebiten.SetWindowTitle("Hello, World!")
	game := NewGame(grid)
	if err := ebiten.RunGame(&game); err != nil {
		log.Fatal(err)
	}
}

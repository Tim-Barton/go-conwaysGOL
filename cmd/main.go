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

	tickCount int
}

func NewGame(grid LifeGrid) Game {
	alive := ebiten.NewImage(300, 300)
	alive.Fill(color.Black)
	dead := ebiten.NewImage(300, 300)
	dead.Fill(color.White)
	return Game{
		alive:     alive,
		dead:      dead,
		grid:      grid,
		tickCount: 1,
	}
}

func (g *Game) Update() error {
	if g.tickCount == 60 {
		g.tickCount = 0
		tickStart := time.Now()
		newGrid := g.grid.Tick()
		tickEnd := time.Now()

		fmt.Printf("Tick took %v\n", tickEnd.Sub(tickStart))
		g.grid = newGrid
		fmt.Println(g.grid.Print())
	}
	g.tickCount++
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

	fmt.Println(grid.Print())

	ebiten.SetWindowSize(900, 900)
	ebiten.SetWindowTitle("Hello, World!")
	game := NewGame(grid)
	if err := ebiten.RunGame(&game); err != nil {
		log.Fatal(err)
	}
}

package main

import "fmt"

type Grid struct {
	status int
}

type LifeGrid struct {
	grid [][]Grid

	rows int
	cols int
}

func (g Grid) Evaluate(neighbours int) int {
	if g.status == 1 && (neighbours == 2 || neighbours == 3) {
		return 1
	} else if g.status == 0 && neighbours == 3 {
		return 1
	}

	return 0
}

func (l LifeGrid) Get(row, col int) (Grid, error) {
	if row < 0 || row > l.rows-1 {
		return Grid{-1}, fmt.Errorf("Nope")
	}
	if col < 0 || col > l.cols-1 {
		return Grid{-1}, fmt.Errorf("Nope")
	}
	return l.grid[row][col], nil
}

func (l LifeGrid) Tick() LifeGrid {
	newGrid := [][]Grid{}
	for row := range l.grid {
		newColumns := []Grid{}
		for col := range l.grid[row] {
			neighbourCount := 0
			for k := row - 1; k <= row+1; k++ {
				for m := col - 1; m <= col+1; m++ {
					if !(k == row && m == col) {
						neighbour, err := l.Get(k, m)
						if err == nil {
							neighbourCount += neighbour.status
						}
					}
				}
			}
			newValue := l.grid[row][col].Evaluate(neighbourCount)
			newColumns = append(newColumns, Grid{newValue})
		}
		newGrid = append(newGrid, newColumns)
	}
	return LifeGrid{grid: newGrid, rows: l.rows, cols: l.cols}
}

func (l *LifeGrid) Set(row, column int, value Grid) {
	l.grid[row][column] = value
}

func (l LifeGrid) Print() string {
	return fmt.Sprintf("%v", l.grid)
}

func NewGrid(rows, cols int) LifeGrid {
	grid := [][]Grid{}
	for i := 0; i < rows; i++ {
		col := make([]Grid, cols)
		grid = append(grid, col)
	}
	return LifeGrid{grid: grid, rows: rows, cols: cols}
}

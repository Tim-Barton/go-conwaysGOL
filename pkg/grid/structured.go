package grid

import (
	"fmt"
	"math/rand"
)

type Grid struct {
	Status int
}

type StructuredLifeGrid struct {
	grid [][]Grid

	rows int
	cols int
}

func (g Grid) Evaluate(neighbours int) int {
	if g.Status == 1 && (neighbours == 2 || neighbours == 3) {
		return 1
	} else if g.Status == 0 && neighbours == 3 {
		return 1
	}

	return 0
}

func (l StructuredLifeGrid) Get(row, col int) (int, error) {
	if row < 0 || row > l.rows-1 {
		return -1, fmt.Errorf("Nope")
	}
	if col < 0 || col > l.cols-1 {
		return -1, fmt.Errorf("Nope")
	}
	return l.grid[row][col].Status, nil
}

func (l StructuredLifeGrid) Tick() LifeGrid {
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
							neighbourCount += neighbour
						}
					}
				}
			}
			newValue := l.grid[row][col].Evaluate(neighbourCount)
			newColumns = append(newColumns, Grid{newValue})
		}
		newGrid = append(newGrid, newColumns)
	}
	return &StructuredLifeGrid{grid: newGrid, rows: l.rows, cols: l.cols}
}

func (l StructuredLifeGrid) Same(newGrid StructuredLifeGrid) bool {
	if l.rows != newGrid.rows || l.cols != newGrid.cols {
		return false
	}
	for row := range l.grid {
		for col := range l.grid[row] {
			if l.grid[row][col] != newGrid.grid[row][col] {
				return false
			}
		}
	}
	return true
}

func (l *StructuredLifeGrid) Set(row, column, value int) {
	l.grid[row][column] = Grid{value}
}

func (l StructuredLifeGrid) Print() string {
	return fmt.Sprintf("%v", l.grid)
}

func (l StructuredLifeGrid) Rows() int {
	return l.rows
}

func (l StructuredLifeGrid) Cols() int {
	return l.cols
}

func (l *StructuredLifeGrid) Randomize() {
	for row := range l.grid {
		for col := range l.grid[row] {
			l.Set(row, col, int(rand.Int31n(2)))
		}
	}
}

func NewStructuredGrid(rows, cols int) StructuredLifeGrid {
	grid := [][]Grid{}
	for i := 0; i < rows; i++ {
		col := make([]Grid, cols)
		grid = append(grid, col)
	}
	return StructuredLifeGrid{grid: grid, rows: rows, cols: cols}
}

package grid

import (
	"fmt"
	"math/rand"
)

type PrimitiveLifeGrid struct {
	grid [][]int

	rows int
	cols int
}

func Evaluate(local, neighbours int) int {
	if local == 1 && (neighbours == 2 || neighbours == 3) {
		return 1
	} else if local == 0 && neighbours == 3 {
		return 1
	}

	return 0
}

func (l PrimitiveLifeGrid) Get(row, col int) (int, error) {
	if row < 0 || row > l.rows-1 {
		return -1, fmt.Errorf("Nope")
	}
	if col < 0 || col > l.cols-1 {
		return -1, fmt.Errorf("Nope")
	}
	return l.grid[row][col], nil
}

func (l PrimitiveLifeGrid) Tick() LifeGrid {
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
			newValue := Evaluate(l.grid[row][col], neighbourCount)
			newColumns = append(newColumns, Grid{newValue})
		}
		newGrid = append(newGrid, newColumns)
	}
	return &StructuredLifeGrid{grid: newGrid, rows: l.rows, cols: l.cols}
}

func (l PrimitiveLifeGrid) Same(newGrid *PrimitiveLifeGrid) bool {
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

func (l *PrimitiveLifeGrid) Set(row, column, value int) {
	l.grid[row][column] = value
}

func (l PrimitiveLifeGrid) Print() string {
	return fmt.Sprintf("%v", l.grid)
}

func (l PrimitiveLifeGrid) Rows() int {
	return l.rows
}

func (l PrimitiveLifeGrid) Cols() int {
	return l.cols
}

func (l *PrimitiveLifeGrid) Randomize() {
	for row := range l.grid {
		for col := range l.grid[row] {
			l.Set(row, col, int(rand.Int31n(2)))
		}
	}
}

func NewPrimitiveGrid(rows, cols int) *PrimitiveLifeGrid {
	grid := [][]int{}
	for i := 0; i < rows; i++ {
		col := make([]int, cols)
		grid = append(grid, col)
	}
	return &PrimitiveLifeGrid{grid: grid, rows: rows, cols: cols}
}

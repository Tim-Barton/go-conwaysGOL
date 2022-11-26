package main

type Grid struct {
	status bool
}

func (g Grid) Tick([]bool) bool {
	return true
}

type LifeGrid struct {
	grid [][]Grid
}

func (l LifeGrid) Tick() LifeGrid {
	newGrid := [][]Grid{}
	for i := range l.grid {
		newColumns := []Grid{}
		for j := range l.grid[i] {
			surrounds := []bool{}
			for k := i - 1; k <= i+1; k++ {
				if k >= 0 && k < len(l.grid) {
					for m := j - 1; m <= j+1; m++ {
						if m >= 0 && m < len(l.grid[i]) {
							if k != i && m != j {
								surrounds = append(surrounds, l.grid[k][m].status)
							}
						}
					}
				}
			}
			newValue := l.grid[i][j].Tick(surrounds)
			newColumns = append(newColumns, Grid{newValue})
		}
		newGrid = append(newGrid, newColumns)
	}
	return LifeGrid{newGrid}
}

func NewGrid(rows, columns int) LifeGrid {
	grid := [][]Grid{}
	for i := 0; i < rows; i++ {
		col := make([]Grid, columns)
		grid = append(grid, col)
	}
	return LifeGrid{grid}
}

package grid

import (
	"testing"
)

func TestNewGrid(t *testing.T) {
	// 3x3 grid all of all false
	grid := NewStructuredGrid(3, 3)
	for i := range grid.grid {
		for j, gridPoint := range grid.grid[i] {
			if gridPoint.Status == 1 {
				t.Errorf("Grid at %d:%d was Alive", i, j)
			}
		}
	}
}

func TestDodgyGrid(t *testing.T) {
	grid := NewStructuredGrid(3, 3)
	grid.Set(0, 0, Grid{1})
	grid.Set(0, 1, Grid{1})
	grid.Set(0, 2, Grid{1})

	newGrid := grid.Tick()

	//old grid should be unchanged
	//fmt.Println(grid.Print())
	for i := range grid.grid {
		for j, gridPoint := range grid.grid[i] {
			if i == 0 {
				if gridPoint.Status == 0 {
					t.Errorf("Grid at %d:%d was Dead", i, j)
				}
			} else {
				if gridPoint.Status == 1 {
					t.Errorf("Grid at %d:%d was Alive", i, j)
				}
			}

		}
	}
	//new grid 0,1 & 1,1 should be Alive
	//fmt.Println(newGrid.Print())
	for i := range newGrid.grid {
		for j, gridPoint := range newGrid.grid[i] {
			if (i == 0 && j == 1) || (i == 1 && j == 1) {
				if gridPoint.Status == 0 {
					t.Errorf("New Grid at %d:%d was Dead", i, j)
				}
			} else {
				if gridPoint.Status == 1 {
					t.Errorf("New Grid at %d:%d was Alive", i, j)
				}
			}

		}
	}
}

func TestSameGrid(t *testing.T) {

	grid := NewStructuredGrid(3, 3)
	grid.Set(0, 0, Grid{1})
	grid.Set(0, 1, Grid{1})
	grid.Set(0, 2, Grid{1})

	if !grid.Same(grid) {
		t.Errorf("Grids should match")
	}

	newGrid1 := NewStructuredGrid(3, 3)
	newGrid1.Set(0, 0, Grid{1})
	newGrid1.Set(0, 1, Grid{1})
	newGrid1.Set(0, 2, Grid{1})

	if !grid.Same(newGrid1) {
		t.Errorf("Grids should match")
	}

	newGrid2 := NewStructuredGrid(3, 2)
	newGrid2.Set(0, 0, Grid{1})
	newGrid2.Set(0, 1, Grid{1})

	if grid.Same(newGrid2) {
		t.Errorf("Grids shouldn't match - different size")
	}

	newGrid3 := NewStructuredGrid(3, 3)
	newGrid3.Set(0, 0, Grid{1})
	newGrid3.Set(0, 1, Grid{1})

	if grid.Same(newGrid2) {
		t.Errorf("Grids shouldn't match - same size, different values")
	}

}

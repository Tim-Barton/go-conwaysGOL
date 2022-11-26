package main

import "testing"

func TestNewGrid (t *testing.T) {
  // 3x3 grid all of all false
	grid := NewGrid(3, 3)
	for i := range grid.grid {
		for j, gridPoint := range grid.grid[i] {
			if gridPoint.status {
				t.Errorf("Grid at %d:%d was True", i, j)
			}
		}
	}
}

func TestDodgyGrid(t *testing.T) {
	// 3x3 grid all of all false
	grid := NewGrid(3, 3)
  grid.grid[0][0] = true
  grid.grid[0][2] = true
  grid.grid[0][0] = true
  grid.grid[0][0] = true
  grid.grid[0][0] = true


	newGrid := grid.Tick()

	//old grid should be unchanged (all False)
	for i := range grid.grid {
		for j, gridPoint := range grid.grid[i] {
			if gridPoint.status {
				t.Errorf("Grid at %d:%d was True", i, j)
			}
		}
	}
	//new Grid should be all True
	for i := range newGrid.grid {
		for j, gridPoint := range newGrid.grid[i] {
			if !gridPoint.status {
				t.Errorf("New Grid at %d:%d was False", i, j)
			}
		}
	}

}

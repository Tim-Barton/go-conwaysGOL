package grid

type LifeGrid interface {
	Get(row, col int) (int, error)

	Tick() LifeGrid

	//	Same(newGrid LifeGrid) bool

	Set(row, column, value int)

	Print() string

	Rows() int

	Cols() int

	Randomize()
}

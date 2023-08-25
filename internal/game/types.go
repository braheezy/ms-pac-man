package game

type Direction int

const (
	Up Direction = iota
	Down
	Left
	Right
)

type Position struct {
	x, y float64
}

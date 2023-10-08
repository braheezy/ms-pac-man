package game

import "github.com/braheezy/ms-pacman/internal/assets"

type Direction int

const (
	Up Direction = iota
	Down
	Left
	Right
)

func (d Direction) String() string {
	names := [...]string{
		"Up",
		"Down",
		"Left",
		"Right",
	}

	if d < Up || d > Right {
		return "UnknownDirection"
	}

	return names[d]
}

type PixelPos struct {
	X, Y float64
}

type WaypointPos struct {
	X, Y int
}

func (wp *WaypointPos) Center() PixelPos {
	return PixelPos{float64(wp.X*assets.TileSize + (assets.TileSize / 2)), float64(wp.Y*assets.TileSize + (assets.TileSize / 2))}
}

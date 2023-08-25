package game

import (
	"github.com/braheezy/ms-pacman/internal/assets"
	"github.com/hajimehoshi/ebiten/v2"
)

type Player struct {
	// The current sprite to show
	image *ebiten.Image
	// The current movement direction
	direction Direction
	// Absolute pixel local to start drawing the sprite
	pixelCoord Position
	// Where on the grid the play is occupying
	tileCoord Position
}

func (p *Player) getNextTileCoord() (int, int) {
	nextX, nextY := p.tileCoord.x, p.tileCoord.y
	switch p.direction {
	case Up:
		nextY -= p.moveSpeed
	case Right:
		nextX += p.moveSpeed
	case Down:
		nextY += p.moveSpeed
	case Left:
		nextX -= p.moveSpeed
	}
	return nextX, nextY
}

func (p *Player) getNextPixelCoord() (int, int) {
	nextX, nextY := p.pixelCoord.x, p.pixelCoord.y
	switch p.direction {
	case Up:
		nextY -= assets.MoveSpeed
	case Right:
		nextX += assets.MoveSpeed
	case Down:
		nextY += assets.MoveSpeed
	case Left:
		nextX -= assets.MoveSpeed
	}
	return nextX, nextY
}

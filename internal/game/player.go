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
	// How fast the player moves in pixels per frame
	moveSpeed float64
}

func (p *Player) getNextPixelCoord() (float64, float64) {
	nextX, nextY := p.pixelCoord.x, p.pixelCoord.y
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

func (p *Player) getNextTileCoord(gridSize int) (int, int) {
	gridX := p.pixelCoord.x / assets.TileSize
	gridY := p.pixelCoord.y / assets.TileSize

	offsetX := int(p.pixelCoord.x) % assets.TileSize
	offsetY := int(p.pixelCoord.y) % assets.TileSize

	coverageX := float64(offsetX) / assets.TileSize
	coverageY := float64(offsetY) / assets.TileSize

	// Determine the tile that the player is mostly covering
	var tileX, tileY float64
	if coverageX >= 0.50 {
		tileX = gridX + 1
	} else {
		tileX = gridX
	}
	if coverageY >= 0.50 {
		tileY = gridY + 1
	} else {
		tileY = gridY
	}

	// Prevent overflows
	gridSizeF := float64(gridSize)
	if tileX > gridSizeF-1 {
		tileX = gridSizeF - 1
	} else if tileX < 0 {
		tileX = 0
	}
	if tileY > gridSizeF-1 {
		tileY = gridSizeF - 1
	} else if tileY < 0 {
		tileY = 0
	}

	return int(tileX), int(tileY)
}

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
	pixelX, pixelY float64
	// Where on the grid the play is occupying
	tileX, tileY int
	// How fast the player moves in pixels per frame
	moveSpeed float64
}

func (p *Player) getNextPixelCoord() (float64, float64) {
	nextX, nextY := p.pixelX, p.pixelY
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
	if nextX < 0 {
		nextX = 0
	} else if int(nextX) > Config.Width {
		nextX = float64(Config.Width)
	}

	if nextY < 0 {
		nextY = 0
	} else if nextY > float64(Config.Height) {
		nextY = float64(Config.Height)
	}
	return nextX, nextY
}

// func (p *Player) getCurrentTileCoord(gridSize int) (int, int) {
// 	gridX := p.pixelX / assets.TileSize
// 	gridY := p.pixelY / assets.TileSize

// 	offsetX := int(p.pixelX) % assets.TileSize
// 	offsetY := int(p.pixelY) % assets.TileSize

// 	coverageX := float64(offsetX) / assets.TileSize
// 	coverageY := float64(offsetY) / assets.TileSize

// 	// Determine the tile that the player is mostly covering
// 	var tileX, tileY float64
// 	if coverageX >= 0.50 {
// 		tileX = gridX + 1
// 	} else {
// 		tileX = gridX
// 	}
// 	if coverageY >= 0.50 {
// 		tileY = gridY + 1
// 	} else {
// 		tileY = gridY
// 	}

// 	// Prevent overflows
// 	gridSizeF := float64(gridSize)
// 	if tileX > gridSizeF-1 {
// 		tileX = gridSizeF - 1
// 	} else if tileX < 0.5 {
// 		tileX = 0
// 	} else if tileX > 0.5 && tileX < 1 {
// 		tileX = 1
// 	}
// 	if tileY > gridSizeF-1 {
// 		tileY = gridSizeF - 1
// 	} else if tileY < 0.5 {
// 		tileY = 1
// 	} else if tileY > 0.5 && tileY < 1 {
// 		tileY = 1
// 	}

// 	return int(tileX), int(tileY)
// }

func (p *Player) getNextTileCoord() (int, int) {
	nextX, nextY := p.tileX, p.tileY
	switch p.direction {
	case Up:
		nextY--
	case Right:
		nextX++
	case Down:
		nextY++
	case Left:
		nextX--
	}

	return int(nextX), int(nextY)
}

func (p *Player) updateTileLocation() {
	// To ensure we pick the right tile, use a pixel in the center of tile
	centerPixelX := p.pixelX + assets.TileSize/2
	centerPixelY := p.pixelY + assets.TileSize/2

	gridX := centerPixelX / assets.TileSize
	gridY := centerPixelY / assets.TileSize

	offsetX := int(centerPixelX) % assets.TileSize
	offsetY := int(centerPixelY) % assets.TileSize

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

	p.tileX = int(tileX)
	p.tileY = int(tileY)
}

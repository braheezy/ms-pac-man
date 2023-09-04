package game

import (
	"image"

	"github.com/braheezy/ms-pacman/internal/assets"
	"github.com/hajimehoshi/ebiten/v2"
)

type Player struct {
	// The current sprite to show
	image *ebiten.Image
	// The current movement currentDirection
	currentDirection Direction
	// The input the user has requested
	requestedDirection Direction
	// Absolute pixel local to start drawing the sprite
	currentPixelPos PixelPos
	// The estimated pixel position of where the player will be
	nextPixelPos PixelPos
	// Where on the grid the play is occupying
	currentTilePos TilePos
	// The tiles the player is about to encounter
	nextTilePos []TilePos
	// How fast the player moves in pixels per frame
	moveSpeed float32
}

func (p *Player) setNextPixelCoord() {
	switch p.currentDirection {
	case Up:
		p.nextPixelPos.Y = p.currentPixelPos.Y - p.moveSpeed
	case Right:
		p.nextPixelPos.X = p.currentPixelPos.X + p.moveSpeed
	case Down:
		p.nextPixelPos.Y = p.currentPixelPos.Y + p.moveSpeed
	case Left:
		p.nextPixelPos.X = p.currentPixelPos.X - p.moveSpeed
	}
	if p.nextPixelPos.X < 0 {
		p.nextPixelPos.X = 0
	} else if p.nextPixelPos.X > float32(Config.Width) {
		p.nextPixelPos.X = float32(Config.Width)
	}

	if p.nextPixelPos.Y < 0 {
		p.nextPixelPos.Y = 0
	} else if p.nextPixelPos.Y > float32(Config.Height) {
		p.nextPixelPos.Y = float32(Config.Height)
	}
}

// func (p *Player) getNextTileDiagCoord(nextX, nextY int) (nextDiagX1 int, nextDiagY1 int, nextDiagX2 int, nextDiagY2 int) {
// 	switch p.currentDirection {
// 	case Up:
// 		fallthrough
// 	case Down:
// 		nextDiagX1 = nextX - 1
// 		nextDiagY1 = nextY
// 		nextDiagX2 = nextX + 1
// 		nextDiagY2 = nextY
// 	case Left:
// 		fallthrough
// 	case Right:
// 		nextDiagX1 = nextX
// 		nextDiagY1 = nextY - 1
// 		nextDiagX2 = nextX
// 		nextDiagY2 = nextY + 1
// 	}

// 	return nextDiagX1, nextDiagY1, nextDiagX2, nextDiagY2
// }

func (p *Player) setNextTiles(bound int) {
	var nextDiagX1, nextDiagY1, nextDiagX2, nextDiagY2 int
	nextX, nextY := p.currentTilePos.X, p.currentTilePos.Y
	switch p.requestedDirection {
	case Up:
		nextY--
		nextDiagX1 = nextX - 1
		nextDiagY1 = nextY
		nextDiagX2 = nextX + 1
		nextDiagY2 = nextY
	case Right:
		nextX++
		nextDiagX1 = nextX
		nextDiagY1 = nextY - 1
		nextDiagX2 = nextX
		nextDiagY2 = nextY + 1
	case Down:
		nextY++
		nextDiagX1 = nextX - 1
		nextDiagY1 = nextY
		nextDiagX2 = nextX + 1
		nextDiagY2 = nextY
	case Left:
		nextX--
		nextDiagX1 = nextX
		nextDiagY1 = nextY - 1
		nextDiagX2 = nextX
		nextDiagY2 = nextY + 1
	}

	if nextX < 0 {
		nextX = 0
	} else if nextX >= bound {
		nextX = bound - 1
	}
	if nextY < 0 {
		nextY = 0
	} else if nextY >= bound {
		nextY = bound - 1
	}

	p.nextTilePos = []TilePos{
		{nextX, nextY},
		{nextDiagX1, nextDiagY1},
		{nextDiagX2, nextDiagY2},
	}
}

func (p *Player) updateTileLocation() {
	// To ensure we pick the right tile, use a pixel in the center of tile
	centerPixelX := p.currentPixelPos.X + assets.TileSize/2
	centerPixelY := p.currentPixelPos.Y + assets.TileSize/2

	gridX := centerPixelX / assets.TileSize
	gridY := centerPixelY / assets.TileSize

	offsetX := int(centerPixelX) % assets.TileSize
	offsetY := int(centerPixelY) % assets.TileSize

	coverageX := float32(offsetX) / assets.TileSize
	coverageY := float32(offsetY) / assets.TileSize

	// Determine the tile that the player is mostly covering
	var tileX, tileY float32
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

	p.currentTilePos.X = int(tileX)
	p.currentTilePos.Y = int(tileY)
}

func (p *Player) Update(grid *[][]assets.Tile, levelImage *ebiten.Image) {
	if ebiten.IsKeyPressed(ebiten.KeyArrowRight) {
		p.requestedDirection = Right
	} else if ebiten.IsKeyPressed(ebiten.KeyArrowLeft) {
		p.requestedDirection = Left
	} else if ebiten.IsKeyPressed(ebiten.KeyArrowUp) {
		p.requestedDirection = Up
	} else if ebiten.IsKeyPressed(ebiten.KeyArrowDown) {
		p.requestedDirection = Down
	}

	p.setNextPixelCoord()
	p.setNextTiles(len(*grid))

	nextTilePos := p.nextTilePos[0]
	nextTile := (*grid)[nextTilePos.Y][nextTilePos.X]

	if nextTile.Type == assets.TileTypeWall {

		if p.requestedDirection != p.currentDirection {
			// They are trying to turn into a wall
			p.currentPixelPos.X, p.currentPixelPos.Y = p.nextPixelPos.X, p.nextPixelPos.Y

			p.updateTileLocation()
			return
		}

		// Perform pixel analysis to allow proper approaches to wall tiles.
		boundingBox := image.Rect(int(p.nextPixelPos.X), int(p.nextPixelPos.Y), int(p.nextPixelPos.X+assets.SpriteSize), int(p.nextPixelPos.Y+assets.SpriteSize))

		for y := boundingBox.Min.Y; y < boundingBox.Max.Y; y++ {
			for x := boundingBox.Min.X; x < boundingBox.Max.X; x++ {
				r, g, b, _ := levelImage.At(x, y).RGBA()
				// Magic numbers avoid pellet color
				if (0 < r && r < 57054) || (0 < g && g < 57054) || (0 < b && b < 65535) {
					return
				}
			}
		}
	} else if p.requestedDirection != p.currentDirection {
		// They are trying to turn into a non-wall, so enable Cornering logic
		println("Cornering")
	}

	p.currentPixelPos.X, p.currentPixelPos.Y = p.nextPixelPos.X, p.nextPixelPos.Y
	p.currentDirection = p.requestedDirection
	p.updateTileLocation()
}

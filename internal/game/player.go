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
		{nextX, nextY},           // In front
		{nextDiagX1, nextDiagY1}, // Lower
		{nextDiagX2, nextDiagY2}, // Upper
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

	// Predict next position
	p.setNextPixelCoord()
	p.setNextTiles(len(*grid))

	if p.requestedDirection != p.currentDirection {
		// Player is trying to turn
		nextTilePos := p.nextTilePos[0]
		lowerDiagonalTilePos := p.nextTilePos[1]
		upperDiagonalTilePos := p.nextTilePos[2]
		nextTile := (*grid)[nextTilePos.Y][nextTilePos.X]
		lowerDiagonalTile := (*grid)[lowerDiagonalTilePos.Y][lowerDiagonalTilePos.X]
		upperDiagonalTile := (*grid)[upperDiagonalTilePos.Y][upperDiagonalTilePos.X]

		nextTileIsWall := nextTile.Type == assets.TileTypeWall
		upperDiagonalTileIsWall := upperDiagonalTile.Type == assets.TileTypeWall
		lowerDiagonalTileIsWall := lowerDiagonalTile.Type == assets.TileTypeWall

		if nextTileIsWall {
			// Player is trying to turn into a wall, so ignore the direction change
			// but do update position.
			p.currentPixelPos.X, p.currentPixelPos.Y = p.nextPixelPos.X, p.nextPixelPos.Y

			p.updateTileLocation()
			return
		}

		preTurn, postTurn := false, false

		// Calculate Pac-Man's position within the current tile.
		// The fractional part represents how far Pac-Man has moved within the tile.
		tileFractionX := float64(int(p.currentPixelPos.X)%assets.TileSize) / float64(assets.TileSize)
		tileFractionY := float64(int(p.currentPixelPos.Y)%assets.TileSize) / float64(assets.TileSize)

		// Calculate the centerline within the current tile.
		centerlineX := 0.5 // Assuming the centerline is at the middle of the tile (0.5).

		// Check if Pac-Man has crossed the centerline.
		if (p.currentDirection == Right && tileFractionX >= centerlineX) ||
			(p.currentDirection == Left && tileFractionX <= centerlineX) ||
			(p.currentDirection == Up && tileFractionY <= centerlineX) ||
			(p.currentDirection == Down && tileFractionY >= centerlineX) {
			// Pac-Man has crossed the centerline, indicating a post-turn state.
			postTurn = true
		} else {
			// Pac-Man is not in a post-turn state, so check if he's in a pre-turn state.
			if (!upperDiagonalTileIsWall && (p.currentDirection == Up || p.currentDirection == Right)) ||
				(!lowerDiagonalTileIsWall && (p.currentDirection == Down || p.currentDirection == Right)) {
				// Pac-Man is in a pre-turn state.
				preTurn = true
			}
		}

		if preTurn {
			println("Pre turn")
			// Modify Pac-Man's velocity for a 45-degree angle
			diagonalSpeed := p.moveSpeed / 1.4142 // 1.4142 is the square root of 2 (for equal diagonal speed)
			// Update Pac-Man's position until he reaches the centerline
			// Assuming `centerlineX` is calculated as in the previous code example.
			if (p.currentDirection == Right && tileFractionX < centerlineX) ||
				(p.currentDirection == Left && tileFractionX > centerlineX) ||
				(p.currentDirection == Up && tileFractionY > centerlineX) ||
				(p.currentDirection == Down && tileFractionY < centerlineX) {
				// Pac-Man is in a pre-turn state, so update his position.
				switch p.currentDirection {
				case Right:
					p.currentPixelPos.X += diagonalSpeed // Move diagonally up and right
					p.currentPixelPos.Y -= diagonalSpeed
				case Down:
					p.currentPixelPos.X += diagonalSpeed // Move diagonally down and right
					p.currentPixelPos.Y += diagonalSpeed
				case Left:
					p.currentPixelPos.X -= diagonalSpeed // Move diagonally down and left
					p.currentPixelPos.Y += diagonalSpeed
				case Up:
					p.currentPixelPos.X -= diagonalSpeed // Move diagonally up and left
					p.currentPixelPos.Y -= diagonalSpeed
				}
			}
		} else if postTurn {
			println("Post turn")
			// Modify Pac-Man's velocity for a 45-degree angle
			diagonalSpeed := p.moveSpeed / 1.4142 // 1.4142 is the square root of 2 (for equal diagonal speed)

			// Update Pac-Man's position until he has completed the turn
			// Assuming `centerlineX` is calculated as in the previous code example.
			if (p.currentDirection == Right && tileFractionX >= centerlineX) ||
				(p.currentDirection == Left && tileFractionX <= centerlineX) ||
				(p.currentDirection == Up && tileFractionY <= centerlineX) ||
				(p.currentDirection == Down && tileFractionY >= centerlineX) {
				// Pac-Man is in a post-turn state, so update his position.
				switch p.currentDirection {
				case Right:
					p.currentPixelPos.X += diagonalSpeed // Move diagonally up and right
					p.currentPixelPos.Y -= diagonalSpeed
				case Down:
					p.currentPixelPos.X += diagonalSpeed // Move diagonally down and right
					p.currentPixelPos.Y += diagonalSpeed
				case Left:
					p.currentPixelPos.X -= diagonalSpeed // Move diagonally down and left
					p.currentPixelPos.Y += diagonalSpeed
				case Up:
					p.currentPixelPos.X -= diagonalSpeed // Move diagonally up and left
					p.currentPixelPos.Y -= diagonalSpeed
				}
			}
		} else {
			p.currentPixelPos.X, p.currentPixelPos.Y = p.nextPixelPos.X, p.nextPixelPos.Y

			p.updateTileLocation()
			return
		}

	} else {
		// Continue in the same direction
		nextTilePos := p.nextTilePos[0]
		nextTile := (*grid)[nextTilePos.Y][nextTilePos.X]
		if nextTile.Type == assets.TileTypeWall {
			// Perform pixel analysis to allow proper approaches to wall tiles.
			boundingBox := image.Rect(int(p.nextPixelPos.X), int(p.nextPixelPos.Y), int(p.nextPixelPos.X+assets.SpriteSize), int(p.nextPixelPos.Y+assets.SpriteSize))

			for y := boundingBox.Min.Y; y < boundingBox.Max.Y; y++ {
				for x := boundingBox.Min.X; x < boundingBox.Max.X; x++ {
					r, g, b, _ := levelImage.At(x, y).RGBA()
					// Magic numbers avoid pellet color
					if (0 < r && r < 57054) || (0 < g && g < 57054) || (0 < b && b < 65535) {
						// Player is up against a wall, stop all movement.
						return
					}
				}
			}
		}
	}

	// Update the position of the player to the next position.
	p.currentPixelPos.X, p.currentPixelPos.Y = p.nextPixelPos.X, p.nextPixelPos.Y
	p.currentDirection = p.requestedDirection
	p.updateTileLocation()
}

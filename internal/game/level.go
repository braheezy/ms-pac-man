package game

import (
	"image/color"
	"log"

	"github.com/braheezy/ms-pacman/internal/assets"
	"github.com/hajimehoshi/ebiten/v2"
)

type Level struct {
	image  *ebiten.Image
	grid   [][]assets.TileType
	name   string
	player Player
}

var levels = []Level{
	{
		name: "Level 1",
		player: Player{
			tileCoord:  Position{x: 13, y: 23},
			pixelCoord: Position{x: 13 * assets.TileSize, y: 23*assets.TileSize - (assets.TileSize / 2)},
			image:      assets.LoadSprite("mspac_Lnorm"),
			direction:  Left,
			moveSpeed:  playerSpeedForLevel(1),
		},
	},
}

func playerSpeedForLevel(level int) float64 {
	if level == 1 {
		return 0.8 * Config.MaxMoveSpeed / float64(ebiten.TPS())
	} else if level >= 2 && level <= 4 {
		return 0.9 * Config.MaxMoveSpeed / float64(ebiten.TPS())
	} else if level >= 5 && level <= 20 {
		return Config.MaxMoveSpeed / float64(ebiten.TPS())
	} else {
		return 0.9 * Config.MaxMoveSpeed / float64(ebiten.TPS())
	}
}

func newDefaultLevel() *Level {
	firstLevel := levels[0]
	levelImage, tiles, err := assets.LoadLevelImage(firstLevel.name)
	if err != nil {
		log.Fatalln(err)
	}
	firstLevel.image = levelImage
	firstLevel.grid = tiles

	return &firstLevel
}

func (l *Level) MovePlayer() {
	if ebiten.IsKeyPressed(ebiten.KeyArrowRight) {
		l.player.direction = Right
	} else if ebiten.IsKeyPressed(ebiten.KeyArrowLeft) {
		l.player.direction = Left
	} else if ebiten.IsKeyPressed(ebiten.KeyArrowUp) {
		l.player.direction = Up
	} else if ebiten.IsKeyPressed(ebiten.KeyArrowDown) {
		l.player.direction = Down
	}

	nextX, nextY := l.player.getNextPixelCoord()
	gridSize := len(l.grid)
	x, y := l.player.getNextTileCoord(gridSize)
	nextTileType := l.grid[y][x]
	if nextTileType == assets.TileTypeWall {
		tileX := x * assets.TileSize
		tileY := y * assets.TileSize

		// Create a new image containing only the specific tile
		tileSprite := ebiten.NewImage(assets.TileSize, assets.TileSize)
		tileSprite.Fill(color.Transparent) // Fill the new image with transparent color

		// Draw the specific tile onto the new image
		op := &ebiten.DrawImageOptions{}
		op.GeoM.Translate(-float64(tileX), -float64(tileY)) // Offset to draw the specific tile
		tileSprite.DrawImage(l.image, op)

		pixels := make([]byte, 4*tileSprite.Bounds().Dx()*tileSprite.Bounds().Dy())
		tileSprite.ReadPixels(pixels)

		// Calculate the pixel position within the tile sprite
		tileOffsetX := int(nextX) % assets.TileSize
		tileOffsetY := int(nextY) % assets.TileSize
		// Check if the pixel at the calculated position in the tile sprite is black (indicating a wall)
		r, g, b, _ := tileSprite.At(tileOffsetX, tileOffsetY).RGBA()
		if (r == 0 || r == 65535) && (g == 0 || g == 65535) && (b == 0 || b == 65535) {
			l.player.pixelCoord.x, l.player.pixelCoord.y = nextX, nextY
			l.player.tileCoord.x, l.player.tileCoord.y = float64(x), float64(y)
		} else {
			// It's a wall, do not update player's position
			print("found wall pixel")
		}
	} else {

		l.player.pixelCoord.x, l.player.pixelCoord.y = nextX, nextY
		l.player.tileCoord.x, l.player.tileCoord.y = float64(x), float64(y)
	}

}

func (l *Level) Update() {
	l.MovePlayer()
}

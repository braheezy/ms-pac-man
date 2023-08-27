package game

import (
	"image"
	"log"

	"github.com/braheezy/ms-pacman/internal/assets"
	"github.com/hajimehoshi/ebiten/v2"
)

type Level struct {
	image  *ebiten.Image
	grid   [][]assets.Tile
	name   string
	player Player
}

var levels = []Level{
	{
		name: "Level 1",
		player: Player{
			tileX:     13,
			tileY:     23,
			pixelX:    13 * assets.TileSize,
			pixelY:    23*assets.TileSize - (assets.TileSize / 2),
			image:     assets.LoadSprite("mspac_Lnorm"),
			direction: Left,
			moveSpeed: playerSpeedForLevel(1),
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
	nextTileX, nextTileY := l.player.getNextTileCoord()
	nextTile := l.grid[nextTileY][nextTileX]

	if nextTile.Type == assets.TileTypeWall {
		boundingBox := image.Rect(int(nextX), int(nextY), int(nextX+assets.SpriteSize), int(nextY+assets.SpriteSize))

		for y := boundingBox.Min.Y; y < boundingBox.Max.Y; y++ {
			for x := boundingBox.Min.X; x < boundingBox.Max.X; x++ {
				r, g, b, _ := l.image.At(x, y).RGBA()
				// Magic numbers avoid pellet color
				if (0 < r && r < 57054) || (0 < g && g < 57054) || (0 < b && b < 65535) {
					return
				}
			}
		}
	}
	l.player.pixelX, l.player.pixelY = nextX, nextY
	l.player.updateTileLocation()
}

func (l *Level) Update() {
	l.MovePlayer()
}

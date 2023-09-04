package game

import (
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
			currentTilePos: TilePos{
				X: 13,
				Y: 23,
			},
			currentPixelPos: PixelPos{
				X: 13 * assets.TileSize,
				Y: 23*assets.TileSize - (assets.TileSize / 2),
			},
			nextPixelPos: PixelPos{
				X: 13 * assets.TileSize,
				Y: 23*assets.TileSize - (assets.TileSize / 2),
			},
			image:              assets.LoadSprite("mspac_Lnorm"),
			currentDirection:   Left,
			requestedDirection: Left,
			moveSpeed:          playerSpeedForLevel(1),
		},
	},
}

func playerSpeedForLevel(level int) float32 {
	if level == 1 {
		return 0.8 * Config.MaxMoveSpeed / float32(ebiten.TPS())
	} else if level >= 2 && level <= 4 {
		return 0.9 * Config.MaxMoveSpeed / float32(ebiten.TPS())
	} else if level >= 5 && level <= 20 {
		return Config.MaxMoveSpeed / float32(ebiten.TPS())
	} else {
		return 0.9 * Config.MaxMoveSpeed / float32(ebiten.TPS())
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

func (l *Level) Update() {
	l.player.Update(&l.grid, l.image)
}

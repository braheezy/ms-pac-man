package game

import (
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

	l.player.pixelCoord.x, l.player.pixelCoord.y = l.player.getNextPixelCoord()

}

func (l *Level) Update() {
	l.MovePlayer()
}

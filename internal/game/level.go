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
		},
	},
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
	var newDirection Direction
	if ebiten.IsKeyPressed(ebiten.KeyArrowRight) {
		newDirection = Right
	} else if ebiten.IsKeyPressed(ebiten.KeyArrowLeft) {
		newDirection = Left
	} else if ebiten.IsKeyPressed(ebiten.KeyArrowUp) {
		newDirection = Up
	} else if ebiten.IsKeyPressed(ebiten.KeyArrowDown) {
		newDirection = Down
	}

	nextX, nextY := l.player.getNextTileCoord()
	if !l.isWallTile(nextX, nextY) {
		l.player.direction = newDirection
	}

	l.player.pixelCoord.x, l.player.pixelCoord.y = l.player.getNextPixelCoord()

	op := &ebiten.DrawImageOptions{}

	op.GeoM.Translate(float64(l.player.pixelCoord.x), float64(l.player.pixelCoord.y))
	l.image.DrawImage(l.player.image, op)
}

func (l *Level) Update() {
	l.MovePlayer()
}

func (l *Level) isWallTile(x, y int) bool {
	return l.grid[x][y] == assets.TileTypeWall
}

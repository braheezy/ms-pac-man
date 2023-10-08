package game

import (
	"log"

	"github.com/braheezy/ms-pacman/internal/assets"
	"github.com/hajimehoshi/ebiten/v2"
)

type Level struct {
	image  *ebiten.Image
	name   string
	player Player
}

var levels = []Level{
	{
		name: "Level 1",
		player: Player{
			currentWaypoint: WaypointPos{
				X: 13,
				Y: 23,
			},
			nextWaypoint: WaypointPos{
				X: 12,
				Y: 23,
			},
			currentPixelPos: PixelPos{
				X: 13*assets.TileSize + (assets.TileSize / 2),
				Y: 23*assets.TileSize + (assets.TileSize / 2),
			},
			image:              assets.LoadSprite("mspac_Lnorm"),
			currentDirection:   Left,
			requestedDirection: Left,
			moveSpeed:          playerSpeedForLevel(1),
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
	return createLevel(0)
}

func createLevel(index int) *Level {
	level := levels[index]

	var err error
	level.image, level.player.grid, err = assets.LoadLevelImage(level.name)
	if err != nil {
		log.Fatalln(err)
	}
	level.player.waypointHeight = len(level.player.grid)
	level.player.waypointWidth = len(level.player.grid[0])

	return &level
}

func (l *Level) Update() {
	l.player.Update()
}

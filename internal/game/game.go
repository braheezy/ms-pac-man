package game

import (
	"github.com/hajimehoshi/ebiten/v2"
)

type Game struct {
	level Level
}

func New() *Game {
	firstLevel := newDefaultLevel()

	return &Game{level: *firstLevel}
}

func (g *Game) Update() error {
	g.level.Update()
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	screen.DrawImage(g.level.image, nil)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return Config.Width, Config.Height
}

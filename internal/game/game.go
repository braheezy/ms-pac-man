package game

import (
	"fmt"

	"github.com/braheezy/ms-pacman/internal/assets"
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
	if ebiten.IsKeyPressed(ebiten.KeyEscape) {
		return fmt.Errorf("user pressed escape")
	}
	g.level.Update()
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	// Draw the level image
	screen.DrawImage(g.level.image, nil)

	halfSpriteSize := assets.SpriteSize / 2.0

	op := &ebiten.DrawImageOptions{}
	// Put the player on the level
	op.GeoM.Translate(
		float64(g.level.player.currentPixelPos.X)-halfSpriteSize,
		float64(g.level.player.currentPixelPos.Y)-halfSpriteSize,
	)
	screen.DrawImage(g.level.player.image, op)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return Config.Width, Config.Height
}

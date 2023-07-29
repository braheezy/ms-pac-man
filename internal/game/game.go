package game

import (
	"github.com/hajimehoshi/ebiten/v2"
)

// Defines Game struct. Game implements ebiten.Game interface. ebiten.Game has necessary functions for an Ebitengine game: Update, Draw and Layout. Let's see them one by one.
type Game struct {
	CurrentLevelImage *ebiten.Image
}

// called every tick. Tick is a time unit for logical updating. The default value is 1/60 [s], then Update is called 60 times per second by default (i.e. an Ebitengine game works in 60 ticks-per-second).
//
//	In general, when updating function returns a non-nil error, the Ebitengine game suspends. As this program never returns a non-nil error, the Ebitengine game never stops unless the user closes the window.
func (g *Game) Update() error {
	return nil
}

// called every frame. Frame is a time unit for rendering and this depends on the display's refresh rate. If the monitor's refresh rate is 60 [Hz], Draw is called 60 times per second.
//
// Draw takes an argument screen, which is a pointer to an ebiten.Image. In Ebitengine, all images like images created from image files, offscreen images (temporary render target), and the screen are represented as ebiten.Image objects. screen argument is the final destination of rendering. The window shows the final state of screen every frame.
func (g *Game) Draw(screen *ebiten.Image) {
	screen.DrawImage(g.CurrentLevelImage, &ebiten.DrawImageOptions{})
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return Config.Width, Config.Height
}

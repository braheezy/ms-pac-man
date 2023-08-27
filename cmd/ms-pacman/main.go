package main

import (
	"log"

	"github.com/braheezy/ms-pacman/internal/game"

	"github.com/hajimehoshi/ebiten/v2"
)

func main() {
	ebiten.SetWindowSize(game.Config.Width, game.Config.Height)
	ebiten.SetWindowTitle("Ms. Pacman")

	g := game.New()

	if err := ebiten.RunGame(g); err != nil && err.Error() != "user pressed escape" {
		log.Fatal(err)
	}
}

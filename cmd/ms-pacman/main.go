package main

import (
	"log"

	"github.com/braheezy/ms-pacman/internal/assets"
	"github.com/braheezy/ms-pacman/internal/game"

	"github.com/hajimehoshi/ebiten/v2"
)

func main() {
	ebiten.SetWindowSize(game.Config.Width, game.Config.Height)
	ebiten.SetWindowTitle("Ms. Pacman")

	img, err := assets.CreateLevelImage("level1")
	if err != nil {
		panic(err)
	}

	if err := ebiten.RunGame(&game.Game{CurrentLevelImage: img}); err != nil {
		log.Fatal(err)
	}
}

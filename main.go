package main

import (
	"log"

	"github.com/akhy/chipwolu/pkg/chip8"
	"github.com/hajimehoshi/ebiten/v2"
)

func main() {
	game := &chip8.Game{
		GameOpts: &chip8.DefaultGameOpts,
		Display:  &chip8.DefaultDisplayOpts,
	}
	game.Init()
	w, h := game.ScreenSize()

	ebiten.SetWindowSize(w, h)
	ebiten.SetWindowTitle("CHIP-8")

	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}

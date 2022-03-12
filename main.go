package main

import (
	"log"

	"github.com/akhy/chipwolu/pkg/chip8"
	"github.com/hajimehoshi/ebiten/v2"
)

func main() {
	screen := chip8.NewScreen(chip8.DefaultScreenOpts)
	cpu := &chip8.CPU{Screen: screen}
	gameOpts := chip8.DefaultGameOpts
	game := chip8.NewGame(cpu, screen, gameOpts)

	ebiten.SetWindowSize(screen.Width*gameOpts.Scale, screen.Height*gameOpts.Scale)

	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}

package main

import (
	"log"

	"github.com/akhy/chipwolu/pkg/chip8"
	"github.com/hajimehoshi/ebiten/v2"
)

func main() {
	emu := &chip8.Emulator{Width: 64, Height: 32}
	emu.Init()
	if err := emu.LoadGameROM("ibm.ch8"); err != nil {
		panic(err)
	}

	gameOpts := chip8.DefaultGameOpts
	game := chip8.NewGame(emu, gameOpts)
	ebiten.SetWindowSize(emu.Width*gameOpts.Scale, emu.Height*gameOpts.Scale)
	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}

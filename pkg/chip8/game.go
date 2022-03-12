package chip8

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

type GameOpts struct {
	Scale   int
	BgColor color.Color
	FgColor color.Color
}

var DefaultGameOpts = &GameOpts{
	Scale:   16,
	BgColor: color.Black,
	FgColor: color.White,
}

type game struct {
	opts *GameOpts
	emu  *Emulator
}

func NewGame(emu *Emulator, opts *GameOpts) ebiten.Game {
	return &game{
		opts: opts,
		emu:  emu,
	}
}

func (g *game) Update() error {
	g.emu.Step()

	return nil
}

func (g *game) Draw(screen *ebiten.Image) {
	stage := ebiten.NewImage(g.emu.Width, g.emu.Height)
	stage.Fill(g.opts.BgColor)
	for x := 0; x < g.emu.Width; x++ {
		for y := 0; y < g.emu.Height; y++ {
			if g.emu.PixelAt(x, y) {
				stage.Set(x, y, g.opts.FgColor)
			}
		}
	}

	op := &ebiten.DrawImageOptions{}
	op.GeoM.Scale(float64(g.opts.Scale), float64(g.opts.Scale))
	screen.DrawImage(stage, op)

	ebitenutil.DebugPrint(screen, "Hello, World!")
}

func (g *game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return g.emu.Width * g.opts.Scale, g.emu.Height * g.opts.Scale
}

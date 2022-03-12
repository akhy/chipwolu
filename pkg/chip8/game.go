package chip8

import (
	"image/color"
	"math/rand"

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
	opts   *GameOpts
	screen *Screen
	cpu    *CPU
}

func NewGame(cpu *CPU, screen *Screen, opts *GameOpts) ebiten.Game {
	return &game{
		opts:   opts,
		cpu:    cpu,
		screen: screen,
	}
}

func (g *game) Init() {
	g.screen.Clear()
}

func (g *game) Update() error {
	x := rand.Intn(g.screen.Width)
	y := rand.Intn(g.screen.Height)
	g.screen.pixel[x][y] = true
	return nil
}

func (g *game) Draw(screen *ebiten.Image) {
	stage := ebiten.NewImage(g.screen.Width, g.screen.Height)
	stage.Fill(g.opts.BgColor)
	for x := 0; x < g.screen.Width; x++ {
		for y := 0; y < g.screen.Height; y++ {
			if g.screen.pixel[x][y] {
				stage.Set(x, y, g.opts.FgColor)
			}
		}
	}

	op := &ebiten.DrawImageOptions{}
	op.GeoM.Scale(float64(g.opts.Scale), float64(g.opts.Scale))
	screen.DrawImage(stage, op)

	ebitenutil.DebugPrint(screen, "Hello, World!")
}

func (g *game) ScreenSize() (width, height int) {
	return g.screen.Width * g.opts.Scale, g.screen.Height * g.opts.Scale
}

func (g *game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return g.ScreenSize()
}

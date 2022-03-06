package chip8

import (
	"image/color"
	"math/rand"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

type GameOpts struct {
	Width  int
	Height int
}

var DefaultGameOpts = GameOpts{
	Width:  64,
	Height: 32,
}

type DisplayOpts struct {
	Scale   int
	BgColor color.Color
	FgColor color.Color
}

var DefaultDisplayOpts = DisplayOpts{
	Scale:   16,
	BgColor: color.Black,
	FgColor: color.White,
}

type Game struct {
	*GameOpts
	Display *DisplayOpts

	screen [][]bool
}

func (g *Game) Init() {
	g.ClearScreen()
}

func (g *Game) ScreenSize() (width, height int) {
	return g.Width * g.Display.Scale, g.Height * g.Display.Scale
}

func (g *Game) ClearScreen() {
	g.screen = make([][]bool, g.Width)
	for x := range g.screen {
		g.screen[x] = make([]bool, g.Height)
	}
}

func (g *Game) Update() error {
	x := rand.Intn(g.Width)
	y := rand.Intn(g.Height)
	g.screen[x][y] = true
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	stage := ebiten.NewImage(g.Width, g.Height)
	stage.Fill(g.Display.BgColor)
	for x := 0; x < g.Width; x++ {
		for y := 0; y < g.Height; y++ {
			if g.screen[x][y] {
				stage.Set(x, y, g.Display.FgColor)
			}
		}
	}

	op := &ebiten.DrawImageOptions{}
	op.GeoM.Scale(float64(g.Display.Scale), float64(g.Display.Scale))
	screen.DrawImage(stage, op)

	ebitenutil.DebugPrint(screen, "Hello, World!")
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return g.ScreenSize()
}

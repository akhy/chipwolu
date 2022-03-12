package chip8

type Screen struct {
	Width  int
	Height int

	pixel [][]bool
}

type ScreenOpts struct {
	Width  int
	Height int
}

var DefaultScreenOpts = ScreenOpts{
	Width:  64,
	Height: 32,
}

func NewScreen(opts ScreenOpts) *Screen {
	screen := &Screen{
		Width:  opts.Width,
		Height: opts.Height,
	}
	screen.Clear()
	return screen
}

func (s *Screen) Clear() {
	s.pixel = make([][]bool, s.Width)

	for x := range s.pixel {
		s.pixel[x] = make([]bool, s.Height)
	}
}

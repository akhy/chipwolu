package chip8

type Machine struct {
	Screen [][]bool
}

func (m *Machine) Init(width, height int) {
	m.Screen = make([][]bool, width, height)
}

func (m *Machine) ClearScreen() {
	for i, c := range m.Screen {
		for j := range c {
			m.Screen[i][j] = false
		}
	}
}

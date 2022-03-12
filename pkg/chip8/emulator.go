package chip8

import (
	"fmt"
	"os"
)

const (
	fontStart = 0x050
	romStart  = 0x200
)

var (
	fontSet = []uint8{
		0xF0, 0x90, 0x90, 0x90, 0xF0, // 0
		0x20, 0x60, 0x20, 0x20, 0x70, // 1
		0xF0, 0x10, 0xF0, 0x80, 0xF0, // 2
		0xF0, 0x10, 0xF0, 0x10, 0xF0, // 3
		0x90, 0x90, 0xF0, 0x10, 0x10, // 4
		0xF0, 0x80, 0xF0, 0x10, 0xF0, // 5
		0xF0, 0x80, 0xF0, 0x90, 0xF0, // 6
		0xF0, 0x10, 0x20, 0x40, 0x40, // 7
		0xF0, 0x90, 0xF0, 0x90, 0xF0, // 8
		0xF0, 0x90, 0xF0, 0x10, 0xF0, // 9
		0xF0, 0x90, 0xF0, 0x90, 0x90, // A
		0xE0, 0x90, 0xE0, 0x90, 0xE0, // B
		0xF0, 0x80, 0x80, 0x80, 0xF0, // C
		0xE0, 0x90, 0x90, 0x90, 0xE0, // D
		0xF0, 0x80, 0xF0, 0x80, 0xF0, // E
		0xF0, 0x80, 0xF0, 0x80, 0x80, // F
	}
)

type Emulator struct {
	Width  int
	Height int

	screen [][]uint8
	mem    [4096]uint8 // 4kB memory
	vx     [16]uint8   // cpu register V0-VF
	key    [16]uint8   // input keys
	stack  [16]uint16  // stack

	pc uint16 // program counter
	ir uint16 // index register
	sp uint8  // stack pointer

	delayTimer uint8 // delay timer
	soundTimer uint8 // sound timer
}

func (e *Emulator) Init() {
	if e.Width == 0 || e.Height == 0 {
		panic("emulator width and hight must be set")
	}

	// set screen
	e.screen = make([][]uint8, e.Width)
	for x := range e.screen {
		e.screen[x] = make([]uint8, e.Height)
	}

	// set the starting position
	e.pc = romStart

	// load the font set
	for i := 0; i < len(fontSet); i++ {
		e.mem[i+fontStart] = fontSet[i]
	}
}

func (e *Emulator) LoadGameROM(path string) error {
	f, err := os.Open(path)
	if err != nil {
		return err
	}
	defer f.Close()

	stat, err := f.Stat()
	if err != nil {
		return err
	}

	// TODO check if rom is too big for the memory

	buffer := make([]uint8, stat.Size())
	if _, err := f.Read(buffer); err != nil {
		return err
	}

	for i := 0; i < len(buffer); i++ {
		e.mem[i+romStart] = buffer[i]
	}

	return nil
}

func (e *Emulator) PixelAt(x, y int) bool {
	return e.screen[x][y] != 0
}

func (e *Emulator) ClearScreen() {
	for x := 0; x < e.Width; x++ {
		for y := 0; y < e.Height; y++ {
			e.screen[x][y] = 0x0
		}
	}
}

func (e *Emulator) Step() {
	// read 2 bytes at once from current PC
	op := uint16(e.mem[e.pc])<<8 | uint16(e.mem[e.pc+1])

	// move pointer forward
	e.pc += 2

	// 1111 1111 0000 0000

	switch op & 0xF000 {
	case 0x0000:
		switch op {
		case 0x00E0: // Clear Screen
			e.ClearScreen()
		default:
			fmt.Printf("unknown instruction: %x", op)
		}

	case 0x1000: // 1nnn: jump to nnn
		e.pc = op & 0x0FFF

	case 0x6000: // 6xnn: set Vx to nn
		idx := (op & 0x0F00) >> 8
		val := uint8(op & 0x00FF)
		e.vx[idx] = val

	case 0x7000: // 7xnn: add Vx with nn
		idx := (op & 0x0F00) >> 8
		val := uint8(op & 0x00FF)
		e.vx[idx] += val

	case 0xA000: // Annn: set I register to nnn
		e.ir = (op & 0x0FFF)

	case 0xD000: // Dxyn: draw sprite at [x,y] with n height
		sx := e.vx[(op&0x0F00)>>8] // starting x
		sy := e.vx[(op&0x00F0)>>4] // starting y

		h := (op & 0x000F) // sprite height
		sprite := e.mem[e.ir : e.ir+h]

		e.vx[0xF] = 0
		for j := uint16(0); j < h; j++ {
			row := sprite[j] // 1010 1010
			for i := uint16(0); i < 8; i++ {
				mask := uint8(0x80 >> i) // 1000 0000 >> [0..8]

				on := (row & mask) == mask // (row & mask) >> i
				x := (uint16(sx) + i) & uint16(e.Width-1)
				y := (uint16(sy) + j) % uint16(e.Height-1)

				if e.screen[x][y] == 1 {
					e.vx[0xF] = 1
				}

				var v uint8
				if on {
					v = 0x01
				}

				e.screen[x][y] ^= v
			}
		}
	default:
		fmt.Printf("unknown instruction: %x", op)
	}
}

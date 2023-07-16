package chip8

import (
	"emulator/app"
	"fmt"
	"os"
)

const ENTRY_POINT = 0x200 // Program starts at 512 - 0010 0000 0000

/*
Fonts are read from left to right and creating the font up to bottom
Exemple for 0 :
1111 0000
1001 0000
1001 0000
1001 0000
1111 0000
*/
var FONTS = [16 * 5]uint8{
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

/*
Data Registers (V): The CHIP-8 has 16 data registers (V0 to VF), each 8 bits in size.
These registers are used to store temporary data values when executing instructions.

Address register (I): The address register is used to store 16-bit memory addresses.
It is mainly used to access data stored in memory.

Program counter (PC): The program register contains the address of the next instruction to be executed.
After each instruction, the PC is incremented to point to the next instruction.

Shift Register (DT) and Sound Register (ST): These registers are used to manage delay shifting and sound effects in the CHIP-8.
The DT register is used for delay counting, while the ST register is used to generate sounds.
*/
type Chip8 struct {
	Ram     [4096]uint8
	Display [64 * 32]bool
	Stack   [12]uint16
	SP      uint16
	Keypad  [16]bool
	V       [16]uint8 // data registers V0-VF
	I       uint16    // register index
	PC      uint16    // program counter
	DT      uint8     // delay timer
	ST      uint8     // sound timer

	inst *instructions
	app  *app.App
}

type instructions struct {
	OpCode uint16
	NNN    uint16
	NN     uint8
	N      uint8
	X      uint8
	Y      uint8
}

func (i *instructions) Init(opCode uint16) {
	i.OpCode = opCode
	i.NNN = i.OpCode & 0x0FFF
	i.NN = uint8(i.OpCode & 0x00FF)
	i.N = uint8(i.OpCode & 0x000F)
	i.X = uint8(i.OpCode >> 8 & 0x000F)
	i.Y = uint8(i.OpCode >> 4 & 0x000F)
}

func NewChip8(rom string) *Chip8 {

	chip := &Chip8{
		PC:   ENTRY_POINT,
		inst: new(instructions),
	}
	copy(chip.Ram[:], FONTS[:])
	return chip
}

/*
https://en.wikipedia.org/wiki/CHIP-8
*/
func (c *Chip8) cycle() error {
	c.inst.Init((uint16(c.Ram[c.PC]) << 8) | uint16(c.Ram[c.PC+1]))
	c.PC += 2
	switch (c.inst.OpCode >> 12) & 0x000F {
	case 0x00:
		if c.inst.NN == 0xE0 {
			// Clear screen
			for i := range c.Display {
				c.Display[i] = false
			}
		} else if c.inst.NN == 0xEE {
			// Subroutine return
			c.SP = c.SP - 1
			c.PC = c.Stack[c.SP]
		}
	case 0x01:
	case 0x02:
		// Call subroutine at NNN
		c.Stack[c.SP] = c.PC
		c.SP++
		c.PC = c.inst.NNN
	case 0x03:
	case 0x04:
	case 0x05:
	case 0x06:
	case 0x09:
	case 0x0A:
	case 0x0B:
	case 0x0C:
	case 0x0D:
	case 0x0E:
	case 0x0F:
	default:

	}
	return nil
}

func (c *Chip8) Update() {
	c.cycle()
}

func (c *Chip8) LoadFile(fileName string) error {
	file, err := os.OpenFile(fileName, os.O_RDONLY, 0777)
	if err != nil {
		return err
	}
	defer file.Close()

	stat, err := file.Stat()
	if err != nil {
		return err
	}

	size := int64(len(c.Ram) - ENTRY_POINT)
	if size < stat.Size() { // program is loaded at 0x200
		return fmt.Errorf("Program size bigger than memory size")
	}

	buffer := make([]byte, stat.Size())
	if _, readErr := file.Read(buffer); readErr != nil {
		return readErr
	}

	for i := 0; i < len(buffer); i++ {
		c.Ram[i+ENTRY_POINT] = buffer[i]
	}

	return nil
}

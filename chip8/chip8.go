package chip8

const ENTRY_POINT = 0x200 // Program starts at 512 - 0010 0000 0000

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
	ram     [4096]uint8
	display [64 * 32]bool
	stack   [12]uint16
	keypad  [16]bool
	V       [16]uint8 // data registers V0-VF
	I       uint16    // register index
	PC      uint16    // program counter
	DT      uint8     // delay timer
	ST      uint8     // sound timer

	font [16 * 5]uint8
	rom  string
}

func NewChip8(rom string) *Chip8 {
	return &Chip8{
		PC:   ENTRY_POINT,
		font: FONTS,
	}
}

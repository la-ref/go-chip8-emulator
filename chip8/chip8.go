package chip8

const ENTRY_POINT = 0x200 // Program starts at 512 - 0010 0000 0000

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

	rom string
}

func NewChip8(rom string) *Chip8 {
	return &Chip8{
		PC: ENTRY_POINT,
	}
}

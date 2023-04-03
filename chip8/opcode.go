package chip8

/*

nnn or addr - A 12-bit value, the lowest 12 bits of the instruction
n or nibble - A 4-bit value, the lowest 4 bits of the instruction
x - A 4-bit value, the lower 4 bits of the high byte of the instruction
y - A 4-bit value, the upper 4 bits of the low byte of the instruction
kk or byte - An 8-bit value, the lowest 8 bits of the instruction
the numbers are identifiers

   00E0 - CLS
   00EE - RET
   1nnn - JP addr
   2nnn - CALL addr
   3xkk - SE Vx, byte
   4xkk - SNE Vx, byte
   5xy0 - SE Vx, Vy
   6xkk - LD Vx, byte
   7xkk - ADD Vx, byte
   8xy0 - LD Vx, Vy
   8xy1 - OR Vx, Vy
   8xy2 - AND Vx, Vy
   8xy3 - XOR Vx, Vy
   8xy4 - ADD Vx, Vy
   8xy5 - SUB Vx, Vy
   8xy6 - SHR Vx {, Vy}
   8xy7 - SUBN Vx, Vy
   8xyE - SHL Vx {, Vy}
   9xy0 - SNE Vx, Vy
   Annn - LD I, addr
   Bnnn - JP V0, addr
   Cxkk - RND Vx, byte
   Dxyn - DRW Vx, Vy, nibble
   Ex9E - SKP Vx
   ExA1 - SKNP Vx
   Fx07 - LD Vx, DT
   Fx0A - LD Vx, K
   Fx15 - LD DT, Vx
   Fx18 - LD ST, Vx
   Fx1E - ADD I, Vx
   Fx29 - LD F, Vx
   Fx33 - LD B, Vx
   Fx55 - LD [I], Vx
   Fx65 - LD Vx, [I]
*/
type opcode uint16

//TakeOpcodeID takes the ID of the instruction of the opcode, taking aside the parameters of the opcode such as NNN, N, X, Y and KK
func (oc opcode) TakeOpcodeID() uint16 {
	first4Bits := (uint16(oc) & uint16(0xF000)) >> 12
	switch first4Bits {
	case uint16(0x0):
		return uint16(oc)

	case uint16(0x8):
		return uint16(oc) & uint16(0xF00F)
	case uint16(0x9):
		return uint16(oc) & uint16(0xF00F)

	case uint16(0xE):
		return uint16(oc) & uint16(0xF0FF)

	case uint16(0xF):
		return uint16(oc) & uint16(0xF0FF)

	default:
		return uint16(oc) & uint16(0xF000)
	}
}

//NNN decodes the opcode and returns the NNN parameter
func (oc opcode) NNN() uint16 {
	return uint16(oc) & uint16(0x0FFF)
}

//KK decodes the opcode and returns the KK parameter
func (oc opcode) KK() uint8 {
	return uint8(uint16(oc) & 0x00FF)
}

//N decodes the opcode and returns the N parameter
func (oc opcode) N() uint16 {
	return uint16(oc) & 0x000F
}

//X decodes the opcode and returns the X parameter
func (oc opcode) X() uint8 {
	return uint8((uint16(oc) & 0x0F00) >> 8)
}

//Y decodes the opcode and returns the Y parameter
func (oc opcode) Y() uint8 {
	return uint8((uint16(oc) & 0x00F0) >> 4)
}

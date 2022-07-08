package chip8

import (
	"math/rand"
)

//I00E0 clears the myMonitor
func (c8 *Chip8) I00E0(){ //CLS
	c8.frameBuffer = [64*32]byte{}
}

//I00EE returns from a subroutine
func (c8 *Chip8) I00EE(){ //RET
	c8.sp--
	c8.pc = c8.stack[c8.sp]
}

//I1NNN Jumps to location nnn
func (c8 *Chip8) I1NNN(){ //JP (ADDR)
	addr :=c8.cOpcode.NNN()
	c8.pc = addr
}

// I2NNN CALL (ADDR)
func (c8 *Chip8) I2NNN(){
	addr := c8.cOpcode.NNN()
	c8.stack[c8.sp] = c8.pc
	c8.sp++
	c8.pc = addr
}

//I3XKK
//Skip next instruction if Vx = kk
func (c8 *Chip8) I3XKK(){ //SE (VX, BYTE)
	vx := c8.registers[c8.cOpcode.X()]
	_byte := c8.cOpcode.KK()
	if vx == _byte{
		c8.pc += 2
	}
}

//I4XKK
//Skip next instruction if Vx != kk
func (c8 *Chip8) I4XKK(){ //SNE (VX, BYTE)
	vx := c8.registers[c8.cOpcode.X()]
	_byte := c8.cOpcode.KK()
	if vx != _byte{
		c8.pc += 2
	}
}

//I5XY0
//Skip next instruction if vx = vy
func (c8 *Chip8) I5XY0(){ //SE (VX, VY)
	vx := c8.registers[c8.cOpcode.X()]
	vy := c8.registers[c8.cOpcode.Y()]

	if vx == vy{
		c8.pc += 2
	}
}

//I6XKK Set vx = kk
func (c8 *Chip8) I6XKK(){ //LV (VX, BYTE)
  	_byte := c8.cOpcode.KK()
	c8.registers[c8.cOpcode.X()] = _byte
}

//I7XKK ADD (VX, BYTE)
func (c8 *Chip8) I7XKK(){
	_byte := c8.cOpcode.KK()
	c8.registers[c8.cOpcode.X()] += _byte
}

//I8XY0 LD (VX, VY) Set Vx = Vy
func (c8 *Chip8) I8XY0(){
	c8.registers[c8.cOpcode.X()] = c8.registers[c8.cOpcode.Y()]
}

//I8XY1 OR(VX, VY)
func (c8 *Chip8) I8XY1(){
	c8.registers[c8.cOpcode.X()] |= c8.registers[c8.cOpcode.Y()]
}

//I8XY2 AND (VX, VY)
func (c8 *Chip8) I8XY2(){
	c8.registers[c8.cOpcode.X()] &= c8.registers[c8.cOpcode.Y()]
}

//I8XY3 XOR (VX, VY)
func (c8 *Chip8) I8XY3(){
	c8.registers[c8.cOpcode.X()] ^= c8.registers[c8.cOpcode.Y()]

}

//I8XY4
//The values of Vx and Vy are added together.
//If the result is greater than 8 bits (i.e., > 255,) VF is set to 1, otherwise 0. Only the lowest 8 bits of the result are kept, and stored in Vx.
func (c8 *Chip8) I8XY4(){ //ADD (VS, VY)
	sum := c8.registers[c8.cOpcode.X()] + c8.registers[c8.cOpcode.Y()]
	c8.registers[c8.cOpcode.X()] = sum &0x00FF

	if sum > 255{
		c8.registers[0xF] = 1
	}else {
		c8.registers[0xF] = 0
	}
}

//f Vx > Vy, then VF is set to 1, otherwise 0. Then Vy is subtracted from Vx, and the results stored in Vx.
func (c8 *Chip8) I8XY5(){ //SUB (VX, VY)
	if c8.registers[c8.cOpcode.X()] > c8.registers[c8.cOpcode.Y()]{
		c8.registers[0xF] = 1
	}else{
		c8.registers[0xF] = 0
	}

	c8.registers[c8.cOpcode.X()] -= c8.registers[c8.cOpcode.Y()]
}

//I8XY6
//If the least-significant bit of Vx is 1, then VF is set to 1, otherwise 0. Then Vx is divided by 2.
func (c8 *Chip8) I8XY6(){ //SHR (VX, {, VY})
	c8.registers[0xF] = c8.registers[c8.cOpcode.X()] & 0x1 // 0x1: 00000001
	c8.registers[c8.cOpcode.X()] >>= 1
}

//I8XY7
//If Vy > Vx, then VF is set to 1, otherwise 0. Then Vx is subtracted from Vy, and the results stored in Vx.
func (c8 *Chip8) I8XY7(){ //SUBN (VX, VY)
	if c8.registers[c8.cOpcode.Y()] > c8.registers[c8.cOpcode.X()]{
		c8.registers[0xF] = 1
	}else{
		c8.registers[0xF] = 0
	}

	c8.registers[c8.cOpcode.X()] = c8.registers[c8.cOpcode.Y()] - c8.registers[c8.cOpcode.X()]

}


//I8XYE
//If the most-significant bit of Vx is 1, then VF is set to 1, otherwise to 0. Then Vx is multiplied by 2.
func (c8 *Chip8) I8XYE(){ //SHL (Vx {, Vy})
	c8.registers[0xF] = c8.registers[c8.cOpcode.X()] & 0x80 //0x80: 10000000
	c8.registers[c8.cOpcode.X()] <<= 1
}

//I9XY0
//Skip next instruction if Vx != Vy.
func (c8 *Chip8) I9XY0(){ //SNE (Vx, Vy)
	if c8.registers[c8.cOpcode.X()] != c8.registers[c8.cOpcode.Y()] {
		c8.pc += 2
	}
}

//IANNN Set I = NNN
func (c8 *Chip8) IANNN(){ // LD I, addr
	c8.i = c8.cOpcode.NNN()
}

//IBNNN Jump to location nnn + V0.
func (c8 *Chip8) IBNNN(){ // JP V0, addr
	c8.pc = uint16(c8.registers[0]) + c8.cOpcode.NNN()
}

//ICXKK Set Vx = random byte AND kk.
func (c8 *Chip8) ICXKK(){ // RND Vx, byte
	c8.registers[c8.cOpcode.X()] = uint8(rand.Intn(256)) & c8.cOpcode.KK()
}

//IDXYN displays n-byte sprite starting at memory location I at (Vx, Vy) and set VF = collision.
func (c8 *Chip8) IDXYN(){ // DRW (Vx, Vy, hSprite)
	vx := c8.registers[c8.cOpcode.X()]
	vy := c8.registers[c8.cOpcode.Y()]

	//If a sprite is attempting to draw outside the bounds of the screen,
	//it wraps around to the other side, that's why we do x0 = vx % width, y0 = vy % height.
	x0 := int(vx % WidthScreen)
	y0 := int(vy % HeightScreen)

	//Each sprite has a width of 8 pixels (represented by a byte), and a height N
	hSprite := int(c8.cOpcode.N())
	i := int(c8.i)
	var _byte byte
	var bit byte

	for y :=0; y < hSprite; y++{
		_byte = c8.memory[i + y]

		//Every bit of each i-byte (0<i<N) of the sprite represents a pixel on the screen which can be ON or OFF.
		//if the sprite pixel is ON, then we check if that pixel is already ON in the FrameBuffer. In that case we set Vf = 1 to indicate a collision
		//then we use a XOR operation, so if the sprite pixel is ON and the display pixel is ON, we set the display pixel to OFF
		//and if the sprite pixel is ON and the display pixel is OFF, we set the display pixel to ON
		for x :=0; x < 8; x++{
			bit = _byte & (0x80 >> x)
			if bit != 0{

				if c8.frameBuffer.CheckOverlap(x0 + x, y0 + y){
					continue
				}else{
					cellFrameBuffer := c8.frameBuffer.Get(x0 + x, y0 + y)
					if *cellFrameBuffer == 1{
						c8.registers[0xF] = 0xFF
					}
					*cellFrameBuffer ^= 0xFF
				}

			}
		}
	}
	c8.MustDraw = true

}

//IEX9E Skip next instruction if key with the value of Vx is pressed.
func (c8 *Chip8) IEX9E(){ //SKP(VX)
	key := c8.registers[c8.cOpcode.X()]
	if c8.Keypad[key] == 1{
		c8.pc += 2
		c8.Keypad[key] = 0
	}

}
//IEXA1 Skip next instruction if key with the value of Vx is not pressed.
func (c8 *Chip8) IEXA1(){ //SKP(VX)
	key := c8.registers[c8.cOpcode.X()]
	if c8.Keypad[key] != 1{
		c8.pc += 2
	}else{
		c8.Keypad[key] = 0
	}
}

//IFX07 Set Vx = delay timer value.
func (c8 *Chip8) IFX07(){ //LD (Vx, DT)
	c8.registers[c8.cOpcode.X()] = c8.delayTimer
}

//IFX0A Wait for a key press, store the value of the key in Vx.
func (c8 *Chip8) IFX0A(){//LD (Vx, K)
	for i, isPress := range c8.Keypad {
		if isPress != 0 {
			c8.registers[c8.cOpcode.X()] = byte(i)
			return
		}
	}
}

//IFX15 Set delay timer = Vx
func (c8 *Chip8) IFX15(){ //LD (DT, Vx)
	c8.delayTimer = c8.registers[c8.cOpcode.X()]
}

//IFX18 Set sound timer = Vx
func (c8 *Chip8) IFX18(){ //LD (ST, Vx)
	c8.soundTimer = c8.registers[c8.cOpcode.X()]
}

//IFX1E Set I = I + Vx
func (c8 *Chip8) IFX1E(){ //ADD (I, Vx)
	c8.i += uint16(c8.registers[c8.cOpcode.X()])
}

//IFX29 Set I = location of sprite for digit Vx.
func (c8 *Chip8) IFX29(){ //LD (F, Vx)
	vx := c8.registers[c8.cOpcode.X()] // c between 0 and F
	c8.i = FontsetStartAddress + uint16(FontSize * vx)
}

//IFX33 Store BCD representation of Vx in memory locations I, I+1, and I+2.
//The interpreter takes the decimal value of Vx,
//and places the hundreds digit in memory at location in I, the tens digit at location I+1, and the ones digit at location I+2.
func (c8 *Chip8) IFX33(){ //LD (B, Vx)
	vx := c8.registers[c8.cOpcode.X()]
	c8.memory[c8.i + 2] = vx  % 10
	c8.memory[c8.i + 1] = (vx / 10) % 10
	c8.memory[c8.i] = (vx / 100) % 10
}

//IFX55 Stores registers V0 through Vx in memory starting at location I.
func (c8 *Chip8) IFX55(){ //LD (I,Vx)
	for  k := 0; k <= int(c8.cOpcode.X()); k++{
		c8.memory[c8.i + uint16(k)] = c8.registers[k]
	}
}

//IFX65 Reads registers V0 through Vx from memory starting at location I.
func (c8 *Chip8) IFX65(){ //LD (Vx, I)
	for  k := 0; k <= int(c8.cOpcode.X()); k++{
		c8.registers[k] = c8.memory[c8.i + uint16(k)]
	}
}
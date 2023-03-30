package chip8

import (
	"errors"
	"github.com/NoetherianRing/Chip-8/monitor"
	"github.com/NoetherianRing/Chip-8/state"
	"os"
)

type Chip8 struct {
	memory    [TotalMemory]byte       //The chip8 has 4096 memory cells.
	registers [NumberOfRegisters]byte //The chip8 has 16 registers.
	pc        uint16                  //ProgramCounter. It's an uint16 to be able to store each of the 4096 memory addresses
	//(There are some memory addresses too large to store in just 8 bits)
	i     uint16              //Index Register. It's used to store memory addresses for use in operations.
	stack [StackLevels]uint16 //A stack is the way for the Chip8 to keep track of the order of execution when it calls into functions.
	//It's an array of length 16 because there are 16 levels of nesting
	sp byte //stack pointer, to keep track of what nesting level the program is at.

	instructions map[uint16]func()
	cOpcode      opcode              //current opcode
	keyPressed   chan byte           //The chip8 has a hex keypad. The channel is acceded by the peripherals and represents the key that was just pressed
	frameBuffer  monitor.FrameBuffer //The Chip8 has a monochromatic screen of 64x32 pixels. Each element of the FrameBuffer represents a pixel
	//Each pixel can be on or off.

	delayTimer byte
	soundTimer byte
	MustDraw   bool
	quit       bool
}

func NewChip8(keyPressed chan byte) (*Chip8, error) {
	c8 := &Chip8{
		memory:       [TotalMemory]byte{},
		registers:    [NumberOfRegisters]byte{},
		pc:           PCStartAddress,
		stack:        [StackLevels]uint16{},
		frameBuffer:  monitor.FrameBuffer{},
		instructions: map[uint16]func(){},
	}

	c8.keyPressed = keyPressed
	c8.instructions[0x00E0] = c8.I00E0
	c8.instructions[0x00EE] = c8.I00EE
	c8.instructions[0x1000] = c8.I1NNN
	c8.instructions[0x2000] = c8.I2NNN
	c8.instructions[0x3000] = c8.I3XKK
	c8.instructions[0x4000] = c8.I4XKK
	c8.instructions[0x5000] = c8.I5XY0
	c8.instructions[0x6000] = c8.I6XKK
	c8.instructions[0x7000] = c8.I7XKK
	c8.instructions[0x8000] = c8.I8XY0
	c8.instructions[0x8001] = c8.I8XY1
	c8.instructions[0x8002] = c8.I8XY2
	c8.instructions[0x8003] = c8.I8XY3
	c8.instructions[0x8004] = c8.I8XY4
	c8.instructions[0x8005] = c8.I8XY5
	c8.instructions[0x8006] = c8.I8XY6
	c8.instructions[0x8007] = c8.I8XY7
	c8.instructions[0x800E] = c8.I8XYE
	c8.instructions[0x9000] = c8.I9XY0
	c8.instructions[0xA000] = c8.IANNN
	c8.instructions[0xB000] = c8.IBNNN
	c8.instructions[0xC000] = c8.ICXKK
	c8.instructions[0xD000] = c8.IDXYN
	c8.instructions[0xE09E] = c8.IEX9E
	c8.instructions[0xE0A1] = c8.IEXA1
	c8.instructions[0xF007] = c8.IFX07
	c8.instructions[0xF00A] = c8.IFX0A
	c8.instructions[0xF015] = c8.IFX15
	c8.instructions[0xF018] = c8.IFX18
	c8.instructions[0xF01E] = c8.IFX1E
	c8.instructions[0xF029] = c8.IFX29
	c8.instructions[0xF033] = c8.IFX33
	c8.instructions[0xF055] = c8.IFX55
	c8.instructions[0xF065] = c8.IFX65
	c8.instructions[0x9001] = c8.I9XY1
	c8.instructions[0x9002] = c8.I9XY2

	return c8, nil
}

//LoadROM is called by an external app running chip8 to load a ROM file into memory
//The amount of memory that is allowed to be used for the ROM File and the addresses to store it are given by the specification of the chip8
func (c8 *Chip8) LoadROM(filename string) error {
	return loadFile(filename, MemoryForROM, PCStartAddress, &c8.memory)
}

//LoadFonts is called by an external app running chip8 to load a font file into memory
func (c8 *Chip8) LoadFonts(filename string) error {
	return loadFile(filename, MemoryForFonts, FontsetStartAddress, &c8.memory)
}

//loadFile loads a file into the chip8 memory
func loadFile(filename string, maxCapacity int, startAddress int, dst *[TotalMemory]byte) error {
	file, err := os.ReadFile(filename)

	if err != nil {
		return err
	}
	if len(file) > maxCapacity {
		errS := "the ROM in '" + filename + "' exceeds the memory capacity of Chip-8."
		return errors.New(errS)
	}
	copy(dst[startAddress:], file[:])

	return nil
}

//fetchOpcode takes half of the opcode from the current position of the program counter, and the other half from program counter + 1
//this is because an opcode has 2 bytes and every memory cell only has 1 byte,
//then our program counter moves two cells forward
func (c8 *Chip8) fetchOpcode() {
	c8.cOpcode = opcode(uint16(c8.memory[c8.pc])<<8 | uint16(c8.memory[c8.pc+1]))
	c8.pc += 2
}

//executeOpcode decodes the ID of the current opcode, and then execute the corresponding instruction
func (c8 *Chip8) executeOpcode() {
	id := c8.cOpcode.TakeOpcodeID()
	if inst, ok := c8.instructions[id]; ok {
		inst()
	}

}

//countBackDelayTimer. The chip8 has a delay timer which decreases in every cycle
func (c8 *Chip8) countBackDelayTimer() {
	if c8.delayTimer != 0 {
		c8.delayTimer--
	}
}

//countBackSoundTimer. The chip8 has a sound timer which decreases in every cycle
func (c8 *Chip8) countBackSoundTimer() {
	if c8.soundTimer != 0 {
		c8.soundTimer--
	}
}

//MustBeep. If the soundTimer != 0, the chip must make a beep.
func (c8 *Chip8) MustBeep() bool {
	if c8.soundTimer != 0 {
		return true
	} else {
		return false
	}
}

//GetFrameBuffer expose the FrameBuffer so it can be read by the monitor peripheral
func (c8 *Chip8) GetFrameBuffer() monitor.FrameBuffer {
	return c8.frameBuffer
}

func (c8 *Chip8) IsClosed() bool {
	return c8.quit
}

//Close can be call for an external app which manages the chip8 to close it
func (c8 *Chip8) Close() {
	c8.quit = true
}

//Cycle can be call for an external app which manages the chip8 with certain frequency
//In every cycle we read, decode and execute the current opcode and we move the program counter by two, the we count back the sound timer and the delay timer
func (c8 *Chip8) Cycle() {
	c8.fetchOpcode()
	c8.executeOpcode()
	c8.countBackSoundTimer()
	c8.countBackDelayTimer()
}

//Dump is used in the debug mode of the app, it dumps the state of the chip8 into a StateChip8 an return it
func (c8 *Chip8) Dump() *state.StateChip8 {
	s := new(state.StateChip8)
	array := [400096]byte{}
	s.Memory = array//c8.memory
	s.Registers = c8.registers
	s.Pc = c8.pc
	s.I = c8.i
	s.Stack = c8.stack
	s.Sp = c8.sp
	s.COpcode = uint16(c8.cOpcode)
	s.FrameBuffer = c8.frameBuffer
	s.DelayTimer = c8.delayTimer
	s.SoundTimer = c8.soundTimer
	s.MustDraw = c8.MustDraw
	s.Quit = c8.quit
	return s

}

package chip8

import "time"

const (
	PCStartAddress      = 0x200 //Memory from 0x200 to 0xFFF is reserved for the ROM File, so the program counter must start at 0x200
	FontsetStartAddress = 0x50  //Originally the memory from 0x000 to 0x1FF was reserved for the Chip8 interpreter, we use the memory from
	//0x050 to 0x0A0 to store Fonts.
	//TotalMemory    = 4096
	TotalMemory    = 400096
	MemoryForROM   = 0xFFF - 0x200 //Amount of memory reserved for ROMs Files
	MemoryForFonts = 0x0A0 - 0x050 //Amount of memory reserved for Fonts
	//There are 16 fonts (0 to F), each is represented by 5 bytes
	NumberOfRegisters  = 16
	StackLevels        = 16
	NumberOfKeys       = 16
	FontSize           = 5                                   //every font is represented by 5 bytes
	Frequency          = time.Second / time.Duration(500)    //The frequency should be 60Hz but that is very slow
	FrequencyDebugMode = time.Second / time.Duration(500000) //In debug mode the frequency must be smaller, because it's a slower mode
	WidthScreen        = 64
	HeightScreen       = 32
	AsciiEscape		   = 0x1B
)

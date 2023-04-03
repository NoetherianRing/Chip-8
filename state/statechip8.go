package state

import (
	"github.com/NoetherianRing/Chip-8/monitor"
)

//StateChip8 is used to save the state of the chip8 on every cycle when the app is running in debug mode,
//allowing it to be easily stored in a file and then compared to an expected result
type StateChip8 struct {
	Memory      [400096]byte
	Registers   [16]byte
	Pc          uint16
	I           uint16
	Stack       [16]uint16
	Sp          byte
	COpcode     uint16
	FrameBuffer monitor.FrameBuffer
	DelayTimer  byte
	SoundTimer  byte
	MustDraw    bool
	Quit        bool
}

package chip8

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestOpcode_TakeOpcodeID(t *testing.T) {
	var oc opcode
	oc = 0x1234
	assert.Equal(t, uint16(0x1000), oc.TakeOpcodeID(), "")
	oc = 0x2234
	assert.Equal(t, uint16(0x2000), oc.TakeOpcodeID(), "")
	oc = 0x8230
	assert.Equal(t, uint16(0x8000), oc.TakeOpcodeID(), "")
	oc = 0x8235
	assert.Equal(t, uint16(0x8005), oc.TakeOpcodeID(), "")
	oc = 0xE2A1
	assert.Equal(t, uint16(0xE0A1), oc.TakeOpcodeID(), "")
	oc = 0xF333
	assert.Equal(t, uint16(0xF033), oc.TakeOpcodeID(), "")
	oc = 0x00E0
	assert.Equal(t, uint16(0x00E0), oc.TakeOpcodeID(), "")
	oc = 0x00EE
	assert.Equal(t, uint16(0x00EE), oc.TakeOpcodeID(), "")

}

func TestOpcode_X(t *testing.T) {
	var oc opcode
	oc = 0x3A22
	assert.Equal(t, uint8(0xA), oc.X(), "")
	oc = 0x8220
	assert.Equal(t, uint8(0x2), oc.X(), "")
	oc = 0xE3A1
	assert.Equal(t, uint8(0x3), oc.X(), "")
	oc = 0xFB33
	assert.Equal(t, uint8(0xB), oc.X(), "")

}

func TestOpcode_Y(t *testing.T) {
	var oc opcode
	oc = 0x3A22
	assert.Equal(t, uint8(0x2), oc.Y(), "")
	oc = 0x82B0
	assert.Equal(t, uint8(0xB), oc.Y(), "")
	oc = 0xD3A1
	assert.Equal(t, uint8(0xA), oc.Y(), "")
	oc = 0x9B30
	assert.Equal(t, uint8(0x3), oc.Y(), "")

}

func TestOpcode_KK(t *testing.T) {
	var oc opcode
	oc = 0x6922
	assert.Equal(t, uint8(0x22), oc.KK(), "")
	oc = 0x7321
	assert.Equal(t, uint8(0x21), oc.KK(), "")
	oc = 0x33A1
	assert.Equal(t, uint8(0xA1), oc.KK(), "")

}

func TestOpcode_NNN(t *testing.T) {
	var oc opcode
	oc = 0x1AAA
	assert.Equal(t, uint16(0xAAA), oc.NNN(), "")
	oc = 0xABFA
	assert.Equal(t, uint16(0xBFA), oc.NNN(), "")
	oc = 0xBFFF
	assert.Equal(t, uint16(0xFFF), oc.NNN(), "")
	oc = 0x100B
	assert.Equal(t, uint16(0x00B), oc.NNN(), "")
	oc = 0x2070
	assert.Equal(t, uint16(0x070), oc.NNN(), "")
	oc = 8304
	assert.Equal(t, uint16(0x070), oc.NNN(), "")

}

func TestOpcode_N(t *testing.T) {
	var oc opcode
	oc = 0xDAFD
	assert.Equal(t, uint16(0xD), oc.N(), "")
	oc = 0xD01A
	assert.Equal(t, uint16(0xA), oc.N(), "")
	oc = 0xD3D0
	assert.Equal(t, uint16(0x0), oc.N(), "")
}

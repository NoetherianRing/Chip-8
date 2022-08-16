package chip8

import (
	"encoding/json"
	"github.com/NoetherianRing/Chip-8/config"
	"github.com/NoetherianRing/Chip-8/state"
	"github.com/stretchr/testify/assert"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"
)

func ObtainConfig() config.Config {
	f, err := os.Open("../config.yml")
	defer f.Close()
	if err != nil {
		panic(err)
	}
	decoder := yaml.NewDecoder(f)
	var cfg config.Config

	_ = decoder.Decode(&cfg)

	return cfg
}

func ObtainExpectedState(path string) []byte {
	absPath, _ := filepath.Abs(path)

	expectedResultROM, _ := ioutil.ReadFile(absPath)

	return expectedResultROM
}

func TestChip8_LoadROM(t *testing.T) {
	cfg := ObtainConfig()
	expectedResultROM1 := ObtainExpectedState(cfg.Test.ExpectedStateROM1)
	c8, err := NewChip8()

	assert.NoError(t, err, "error in NewChip8")

	var expected []state.StateChip8
	_ = json.Unmarshal(expectedResultROM1, &expected)

	absPath, _ := filepath.Abs(cfg.Test.ROM1)

	err = c8.LoadROM(absPath)
	assert.NoError(t, err, "error in LoadROM")

	assert.Equal(t, expected[0].Memory[0x200:], c8.memory[0x200:], "ROM1")

}

func TestChip8_LoadFonts(t *testing.T) {
	cfg := ObtainConfig()
	expectedResultROM1 := ObtainExpectedState(cfg.Test.ExpectedStateROM1)
	c8, err := NewChip8()

	assert.NoError(t, err, "error in NewChip8")
	var expected []state.StateChip8
	_ = json.Unmarshal(expectedResultROM1, &expected)
	absPath, _ := filepath.Abs(cfg.Test.FONT)

	err = c8.LoadFonts(absPath)
	assert.NoError(t, err, "error in LoadFonts")
	assert.Equal(t, expected[0].Memory[:0x200], c8.memory[:0x200], "ROM3")
}

func TestChip8_Cycle(t *testing.T) {
	cfg := ObtainConfig()
	testCycle(t, cfg.Test.ExpectedStateROM1, cfg.Test.ROM1, cfg.Test.FONT)
	testCycle(t, cfg.Test.ExpectedStateROM2, cfg.Test.ROM2, cfg.Test.FONT)
	testCycle(t, cfg.Test.ExpectedStateROM3, cfg.Test.ROM3, cfg.Test.FONT)

}

func testCycle(t *testing.T, pathExpectedFile string, pathRom string, pathFont string) {
	expectedResultROM := ObtainExpectedState(pathExpectedFile)
	c8, err := NewChip8()
	assert.NoError(t, err, "error in NewChip8")

	absPathROM, _ := filepath.Abs(pathRom)
	err = c8.LoadROM(absPathROM)
	assert.NoError(t, err, "error in LoadROM")

	absPathFonts, _ := filepath.Abs(pathFont)
	err = c8.LoadFonts(absPathFonts)
	assert.NoError(t, err, "error in LoadFonts")

	var expected []state.StateChip8
	_ = json.Unmarshal(expectedResultROM, &expected)

	for i := 0; i < len(expected); i++ {
		c8.Cycle()
		assert.Equal(t, expected[i].Memory, c8.memory, "ERROR IN MEMORY")
		assert.Equal(t, expected[i].Registers, c8.registers, "ERROR IN REGISTERS")
		assert.Equal(t, expected[i].Pc, c8.pc, "ERROR IN PC")
		assert.Equal(t, expected[i].I, c8.i, "ERROR IN I")
		assert.Equal(t, expected[i].Stack, c8.stack, "ERROR IN STACK")
		assert.Equal(t, expected[i].Sp, c8.sp, "ERROR IN SP")
		assert.Equal(t, expected[i].COpcode, uint16(c8.cOpcode), "ERROR IN COPCODE")
		assert.Equal(t, expected[i].Keypad, c8.Keypad, "ERROR IN KEYPAD")
		assert.Equal(t, expected[i].FrameBuffer, c8.frameBuffer, "ERROR IN FRAMEBUFFER")
		assert.Equal(t, expected[i].DelayTimer, c8.delayTimer, "ERROR IN DELAYTIMER")
		assert.Equal(t, expected[i].SoundTimer, c8.soundTimer, "ERROR IN SOUNDTIMER")

	}
}

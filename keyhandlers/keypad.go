package keyhandlers

import (
	"github.com/faiface/pixel/pixelgl"
	"time"
)

type HexKeypad [16]byte

type KeyHandler interface {
	ExecuteInputs(chan bool)
}
type keyHandler struct {
	*pixelgl.Window
	cmd *Cmd
}

type Cmd map[pixelgl.Button]func()

//KeyboardToKeypad maps the keys a computer keyboard to the keys of a chip8 keypad, following the conversion:
//	Computer Keyboard  Keypad
//	 |1|2|3|4|        |1|2|3|C|
//	|Q|W|E|R|         |4|5|6|D|
//	|A|S|D|F|         |7|8|9|E|
//	|Z|X|C|V|         |A|0|B|F|
var KeyboardToKeypad = map[pixelgl.Button]byte{
	pixelgl.Key1: 1,
	pixelgl.Key2: 2,
	pixelgl.Key3: 3,
	pixelgl.Key4: 0xC,
	pixelgl.KeyQ: 4,
	pixelgl.KeyW: 5,
	pixelgl.KeyE: 6,
	pixelgl.KeyR: 0xD,
	pixelgl.KeyA: 7,
	pixelgl.KeyS: 8,
	pixelgl.KeyD: 9,
	pixelgl.KeyF: 0xE,
	pixelgl.KeyZ: 0xA,
	pixelgl.KeyX: 0,
	pixelgl.KeyC: 0xB,
	pixelgl.KeyV: 0xF,
}

//NewKeyHandler receives a Window to embed, and a map with keys to handler
func NewKeyHandler(window *pixelgl.Window, cmd *Cmd) KeyHandler {
	keyHandler := new(keyHandler)
	keyHandler.Window = window
	keyHandler.cmd = cmd
	return keyHandler
}

//ExecuteInputs checks which keys of the command map have been pressed and executes them
func (kHandler *keyHandler) ExecuteInputs(out chan bool) {
	for true{
		for key, c := range *kHandler.cmd {
			if kHandler.JustPressed(key) {
				c()

				out <- true
				time.Sleep((time.Second / time.Duration(500))*2)
			}

		}
	}

}

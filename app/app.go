package app

import (
	"encoding/json"
	"github.com/NoetherianRing/Chip-8/chip8"
	"github.com/NoetherianRing/Chip-8/config"
	"github.com/NoetherianRing/Chip-8/keyhandlers"
	"github.com/NoetherianRing/Chip-8/monitor"
	"github.com/NoetherianRing/Chip-8/state"
	"github.com/faiface/beep"
	"github.com/faiface/beep/mp3"
	"github.com/faiface/beep/speaker"
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
	"io/ioutil"
	"os"
	"path/filepath"
	"time"
)

type App struct {
	c8           *chip8.Chip8
	keypad       keyhandlers.KeyHandler
	keyboard     keyhandlers.KeyHandler
	m            monitor.Monitor
	beepFile     *os.File
	beepStreamer beep.StreamSeekCloser
	cfg          config.Config
	window       *pixelgl.Window
	channel 	 chan byte
}

//NewApp instantiates the App in which the chip8 is going to run.
//It contains a chip8, a configuration, a beepFile and a beepStream to manage the sound,
//a pixelgl window which is used for all the peripherals,
//and the peripherals: a monitor(m) which draws the FrameBuffer of the chip 8, a keypad  which manages all the inputs of the chip8,
//and a keyboard which manages the inputs of the app (in this case we only use it to quit when we press Esc., a key which is not used by chip8 ROM files).
//NewApp also load the fonts file given in the configuration.
func NewApp(cfg config.Config) (*App, error) {
	myApp := new(App)
	var err error
	myApp.cfg = cfg
	myApp.channel = make(chan byte, 4)
	myApp.c8, err = chip8.NewChip8(myApp.channel)
	if err != nil {
		return nil, err
	}

	cfgPixel := pixelgl.WindowConfig{
		Title:       "Chip-8",
		Bounds:      pixel.R(0, 0, monitor.WidthScreen, monitor.HeightScreen),
		VSync:       true,
		Undecorated: true,
	}

	myApp.window, err = pixelgl.NewWindow(cfgPixel)
	if err != nil {
		return nil, err
	}

	absPathBeep, err := filepath.Abs(cfg.Paths.Beep)

	if err != nil {
		panic(err)
	}
	myApp.m = monitor.NewMonitor(myApp.window)
	myApp.beepFile, err = os.Open(absPathBeep)
	if err != nil {
		return nil, err
	}

	var format beep.Format
	myApp.beepStreamer, format, err = mp3.Decode(myApp.beepFile)

	if err != nil {
		return nil, err
	}

	_ = speaker.Init(
		format.SampleRate,
		format.SampleRate.N(time.Second/10),
	)


	cmdKeypad := make(keyhandlers.Cmd)

	for k, v := range keyhandlers.KeyboardToKeypad {
		newKey := v
		cmdKeypad[k] = func() {
				myApp.channel <- newKey
		}
	}
	myApp.keypad = keyhandlers.NewKeyHandler(myApp.window, &cmdKeypad)

	cmdKeyboard := make(keyhandlers.Cmd)
	cmdKeyboard[pixelgl.KeyEscape] = func() {
		myApp.c8.Close()
		defer myApp.beepFile.Close()
		defer myApp.beepStreamer.Close()
		myApp.channel <- chip8.AsciiEscape

	}
	myApp.keyboard = keyhandlers.NewKeyHandler(myApp.window, &cmdKeyboard)

	absPathFonts, err := filepath.Abs(cfg.Paths.Fonts)
	if err != nil {
		panic(err)
	}
	err = myApp.c8.LoadFonts(absPathFonts)

	if err != nil {
		panic(err)
	}

	return myApp, nil
}

//Run loads the ROM given in the configuration into the chip8,
//then runs the chip8 making a distinction if the configuration indicates whether the application should run in debug mode.
func (myApp *App) Run() {
	absPathRom, err := filepath.Abs(myApp.cfg.Paths.Rom)
	if err != nil {
		panic(err)
	}

	err = myApp.c8.LoadROM(absPathRom)
	if err != nil {
		panic(err)
	}
	if myApp.cfg.Debug.On == "true" {
		myApp.debugChip8()
	} else {
		myApp.runChip8()
	}
}

//runChip8 executes the chip8 Cycle with a certain frequency and manages the peripherals.
func (myApp *App) runChip8() {
	go myApp.keyboard.ExecuteInputs()
	go myApp.keypad.ExecuteInputs()
	go myApp.cycle()
	myApp.update()


}

func (myApp *App) cycle(){

	clock := time.NewTicker(chip8.Frequency)
	for !myApp.c8.IsClosed() {
		select {
		case <-clock.C:
			{
				myApp.c8.Cycle()


			}
		}

	}
}

//debugChip8 does the same that runChip8 with the distinction it use another frequency and save the state of the chip8 in every cycle
//to store it into a json file
func (myApp *App) debugChip8() {

	_, err := os.Create(myApp.cfg.Debug.File)

	if err != nil {
		panic(err)
	}


	go myApp.cycleDebug()
	for {
		myApp.update()

	}

}
func (myApp *App) cycleDebug(){
	clock := time.NewTicker(chip8.FrequencyDebugMode)
	var sChip8 []state.StateChip8

	for {
		select {
		case <-clock.C:
			{
				myApp.c8.Cycle()
				if myApp.c8.IsClosed() {
					return
				}
				sChip8 = append(sChip8, *myApp.c8.Dump())
				stateBytes, err := json.Marshal(sChip8)

				if err != nil {
					panic(err)
				}

				err = ioutil.WriteFile(myApp.cfg.Debug.File, stateBytes, 0644)
				if err != nil {
					panic(err)
				}
			}
		}
	}
}

//update draws and beeps if it's needed, and executes the inputs.
func (myApp *App) update() {
	clock := time.NewTicker(chip8.Frequency)

	for !myApp.c8.IsClosed() {
		select {
		case <-clock.C:
			{
				if myApp.c8.MustDraw {
					myApp.c8.MustDraw = false

					myApp.m.ToDraw(myApp.c8.GetFrameBuffer())
				}
				if myApp.c8.MustBeep() {
					speaker.Play(myApp.beepStreamer)
					_ = myApp.beepStreamer.Seek(0)

				}

				myApp.window.Update()

			}

		}
	}
}
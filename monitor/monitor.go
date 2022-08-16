package monitor

import (
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/imdraw"
	"github.com/faiface/pixel/pixelgl"
	"golang.org/x/image/colornames"
)

const (
	SidePixel    = 16
	WidthScreen  = SidePixel * 64
	HeightScreen = SidePixel * 32
)

type FrameBuffer [64 * 32]byte

//CheckOverlap checks if the given x, y are outside the bounds of the array
func (f *FrameBuffer) CheckOverlap(x, y int) bool {
	if y*64+x > len(f) {
		return true
	} else {
		return false
	}
}

//Get returns the cell of the FrameBuffer corresponded to the (x,y) coordinate of the 64x32 display
func (f *FrameBuffer) Get(x, y int) *byte {
	return &f[y*64+x]
}

type Monitor interface {
	ToDraw(buffer FrameBuffer)
}

type monitor struct {
	*pixelgl.Window
}

func NewMonitor(window *pixelgl.Window) Monitor {
	m := new(monitor)
	m.Window = window
	return m
}

//ToDraw reads the FrameBuffer of the chip8.
//Every element in FrameBuffer represents a pixel on the screen which can be on or off.
//If it's on ToDraw draws a 16x16 "pixel" on the screen
func (m *monitor) ToDraw(buffer FrameBuffer) {
	m.Clear(colornames.Black)
	imd := imdraw.New(nil)
	imd.Color = pixel.RGB(1, 1, 1)

	//Chip8 has a coordinate system in which the (0,0) is at the upper left corner of the screen
	//Pixelgls a coordinate system in which the (0,0) is at the lower left corner of the screen
	//that's why we get the element (x, 31-y) of the buffer instead of the element (x,y)
	for y := 0; y < 32; y++ {
		for x := 0; x < 64; x++ {
			if *buffer.Get(x, 31-y) != 0 {
				imd.Push(pixel.V(SidePixel*float64(x), SidePixel*float64(y)))
				imd.Push(pixel.V(SidePixel*float64(x)+SidePixel, SidePixel*float64(y)+SidePixel))
				imd.Rectangle(0)
			}
		}
	}
	imd.Draw(m)

}

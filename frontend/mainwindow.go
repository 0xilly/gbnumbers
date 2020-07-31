package frontend

import (
	"log"

	"github.com/veandco/go-sdl2/sdl"
)

func getPixelRes() int {
	return 10
}

func getWindowSize() (int32, int32) {
	return 500, 500
}

type draw struct {
	*sdl.Surface
	size int32
}

type pixel struct {
	size, x, y int32
	color      uint32
}

// InitSdl initializes sdl
func InitSdl() {
	if err := sdl.Init(sdl.INIT_EVERYTHING); err != nil {
		log.Panic(err)
	}

	defer sdl.Quit()

	width, height := getWindowSize()

	window, err := sdl.CreateWindow("Gb Numbers", sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED, width, height, sdl.WINDOW_SHOWN)

	if err != nil {
		log.Panic(err)
	}

	defer window.Destroy()

	surface, err := window.GetSurface()

	if err != nil {
		log.Panicln(err)
	}

	dr := draw{surface, 10}

	dr.drawPixel(0, 0, 0xFFFFFF)

	for i := 0; i < 51; i++ {
		dr.drawPixel(i*10, i*10, 0xFFFFFF)
	}

	window.UpdateSurface()

	running := true

	for running {
		for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
			switch event.(type) {
			case *sdl.QuitEvent:
				running = false
				break
			}
		}
	}

}

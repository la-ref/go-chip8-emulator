package main

import "github.com/veandco/go-sdl2/sdl"

const (
	WIN_TITLE  = "CHIP8 Emulator"
	WIN_WIDTH  = 64
	WIN_HEIGHT = 32
)

func main() {
	config := NewAppConfig(WIN_TITLE, WIN_HEIGHT, WIN_WIDTH)
	app, err := NewApp(config)
	if err != nil {
		panic(err)
	}
	defer app.HandleQuit()
	window := app.window
	surface, err := window.GetSurface()
	if err != nil {
		panic(err)
	}
	surface.FillRect(nil, 0)

	rect := sdl.Rect{0, 0, 200, 200}
	colour := sdl.Color{R: 255, G: 0, B: 255, A: 255} // purple
	pixel := sdl.MapRGBA(surface.Format, colour.R, colour.G, colour.B, colour.A)
	surface.FillRect(&rect, pixel)
	window.UpdateSurface()

	running := true
	for running {
		for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
			switch event.(type) {
			case *sdl.QuitEvent:
				println("Quit")
				running = false
				break
			}
		}
	}
}

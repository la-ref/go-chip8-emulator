package main

import (
	application "emulator/app"
	config2 "emulator/config"
	"fmt"
	"github.com/veandco/go-sdl2/sdl"
	"time"
)

const (
	WIN_TITLE    = "CHIP8 Emulator"
	WIN_WIDTH    = 64
	WIN_HEIGHT   = 32
	FG_COLOR     = 0xffffffff
	BG_COLOR     = 0x00000000
	TARGET_FPS   = 60
	TARGET_FRAME = 1000 / TARGET_FPS
)

func main() {
	config := config2.NewAppConfig(WIN_TITLE, WIN_HEIGHT, WIN_WIDTH, FG_COLOR, BG_COLOR)
	app, err := application.NewApp(config)
	if err != nil {
		panic(err)
	}
	defer app.HandleQuit()
	window := app.GetWindow()
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

	previousTime := sdl.GetTicks()
	running := true
	for running {
		time.Sleep(TARGET_FRAME)
		currentTime := sdl.GetTicks()
		deltaTime := currentTime - previousTime
		fmt.Println(currentTime, previousTime, deltaTime)
		previousTime = currentTime
		for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
			switch event.(type) {
			case *sdl.QuitEvent:
				println("Quit")
				running = false
				break
			}
		}
		frameTime := sdl.GetTicks() - currentTime
		if frameTime < TARGET_FRAME {
			sdl.Delay(TARGET_FRAME - frameTime)
		}
	}
}

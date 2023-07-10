package main

import (
	application "emulator/app"
	config2 "emulator/config"
	"emulator/utils"
	"fmt"
	"github.com/veandco/go-sdl2/sdl"
)

const (
	WIN_TITLE    = "CHIP8 Emulator"
	WIN_WIDTH    = 64
	WIN_HEIGHT   = 32
	SCALE_FACTOR = 20 // Scale for each pixel
	FG_COLOR     = 0xffffffff
	BG_COLOR     = 0x00000000
	TARGET_FPS   = 60
	TARGET_FRAME = 1000 / TARGET_FPS
)

func handleEvents(app *application.App) {
	for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
		switch e := event.(type) {
		case *sdl.QuitEvent:
			println("Quit")
			app.SetState(utils.STOPPED)
			break
		case *sdl.KeyboardEvent:
			if e.Type == sdl.KEYDOWN {
				keyCode := e.Keysym.Sym
				fmt.Println(keyCode)
			}

		}
	}
}

func main() {
	config := config2.NewAppConfig(WIN_TITLE, WIN_HEIGHT, WIN_WIDTH, SCALE_FACTOR, FG_COLOR, BG_COLOR)
	app, err := application.NewApp(config)
	if err != nil {
		panic(err)
	}
	defer app.HandleQuit()

	err = app.ClearScreen()
	if err != nil {
		return
	}

	previousTime := sdl.GetTicks()
	for app.GetState() != utils.STOPPED {

		handleEvents(app)

		currentTime := sdl.GetTicks()
		deltaTime := currentTime - previousTime

		app.Update(deltaTime)

		previousTime = currentTime
		frameTime := sdl.GetTicks() - previousTime
		if frameTime < TARGET_FRAME {
			sdl.Delay(TARGET_FRAME - frameTime)
		}
	}
}

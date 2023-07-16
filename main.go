package main

import (
	application "emulator/app"
	chip "emulator/chip8"
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
				switch e.Keysym.Sym {
				case sdl.K_ESCAPE:
					app.SetState(utils.STOPPED)
					break
				case sdl.K_SPACE:
					if app.GetState() == utils.RUNNING {
						app.SetState(utils.PAUSED)
					} else {
						app.SetState(utils.RUNNING)
					}
				}
			}

		}
	}
}

func main() {
	config := config2.NewAppConfig(WIN_TITLE, WIN_HEIGHT, WIN_WIDTH, SCALE_FACTOR, FG_COLOR, BG_COLOR)
	fmt.Println(config.GetWinWidth(), WIN_WIDTH)
	chip8, err := chip.NewChip8("./rom/IBM.ch8", config)
	if err != nil {
		panic(err)
	}
	app, err := application.NewApp(config, chip8)
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
		app.Draw()

		previousTime = currentTime
		frameTime := sdl.GetTicks() - previousTime
		if frameTime < TARGET_FRAME {
			sdl.Delay(TARGET_FRAME - frameTime)
		}
	}
}

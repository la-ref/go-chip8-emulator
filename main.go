package main

import (
	application "emulator/app"
	chip "emulator/chip8"
	config2 "emulator/config"
	"emulator/utils"
	"github.com/veandco/go-sdl2/sdl"
)

const (
	VERSION      = chip.CHIP8 // chip8 version type
	WIN_TITLE    = "CHIP8 Emulator"
	WIN_WIDTH    = 64
	WIN_HEIGHT   = 32
	SCALE_FACTOR = 6 // Scale for each pixel

	FG_COLOR = 0xffffffff // white foreground color
	BG_COLOR = 0x00000000 // black background color

	TARGET_FPS   = 60
	TARGET_FRAME = 1000 / TARGET_FPS
	CLOCK_RATE   = 700 // number of chip8 instruction cycles per second

	VOLUME           = 100 // volume [0;100]
	SQUARE_WAVE_FREQ = 440
	SAMPLE_RATE      = 44100

	KEYBOARD = "AZERTY"

	ROM = "./rom/tetris.ch8"
)

func main() {
	config := config2.NewAppConfig(uint8(VERSION), WIN_TITLE, ROM, WIN_HEIGHT, WIN_WIDTH, SCALE_FACTOR, FG_COLOR, BG_COLOR, CLOCK_RATE, SQUARE_WAVE_FREQ, SAMPLE_RATE, VOLUME, KEYBOARD)
	chip8, err := chip.NewChip8(config)
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

		app.HandleEvents()

		if app.GetState() != utils.PAUSED {
			currentTime := sdl.GetTicks()
			deltaTime := currentTime - previousTime

			app.Update(float32(deltaTime) / 1000)
			app.Draw()

			previousTime = currentTime
			frameTime := sdl.GetTicks() - previousTime
			if frameTime < TARGET_FRAME {
				sdl.Delay(TARGET_FRAME - frameTime)
			}
		}
	}
}

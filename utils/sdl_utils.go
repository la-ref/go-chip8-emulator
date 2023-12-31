package utils

import (
	config "emulator/config"
	"github.com/veandco/go-sdl2/sdl"
)

type State int

const (
	STOPPED State = iota
	RUNNING
	PAUSED
)

func InitSdl() error {
	err := sdl.Init(sdl.INIT_VIDEO | sdl.INIT_AUDIO | sdl.INIT_TIMER)
	if err != nil {
		return err
	}
	return nil
}

func CreateWindow(config *config.AppConfig) (*sdl.Window, error) {
	window, err := sdl.CreateWindow(config.GetWinTitle(), sdl.WINDOWPOS_CENTERED, sdl.WINDOWPOS_CENTERED,
		config.GetWinWidth(), config.GetWinHeight(), sdl.WINDOW_SHOWN)
	if err != nil {
		return nil, err
	}
	return window, nil
}

func CreateRenderer(window *sdl.Window) (*sdl.Renderer, error) {
	renderer, err := sdl.CreateRenderer(window, -1, sdl.RENDERER_ACCELERATED)
	if err != nil {
		return nil, err
	}
	return renderer, nil
}

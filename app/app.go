package app

import (
	"emulator/chip8"
	conf "emulator/config"
	utils "emulator/utils"
	"github.com/veandco/go-sdl2/sdl"
)

type App struct {
	window   *sdl.Window
	renderer *sdl.Renderer
	config   *conf.AppConfig
	state    utils.State

	chip8   *chip8.Chip8
	keybind map[sdl.Keycode]byte
}

func NewApp(config *conf.AppConfig, chip *chip8.Chip8) (*App, error) {
	err := utils.InitSdl()
	if err != nil {
		return nil, err
	}
	window, err := utils.CreateWindow(config)
	if err != nil {
		return nil, err
	}
	renderer, err := utils.CreateRenderer(window)
	if err != nil {
		return nil, err
	}
	app := &App{
		window:   window,
		config:   config,
		renderer: renderer,
		state:    utils.RUNNING,
		chip8:    chip,
		keybind:  chip8.AzertyKeyMap,
	}
	return app, nil
}

func (a *App) GetWindow() *sdl.Window {
	return a.window
}

func (a *App) GetRenderer() *sdl.Renderer {
	return a.renderer
}

func (a *App) GetConfig() *conf.AppConfig {
	return a.config
}

func (a *App) GetState() utils.State {
	return a.state
}

func (a *App) SetState(state utils.State) {
	a.state = state
}

func (a *App) HandleQuit() {
	a.renderer.Destroy()
	a.window.Destroy()
	sdl.Quit()
}

func (a *App) ClearScreen() error {
	r, g, b, alpha := utils.BytesToRGBA(a.config.GetBgColor())
	err := a.renderer.SetDrawColor(r, g, b, alpha)
	if err != nil {
		return err
	}
	err = a.renderer.Clear()
	if err != nil {
		return err
	}
	return nil
}

func (a *App) Draw() {
	_ = a.ClearScreen()
	a.chip8.Draw(a.renderer)
	a.renderer.Present()
}

func (a *App) Update(dt float32) {
	a.chip8.Update(dt)
}

func (a *App) handleInputs(e sdl.Keycode, status bool) bool {
	switch e {
	case sdl.K_ESCAPE:
		a.SetState(utils.STOPPED)
		return true
	case sdl.K_SPACE:
		if a.GetState() == utils.RUNNING {
			a.SetState(utils.PAUSED)
		} else {
			a.SetState(utils.RUNNING)
		}
	default:
		a.chip8.Keypad[a.keybind[e]] = status
	}
	return false
}

func (a *App) HandleEvents() {
	for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
		switch e := event.(type) {
		case *sdl.QuitEvent:
			a.SetState(utils.STOPPED)
			break
		case *sdl.KeyboardEvent:
			if e.Type == sdl.KEYDOWN {
				if a.handleInputs(e.Keysym.Sym, true) {
					break
				}
			} else if e.Type == sdl.KEYUP {
				if a.handleInputs(e.Keysym.Sym, false) {
					break
				}
			}

		}
	}
}

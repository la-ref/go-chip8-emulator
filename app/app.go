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

	chip8 *chip8.Chip8
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
	a.chip8.Draw(a.renderer)
	a.renderer.Present()
}

func (a *App) Update(dt uint32) {
	a.chip8.Update(dt)
}

func (a *App) handleInputs(e sdl.Keycode) bool {
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
	case sdl.K_1:
		a.chip8.Keypad[0x1] = true
	case sdl.K_2:
		a.chip8.Keypad[0x2] = true
	case sdl.K_3:
		a.chip8.Keypad[0x3] = true
	case sdl.K_4:
		a.chip8.Keypad[0xC] = true

	case sdl.K_a:
		a.chip8.Keypad[0x4] = true
	case sdl.K_z:
		a.chip8.Keypad[0x5] = true
	case sdl.K_e:
		a.chip8.Keypad[0x6] = true
	case sdl.K_r:
		a.chip8.Keypad[0xD] = true

	case sdl.K_q:
		a.chip8.Keypad[0x7] = true
	case sdl.K_s:
		a.chip8.Keypad[0x8] = true
	case sdl.K_d:
		a.chip8.Keypad[0x9] = true
	case sdl.K_f:
		a.chip8.Keypad[0xE] = true

	case sdl.K_w:
		a.chip8.Keypad[0xA] = true
	case sdl.K_x:
		a.chip8.Keypad[0x0] = true
	case sdl.K_c:
		a.chip8.Keypad[0xB] = true
	case sdl.K_v:
		a.chip8.Keypad[0xF] = true
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
				if a.handleInputs(e.Keysym.Sym) {
					break
				}
			}

		}
	}
}

package app

import (
	conf "emulator/config"
	utils "emulator/utils"
	"github.com/veandco/go-sdl2/sdl"
)

type App struct {
	window   *sdl.Window
	renderer *sdl.Renderer
	config   *conf.AppConfig
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

func (a *App) Update() {
	a.renderer.Present()
}

func NewApp(config *conf.AppConfig) (*App, error) {
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
	}
	return app, nil
}

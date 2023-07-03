package main

import (
	"github.com/veandco/go-sdl2/sdl"
)

type App struct {
	window   *sdl.Window
	renderer *sdl.Renderer
	config   *AppConfig
}

type AppConfig struct {
	winHeight int32
	winWidth  int32
	winTitle  string
}

func (a *App) HandleQuit() {
	a.renderer.Destroy()
	a.window.Destroy()
	sdl.Quit()
}

func NewAppConfig(title string, height, width int32) *AppConfig {
	return &AppConfig{
		winHeight: height,
		winWidth:  width,
		winTitle:  title,
	}
}

func NewApp(config *AppConfig) (*App, error) {
	err := initSdl()
	if err != nil {
		return nil, err
	}
	window, err := createWindow(config)
	if err != nil {
		return nil, err
	}
	renderer, err := createRenderer(window)
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

func initSdl() error {
	err := sdl.Init(sdl.INIT_VIDEO | sdl.INIT_AUDIO | sdl.INIT_TIMER)
	if err != nil {
		return err
	}
	return nil
}

func createWindow(config *AppConfig) (*sdl.Window, error) {
	window, err := sdl.CreateWindow(config.winTitle, sdl.WINDOWPOS_CENTERED, sdl.WINDOWPOS_CENTERED,
		config.winWidth, config.winHeight, sdl.WINDOW_SHOWN)
	if err != nil {
		return nil, err
	}
	return window, nil
}

func createRenderer(window *sdl.Window) (*sdl.Renderer, error) {
	renderer, err := sdl.CreateRenderer(window, -1, sdl.RENDERER_ACCELERATED)
	if err != nil {
		return nil, err
	}
	return renderer, nil
}

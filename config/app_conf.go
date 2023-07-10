package config

type AppConfig struct {
	winHeight int32
	winWidth  int32
	winTitle  string
	fgColor   uint32
	bgColor   uint32
}

func (a AppConfig) GetWinHeight() int32 {
	return a.winHeight
}

func (a AppConfig) GetWinWidth() int32 {
	return a.winWidth
}

func (a AppConfig) GetWinTitle() string {
	return a.winTitle
}

func (a AppConfig) GetFgColor() uint32 {
	return a.fgColor
}

func (a AppConfig) GetBgColor() uint32 {
	return a.bgColor
}

func NewAppConfig(title string, height, width, scale int32, fg, bg uint32) *AppConfig {
	return &AppConfig{
		winHeight: height * scale,
		winWidth:  width * scale,
		winTitle:  title,
		fgColor:   fg,
		bgColor:   bg,
	}
}

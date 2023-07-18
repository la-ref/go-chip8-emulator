package config

type AppConfig struct {
	version     uint8
	winHeight   int32
	winWidth    int32
	scale       int32
	winTitle    string
	fgColor     uint32
	bgColor     uint32
	clockRate   uint32
	rom         string
	audioConfig *AudioConfig
}

func (a *AppConfig) GetVersion() uint8 {
	return a.version
}

func (a *AppConfig) GetWinHeight() int32 {
	return a.winHeight
}

func (a *AppConfig) GetWinWidth() int32 {
	return a.winWidth
}

func (a *AppConfig) GetScale() int32 {
	return a.scale
}

func (a *AppConfig) GetWinTitle() string {
	return a.winTitle
}

func (a *AppConfig) GetFgColor() uint32 {
	return a.fgColor
}

func (a *AppConfig) GetBgColor() uint32 {
	return a.bgColor
}

func (a *AppConfig) GetClockRate() uint32 {
	return a.clockRate
}

func (a *AppConfig) GetAudioConfig() *AudioConfig {
	return a.audioConfig
}

func (a *AppConfig) GetRom() string {
	return a.rom
}

func NewAppConfig(version uint8, title, rom string, height, width, scale int32, fg, bg, cr, freq, sample uint32, volume uint8) *AppConfig {
	return &AppConfig{
		version:     version,
		winHeight:   height * scale,
		winWidth:    width * scale,
		scale:       scale,
		winTitle:    title,
		fgColor:     fg,
		bgColor:     bg,
		clockRate:   cr,
		audioConfig: NewAudioConfig(volume, freq, sample),
		rom:         rom,
	}
}

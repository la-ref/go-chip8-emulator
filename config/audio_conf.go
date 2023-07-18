package config

type AudioConfig struct {
	volume         uint8
	squareWaveFreq uint32
}

func NewAudioConfig(vol uint8, freq uint32) *AudioConfig {
	return &AudioConfig{
		volume:         vol,
		squareWaveFreq: freq,
	}
}

func (a *AudioConfig) GetVolume() uint8 {
	return a.volume
}

func (a *AudioConfig) GetSquareWaveFreq() uint32 {
	return a.squareWaveFreq
}

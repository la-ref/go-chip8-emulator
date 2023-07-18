package config

// Defined also in audio/audio.h
type AudioConfig struct {
	volume         uint8
	squareWaveFreq uint32
	sampleRate     uint32
}

func NewAudioConfig(vol uint8, freq, sample uint32) *AudioConfig {
	return &AudioConfig{
		volume:         vol,
		squareWaveFreq: freq,
		sampleRate:     sample,
	}
}

func (a *AudioConfig) GetVolume() uint8 {
	return a.volume
}

func (a *AudioConfig) GetSquareWaveFreq() uint32 {
	return a.squareWaveFreq
}

func (a *AudioConfig) GetSampleRate() uint32 {
	return a.sampleRate
}

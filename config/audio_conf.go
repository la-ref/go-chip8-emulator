package config

// #include <stdint.h>
// #include <stdlib.h>
// typedef struct {
//     uint8_t volume;
//     uint32_t squareWaveFreq;
//     uint32_t sampleRate;
// } AudioConfig;
import "C"

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

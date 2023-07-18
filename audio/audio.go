package audio

// #include "audio.h"
import "C"
import (
	conf "emulator/config"
	"github.com/veandco/go-sdl2/sdl"
	"reflect"
	"unsafe"
)

type Audio struct {
	WantSpec *sdl.AudioSpec
	HaveSpec *sdl.AudioSpec
	Device   sdl.AudioDeviceID
}

type UserData struct {
	Config   conf.AudioConfig
	HaveSpec sdl.AudioSpec
}

func NewAudio(audioConf *conf.AudioConfig) (*Audio, error) {
	audio := &Audio{
		WantSpec: &sdl.AudioSpec{
			Freq:     int32(audioConf.GetSampleRate()),
			Format:   sdl.AUDIO_S16LSB, // little endian
			Channels: 1,
			Samples:  512,
		},
	}

	cAudioConf := (*C.AudioConfig)(C.malloc(C.size_t(unsafe.Sizeof(C.AudioConfig{}))))
	cAudioConf.volume = C.uint8_t(audioConf.GetVolume())
	cAudioConf.squareWaveFreq = C.uint32_t(audioConf.GetSquareWaveFreq())
	cAudioConf.sampleRate = C.uint32_t(audioConf.GetSampleRate())
	defer C.free(unsafe.Pointer(cAudioConf))

	audio.WantSpec.UserData = unsafe.Pointer(cAudioConf)
	audio.WantSpec.Callback = sdl.AudioCallback(C.AudioCallback)
	deviceId, err := sdl.OpenAudioDevice("", false, audio.WantSpec, audio.HaveSpec, 0)
	if err != nil || deviceId == 0 {
		return nil, err
	}
	audio.Device = deviceId
	return audio, nil
}

//export AudioCallback
func AudioCallback(userdata unsafe.Pointer, _stream *C.uint8_t, _length C.int) {
	data := (*conf.AudioConfig)(userdata)
	length := int(_length) / 2
	header := reflect.SliceHeader{Data: uintptr(unsafe.Pointer(_stream)), Len: length, Cap: length}
	buf := *(*[]int16)(unsafe.Pointer(&header))

	squareWavePeriod := data.GetSampleRate() / data.GetSquareWaveFreq()
	halfSquareWavePeriod := squareWavePeriod / 2
	for i := range buf {
		vol := 3000 * int16(data.GetVolume()) / 100
		if (i/int(halfSquareWavePeriod))%2 == 0 {
			buf[i] = vol
		} else {
			buf[i] = -vol
		}
	}
}

func (a *Audio) HandleQuit() {
	sdl.CloseAudioDevice(a.Device)
}

func (a *Audio) PlayAudio() {
	sdl.PauseAudioDevice(a.Device, false)
}

func (a *Audio) PauseAudio() {
	sdl.PauseAudioDevice(a.Device, true)
}

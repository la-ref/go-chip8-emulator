package audio

// typedef unsigned char Uint8;
// void AudioCallback(void *userdata, Uint8 *stream, int len);
import "C"
import (
	"github.com/veandco/go-sdl2/sdl"
	"unsafe"
)

type Audio struct {
	WantSpec *sdl.AudioSpec
	HaveSpec *sdl.AudioSpec
	Device   sdl.AudioDeviceID
}

func NewAudio() (*Audio, error) {
	audio := &Audio{
		WantSpec: &sdl.AudioSpec{
			Freq:     44100,
			Format:   sdl.AUDIO_S16LSB, // little endian
			Channels: 1,
			Samples:  4096,
		},
	}
	audio.WantSpec.Callback = sdl.AudioCallback(C.AudioCallback)
	deviceId, err := sdl.OpenAudioDevice("", false, audio.WantSpec, audio.HaveSpec, 0)
	if err != nil || deviceId == 0 {
		return nil, err
	}
	audio.Device = deviceId
	return audio, nil
}

//export AudioCallback
func AudioCallback(userdata unsafe.Pointer, stream *C.Uint8, length C.int) {
	_ = int(length)
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

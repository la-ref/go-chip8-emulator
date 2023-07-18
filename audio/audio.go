package audio

// typedef unsigned char Uint8;
// void AudioCallback(void *userdata, Uint8 *stream, int len);
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

func NewAudio(audioConf *conf.AudioConfig) (*Audio, error) {
	audio := &Audio{
		WantSpec: &sdl.AudioSpec{
			Freq:     44100,
			Format:   sdl.AUDIO_S16LSB, // little endian
			Channels: 1,
			Samples:  4096,
			UserData: unsafe.Pointer(audioConf),
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
func AudioCallback(userdata unsafe.Pointer, _stream *C.Uint8, _length C.int) {
	data := (*conf.AudioConfig)(userdata)

	length := int(_length) / 2
	header := reflect.SliceHeader{Data: uintptr(unsafe.Pointer(_stream)), Len: length, Cap: length}
	buf := *(*[]int16)(unsafe.Pointer(&header))

	for i := range buf {

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

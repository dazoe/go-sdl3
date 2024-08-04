package sdl

/*
#include <SDL3/SDL_audio.h>
#include "audio.h"
#include "stdlib.h"
*/
import "C"
import (
	"fmt"
	"unsafe"
)

// # CategoryAudio
// (https://wiki.libsdl.org/SDL3/CategoryAudio)

// masks for different parts of SDL_AudioFormat.
const AUDIO_MASK_BITSIZE = C.SDL_AUDIO_MASK_BITSIZE
const AUDIO_MASK_FLOAT = C.SDL_AUDIO_MASK_FLOAT
const AUDIO_MASK_BIG_ENDIAN = C.SDL_AUDIO_MASK_BIG_ENDIAN
const AUDIO_MASK_SIGNED = C.SDL_AUDIO_MASK_SIGNED

// func DEFINE_AUDIO_FORMAT(signed, bigendian, float, size uint16) uint16 {
// 	return (((uint16)(signed) << 15) | ((uint16)(bigendian) << 12) | ((uint16)(float) << 8) | ((size) & AUDIO_MASK_BITSIZE))
// }

// Audio format.
// (https://wiki.libsdl.org/SDL3/SDL_AudioFormat)
type AudioFormat C.SDL_AudioFormat

func (f *AudioFormat) cptr() *C.SDL_AudioFormat {
	return (*C.SDL_AudioFormat)(unsafe.Pointer(f))
}
func (f AudioFormat) c() C.SDL_AudioFormat {
	return (C.SDL_AudioFormat)(f)
}

const (
	AUDIO_U8 = AudioFormat(C.SDL_AUDIO_U8) /**< Unsigned 8-bit samples */
	/* SDL_DEFINE_AUDIO_FORMAT(0, 0, 0, 8), */
	AUDIO_S8 = AudioFormat(C.SDL_AUDIO_S8) /**< Signed 8-bit samples */
	/* SDL_DEFINE_AUDIO_FORMAT(1, 0, 0, 8), */
	AUDIO_S16LE = AudioFormat(C.SDL_AUDIO_S16LE) /**< Signed 16-bit samples */
	/* SDL_DEFINE_AUDIO_FORMAT(1, 0, 0, 16), */
	AUDIO_S16BE = AudioFormat(C.SDL_AUDIO_S16BE) /**< As above, but big-endian byte order */
	/* SDL_DEFINE_AUDIO_FORMAT(1, 1, 0, 16), */
	AUDIO_S32LE = AudioFormat(C.SDL_AUDIO_S32LE) /**< 32-bit integer samples */
	/* SDL_DEFINE_AUDIO_FORMAT(1, 0, 0, 32), */
	AUDIO_S32BE = AudioFormat(C.SDL_AUDIO_S32BE) /**< As above, but big-endian byte order */
	/* SDL_DEFINE_AUDIO_FORMAT(1, 1, 0, 32), */
	AUDIO_F32LE = AudioFormat(C.SDL_AUDIO_F32LE) /**< 32-bit floating point samples */
	/* SDL_DEFINE_AUDIO_FORMAT(1, 0, 1, 32), */
	AUDIO_F32BE = AudioFormat(C.SDL_AUDIO_F32BE) /**< As above, but big-endian byte order */
	/* SDL_DEFINE_AUDIO_FORMAT(1, 1, 1, 32), */

	AUDIO_S16 = AudioFormat(C.SDL_AUDIO_S16)
	AUDIO_S32 = AudioFormat(C.SDL_AUDIO_S32)
	AUDIO_F32 = AudioFormat(C.SDL_AUDIO_F32)
)

// Retrieve the size, in bits, from an SDL_AudioFormat.
// (https://wiki.libsdl.org/SDL3/SDL_AUDIO_BITSIZE)
func (x AudioFormat) Bitsize() uint8 {
	return uint8(x & AUDIO_MASK_BITSIZE)
}

// Retrieve the size, in bytes, from an SDL_AudioFormat.
// (https://wiki.libsdl.org/SDL3/SDL_AUDIO_BYTESIZE)
func (x AudioFormat) Bytesize() uint8 {
	return x.Bitsize() / 8
}

// Determine if an SDL_AudioFormat represents floating point data.
// (https://wiki.libsdl.org/SDL3/SDL_AUDIO_ISFLOAT)
func (x AudioFormat) IsFloat() bool {
	return (x & AUDIO_MASK_FLOAT) != 0
}

// Determine if an SDL_AudioFormat represents bigendian data.
// (https://wiki.libsdl.org/SDL3/SDL_AUDIO_ISBIGENDIAN)
func (x AudioFormat) IsBigEndian() bool {
	return (x & AUDIO_MASK_BIG_ENDIAN) != 0
}

// Determine if an SDL_AudioFormat represents littleendian data.
// (https://wiki.libsdl.org/SDL3/SDL_AUDIO_ISLITTLEENDIAN)
func (x AudioFormat) IsLittleEndian() bool {
	return !x.IsBigEndian()
}

// Determine if an SDL_AudioFormat represents signed data.
// (https://wiki.libsdl.org/SDL3/SDL_AUDIO_ISSIGNED)
func (x AudioFormat) IsSigned() bool {
	return (x & AUDIO_MASK_SIGNED) != 0
}

// Determine if an SDL_AudioFormat represents integer data.
// (https://wiki.libsdl.org/SDL3/SDL_AUDIO_ISINT)
func (x AudioFormat) IsInt() bool {
	return !x.IsFloat()
}

// Determine if an SDL_AudioFormat represents unsigned data.
// (https://wiki.libsdl.org/SDL3/SDL_AUDIO_ISUNSIGNED)
func (x AudioFormat) IsUnsigned() bool {
	return !x.IsSigned()
}

// SDL Audio Device instance IDs.
// (https://wiki.libsdl.org/SDL3/SDL_AudioDeviceID)
type AudioDeviceID C.SDL_AudioDeviceID

func (id *AudioDeviceID) cptr() *C.SDL_AudioDeviceID {
	return (*C.SDL_AudioDeviceID)(unsafe.Pointer(id))
}
func (id AudioDeviceID) c() C.SDL_AudioDeviceID {
	return (C.SDL_AudioDeviceID)(id)
}

// A value used to request a default playback audio device.
// (https://wiki.libsdl.org/SDL3/SDL_AUDIO_DEVICE_DEFAULT_PLAYBACK)
const AUDIO_DEVICE_DEFAULT_PLAYBACK = AudioDeviceID(C.SDL_AUDIO_DEVICE_DEFAULT_PLAYBACK)

// A value used to request a default recording audio device.
// (https://wiki.libsdl.org/SDL3/SDL_AUDIO_DEVICE_DEFAULT_RECORDING)
const AUDIO_DEVICE_DEFAULT_RECORDING = AudioDeviceID(C.SDL_AUDIO_DEVICE_DEFAULT_RECORDING)

// Format specifier for audio data.
// (https://wiki.libsdl.org/SDL3/SDL_AudioSpec)
type AudioSpec struct {
	Format   AudioFormat /**< Audio data format */
	Channels int32       /**< Number of channels: 1 mono, 2 stereo, etc */
	Freq     int32       /**< sample rate: sample frames per second */
}

func (spec *AudioSpec) cptr() *C.SDL_AudioSpec {
	return (*C.SDL_AudioSpec)(unsafe.Pointer(spec))
}

// Calculate the size of each audio frame (in bytes) from an SDL_AudioSpec.
// (https://wiki.libsdl.org/SDL3/SDL_AUDIO_FRAMESIZE)
func (x AudioSpec) Framesize() uint32 {
	return uint32(x.Channels) * uint32(x.Format.Bytesize())
}

// The opaque handle that represents an audio stream.
// (https://wiki.libsdl.org/SDL3/SDL_AudioStream)
type AudioStream C.SDL_AudioStream

func (s *AudioStream) cptr() *C.SDL_AudioStream {
	return (*C.SDL_AudioStream)(unsafe.Pointer(s))
}

/* Function prototypes */

// Use this function to get the number of built-in audio drivers.
// (https://wiki.libsdl.org/SDL3/SDL_GetNumAudioDrivers)
func GetNumAudioDrivers() int {
	return int(C.SDL_GetNumAudioDrivers())
}

// Use this function to get the name of a built in audio driver.
// (https://wiki.libsdl.org/SDL3/SDL_GetAudioDriver)
func GetAudioDriver(index int) string {
	return C.GoString(C.SDL_GetAudioDriver(C.int(index)))
}

// Get the name of the current audio driver.
// (https://wiki.libsdl.org/SDL3/SDL_GetCurrentAudioDriver)
func GetCurrentAudioDriver() string {
	return C.GoString(C.SDL_GetCurrentAudioDriver())
}

// Get a list of currently-connected audio playback devices.
// (https://wiki.libsdl.org/SDL3/SDL_GetAudioPlaybackDevices)
func GetAudioPlaybackDevices() ([]AudioDeviceID, error) {
	var count C.int
	ret := C.SDL_GetAudioPlaybackDevices(&count)
	if ret != nil {
		return nil, GetError()
	}
	return unsafe.Slice((*AudioDeviceID)(ret), int(count)), nil
}

// Get a list of currently-connected audio recording devices.
// (https://wiki.libsdl.org/SDL3/SDL_GetAudioRecordingDevices)
func GetAudioRecordingDevices() ([]AudioDeviceID, error) {
	var count C.int
	ret := C.SDL_GetAudioRecordingDevices(&count)
	if ret != nil {
		return nil, GetError()
	}
	return unsafe.Slice((*AudioDeviceID)(ret), int(count)), nil
}

// Get the human-readable name of a specific audio device.
// (https://wiki.libsdl.org/SDL3/SDL_GetAudioDeviceName)
func GetAudioDeviceName(devid AudioDeviceID) (name string, err error) {
	name = C.GoString(C.SDL_GetAudioDeviceName(devid.c()))
	if len(name) == 0 {
		err = GetError()
	}
	return
}

// Get the current audio format of a specific audio device.
// (https://wiki.libsdl.org/SDL3/SDL_GetAudioDeviceFormat)
func (devid AudioDeviceID) GetFormat() (spec AudioSpec, sample_frames int32, err error) {
	ret := C.SDL_GetAudioDeviceFormat(devid.c(), spec.cptr(), (*C.int)(&sample_frames))
	if ret != 0 {
		err = GetError()
	}
	return
}

// Get the current channel map of an audio device.
// (https://wiki.libsdl.org/SDL3/SDL_GetAudioDeviceChannelMap)
func GetAudioDeviceChannelMap(devid AudioDeviceID) (channelMap []int32, err error) {
	var count C.int
	ret := C.SDL_GetAudioDeviceChannelMap(devid.c(), &count)
	if ret == nil {
		err = GetError()
	}
	return unsafe.Slice((*int32)(ret), int(count)), err
}

// Open a specific audio device.
// (https://wiki.libsdl.org/SDL3/SDL_OpenAudioDevice)
func OpenAudioDevice(devid AudioDeviceID, spec *AudioSpec) (vid AudioDeviceID, err error) {
	vid = AudioDeviceID(C.SDL_OpenAudioDevice(devid.c(), spec.cptr()))
	if vid == 0 {
		err = GetError()
	}
	return
}

// Use this function to pause audio playback on a specified device.
// (https://wiki.libsdl.org/SDL3/SDL_PauseAudioDevice)
func PauseAudioDevice(devid AudioDeviceID) (err error) {
	ret := C.SDL_PauseAudioDevice(devid.c())
	if ret != 0 {
		err = GetError()
	}
	return
}

// Use this function to unpause audio playback on a specified device.
// (https://wiki.libsdl.org/SDL3/SDL_ResumeAudioDevice)
func ResumeAudioDevice(devid AudioDeviceID) (err error) {
	ret := C.SDL_ResumeAudioDevice(devid.c())
	if ret != 0 {
		err = GetError()
	}
	return
}

// Use this function to query if an audio device is paused.
// (https://wiki.libsdl.org/SDL3/SDL_AudioDevicePaused)
func AudioDevicePaused(devid AudioDeviceID) bool {
	return C.SDL_AudioDevicePaused(devid.c()) != 0
}

// Get the gain of an audio device.
// (https://wiki.libsdl.org/SDL3/SDL_GetAudioDeviceGain)
func GetAudioDeviceGain(devid AudioDeviceID) (gain float32, err error) {
	gain = float32(C.SDL_GetAudioDeviceGain(devid.c()))
	if gain < 0 {
		err = GetError()
	}
	return
}

// Change the gain of an audio device.
// (https://wiki.libsdl.org/SDL3/SDL_SetAudioDeviceGain)
func SetAudioDeviceGain(devid AudioDeviceID, gain float32) (err error) {
	ret := C.SDL_SetAudioDeviceGain(devid.c(), C.float(gain))
	if ret != 0 {
		err = GetError()
	}
	return
}

// Close a previously-opened audio device.
// (https://wiki.libsdl.org/SDL3/SDL_CloseAudioDevice)
func (devid AudioDeviceID) Close() {
	C.SDL_CloseAudioDevice(devid.c())
}

// Bind a list of audio streams to an audio device.
// (https://wiki.libsdl.org/SDL3/SDL_BindAudioStreams)
func BindAudioStreams(devid AudioDeviceID, streams []*AudioStream, num_streams int) (err error) {
	data := unsafe.SliceData(streams)
	ret := C.SDL_BindAudioStreams(devid.c(), (**C.SDL_AudioStream)(unsafe.Pointer(data)), C.int(num_streams))
	if ret != 0 {
		err = GetError()
	}
	return
}

// Bind a single audio stream to an audio device.
// (https://wiki.libsdl.org/SDL3/SDL_BindAudioStream)
func BindAudioStream(devid AudioDeviceID, stream *AudioStream) (err error) {
	ret := C.SDL_BindAudioStream(devid.c(), stream.cptr())
	if ret != 0 {
		err = GetError()
	}
	return
}

// Unbind a list of audio streams from their audio devices.
// (https://wiki.libsdl.org/SDL3/SDL_UnbindAudioStreams)
func UnbindAudioStreams(streams []*AudioStream) {
	data := unsafe.SliceData(streams)
	C.SDL_UnbindAudioStreams((**C.SDL_AudioStream)(unsafe.Pointer(data)), C.int(len(streams)))
}

// Unbind a single audio stream from its audio device.
// (https://wiki.libsdl.org/SDL3/SDL_UnbindAudioStream)
func UnbindAudioStream(stream *AudioStream) {
	C.SDL_UnbindAudioStream(stream.cptr())
}

// Query an audio stream for its currently-bound device.
// (https://wiki.libsdl.org/SDL3/SDL_GetAudioStreamDevice)
func GetAudioStreamDevice(stream *AudioStream) AudioDeviceID {
	return AudioDeviceID(C.SDL_GetAudioStreamDevice(stream.cptr()))
}

// Create a new audio stream.
// (https://wiki.libsdl.org/SDL3/SDL_CreateAudioStream)
func CreateAudioStream(src_spec *AudioSpec, dst_spec *AudioSpec) (stream *AudioStream, err error) {
	stream = (*AudioStream)(C.SDL_CreateAudioStream(src_spec.cptr(), dst_spec.cptr()))
	if stream == nil {
		err = GetError()
	}
	return
}

// Get the properties associated with an audio stream.
// (https://wiki.libsdl.org/SDL3/SDL_GetAudioStreamProperties)
func GetAudioStreamProperties(stream *AudioStream) (properties PropertiesID, err error) {
	properties = PropertiesID(C.SDL_GetAudioStreamProperties(stream.cptr()))
	if properties == 0 {
		err = GetError()
	}
	return
}

// Query the current format of an audio stream.
// (https://wiki.libsdl.org/SDL3/SDL_GetAudioStreamFormat)
func GetAudioStreamFormat(stream *AudioStream) (src_spec AudioSpec, dst_spec AudioSpec, err error) {
	ret := C.SDL_GetAudioStreamFormat(stream.cptr(), src_spec.cptr(), dst_spec.cptr())
	if ret != 0 {
		err = GetError()
	}
	return
}

// Change the input and output formats of an audio stream.
// (https://wiki.libsdl.org/SDL3/SDL_SetAudioStreamFormat)
func SetAudioStreamFormat(stream *AudioStream, src_spec AudioSpec, dst_spec AudioSpec) (err error) {
	ret := C.SDL_SetAudioStreamFormat(stream.cptr(), src_spec.cptr(), dst_spec.cptr())
	if ret != 0 {
		err = GetError()
	}
	return
}

// Get the frequency ratio of an audio stream.
// (https://wiki.libsdl.org/SDL3/SDL_GetAudioStreamFrequencyRatio)
func GetAudioStreamFrequencyRatio(stream *AudioStream) (ratio float32, err error) {
	ratio = float32(C.SDL_GetAudioStreamFrequencyRatio(stream.cptr()))
	if ratio == 0 {
		err = GetError()
	}
	return
}

// Change the frequency ratio of an audio stream.
// (https://wiki.libsdl.org/SDL3/SDL_SetAudioStreamFrequencyRatio)
func SetAudioStreamFrequencyRatio(stream *AudioStream, ratio float32) (err error) {
	ret := C.SDL_SetAudioStreamFrequencyRatio(stream.cptr(), C.float(ratio))
	if ret != 0 {
		err = GetError()
	}
	return
}

// Get the gain of an audio stream.
// (https://wiki.libsdl.org/SDL3/SDL_GetAudioStreamGain)
func GetAudioStreamGain(stream *AudioStream) (gain float32, err error) {
	gain = float32(C.SDL_GetAudioStreamGain(stream.cptr()))
	if gain == -1.0 {
		err = GetError()
	}
	return
}

// Change the gain of an audio stream.
// (https://wiki.libsdl.org/SDL3/SDL_SetAudioStreamGain)
func SetAudioStreamGain(stream *AudioStream, gain float32) (err error) {
	ret := C.SDL_SetAudioStreamGain(stream.cptr(), C.float(gain))
	if ret != 0 {
		err = GetError()
	}
	return
}

// Get the current input channel map of an audio stream.
// (https://wiki.libsdl.org/SDL3/SDL_GetAudioStreamInputChannelMap)
func GetAudioStreamInputChannelMap(stream *AudioStream) []int32 {
	var count C.int
	ret := C.SDL_GetAudioStreamInputChannelMap(stream.cptr(), &count)
	return unsafe.Slice((*int32)(ret), int(count))
}

// Get the current output channel map of an audio stream.
// (https://wiki.libsdl.org/SDL3/SDL_GetAudioStreamOutputChannelMap)
func GetAudioStreamOutputChannelMap(stream *AudioStream) []int32 {
	var count C.int
	ret := C.SDL_GetAudioStreamOutputChannelMap(stream.cptr(), &count)
	return unsafe.Slice((*int32)(ret), int(count))
}

// Set the current input channel map of an audio stream.
// (https://wiki.libsdl.org/SDL3/SDL_SetAudioStreamInputChannelMap)
func SetAudioStreamInputChannelMap(stream *AudioStream, chmap []int32) (err error) {
	data := unsafe.SliceData(chmap)
	ret := C.SDL_SetAudioStreamInputChannelMap(stream.cptr(), (*C.int)(unsafe.Pointer(data)), C.int(len(chmap)))
	if ret != 0 {
		err = GetError()
	}
	return
}

// Set the current output channel map of an audio stream.
// (https://wiki.libsdl.org/SDL3/SDL_SetAudioStreamOutputChannelMap)
func SetAudioStreamOutputChannelMap(stream *AudioStream, chmap []int32) (err error) {
	data := unsafe.SliceData(chmap)
	ret := C.SDL_SetAudioStreamOutputChannelMap(stream.cptr(), (*C.int)(unsafe.Pointer(data)), C.int(len(chmap)))
	if ret != 0 {
		err = GetError()
	}
	return
}

// Add data to the stream.
// (https://wiki.libsdl.org/SDL3/SDL_PutAudioStreamData)
func (stream *AudioStream) PutData(buf []byte) (err error) {
	data := unsafe.SliceData(buf)
	ret := C.SDL_PutAudioStreamData(stream.cptr(), unsafe.Pointer(data), C.int(len(buf)))
	if ret != 0 {
		err = GetError()
	}
	return
}

// Get converted/resampled data from the stream.
// (https://wiki.libsdl.org/SDL3/SDL_GetAudioStreamData)
func (stream *AudioStream) GetData(buf []byte) (n int, err error) {
	data := unsafe.SliceData(buf)
	n = int(C.SDL_GetAudioStreamData(stream.cptr(), unsafe.Pointer(data), C.int(len(buf))))
	if n < 0 {
		err = GetError()
	}
	return
}

// Get the number of converted/resampled bytes available.
// (https://wiki.libsdl.org/SDL3/SDL_GetAudioStreamAvailable)
func (stream *AudioStream) GetAvailable() int {
	return int(C.SDL_GetAudioStreamAvailable(stream.cptr()))
}

// Get the number of bytes currently queued.
// (https://wiki.libsdl.org/SDL3/SDL_GetAudioStreamQueued)
func (stream *AudioStream) GetQueued() int {
	return int(C.SDL_GetAudioStreamQueued(stream.cptr()))
}

// Tell the stream that you're done sending data, and anything being buffered
// should be converted/resampled and made available immediately.
// (https://wiki.libsdl.org/SDL3/SDL_FlushAudioStream)
func (stream *AudioStream) Flush() (err error) {
	ret := C.SDL_FlushAudioStream(stream.cptr())
	if ret != 0 {
		err = GetError()
	}
	return
}

// Clear any pending data in the stream.
// (https://wiki.libsdl.org/SDL3/SDL_ClearAudioStream)
func (stream *AudioStream) Clear() (err error) {
	ret := C.SDL_ClearAudioStream(stream.cptr())
	if ret != 0 {
		err = GetError()
	}
	return
}

// Use this function to pause audio playback on the audio device associated
// with an audio stream.
// (https://wiki.libsdl.org/SDL3/SDL_PauseAudioStreamDevice)
func (stream *AudioStream) PauseDevice() (err error) {
	ret := C.SDL_PauseAudioStreamDevice(stream.cptr())
	if ret != 0 {
		err = GetError()
	}
	return
}

// Use this function to unpause audio playback on the audio device associated
// with an audio stream.
// (https://wiki.libsdl.org/SDL3/SDL_ResumeAudioStreamDevice)
func (stream *AudioStream) ResumeDevice() (err error) {
	ret := C.SDL_ResumeAudioStreamDevice(stream.cptr())
	if ret != 0 {
		err = GetError()
	}
	return
}

// Lock an audio stream for serialized access.
// (https://wiki.libsdl.org/SDL3/SDL_LockAudioStream)
func LockAudioStream(stream *AudioStream) (err error) {
	ret := C.SDL_LockAudioStream(stream.cptr())
	if ret != 0 {
		err = GetError()
	}
	return
}

// Unlock an audio stream for serialized access.
// (https://wiki.libsdl.org/SDL3/SDL_UnlockAudioStream)
func UnlockAudioStream(stream *AudioStream) (err error) {
	ret := C.SDL_UnlockAudioStream(stream.cptr())
	if ret != 0 {
		err = GetError()
	}
	return
}

// A callback that fires when data passes through an SDL_AudioStream.
// (https://wiki.libsdl.org/SDL3/SDL_AudioStreamCallback)
type AudioStreamCallback func(userdata any, stream *AudioStream, additional_amount, total_amount int)

type audioStreamCallbackData struct {
	callback AudioStreamCallback
	userdata any
}

var audioStreamGetCallbacks = map[*AudioStream]audioStreamCallbackData{}

//export goAudioStreamGetCallback
func goAudioStreamGetCallback(stream *C.SDL_AudioStream, additional_amount, total_amount C.int) {
	cd, ok := audioStreamGetCallbacks[(*AudioStream)(stream)]
	if !ok {
		fmt.Printf("No get callback for stream %p\n", stream)
		return
	}
	cd.callback(cd.userdata, (*AudioStream)(stream), int(additional_amount), int(total_amount))
}

// Set a callback that runs when data is requested from an audio stream.
// (https://wiki.libsdl.org/SDL3/SDL_SetAudioStreamGetCallback)
func (stream *AudioStream) SetAudioStreamGetCallback(callback AudioStreamCallback, userdata any) (err error) {
	var cCallback C.SDL_AudioStreamCallback = nil
	delete(audioStreamGetCallbacks, stream)
	if callback != nil {
		audioStreamGetCallbacks[stream] = audioStreamCallbackData{callback, userdata}
		cCallback = C.SDL_AudioStreamCallback(C.cgoAudioStreamGetCallback)
	}
	ret := C.SDL_SetAudioStreamGetCallback(stream.cptr(), cCallback, nil)
	if ret < 0 {
		delete(audioStreamGetCallbacks, stream)
		err = GetError()
	}
	return
}

var audioStreamPutCallbacks = map[*AudioStream]audioStreamCallbackData{}

//export goAudioStreamPutCallback
func goAudioStreamPutCallback(stream *C.SDL_AudioStream, additional_amount, total_amount C.int) {
	cd, ok := audioStreamGetCallbacks[(*AudioStream)(stream)]
	if !ok {
		fmt.Printf("No put callback for stream %p\n", stream)
		return
	}
	cd.callback(cd.userdata, (*AudioStream)(stream), int(additional_amount), int(total_amount))
}

// Set a callback that runs when data is added to an audio stream.
// (https://wiki.libsdl.org/SDL3/SDL_SetAudioStreamPutCallback)
func (stream *AudioStream) SetAudioStreamPutCallback(callback AudioStreamCallback, userdata any) (err error) {
	var cCallback C.SDL_AudioStreamCallback = nil
	delete(audioStreamPutCallbacks, stream)
	if callback != nil {
		audioStreamPutCallbacks[stream] = audioStreamCallbackData{callback, userdata}
		cCallback = C.SDL_AudioStreamCallback(C.cgoAudioStreamPutCallback)
	}
	ret := C.SDL_SetAudioStreamPutCallback(stream.cptr(), cCallback, nil)
	if ret < 0 {
		delete(audioStreamPutCallbacks, stream)
		err = GetError()
	}
	return
}

// Free an audio stream.
// (https://wiki.libsdl.org/SDL3/SDL_DestroyAudioStream)
func (stream *AudioStream) Destroy() {
	C.SDL_DestroyAudioStream(stream.cptr())
}

// When OpenAudioDeviceStream is used with a callback we don't know if it is a Put or Get callback.
var audioStreamUnknownCallbacks = map[*AudioStream]audioStreamCallbackData{}

//export goAudioStreamUnknownCallback
func goAudioStreamUnknownCallback(stream *C.SDL_AudioStream, additional_amount, total_amount C.int) {
	cd, ok := audioStreamUnknownCallbacks[(*AudioStream)(stream)]
	if !ok {
		fmt.Printf("No unknown callback for stream %p\n", stream)
		return
	}
	cd.callback(cd.userdata, (*AudioStream)(stream), int(additional_amount), int(total_amount))
}

// Convenience function for straightforward audio init for the common case.
// (https://wiki.libsdl.org/SDL3/SDL_OpenAudioDeviceStream)
func OpenAudioDeviceStream(devid AudioDeviceID, spec *AudioSpec, callback AudioStreamCallback, userdata any) (stream *AudioStream, err error) {
	var cCallback C.SDL_AudioStreamCallback = nil
	if callback != nil {
		cCallback = C.SDL_AudioStreamCallback(C.cgoAudioStreamUnknownCallback)
	}
	stream = (*AudioStream)(C.SDL_OpenAudioDeviceStream(devid.c(), spec.cptr(), cCallback, nil))
	if stream == nil {
		err = GetError()
	} else if callback != nil {
		audioStreamUnknownCallbacks[stream] = audioStreamCallbackData{callback, userdata}
	}
	return
}

// A callback that fires when data is about to be fed to an audio device.
// (https://wiki.libsdl.org/SDL3/SDL_AudioPostmixCallback)
type AudioPostmixCallback func(userdata any, spec *AudioSpec, buffer []float32)

type audioPostmixCallbackData struct {
	callback AudioPostmixCallback
	userdata any
}

var audioPostmixCallbacks = map[AudioDeviceID]audioPostmixCallbackData{}

//export goAudioPostmixCallback
func goAudioPostmixCallback(userdata unsafe.Pointer, spec *C.SDL_AudioSpec, buffer *C.float, buflen C.int) {
	cd, ok := audioPostmixCallbacks[AudioDeviceID(uintptr(userdata))]
	if !ok {
		fmt.Printf("No postmix callback for device %p\n", userdata)
		return
	}
	cd.callback(cd.userdata, (*AudioSpec)(unsafe.Pointer(spec)), unsafe.Slice((*float32)(buffer), int(buflen)))
}

// Set a callback that fires when data is about to be fed to an audio device.
// (https://wiki.libsdl.org/SDL3/SDL_SetAudioPostmixCallback)
func SetAudioPostmixCallback(devid AudioDeviceID, callback AudioPostmixCallback, userdata any) (err error) {
	var cCallback C.SDL_AudioPostmixCallback = nil
	delete(audioPostmixCallbacks, devid)
	if callback != nil {
		audioPostmixCallbacks[devid] = audioPostmixCallbackData{callback, userdata}
		cCallback = C.SDL_AudioPostmixCallback(C.cgoAudioPostmixCallback)
	}
	ret := C.SDL_SetAudioPostmixCallback(devid.c(), cCallback, unsafe.Pointer(uintptr(devid)))
	if ret != 0 {
		delete(audioPostmixCallbacks, devid)
		err = GetError()
	}
	return
}

// Load the audio data of a WAVE file into memory.
// Returned data must be freed with sdl.Free some how.
// (https://wiki.libsdl.org/SDL3/SDL_LoadWAV_IO)
func (src *IOStream) LoadWAV(closeio bool) (spec *AudioSpec, data []uint8, err error) {
	// TODO This needs work. I'm thinking about createing a WaveBuffer type
	// So i can create a func (buf *WaveBuffer) Free() { C.SDL_Free(buf.cptr()) }
	var count C.Uint32
	var cdata *C.Uint8
	ret := C.SDL_LoadWAV_IO(src.cptr(), sdlBool(closeio), spec.cptr(), &cdata, &count)
	if ret != 0 {
		err = GetError()
		return
	}
	data = unsafe.Slice((*uint8)(unsafe.Pointer(cdata)), count)
	return
}

// Loads a WAV from a file path.
// (https://wiki.libsdl.org/SDL3/SDL_LoadWAV)
func LoadWAV(path string) (spec AudioSpec, data []uint8, err error) {
	// TODO This needs work. I'm thinking about createing a WaveBuffer type
	// So i can create a func (buf *WaveBuffer) Free() { C.SDL_Free(buf.cptr()) }
	cPath := C.CString(path)
	defer C.free(unsafe.Pointer(cPath))
	var count C.Uint32
	var cdata *C.Uint8
	ret := C.SDL_LoadWAV(C.CString(path), spec.cptr(), &cdata, &count)
	if ret != 0 {
		err = GetError()
		return
	}
	data = unsafe.Slice((*uint8)(unsafe.Pointer(cdata)), count)
	return
}

// Mix audio data in a specified format.
// (https://wiki.libsdl.org/SDL3/SDL_MixAudio)
func MixAudio(dst []uint8, src []uint8, format AudioFormat, volume float32) (err error) {
	dstPtr := unsafe.SliceData(dst)
	srcPtr := unsafe.SliceData(src)
	ret := C.SDL_MixAudio((*C.Uint8)(dstPtr), (*C.Uint8)(srcPtr), format.c(), C.Uint32(len(dst)), C.float(volume))
	if ret != 0 {
		err = GetError()
	}
	return
}

// Convert some audio data of one format to another format.
// (https://wiki.libsdl.org/SDL3/SDL_ConvertAudioSamples)
func ConvertAudioSamples(srcSpec *AudioSpec, src []uint8, dstSpec *AudioSpec) (dst []uint8, err error) {
	srcPtr := unsafe.SliceData(src)
	var count C.int
	var cdata *C.Uint8
	ret := C.SDL_ConvertAudioSamples(srcSpec.cptr(), (*C.Uint8)(srcPtr), C.int(len(src)), dstSpec.cptr(), &cdata, &count)
	if ret != 0 {
		err = GetError()
		return
	}
	return unsafe.Slice((*uint8)(cdata), count), nil
}

// Get the appropriate memset value for silencing an audio format.
// (https://wiki.libsdl.org/SDL3/SDL_GetSilenceValueForFormat)
func GetSilenceValueForFormat(format AudioFormat) int {
	return int(C.SDL_GetSilenceValueForFormat(format.c()))
}

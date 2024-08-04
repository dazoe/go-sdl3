package sdl

//#include <SDL3/SDL_touch.h>
import "C"
import "unsafe"

// # CategoryTouch
// (https://wiki.libsdl.org/SDL3/CategoryTouch)

type TouchID C.SDL_TouchID

func (t TouchID) c() C.SDL_TouchID {
	return C.SDL_TouchID(t)
}

type FingerID C.SDL_FingerID

func (f FingerID) c() C.SDL_FingerID {
	return C.SDL_FingerID(f)
}

type TouchDeviceType C.SDL_TouchDeviceType

const (
	TOUCH_DEVICE_INVALID           = TouchDeviceType(C.SDL_TOUCH_DEVICE_INVALID)
	TOUCH_DEVICE_DIRECT            = TouchDeviceType(C.SDL_TOUCH_DEVICE_DIRECT)
	TOUCH_DEVICE_INDIRECT_ABSOLUTE = TouchDeviceType(C.SDL_TOUCH_DEVICE_INDIRECT_ABSOLUTE)
	TOUCH_DEVICE_INDIRECT_RELATIVE = TouchDeviceType(C.SDL_TOUCH_DEVICE_INDIRECT_RELATIVE)
)

// Data about a single finger in a multitouch event.
// (https://wiki.libsdl.org/SDL3/SDL_Finger)
type Finger struct {
	Id       FingerID /**< the finger ID */
	X        float32  /**< the x-axis location of the touch event, normalized (0...1) */
	Y        float32  /**< the y-axis location of the touch event, normalized (0...1) */
	Pressure float32  /**< the quantity of pressure applied, normalized (0...1) */
}

/* Used as the device ID for mouse events simulated with touch input */
const TOUCH_MOUSEID = MouseID(C.SDL_TOUCH_MOUSEID)

/* Used as the SDL_TouchID for touch events simulated with mouse input */
const MOUSE_TOUCHID = TouchID(C.SDL_MOUSE_TOUCHID)

// Get a list of registered touch devices.
// (https://wiki.libsdl.org/SDL3/SDL_GetTouchDevices)
func GetTouchDevices() ([]TouchID, error) {
	var count C.int
	ret := C.SDL_GetTouchDevices(&count)
	if ret == nil {
		return nil, GetError()
	}
	return unsafe.Slice((*TouchID)(unsafe.Pointer(ret)), count), nil
}

// Get the touch device name as reported from the driver.
// (https://wiki.libsdl.org/SDL3/SDL_GetTouchDeviceName)
func GetTouchDeviceName(touchID TouchID) (string, error) {
	ret := C.SDL_GetTouchDeviceName(touchID.c())
	if ret == nil {
		return "", GetError()
	}
	return C.GoString(ret), nil
}

// Get the type of the given touch device.
// (https://wiki.libsdl.org/SDL3/SDL_GetTouchDeviceType)
func GetTouchDeviceType(touchID TouchID) TouchDeviceType {
	return TouchDeviceType(C.SDL_GetTouchDeviceType(touchID.c()))
}

// Get a list of active fingers for a given touch device.
// (https://wiki.libsdl.org/SDL3/SDL_GetTouchFingers)
func GetTouchFingers(touchID TouchID) ([]*Finger, error) {
	var count C.int
	ret := C.SDL_GetTouchFingers(touchID.c(), &count)
	if ret == nil {
		return nil, GetError()
	}
	return unsafe.Slice((**Finger)(unsafe.Pointer(ret)), count), nil
}

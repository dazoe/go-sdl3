package sdl

//#include <SDL3/SDL_camera.h>
import "C"
import "unsafe"

// # CategoryCamera
// (https://wiki.libsdl.org/SDL3/CategoryCamera)

// This is a unique ID for a camera device for the time it is connected to the
// system, and is never reused for the lifetime of the application.
// (https://wiki.libsdl.org/SDL3/SDL_CameraID)
type CameraID C.SDL_CameraID

func (c CameraID) c() C.SDL_CameraID {
	return (C.SDL_CameraID)(c)
}

// The opaque structure used to identify an opened SDL camera.
// (https://wiki.libsdl.org/SDL3/SDL_Camera)
type Camera C.SDL_Camera

func (c *Camera) cptr() *C.SDL_Camera {
	return (*C.SDL_Camera)(c)
}

// The details of an output format for a camera device.
// (https://wiki.libsdl.org/SDL3/SDL_CameraSpec)
type CameraSpec struct {
	Format               PixelFormat
	Colorspace           Colorspace
	Width                int32
	Height               int32
	FramerateNumerator   int32
	FramerateDenominator int32
}

func (c *CameraSpec) cptr() *C.SDL_CameraSpec {
	return (*C.SDL_CameraSpec)(unsafe.Pointer(c))
}

// The position of camera in relation to system device.
type CameraPosition C.SDL_CameraPosition

const (
	CAMERA_POSITION_UNKNOWN      = CameraPosition(C.SDL_CAMERA_POSITION_UNKNOWN)
	CAMERA_POSITION_FRONT_FACING = CameraPosition(C.SDL_CAMERA_POSITION_FRONT_FACING)
	CAMERA_POSITION_BACK_FACING  = CameraPosition(C.SDL_CAMERA_POSITION_BACK_FACING)
)

// Use this function to get the number of built-in camera drivers.
// (https://wiki.libsdl.org/SDL3/SDL_GetNumCameraDrivers)
func GetNumCameraDrivers() int32 {
	return int32(C.SDL_GetNumCameraDrivers())
}

// Use this function to get the name of a built in camera driver.
// (https://wiki.libsdl.org/SDL3/SDL_GetCameraDriver)
func GetCameraDriver(index int32) string {
	return C.GoString(C.SDL_GetCameraDriver(C.int(index)))
}

// Get the name of the current camera driver.
// (https://wiki.libsdl.org/SDL3/SDL_GetCurrentCameraDriver)
func GetCurrentCameraDriver() string {
	return C.GoString(C.SDL_GetCurrentCameraDriver())
}

// Get a list of currently connected camera devices.
// (https://wiki.libsdl.org/SDL3/SDL_GetCameras)
func GetCameras() ([]CameraID, error) {
	var count C.int
	ret := C.SDL_GetCameras(&count)
	if ret == nil {
		return nil, GetError()
	}
	return unsafe.Slice((*CameraID)(ret), int(count)), nil
}

// Get the list of native formats/sizes a camera supports.
// (https://wiki.libsdl.org/SDL3/SDL_GetCameraSupportedFormats)
func GetCameraSupportedFormats(devid CameraID) ([]*CameraSpec, error) {
	var count C.int
	ret := C.SDL_GetCameraSupportedFormats(devid.c(), &count)
	if ret == nil {
		return nil, GetError()
	}
	return unsafe.Slice((**CameraSpec)(unsafe.Pointer(ret)), int(count)), nil
}

// Get the human-readable device name for a camera.
// (https://wiki.libsdl.org/SDL3/SDL_GetCameraName)
func GetCameraName(devid CameraID) (string, error) {
	ret := C.SDL_GetCameraName(devid.c())
	if ret == nil {
		return "", GetError()
	}
	return C.GoString(ret), nil
}

// Get the position of the camera in relation to the system.
// (https://wiki.libsdl.org/SDL3/SDL_GetCameraPosition)
func GetCameraPosition(devid CameraID) CameraPosition {
	return CameraPosition(C.SDL_GetCameraPosition(devid.c()))
}

// Open a video recording device (a "camera").
// (https://wiki.libsdl.org/SDL3/SDL_OpenCamera)
func OpenCamera(devid CameraID, spec CameraSpec) (*Camera, error) {
	ret := C.SDL_OpenCamera(devid.c(), spec.cptr())
	if ret == nil {
		return nil, GetError()
	}
	return (*Camera)(ret), nil
}

// Query if camera access has been approved by the user.
// (https://wiki.libsdl.org/SDL3/SDL_GetCameraPermissionState)
func GetCameraPermissionState(camera *Camera) int32 {
	return int32(C.SDL_GetCameraPermissionState(camera.cptr()))
}

// Get the instance ID of an opened camera.
// (https://wiki.libsdl.org/SDL3/SDL_GetCameraID)
func GetCameraID(camera *Camera) (CameraID, error) {
	ret := C.SDL_GetCameraID(camera.cptr())
	if ret == 0 {
		return 0, GetError()
	}
	return CameraID(ret), nil
}

// Get the properties associated with an opened camera.
// (https://wiki.libsdl.org/SDL3/SDL_GetCameraProperties)
func GetCameraProperties(camera *Camera) (PropertiesID, error) {
	ret := C.SDL_GetCameraProperties(camera.cptr())
	if ret == 0 {
		return 0, GetError()
	}
	return PropertiesID(ret), nil
}

// Get the spec that a camera is using when generating images.
// (https://wiki.libsdl.org/SDL3/SDL_GetCameraFormat)
func GetCameraFormat(camera *Camera) (spec CameraSpec, err error) {
	ret := C.SDL_GetCameraFormat(camera.cptr(), spec.cptr())
	if ret < 0 {
		err = GetError()
	}
	return
}

// Acquire a frame.
// (https://wiki.libsdl.org/SDL3/SDL_AcquireCameraFrame)
func AcquireCameraFrame(camera *Camera) (surface *Surface, timestamp uint64, err error) {
	surface = (*Surface)(unsafe.Pointer(C.SDL_AcquireCameraFrame(camera.cptr(), (*C.Uint64)(unsafe.Pointer(&timestamp)))))
	if surface == nil {
		err = GetError()
	}
	return
}

// Release a frame of video acquired from a camera.
// (https://wiki.libsdl.org/SDL3/SDL_ReleaseCameraFrame)
func ReleaseCameraFrame(camera *Camera, surface *Surface) error {
	ret := C.SDL_ReleaseCameraFrame(camera.cptr(), surface.cptr())
	if ret < 0 {
		return GetError()
	}
	return nil
}

// Use this function to shut down camera processing and close the camera
// device.
// (https://wiki.libsdl.org/SDL3/SDL_CloseCamera)
func CloseCamera(camera *Camera) {
	C.SDL_CloseCamera(camera.cptr())
}

package sdl

/*
#include <SDL3/SDL_init.h>
#include <stdlib.h>
*/
import "C"
import "unsafe"

type InitFlags C.SDL_InitFlags

const (
	INIT_AUDIO    = InitFlags(C.SDL_INIT_AUDIO)    /**< `SDL_INIT_AUDIO` implies `SDL_INIT_EVENTS` */
	INIT_VIDEO    = InitFlags(C.SDL_INIT_VIDEO)    /**< `SDL_INIT_VIDEO` implies `SDL_INIT_EVENTS` */
	INIT_JOYSTICK = InitFlags(C.SDL_INIT_JOYSTICK) /**< `SDL_INIT_JOYSTICK` implies `SDL_INIT_EVENTS`, should be initialized on the same thread as SDL_INIT_VIDEO on Windows if you don't set SDL_HINT_JOYSTICK_THREAD */
	INIT_HAPTIC   = InitFlags(C.SDL_INIT_HAPTIC)
	INIT_GAMEPAD  = InitFlags(C.SDL_INIT_GAMEPAD) /**< `SDL_INIT_GAMEPAD` implies `SDL_INIT_JOYSTICK` */
	INIT_EVENTS   = InitFlags(C.SDL_INIT_EVENTS)
	INIT_SENSOR   = InitFlags(C.SDL_INIT_SENSOR) /**< `SDL_INIT_SENSOR` implies `SDL_INIT_EVENTS` */
	INIT_CAMERA   = InitFlags(C.SDL_INIT_CAMERA) /**< `SDL_INIT_CAMERA` implies `SDL_INIT_EVENTS` */
)

// Init Initialize the SDL library
// (https://wiki.libsdl.org/SDL3/SDL_Init)
func Init(flags InitFlags) error {
	if !C.SDL_Init(C.SDL_InitFlags(flags)) {
		return GetError()
	}
	return nil
}

// InitSubSystem Compatibility function to initialize the SDL library
// (https://wiki.libsdl.org/SDL3/SDL_InitSubSystem)
func InitSubSystem(flag InitFlags) error {
	if !C.SDL_InitSubSystem(C.SDL_InitFlags(flag)) {
		return GetError()
	}
	return nil
}

// QuitSubSystem Shut down specific SDL subsystems
// (https://wiki.libsdl.org/SDL3/SDL_QuitSubSystem)
func QuitSubSystem(flag InitFlags) {
	C.SDL_QuitSubSystem(C.SDL_InitFlags(flag))
}

// WasInit Get a mask of the specified subsystems which are currently initialized
// (https://wiki.libsdl.org/SDL3/SDL_WasInit)
func WasInit(flags InitFlags) InitFlags {
	return InitFlags(C.SDL_WasInit(C.SDL_InitFlags(flags)))
}

// Clean up all initialized subsystems
// (https://wiki.libsdl.org/SDL3/SDL_Quit)
func Quit() {
	C.SDL_Quit()
}

// TODO: needs tested
// Specify basic metadata about your app.
// (https://wiki.libsdl.org/SDL3/SDL_SetAppMetadata)
func SetAppMetadata(appname, appversion, appidentifier string) error {
	cappname := C.CString(appname)
	defer C.free(unsafe.Pointer(cappname))
	cappversion := C.CString(appversion)
	defer C.free(unsafe.Pointer(cappversion))
	cappidentifier := C.CString(appidentifier)
	defer C.free(unsafe.Pointer(cappidentifier))
	ret := C.SDL_SetAppMetadata(cappname, cappversion, cappidentifier)
	if !ret {
		return GetError()
	}
	return nil
}

// Specify metadata about your app through a set of properties.
// (https://wiki.libsdl.org/SDL3/SDL_SetAppMetadataProperty)
func SetAppMetadataProperty(name, value string) error {
	cname := C.CString(name)
	defer C.free(unsafe.Pointer(cname))
	var cvalue *C.char
	if len(value) == 0 {
		cvalue = C.CString(value)
		defer C.free(unsafe.Pointer(cvalue))
	}
	ret := C.SDL_SetAppMetadataProperty(cname, cvalue)
	if !ret {
		return GetError()
	}
	return nil
}

const (
	PROP_APP_METADATA_NAME_STRING       = "SDL.app.metadata.name"
	PROP_APP_METADATA_VERSION_STRING    = "SDL.app.metadata.version"
	PROP_APP_METADATA_IDENTIFIER_STRING = "SDL.app.metadata.identifier"
	PROP_APP_METADATA_CREATOR_STRING    = "SDL.app.metadata.creator"
	PROP_APP_METADATA_COPYRIGHT_STRING  = "SDL.app.metadata.copyright"
	PROP_APP_METADATA_URL_STRING        = "SDL.app.metadata.url"
	PROP_APP_METADATA_TYPE_STRING       = "SDL.app.metadata.type"
)

// Get metadata about your app.
// (https://wiki.libsdl.org/SDL3/SDL_GetAppMetadataProperty)
func GetAppMetadataProperty(name string) string {
	cname := C.CString(name)
	defer C.free(unsafe.Pointer(cname))
	return C.GoString(C.SDL_GetAppMetadataProperty(cname))
}

package sdl

//#include <SDL3/SDL_init.h>
import "C"

type InitFlags C.SDL_InitFlags

const (
	INIT_TIMER    = InitFlags(C.SDL_INIT_TIMER)
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
	if C.SDL_Init(C.SDL_InitFlags(flags)) != 0 {
		return GetError()
	}
	return nil
}

// InitSubSystem Compatibility function to initialize the SDL library
// (https://wiki.libsdl.org/SDL3/SDL_InitSubSystem)
func InitSubSystem(flag InitFlags) error {
	if C.SDL_InitSubSystem(C.SDL_InitFlags(flag)) != 0 {
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

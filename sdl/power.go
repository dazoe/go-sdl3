package sdl

//#include <SDL3/SDL_power.h>
import "C"

// # CategoryPower
// (https://wiki.libsdl.org/SDL3/CategoryPower)

// The basic state for the system's power supply.
// (https://wiki.libsdl.org/SDL3/SDL_PowerState)
type PowerState C.SDL_PowerState

const (
	POWERSTATE_ERROR      = PowerState(C.SDL_POWERSTATE_ERROR)      /**< error determining power status */
	POWERSTATE_UNKNOWN    = PowerState(C.SDL_POWERSTATE_UNKNOWN)    /**< cannot determine power status */
	POWERSTATE_ON_BATTERY = PowerState(C.SDL_POWERSTATE_ON_BATTERY) /**< Not plugged in, running on the battery */
	POWERSTATE_NO_BATTERY = PowerState(C.SDL_POWERSTATE_NO_BATTERY) /**< Plugged in, no battery available */
	POWERSTATE_CHARGING   = PowerState(C.SDL_POWERSTATE_CHARGING)   /**< Plugged in, charging battery */
	POWERSTATE_CHARGED    = PowerState(C.SDL_POWERSTATE_CHARGED)    /**< Plugged in, battery charged */
)

// Get the current power supply details.
// (https://wiki.libsdl.org/SDL3/SDL_GetPowerInfo)
func GetPowerInfo() (seconds, percent int32, state PowerState, err error) {
	state = PowerState(C.SDL_GetPowerInfo((*C.int)(&seconds), (*C.int)(&percent)))
	if state == POWERSTATE_ERROR {
		err = GetError()
	}
	return
}

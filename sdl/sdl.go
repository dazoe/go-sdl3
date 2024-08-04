package sdl

import "C"
import "runtime"

func init() {
	// Stick to the main thread
	runtime.LockOSThread()
}

// Utility functions....
func sdlBool(b bool) C.int {
	if b {
		return 1
	}
	return 0
}

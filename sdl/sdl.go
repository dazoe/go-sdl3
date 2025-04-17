package sdl

import "C"
import "runtime"

func init() {
	// Stick to the main thread
	runtime.LockOSThread()
}

// TODO: still needed?
// Utility functions....
func sdlBool(b bool) int {
	if b {
		return 1
	}
	return 0
}

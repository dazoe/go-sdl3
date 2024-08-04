package sdl

//#include<SDL3/SDL_stdinc.h>
import "C"

//TODO incomplete

// For some reason C.SDL_FLT_EPSILON evaluated to 0.0 it should be
const FLT_EPSILON = 1.1920928955078125e-07

func GetNumAllocations() int {
	return int(C.SDL_GetNumAllocations())
}

type FunctionPointer C.SDL_FunctionPointer

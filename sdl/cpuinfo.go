package sdl

//#include <SDL3/SDL_cpuinfo.h>
import "C"

// # CategoryCPUInfo
// (https://wiki.libsdl.org/SDL3/CategoryCPUInfo)

// Get the number of logical CPU cores available.
// (https://wiki.libsdl.org/SDL3/SDL_GetNumLogicalCPUCores)
func GetNumLogicalCPUCores() int {
	return int(C.SDL_GetNumLogicalCPUCores())
}

// Determine the L1 cache line size of the CPU.
// (https://wiki.libsdl.org/SDL3/SDL_GetCPUCacheLineSize)
func GetCPUCacheLineSize() int {
	return int(C.SDL_GetCPUCacheLineSize())
}

// Determine whether the CPU has AltiVec features.
// (https://wiki.libsdl.org/SDL3/SDL_HasAltiVec)
func HasAltiVec() bool {
	return bool(C.SDL_HasAltiVec())
}

// Determine whether the CPU has MMX features.
// (https://wiki.libsdl.org/SDL3/SDL_HasMMX)
func HasMMX() bool {
	return bool(C.SDL_HasMMX())
}

// Determine whether the CPU has SSE features.
// (https://wiki.libsdl.org/SDL3/SDL_HasSSE)
func HasSSE() bool {
	return bool(C.SDL_HasSSE())
}

// Determine whether the CPU has SSE2 features.
// (https://wiki.libsdl.org/SDL3/SDL_HasSSE2)
func HasSSE2() bool {
	return bool(C.SDL_HasSSE2())
}

// Determine whether the CPU has SSE3 features.
// (https://wiki.libsdl.org/SDL3/SDL_HasSSE3)
func HasSSE3() bool {
	return bool(C.SDL_HasSSE3())
}

// Determine whether the CPU has SSE4.1 features.
// (https://wiki.libsdl.org/SDL3/SDL_HasSSE41)
func HasSSE41() bool {
	return bool(C.SDL_HasSSE41())
}

// Determine whether the CPU has SSE4.2 features.
// (https://wiki.libsdl.org/SDL3/SDL_HasSSE42)
func HasSSE42() bool {
	return bool(C.SDL_HasSSE42())
}

// Determine whether the CPU has AVX features.
// (https://wiki.libsdl.org/SDL3/SDL_HasAVX)
func HasAVX() bool {
	return bool(C.SDL_HasAVX())
}

// Determine whether the CPU has AVX2 features.
// (https://wiki.libsdl.org/SDL3/SDL_HasAVX2)
func HasAVX2() bool {
	return bool(C.SDL_HasAVX2())
}

// Determine whether the CPU has AVX-512F (foundation) features.
// (https://wiki.libsdl.org/SDL3/SDL_HasAVX512F)
func HasAVX512F() bool {
	return bool(C.SDL_HasAVX512F())
}

// Determine whether the CPU has ARM SIMD (ARMv6) features.
// (https://wiki.libsdl.org/SDL3/SDL_HasARMSIMD)
func HasARMSIMD() bool {
	return bool(C.SDL_HasARMSIMD())
}

// Determine whether the CPU has NEON (ARM SIMD) features.
// (https://wiki.libsdl.org/SDL3/SDL_HasNEON)
func HasNEON() bool {
	return bool(C.SDL_HasNEON())
}

// Determine whether the CPU has LSX (LOONGARCH SIMD) features.
// (https://wiki.libsdl.org/SDL3/SDL_HasLSX)
func HasLSX() bool {
	return bool(C.SDL_HasLSX())
}

// Determine whether the CPU has LASX (LOONGARCH SIMD) features.
// (https://wiki.libsdl.org/SDL3/SDL_HasLASX)
func HasLASX() bool {
	return bool(C.SDL_HasLASX())
}

// Get the amount of RAM configured in the system.
// (https://wiki.libsdl.org/SDL3/SDL_GetSystemRAM)
func GetSystemRAM() int {
	return int(C.SDL_GetSystemRAM())
}

// Report the alignment this system needs for SIMD allocations.
// (https://wiki.libsdl.org/SDL3/SDL_GetSIMDAlignment)
func GetSIMDAlignment() int {
	return int(C.SDL_GetSIMDAlignment())
}

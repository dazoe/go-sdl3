package sdl

/*
#include <SDL3/SDL_atomic.h>
*/
import "C"
import (
	"runtime"
	"unsafe"
)

// # CategoryAtomic
// (https://wiki.libsdl.org/SDL3/CategoryAtomic)

// An atomic spinlock.
// (https://wiki.libsdl.org/SDL3/SDL_SpinLock)
type SpinLock C.SDL_SpinLock

func (lock *SpinLock) cptr() *C.SDL_SpinLock {
	return (*C.SDL_SpinLock)(unsafe.Pointer(lock))
}

// Try to lock a spin lock by setting it to a non-zero value.
// (https://wiki.libsdl.org/SDL3/SDL_TryLockSpinlock)
func (lock *SpinLock) TryLock() bool {
	return bool(C.SDL_TryLockSpinlock(lock.cptr()))
}

// Lock a spin lock by setting it to a non-zero value.
// (https://wiki.libsdl.org/SDL3/SDL_LockSpinlock)
func (lock *SpinLock) Lock() {
	C.SDL_LockSpinlock(lock.cptr())
}

// Unlock a spin lock by setting it to 0.
// (https://wiki.libsdl.org/SDL3/SDL_UnlockSpinlock)
func (lock *SpinLock) Unlock() {
	C.SDL_UnlockSpinlock(lock.cptr())
}

// Mark a compiler barrier.
// (https://wiki.libsdl.org/SDL3/SDL_CompilerBarrier)
func CompilerBarrier() {
	// No idea what go's equivalent to this is.
	// C.SDL_CompilerBarrier()
}

// Insert a memory release barrier.
// (https://wiki.libsdl.org/SDL3/SDL_MemoryBarrierReleaseFunction)
func MemoryBarrierReleaseFunction() {
	C.SDL_MemoryBarrierReleaseFunction()
}

// Insert a memory acquire barrier.
// (https://wiki.libsdl.org/SDL3/SDL_MemoryBarrierAcquireFunction)
func MemoryBarrierAcquireFunction() {
	C.SDL_MemoryBarrierAcquireFunction()
}

// !!! FIXME: this should have documentation! */
// (https://wiki.libsdl.org/SDL3/SDL_MemoryBarrierRelease)
func MemoryBarrierRelease() {
	// No idea what go's equivalent to this is.
	// C.SDL_MemoryBarrierRelease()
}
func MemoryBarrierAcquire() {
	// No idea what go's equivalent to this is.
	// C.SDL_MemoryBarrierAcquire()
}

// A macro to insert a CPU-specific "pause" instruction into the program.
// (https://wiki.libsdl.org/SDL3/SDL_CPUPauseInstruction)
func CPUPauseInstruction() {
	// Maybe a runtime.Gosched() is what should be used here?
	runtime.Gosched()
}

// A type representing an atomic integer value.
// (https://wiki.libsdl.org/SDL3/SDL_AtomicInt)
// type AtomicInt C.SDL_AtomicInt
type AtomicInt struct {
	Value int32
}

func (a *AtomicInt) cptr() *C.SDL_AtomicInt {
	return (*C.SDL_AtomicInt)(unsafe.Pointer(a))
}

// Set an atomic variable to a new value if it is currently an old value.
// (https://wiki.libsdl.org/SDL3/SDL_CompareAndSwapAtomicInt)
func (a *AtomicInt) CompareAndSwap(oldval, newval int) bool {
	return bool(C.SDL_CompareAndSwapAtomicInt(a.cptr(), C.int(oldval), C.int(newval)))
}

// Set an atomic variable to a value.
// (https://wiki.libsdl.org/SDL3/SDL_SetAtomicInt)
func (a *AtomicInt) Set(v int32) int32 {
	return int32(C.SDL_SetAtomicInt(a.cptr(), C.int(v)))
}

// Get the value of an atomic variable.
// (https://wiki.libsdl.org/SDL3/SDL_GetAtomicInt)
func (a *AtomicInt) Get() int32 {
	return int32(C.SDL_GetAtomicInt(a.cptr()))
}

// Add to an atomic variable.
// (https://wiki.libsdl.org/SDL3/SDL_AddAtomicInt)
func (a *AtomicInt) Add(v int32) int32 {
	return int32(C.SDL_AddAtomicInt(a.cptr(), C.int(v)))
}

// Increment an atomic variable used as a reference count.
// (https://wiki.libsdl.org/SDL3/SDL_AtomicIncRef)
func (a *AtomicInt) IncRef() int32 {
	return int32(C.SDL_AddAtomicInt(a.cptr(), 1))
}

// Decrement an atomic variable used as a reference count.
// (https://wiki.libsdl.org/SDL3/SDL_AtomicDecRef)
func (a *AtomicInt) DecRef() bool {
	return C.SDL_AddAtomicInt(a.cptr(), -1) == 1
}

// A type representing an atomic unsigned 32-bit value.
// (https://wiki.libsdl.org/SDL3/SDL_AtomicU32)
// type AtomicU32 C.SDL_AtomicU32
type AtomicU32 struct {
	Value uint32
}

func (a *AtomicU32) cptr() *C.SDL_AtomicU32 {
	return (*C.SDL_AtomicU32)(unsafe.Pointer(a))
}

// Set an atomic variable to a new value if it is currently an old value.
// (https://wiki.libsdl.org/SDL3/SDL_CompareAndSwapAtomicU32)
func (a *AtomicU32) CompareAndSwap(oldval, newval uint32) bool {
	return bool(C.SDL_CompareAndSwapAtomicU32(a.cptr(), C.Uint32(oldval), C.Uint32(newval)))
}

// Set an atomic variable to a value.
// (https://wiki.libsdl.org/SDL3/SDL_SetAtomicU32)
func (a *AtomicU32) Set(v uint32) uint32 {
	return uint32(C.SDL_SetAtomicU32(a.cptr(), C.Uint32(v)))
}

// Get the value of an atomic variable.
// (https://wiki.libsdl.org/SDL3/SDL_GetAtomicU32)
func (a *AtomicU32) Get() uint32 {
	return uint32(C.SDL_GetAtomicU32(a.cptr()))
}

// Set a pointer to a new value if it is currently an old value.
// (https://wiki.libsdl.org/SDL3/SDL_CompareAndSwapAtomicPointer)
func CompareAndSwapAtomicPointer(ptr *unsafe.Pointer, oldval, newval unsafe.Pointer) bool {
	return bool(C.SDL_CompareAndSwapAtomicPointer(ptr, oldval, newval))
}

// Set a pointer to a value atomically.
// (https://wiki.libsdl.org/SDL3/SDL_SetAtomicPointer)
func SetAtomicPointer(ptr *unsafe.Pointer, v unsafe.Pointer) unsafe.Pointer {
	return unsafe.Pointer(C.SDL_SetAtomicPointer(ptr, v))
}

// Get the value of a pointer atomically.
// (https://wiki.libsdl.org/SDL3/SDL_GetAtomicPointer)
func GetAtomicPointer(ptr *unsafe.Pointer) unsafe.Pointer {
	return unsafe.Pointer(C.SDL_GetAtomicPointer(ptr))
}

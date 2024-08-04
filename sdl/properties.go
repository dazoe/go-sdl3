package sdl

/*
#include <SDL3/SDL_properties.h>
#include "properties.h"
#include "stdlib.h"
*/
import "C"
import "unsafe"

// # CategoryProperties
// (https://wiki.libsdl.org/SDL3/CategoryProperties)

// SDL properties ID
// (https://wiki.libsdl.org/SDL3/SDL_PropertiesID)
type PropertiesID C.SDL_PropertiesID

func (id PropertiesID) c() C.SDL_PropertiesID {
	return C.SDL_PropertiesID(id)
}

// SDL property type
// (https://wiki.libsdl.org/SDL3/SDL_PropertyType)
type PropertyType C.SDL_PropertyType

const (
	PROPERTY_TYPE_INVALID = C.SDL_PROPERTY_TYPE_INVALID
	PROPERTY_TYPE_POINTER = C.SDL_PROPERTY_TYPE_POINTER
	PROPERTY_TYPE_STRING  = C.SDL_PROPERTY_TYPE_STRING
	PROPERTY_TYPE_NUMBER  = C.SDL_PROPERTY_TYPE_NUMBER
	PROPERTY_TYPE_FLOAT   = C.SDL_PROPERTY_TYPE_FLOAT
	PROPERTY_TYPE_BOOLEAN = C.SDL_PROPERTY_TYPE_BOOLEAN
)

// Get the global SDL properties.
// (https://wiki.libsdl.org/SDL3/SDL_GetGlobalProperties)
func GetGlobalProperties() (PropertiesID, error) {
	ret := PropertiesID(C.SDL_GetGlobalProperties())
	if ret == 0 {
		return 0, GetError()
	}
	return PropertiesID(ret), nil
}

// Create a group of properties.
// (https://wiki.libsdl.org/SDL3/SDL_CreateProperties)
func CreateProperties() (PropertiesID, error) {
	ret := PropertiesID(C.SDL_CreateProperties())
	if ret == 0 {
		return 0, GetError()
	}
	return PropertiesID(ret), nil
}

// Copy a group of properties.
// (https://wiki.libsdl.org/SDL3/SDL_CopyProperties)
func (pid PropertiesID) Copy(dst PropertiesID) (PropertiesID, error) {
	ret := PropertiesID(C.SDL_CopyProperties(pid.c(), dst.c()))
	if ret < 0 {
		return 0, GetError()
	}
	return PropertiesID(ret), nil
}

// Lock a group of properties.
// (https://wiki.libsdl.org/SDL3/SDL_LockProperties)
func (pid PropertiesID) Lock() error {
	ret := C.SDL_LockProperties(pid.c())
	if ret < 0 {
		return GetError()
	}
	return nil
}

// Unlock a group of properties.
// (https://wiki.libsdl.org/SDL3/SDL_UnlockProperties)
func (pid PropertiesID) Unlock() {
	C.SDL_UnlockProperties(pid.c())
}

// A callback used to free resources when a property is deleted.
// (https://wiki.libsdl.org/SDL3/SDL_CleanupPropertyCallback)
type CleanupPropertyCallback func(value unsafe.Pointer)

//export goCleanupPropertyCallback
func goCleanupPropertyCallback(userdata unsafe.Pointer, value unsafe.Pointer) {
	callback := (*CleanupPropertyCallback)(userdata)
	(*callback)(value)
}

// Set a pointer property in a group of properties with a cleanup function
// that is called when the property is deleted.
// (https://wiki.libsdl.org/SDL3/SDL_SetPointerPropertyWithCleanup)
func (pid PropertiesID) SetPointerWithCleanup(name string, value unsafe.Pointer, cleanup CleanupPropertyCallback) error {
	cname := C.CString(name)
	defer C.free(unsafe.Pointer(cname))
	ret := C.SDL_SetPointerPropertyWithCleanup(pid.c(), cname, value, C.SDL_CleanupPropertyCallback(C.cgoCleanupPropertyCallback), unsafe.Pointer(&cleanup))
	if ret < 0 {
		return GetError()
	}
	return nil
}

// Set a pointer property in a group of properties.
// (https://wiki.libsdl.org/SDL3/SDL_SetPointerProperty)
func (pid PropertiesID) SetPointer(name string, value unsafe.Pointer) error {
	cname := C.CString(name)
	defer C.free(unsafe.Pointer(cname))
	ret := C.SDL_SetPointerProperty(pid.c(), cname, value)
	if ret < 0 {
		return GetError()
	}
	return nil
}

// Set a string property in a group of properties.
// (https://wiki.libsdl.org/SDL3/SDL_SetStringProperty)
func (pid PropertiesID) SetString(name string, value string) error {
	cname := C.CString(name)
	defer C.free(unsafe.Pointer(cname))
	cvalue := C.CString(value)
	defer C.free(unsafe.Pointer(cvalue))
	ret := C.SDL_SetStringProperty(pid.c(), cname, cvalue)
	if ret < 0 {
		return GetError()
	}
	return nil
}

// Set an integer property in a group of properties.
// (https://wiki.libsdl.org/SDL3/SDL_SetNumberProperty)
func (pid PropertiesID) SetNumber(name string, value int64) error {
	cname := C.CString(name)
	defer C.free(unsafe.Pointer(cname))
	ret := C.SDL_SetNumberProperty(pid.c(), cname, C.Sint64(value))
	if ret < 0 {
		return GetError()
	}
	return nil
}

// Set a floating point property in a group of properties.
// (https://wiki.libsdl.org/SDL3/SDL_SetFloatProperty)
func (pid PropertiesID) SetFloat(name string, value float32) error {
	cname := C.CString(name)
	defer C.free(unsafe.Pointer(cname))
	ret := C.SDL_SetFloatProperty(pid.c(), cname, C.float(value))
	if ret < 0 {
		return GetError()
	}
	return nil
}

// Set a boolean property in a group of properties.
// (https://wiki.libsdl.org/SDL3/SDL_SetBooleanProperty)
func (pid PropertiesID) SetBoolean(name string, value bool) error {
	cname := C.CString(name)
	defer C.free(unsafe.Pointer(cname))
	ret := C.SDL_SetBooleanProperty(pid.c(), cname, sdlBool(value))
	if ret < 0 {
		return GetError()
	}
	return nil
}

// Return whether a property exists in a group of properties.
// (https://wiki.libsdl.org/SDL3/SDL_HasProperty)
func (pid PropertiesID) Has(name string) bool {
	cname := C.CString(name)
	defer C.free(unsafe.Pointer(cname))
	ret := C.SDL_HasProperty(pid.c(), cname)
	return ret != 0
}

// Get the type of a property in a group of properties.
// (https://wiki.libsdl.org/SDL3/SDL_GetPropertyType)
func (pid PropertiesID) GetType(name string) PropertyType {
	cname := C.CString(name)
	defer C.free(unsafe.Pointer(cname))
	ret := C.SDL_GetPropertyType(pid.c(), cname)
	return PropertyType(ret)
}

// Get a pointer property from a group of properties.
// (https://wiki.libsdl.org/SDL3/SDL_GetPointerProperty)
func (pid PropertiesID) GetPointer(name string, defaultValue unsafe.Pointer) unsafe.Pointer {
	cname := C.CString(name)
	defer C.free(unsafe.Pointer(cname))
	return C.SDL_GetPointerProperty(pid.c(), cname, defaultValue)
}

// Get a string property from a group of properties.
// (https://wiki.libsdl.org/SDL3/SDL_GetStringProperty)
func (pid PropertiesID) GetString(name string, defaultValue string) string {
	cname := C.CString(name)
	defer C.free(unsafe.Pointer(cname))
	cvalue := C.CString(defaultValue)
	defer C.free(unsafe.Pointer(cvalue))
	return C.GoString(C.SDL_GetStringProperty(pid.c(), cname, cvalue))
}

// Get a number property from a group of properties.
// (https://wiki.libsdl.org/SDL3/SDL_GetNumberProperty)
func (pid PropertiesID) GetNumber(name string, defaultValue int64) int64 {
	cname := C.CString(name)
	defer C.free(unsafe.Pointer(cname))
	return int64(C.SDL_GetNumberProperty(pid.c(), cname, C.Sint64(defaultValue)))
}

// Get a floating point property from a group of properties.
// (https://wiki.libsdl.org/SDL3/SDL_GetFloatProperty)
func (pid PropertiesID) GetFloat(name string, defaultValue float32) float32 {
	cname := C.CString(name)
	defer C.free(unsafe.Pointer(cname))
	return float32(C.SDL_GetFloatProperty(pid.c(), cname, C.float(defaultValue)))
}

// Get a boolean property from a group of properties.
// (https://wiki.libsdl.org/SDL3/SDL_GetBooleanProperty)
func (pid PropertiesID) GetBoolean(name string, defaultValue bool) bool {
	cname := C.CString(name)
	defer C.free(unsafe.Pointer(cname))
	return C.SDL_GetBooleanProperty(pid.c(), cname, sdlBool(defaultValue)) != 0
}

// Clear a property from a group of properties.
// (https://wiki.libsdl.org/SDL3/SDL_ClearProperty)
func (pid PropertiesID) Clear(name string) error {
	cname := C.CString(name)
	defer C.free(unsafe.Pointer(cname))
	ret := C.SDL_ClearProperty(pid.c(), cname)
	if ret < 0 {
		return GetError()
	}
	return nil
}

// A callback used to enumerate all the properties in a group of properties.
// (https://wiki.libsdl.org/SDL3/SDL_EnumeratePropertiesCallback)
type EnumeratePropertiesCallback func(props PropertiesID, name string)

//export goEnumeratePropertiesCallback
func goEnumeratePropertiesCallback(userdata unsafe.Pointer, props PropertiesID, name *C.char) {
	callback := *(*EnumeratePropertiesCallback)(userdata)
	callback(PropertiesID(props), C.GoString(name))
}

// Enumerate the properties contained in a group of properties.
// (https://wiki.libsdl.org/SDL3/SDL_EnumerateProperties)
func (pid PropertiesID) Enumerate(callback EnumeratePropertiesCallback) error {
	ret := C.SDL_EnumerateProperties(pid.c(), C.SDL_EnumeratePropertiesCallback(C.cgoEnumeratePropertiesCallback), unsafe.Pointer(&callback))
	if ret < 0 {
		return GetError()
	}
	return nil
}

// Destroy a group of properties.
// (https://wiki.libsdl.org/SDL3/SDL_DestroyProperties)
func (pid PropertiesID) Destroy() {
	C.SDL_DestroyProperties(pid.c())
}

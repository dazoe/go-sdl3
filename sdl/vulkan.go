package sdl

/*
#include <SDL3/SDL_vulkan.h>
#include "stdlib.h"
*/
import "C"
import (
	"reflect"
	"unsafe"
)

/**
 * # CategoryVulkan
 *
 * Functions for creating Vulkan surfaces on SDL windows.
 *
 * (https://wiki.libsdl.org/SDL3/CategoryVulkan)
 */

// VK_DEFINE_HANDLE(VkInstance)
// VK_DEFINE_HANDLE(VkPhysicalDevice)
// VK_DEFINE_NON_DISPATCHABLE_HANDLE(VkSurfaceKHR)
// struct VkAllocationCallbacks;

/**
 *  Vulkan support functions
 */

// Dynamically load the Vulkan loader library.
// (https://wiki.libsdl.org/SDL3/SDL_Vulkan_LoadLibrary)
func Vulkan_LoadLibrary(path string) error {
	cpath := C.CString(path)
	defer C.free(unsafe.Pointer(cpath))
	if !C.SDL_Vulkan_LoadLibrary(cpath) {
		return GetError()
	}
	return nil
}

// Get the address of the `vkGetInstanceProcAddr` function.
// (https://wiki.libsdl.org/SDL3/SDL_Vulkan_GetVkGetInstanceProcAddr)
func Vulkan_GetVkGetInstanceProcAddr() unsafe.Pointer {
	return unsafe.Pointer(C.SDL_Vulkan_GetVkGetInstanceProcAddr())
}

// Unload the Vulkan library previously loaded by SDL_Vulkan_LoadLibrary().
// (https://wiki.libsdl.org/SDL3/SDL_Vulkan_UnloadLibrary)
func Vulkan_UnloadLibrary() {
	C.SDL_Vulkan_UnloadLibrary()
}

// Get the Vulkan instance extensions needed for vkCreateInstance.
// (https://wiki.libsdl.org/SDL3/SDL_Vulkan_GetInstanceExtensions)
func Vulkan_GetInstanceExtensions() ([]string, error) {
	var count C.Uint32
	ret := C.SDL_Vulkan_GetInstanceExtensions(&count)
	if ret == nil {
		return nil, GetError()
	}
	cSlice := unsafe.Slice(ret, count)
	goSlice := make([]string, count)
	for i := 0; i < int(count); i++ {
		goSlice[i] = C.GoString(cSlice[i])
	}
	return goSlice, nil
}

// Create a Vulkan rendering surface for a window.
// (https://wiki.libsdl.org/SDL3/SDL_Vulkan_CreateSurface)
func (w *Window) Vulkan_CreateSurface(vkInstance any, vkAllocationCallbacks any) (surface unsafe.Pointer, err error) {
	if vkInstance == nil {
		return nil, errorSDL("instance cannot be nil")
	}
	var instance C.VkInstance
	val := reflect.ValueOf(vkInstance)
	if val.Kind() != reflect.Pointer {
		return nil, errorSDL("invalid parameter: vkInstance should be a pointer")
	}
	instance = C.VkInstance(unsafe.Pointer(val.Pointer()))
	var allocationCallbacks *C.struct_VkAllocationCallbacks
	if vkAllocationCallbacks != nil {
		val = reflect.ValueOf(vkAllocationCallbacks)
		if val.Kind() != reflect.Pointer {
			return nil, errorSDL("invalid parameter: vkAllocationCallbacks should be a pointer")
		}
		allocationCallbacks = (*C.struct_VkAllocationCallbacks)(unsafe.Pointer(val.Pointer()))
	}
	var vkSurface C.VkSurfaceKHR
	if !C.SDL_Vulkan_CreateSurface(w.cptr(), instance, allocationCallbacks, &vkSurface) {
		return nil, GetError()
	}
	return unsafe.Pointer(&vkSurface), nil
}

// Destroy the Vulkan rendering surface of a window.
// (https://wiki.libsdl.org/SDL3/SDL_Vulkan_DestroySurface)
func Vulkan_DestroySurface(vkInstance any, surface unsafe.Pointer, vkAllocationCallbacks any) error {
	if vkInstance == nil {
		return errorSDL("instance cannot be nil")
	}
	val := reflect.ValueOf(vkInstance)
	if val.Kind() != reflect.Pointer {
		return errorSDL("invalid parameter: vkInstance should be a pointer")
	}
	instance := C.VkInstance(unsafe.Pointer(val.Pointer()))
	var allocationCallbacks *C.struct_VkAllocationCallbacks
	if vkAllocationCallbacks != nil {
		val = reflect.ValueOf(vkAllocationCallbacks)
		if val.Kind() != reflect.Pointer {
			return errorSDL("invalid parameter: vkAllocationCallbacks should be a pointer")
		}
		allocationCallbacks = (*C.struct_VkAllocationCallbacks)(unsafe.Pointer(val.Pointer()))
	}
	C.SDL_Vulkan_DestroySurface(instance, C.VkSurfaceKHR(surface), allocationCallbacks)
	return nil
}

// Query support for presentation via a given physical device and queue family.
// (https://wiki.libsdl.org/SDL3/SDL_Vulkan_GetPresentationSupport)
func Vulkan_GetPresentationSupport(vkInstance any, vkPhysicalDevice any, queueFamilyIndex uint32) (bool, error) {
	if vkInstance == nil {
		return false, errorSDL("invalid parameter: vkInstance cannot be nil")
	}
	val := reflect.ValueOf(vkInstance)
	if val.Kind() != reflect.Pointer {
		return false, errorSDL("invalid parameter: vkInstance should be a pointer")
	}
	instance := C.VkInstance(unsafe.Pointer(val.Pointer()))

	if vkPhysicalDevice == nil {
		return false, errorSDL("invalid parameter: vkPhysicalDevice cannot be nil")
	}
	val = reflect.ValueOf(vkPhysicalDevice)
	if val.Kind() != reflect.Pointer {
		return false, errorSDL("invalid parameter: vkPhysicalDevice should be a pointer")
	}
	physicalDevice := C.VkPhysicalDevice(unsafe.Pointer(val.Pointer()))

	return bool(C.SDL_Vulkan_GetPresentationSupport(instance, physicalDevice, C.Uint32(queueFamilyIndex))), nil
}

package sdl

// #include <SDL3/SDL_locale.h>
import "C"
import "unsafe"

// # CategoryLocale
// (https://wiki.libsdl.org/SDL3/CategoryLocale)

// A struct to provide locale data.
// (https://wiki.libsdl.org/SDL3/SDL_Locale)
type Locale struct {
	Language string
	Country  string
}

// Report the user's preferred locale.
// (https://wiki.libsdl.org/SDL3/SDL_GetPreferredLocales)
func GetPreferredLocales() (locales []Locale, err error) {
	var count C.int
	ret := C.SDL_GetPreferredLocales(&count)
	if ret == nil {
		return nil, GetError()
	}
	defer C.SDL_free(unsafe.Pointer(ret))
	cSlice := unsafe.Slice((**C.SDL_Locale)(ret), int(count))
	result := make([]Locale, int(count))
	for i, cLocale := range cSlice {
		result[i].Language = C.GoString(cLocale.language)
		result[i].Country = C.GoString(cLocale.country)
	}
	return result, nil
}

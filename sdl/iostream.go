package sdl

/*
#include <SDL3/SDL_iostream.h>

size_t SDL_IOprintfGo(SDL_IOStream *stream, const char* str) {
	return SDL_IOprintf(stream, str);
}
*/
import "C"
import (
	"fmt"
	"unsafe"
)

// SDL_IOStream status, set by a read or write operation.
// (https://wiki.libsdl.org/SDL3/SDL_IOStatus)
type IOStatus C.SDL_IOStatus

const (
	IO_STATUS_READY     = IOStatus(C.SDL_IO_STATUS_READY)     /**< Everything is ready (no errors and not EOF). */
	IO_STATUS_ERROR     = IOStatus(C.SDL_IO_STATUS_ERROR)     /**< Read or write I/O error */
	IO_STATUS_EOF       = IOStatus(C.SDL_IO_STATUS_EOF)       /**< End of file */
	IO_STATUS_NOT_READY = IOStatus(C.SDL_IO_STATUS_NOT_READY) /**< Non blocking I/O, not ready */
	IO_STATUS_READONLY  = IOStatus(C.SDL_IO_STATUS_READONLY)  /**< Tried to write a read-only buffer */
	IO_STATUS_WRITEONLY = IOStatus(C.SDL_IO_STATUS_WRITEONLY) /**< Tried to read a write-only buffer */
)

// Possible `whence` values for SDL_IOStream seeking.
// (https://wiki.libsdl.org/SDL3/SDL_IOStatus)
type IOWhence C.SDL_IOWhence

const (
	IO_SEEK_SET = IOWhence(C.SDL_IO_SEEK_SET) /**< Seek from the beginning of data */
	IO_SEEK_CUR = IOWhence(C.SDL_IO_SEEK_CUR) /**< Seek relative to current read point */
	IO_SEEK_END = IOWhence(C.SDL_IO_SEEK_END) /**< Seek relative to the end of data */
)

// TODO: needs implemented
// The function pointers that drive an SDL_IOStream.
// (https://wiki.libsdl.org/SDL3/SDL_IOStatus)
//  typedef struct SDL_IOStreamInterface {
// 	 Sint64 (SDLCALL *size)(void *userdata);
// 	 Sint64 (SDLCALL *seek)(void *userdata, Sint64 offset, SDL_IOWhence whence);
// 	 size_t (SDLCALL *read)(void *userdata, void *ptr, size_t size, SDL_IOStatus *status);
// 	 size_t (SDLCALL *write)(void *userdata, const void *ptr, size_t size, SDL_IOStatus *status);
// 	 int (SDLCALL *close)(void *userdata);
//  } SDL_IOStreamInterface;

// The read/write operation structure.
// (https://wiki.libsdl.org/SDL3/SDL_IOStatus)
type IOStream C.SDL_IOStream

func (stream *IOStream) cptr() *C.SDL_IOStream {
	return (*C.SDL_IOStream)(stream)
}

// Use this function to create a new SDL_IOStream structure for reading from
// and/or writing to a named file.
// (https://wiki.libsdl.org/SDL3/SDL_IOFromFile)
func IOFromFile(file, mode string) (stream *IOStream, err error) {
	stream = (*IOStream)(C.SDL_IOFromFile(C.CString(file), C.CString(mode)))
	if stream == nil {
		err = GetError()
	}
	return
}

const (
	PROP_IOSTREAM_WINDOWS_HANDLE_POINTER = "SDL.iostream.windows.handle"
	PROP_IOSTREAM_STDIO_FILE_POINTER     = "SDL.iostream.stdio.file"
	PROP_IOSTREAM_ANDROID_AASSET_POINTER = "SDL.iostream.android.aasset"
)

// Use this function to prepare a read-write memory buffer for use with
// SDL_IOStream.
// (https://wiki.libsdl.org/SDL3/SDL_IOFromMem)
func IOFromMem(mem []byte) (stream *IOStream, err error) {
	data := unsafe.SliceData(mem)
	stream = (*IOStream)(C.SDL_IOFromMem(unsafe.Pointer(data), C.size_t(len(mem))))
	if stream == nil {
		err = GetError()
	}
	return
}

// Use this function to prepare a read-only memory buffer for use with
// SDL_IOStream.
// (https://wiki.libsdl.org/SDL3/SDL_IOFromConstMem)
func IOFromConstMem(mem []byte) (stream *IOStream, err error) {
	data := unsafe.SliceData(mem)
	stream = (*IOStream)(C.SDL_IOFromConstMem(unsafe.Pointer(data), C.size_t(len(mem))))
	if stream == nil {
		err = GetError()
	}
	return
}

// Use this function to create an SDL_IOStream that is backed by dynamically
// allocated memory.
// (https://wiki.libsdl.org/SDL3/SDL_IOFromDynamicMem)
func IOFromDynamicMem() (stream *IOStream, err error) {
	stream = (*IOStream)(C.SDL_IOFromDynamicMem())
	if stream == nil {
		err = GetError()
	}
	return
}

const (
	PROP_IOSTREAM_DYNAMIC_MEMORY_POINTER   = "SDL.iostream.dynamic.memory"
	PROP_IOSTREAM_DYNAMIC_CHUNKSIZE_NUMBER = "SDL.iostream.dynamic.chunksize"
)

// Create a custom SDL_IOStream.
// (https://wiki.libsdl.org/SDL3/SDL_OpenIO)
// func OpenIO(const SDL_IOStreamInterface *iface, void *userdata) SDL_IOStream * {
// 	ret := C.SDL_OpenIO(iface, userdata)
// }

// Close and free an allocated SDL_IOStream structure.
// (https://wiki.libsdl.org/SDL3/SDL_CloseIO)
func (stream *IOStream) Close() (err error) {
	ret := C.SDL_CloseIO(stream.cptr())
	if ret != 0 {
		err = GetError()
	}
	return
}

// Get the properties associated with an SDL_IOStream.
// (https://wiki.libsdl.org/SDL3/SDL_GetIOProperties)
func (stream *IOStream) GetProperties() (id PropertiesID, err error) {
	id = PropertiesID(C.SDL_GetIOProperties(stream.cptr()))
	if id != 0 {
		err = GetError()
	}
	return
}

// Query the stream status of an SDL_IOStream.
// (https://wiki.libsdl.org/SDL3/SDL_GetIOStatus)
func (stream *IOStream) GetStatus() IOStatus {
	return IOStatus(C.SDL_GetIOStatus(stream.cptr()))
}

// Use this function to get the size of the data stream in an SDL_IOStream.
// (https://wiki.libsdl.org/SDL3/SDL_GetIOSize)
func (stream *IOStream) GetSize() int64 {
	return int64(C.SDL_GetIOSize(stream.cptr()))
}

// Seek within an SDL_IOStream data stream.
// (https://wiki.libsdl.org/SDL3/SDL_SeekIO)
func (stream *IOStream) Seek(offset int64, whence IOWhence) (n int64, err error) {
	n = int64(C.SDL_SeekIO(stream.cptr(), C.Sint64(offset), C.SDL_IOWhence(whence)))
	if n < 0 {
		err = GetError()
	}
	return
}

// Determine the current read/write offset in an SDL_IOStream data stream.
// (https://wiki.libsdl.org/SDL3/SDL_TellIO)
func (stream *IOStream) Tell() int64 {
	return int64(C.SDL_TellIO(stream.cptr()))
}

func (stream *IOStream) errFromStatus() error {
	status := stream.GetStatus()
	switch status {
	case IO_STATUS_ERROR:
		return GetError()
	case IO_STATUS_EOF:
		return ErrorEOF
	case IO_STATUS_NOT_READY:
		return errorSDL("stream not ready")
	case IO_STATUS_READONLY:
		return errorSDL("tried to write a read-only buffer")
	case IO_STATUS_WRITEONLY:
		return errorSDL("tried to read a write-only buffer")
	}
	return nil
}

// Read from a data source.
// (https://wiki.libsdl.org/SDL3/SDL_ReadIO)
func (stream *IOStream) Read(buf []byte) (n int, err error) {
	ptr := unsafe.Pointer(unsafe.SliceData(buf))
	n = int(C.SDL_ReadIO(stream.cptr(), ptr, C.size_t(len(buf))))
	if n == 0 {
		err = stream.errFromStatus()
	}
	return
}

// Write to an SDL_IOStream data stream.
// (https://wiki.libsdl.org/SDL3/SDL_WriteIO)
func (stream *IOStream) Write(buf []byte) (n int, err error) {
	ptr := unsafe.Pointer(unsafe.SliceData(buf))
	n = int(C.SDL_WriteIO(stream.cptr(), ptr, C.size_t(len(buf))))
	if n < len(buf) {
		err = stream.errFromStatus()
	}
	return
}

// Print to an SDL_IOStream data stream.
// (https://wiki.libsdl.org/SDL3/SDL_IOprintf)
func (stream *IOStream) Printf(str string, a ...any) (n int, err error) {
	str = fmt.Sprintf(str, a)
	n = int(C.SDL_IOprintfGo(stream.cptr(), C.CString(str)))
	if n == 0 {
		err = GetError()
	}
	return
}

// // Print to an SDL_IOStream data stream.
// // (https://wiki.libsdl.org/SDL3/SDL_IOvprintf)
// func IOvprintf(SDL_IOStream *context, SDL_PRINTF_FORMAT_STRING const char *fmt, va_list ap) SDL_PRINTF_VARARG_FUNCV(2) size_t {
// 	ret := C.SDL_IOvprintf(context, fmt, SDL_PRINTF_VARARG_FUNCV(2)
// }

// // Load all the data from an SDL data stream.
// // (https://wiki.libsdl.org/SDL3/SDL_LoadFile_IO)
// func LoadFile_IO(SDL_IOStream *src, size_t *datasize, SDL_bool closeio) void * {
// 	ret := C.SDL_LoadFile_IO(src, datasize, closeio)
// }

// // Load all the data from a file path.
// // (https://wiki.libsdl.org/SDL3/SDL_LoadFile)
// func LoadFile(const char *file, size_t *datasize) void * {
// 	ret := C.SDL_LoadFile(file, datasize)
// }

// //  \name Read endian functions
// // (https://wiki.libsdl.org/SDL3/SDL_ReadU8)
// func ReadU8(SDL_IOStream *src, Uint8 *value) SDL_bool {
// 	ret := C.SDL_ReadU8(src, value)
// }

// // Use this function to read a signed byte from an SDL_IOStream.
// // (https://wiki.libsdl.org/SDL3/SDL_ReadS8)
// func ReadS8(SDL_IOStream *src, Sint8 *value) SDL_bool {
// 	ret := C.SDL_ReadS8(src, value)
// }

// // Use this function to read 16 bits of little-endian data from an
// // SDL_IOStream and return in native format.
// // (https://wiki.libsdl.org/SDL3/SDL_ReadU16LE)
// func ReadU16LE(SDL_IOStream *src, Uint16 *value) SDL_bool {
// 	ret := C.SDL_ReadU16LE(src, value)
// }

// // Use this function to read 16 bits of little-endian data from an
// // SDL_IOStream and return in native format.
// // (https://wiki.libsdl.org/SDL3/SDL_ReadS16LE)
// func ReadS16LE(SDL_IOStream *src, Sint16 *value) SDL_bool {
// 	ret := C.SDL_ReadS16LE(src, value)
// }

// // Use this function to read 16 bits of big-endian data from an SDL_IOStream
// // and return in native format.
// // (https://wiki.libsdl.org/SDL3/SDL_ReadU16BE)
// func ReadU16BE(SDL_IOStream *src, Uint16 *value) SDL_bool {
// 	ret := C.SDL_ReadU16BE(src, value)
// }

// // Use this function to read 16 bits of big-endian data from an SDL_IOStream
// // and return in native format.
// // (https://wiki.libsdl.org/SDL3/SDL_ReadS16BE)
// func ReadS16BE(SDL_IOStream *src, Sint16 *value) SDL_bool {
// 	ret := C.SDL_ReadS16BE(src, value)
// }

// // Use this function to read 32 bits of little-endian data from an
// // SDL_IOStream and return in native format.
// // (https://wiki.libsdl.org/SDL3/SDL_ReadU32LE)
// func ReadU32LE(SDL_IOStream *src, Uint32 *value) SDL_bool {
// 	ret := C.SDL_ReadU32LE(src, value)
// }

// // Use this function to read 32 bits of little-endian data from an
// // SDL_IOStream and return in native format.
// // (https://wiki.libsdl.org/SDL3/SDL_ReadS32LE)
// func ReadS32LE(SDL_IOStream *src, Sint32 *value) SDL_bool {
// 	ret := C.SDL_ReadS32LE(src, value)
// }

// // Use this function to read 32 bits of big-endian data from an SDL_IOStream
// // and return in native format.
// // (https://wiki.libsdl.org/SDL3/SDL_ReadU32BE)
// func ReadU32BE(SDL_IOStream *src, Uint32 *value) SDL_bool {
// 	ret := C.SDL_ReadU32BE(src, value)
// }

// // Use this function to read 32 bits of big-endian data from an SDL_IOStream
// // and return in native format.
// // (https://wiki.libsdl.org/SDL3/SDL_ReadS32BE)
// func ReadS32BE(SDL_IOStream *src, Sint32 *value) SDL_bool {
// 	ret := C.SDL_ReadS32BE(src, value)
// }

// // Use this function to read 64 bits of little-endian data from an
// // SDL_IOStream and return in native format.
// // (https://wiki.libsdl.org/SDL3/SDL_ReadU64LE)
// func ReadU64LE(SDL_IOStream *src, Uint64 *value) SDL_bool {
// 	ret := C.SDL_ReadU64LE(src, value)
// }

// // Use this function to read 64 bits of little-endian data from an
// // SDL_IOStream and return in native format.
// // (https://wiki.libsdl.org/SDL3/SDL_ReadS64LE)
// func ReadS64LE(SDL_IOStream *src, Sint64 *value) SDL_bool {
// 	ret := C.SDL_ReadS64LE(src, value)
// }

// // Use this function to read 64 bits of big-endian data from an SDL_IOStream
// // and return in native format.
// // (https://wiki.libsdl.org/SDL3/SDL_ReadU64BE)
// func ReadU64BE(SDL_IOStream *src, Uint64 *value) SDL_bool {
// 	ret := C.SDL_ReadU64BE(src, value)
// }

// // Use this function to read 64 bits of big-endian data from an SDL_IOStream
// // and return in native format.
// // (https://wiki.libsdl.org/SDL3/SDL_ReadS64BE)
// func ReadS64BE(SDL_IOStream *src, Sint64 *value) SDL_bool {
// 	ret := C.SDL_ReadS64BE(src, value)
// }

// //  \name Write endian functions
// // (https://wiki.libsdl.org/SDL3/SDL_WriteU8)
// func WriteU8(SDL_IOStream *dst, Uint8 value) SDL_bool {
// 	ret := C.SDL_WriteU8(dst, value)
// }

// // Use this function to write a signed byte to an SDL_IOStream.
// // (https://wiki.libsdl.org/SDL3/SDL_WriteS8)
// func WriteS8(SDL_IOStream *dst, Sint8 value) SDL_bool {
// 	ret := C.SDL_WriteS8(dst, value)
// }

// // Use this function to write 16 bits in native format to an SDL_IOStream as
// // little-endian data.
// // (https://wiki.libsdl.org/SDL3/SDL_WriteU16LE)
// func WriteU16LE(SDL_IOStream *dst, Uint16 value) SDL_bool {
// 	ret := C.SDL_WriteU16LE(dst, value)
// }

// // Use this function to write 16 bits in native format to an SDL_IOStream as
// // little-endian data.
// // (https://wiki.libsdl.org/SDL3/SDL_WriteS16LE)
// func WriteS16LE(SDL_IOStream *dst, Sint16 value) SDL_bool {
// 	ret := C.SDL_WriteS16LE(dst, value)
// }

// // Use this function to write 16 bits in native format to an SDL_IOStream as
// // big-endian data.
// // (https://wiki.libsdl.org/SDL3/SDL_WriteU16BE)
// func WriteU16BE(SDL_IOStream *dst, Uint16 value) SDL_bool {
// 	ret := C.SDL_WriteU16BE(dst, value)
// }

// // Use this function to write 16 bits in native format to an SDL_IOStream as
// // big-endian data.
// // (https://wiki.libsdl.org/SDL3/SDL_WriteS16BE)
// func WriteS16BE(SDL_IOStream *dst, Sint16 value) SDL_bool {
// 	ret := C.SDL_WriteS16BE(dst, value)
// }

// // Use this function to write 32 bits in native format to an SDL_IOStream as
// // little-endian data.
// // (https://wiki.libsdl.org/SDL3/SDL_WriteU32LE)
// func WriteU32LE(SDL_IOStream *dst, Uint32 value) SDL_bool {
// 	ret := C.SDL_WriteU32LE(dst, value)
// }

// // Use this function to write 32 bits in native format to an SDL_IOStream as
// // little-endian data.
// // (https://wiki.libsdl.org/SDL3/SDL_WriteS32LE)
// func WriteS32LE(SDL_IOStream *dst, Sint32 value) SDL_bool {
// 	ret := C.SDL_WriteS32LE(dst, value)
// }

// // Use this function to write 32 bits in native format to an SDL_IOStream as
// // big-endian data.
// // (https://wiki.libsdl.org/SDL3/SDL_WriteU32BE)
// func WriteU32BE(SDL_IOStream *dst, Uint32 value) SDL_bool {
// 	ret := C.SDL_WriteU32BE(dst, value)
// }

// // Use this function to write 32 bits in native format to an SDL_IOStream as
// // big-endian data.
// // (https://wiki.libsdl.org/SDL3/SDL_WriteS32BE)
// func WriteS32BE(SDL_IOStream *dst, Sint32 value) SDL_bool {
// 	ret := C.SDL_WriteS32BE(dst, value)
// }

// // Use this function to write 64 bits in native format to an SDL_IOStream as
// // little-endian data.
// // (https://wiki.libsdl.org/SDL3/SDL_WriteU64LE)
// func WriteU64LE(SDL_IOStream *dst, Uint64 value) SDL_bool {
// 	ret := C.SDL_WriteU64LE(dst, value)
// }

// // Use this function to write 64 bits in native format to an SDL_IOStream as
// // little-endian data.
// // (https://wiki.libsdl.org/SDL3/SDL_WriteS64LE)
// func WriteS64LE(SDL_IOStream *dst, Sint64 value) SDL_bool {
// 	ret := C.SDL_WriteS64LE(dst, value)
// }

// // Use this function to write 64 bits in native format to an SDL_IOStream as
// // big-endian data.
// // (https://wiki.libsdl.org/SDL3/SDL_WriteU64BE)
// func WriteU64BE(SDL_IOStream *dst, Uint64 value) SDL_bool {
// 	ret := C.SDL_WriteU64BE(dst, value)
// }

// // Use this function to write 64 bits in native format to an SDL_IOStream as
// // big-endian data.
// // (https://wiki.libsdl.org/SDL3/SDL_WriteS64BE)
// func WriteS64BE(SDL_IOStream *dst, Sint64 value) SDL_bool {
// 	ret := C.SDL_WriteS64BE(dst, value)
// }

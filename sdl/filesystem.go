package sdl

/*
#include <SDL3/SDL_filesystem.h>
#include <stdlib.h>
*/
import "C"
import "unsafe"

// TODO: basically this whole file needs tested.
// # CategoryFilesystem
// (https://wiki.libsdl.org/CategoryFilesystem)

// Get the directory where the application was run from.
// (https://wiki.libsdl.org/SDL_GetBasePath)
func GetBasePath() (string, error) {
	ret := C.SDL_GetBasePath()
	if ret == nil {
		return "", GetError()
	}
	return C.GoString(ret), nil
}

// Get the user-and-app-specific path where files can be written.
// (https://wiki.libsdl.org/SDL_GetPrefPath)
func GetPrefPath(org, app string) (string, error) {
	cOrg := C.CString(org)
	defer C.free(unsafe.Pointer(cOrg))
	cApp := C.CString(app)
	defer C.free(unsafe.Pointer(cApp))
	ret := C.SDL_GetPrefPath(cOrg, cApp)
	if ret == nil {
		return "", GetError()
	}
	defer C.SDL_free(unsafe.Pointer(ret))
	return C.GoString(ret), nil
}

// The type of the OS-provided default folder for a specific purpose.
// (https://wiki.libsdl.org/SDL_Folder)
type Folder C.SDL_Folder

const (
	FOLDER_HOME        = Folder(C.SDL_FOLDER_HOME)        /**< The folder which contains all of the current user's data, preferences, and documents. It usually contains most of the other folders. If a requested folder does not exist, the home folder can be considered a safe fallback to store a user's documents. */
	FOLDER_DESKTOP     = Folder(C.SDL_FOLDER_DESKTOP)     /**< The folder of files that are displayed on the desktop. Note that the existence of a desktop folder does not guarantee that the system does show icons on its desktop; certain GNU/Linux distros with a graphical environment may not have desktop icons. */
	FOLDER_DOCUMENTS   = Folder(C.SDL_FOLDER_DOCUMENTS)   /**< User document files, possibly application-specific. This is a good place to save a user's projects. */
	FOLDER_DOWNLOADS   = Folder(C.SDL_FOLDER_DOWNLOADS)   /**< Standard folder for user files downloaded from the internet. */
	FOLDER_MUSIC       = Folder(C.SDL_FOLDER_MUSIC)       /**< Music files that can be played using a standard music player (mp3, ogg...). */
	FOLDER_PICTURES    = Folder(C.SDL_FOLDER_PICTURES)    /**< Image files that can be displayed using a standard viewer (png, jpg...). */
	FOLDER_PUBLICSHARE = Folder(C.SDL_FOLDER_PUBLICSHARE) /**< Files that are meant to be shared with other users on the same computer. */
	FOLDER_SAVEDGAMES  = Folder(C.SDL_FOLDER_SAVEDGAMES)  /**< Save files for games. */
	FOLDER_SCREENSHOTS = Folder(C.SDL_FOLDER_SCREENSHOTS) /**< Application screenshots. */
	FOLDER_TEMPLATES   = Folder(C.SDL_FOLDER_TEMPLATES)   /**< Template files to be used when the user requests the desktop environment to create a new file in a certain folder, such as "New Text File.txt".  Any file in the Templates folder can be used as a starting point for a new file. */
	FOLDER_VIDEOS      = Folder(C.SDL_FOLDER_VIDEOS)      /**< Video files that can be played using a standard video player (mp4, webm...). */
	FOLDER_COUNT       = Folder(C.SDL_FOLDER_COUNT)       /**< Total number of types in this enum, not a folder type by itself. */
)

// Finds the most suitable user folder for a specific purpose.
// (https://wiki.libsdl.org/SDL_GetUserFolder)
func GetUserFolder(folder Folder) (string, error) {
	ret := C.SDL_GetUserFolder(C.SDL_Folder(folder))
	if ret == nil {
		return "", GetError()
	}
	return C.GoString(ret), nil
}

// Abstract filesystem interface
type PathType C.SDL_PathType

const (
	PATHTYPE_NONE      = PathType(C.SDL_PATHTYPE_NONE)      /**< path does not exist */
	PATHTYPE_FILE      = PathType(C.SDL_PATHTYPE_FILE)      /**< a normal file */
	PATHTYPE_DIRECTORY = PathType(C.SDL_PATHTYPE_DIRECTORY) /**< a directory */
	PATHTYPE_OTHER     = PathType(C.SDL_PATHTYPE_OTHER)     /**< something completely different like a device node (not a symlink, those are always followed) */
)

type PathInfo struct {
	Type       PathType /**< the path type */
	Size       uint64   /**< the file size in bytes */
	CreateTime Time     /**< the time when the path was created */
	ModifyTime Time     /**< the last time the path was modified */
	AccessTime Time     /**< the last time the path was read */
}

func (p *PathInfo) cptr() *C.SDL_PathInfo {
	return (*C.SDL_PathInfo)(unsafe.Pointer(p))
}

// Flags for path matching
// (https://wiki.libsdl.org/SDL_GlobFlags)
type GlobFlags C.SDL_GlobFlags

const (
	GLOB_CASEINSENSITIVE = GlobFlags(C.SDL_GLOB_CASEINSENSITIVE)
)

// Create a directory, and any missing parent directories.
// (https://wiki.libsdl.org/SDL_CreateDirectory)
func CreateDirectory(path string) error {
	cPath := C.CString(path)
	defer C.free(unsafe.Pointer(cPath))
	if !C.SDL_CreateDirectory(cPath) {
		return GetError()
	}
	return nil
}

// Possible results from an enumeration callback.
// (https://wiki.libsdl.org/SDL_EnumerationResult)
type EnumerationResult C.SDL_EnumerationResult

const (
	ENUM_CONTINUE = EnumerationResult(C.SDL_ENUM_CONTINUE) /**< Value that requests that enumeration continue. */
	ENUM_SUCCESS  = EnumerationResult(C.SDL_ENUM_SUCCESS)  /**< Value that requests that enumeration stop, successfully. */
	ENUM_FAILURE  = EnumerationResult(C.SDL_ENUM_FAILURE)  /**< Value that requests that enumeration stop, as a failure. */

)

// Callback for directory enumeration.
// (https://wiki.libsdl.org/SDL_EnumerateDirectoryCallback)
type EnumerateDirectoryCallback func(userdata unsafe.Pointer, dirname string, fname string) EnumerationResult

// Enumerate a directory through a callback function.
// (https://wiki.libsdl.org/SDL_EnumerateDirectory)
func EnumerateDirectory(path string, callback EnumerateDirectoryCallback, userdata unsafe.Pointer) error {
	panic("not implemented")
}

//  extern SDL_DECLSPEC bool SDLCALL SDL_EnumerateDirectory(const char *path, SDL_EnumerateDirectoryCallback callback, void *userdata);

// Remove a file or an empty directory.
// (https://wiki.libsdl.org/SDL_RemovePath)
func RemovePath(path string) error {
	cPath := C.CString(path)
	defer C.free(unsafe.Pointer(cPath))
	if !C.SDL_RemovePath(cPath) {
		return GetError()
	}
	return nil
}

// Rename a file or directory.
// (https://wiki.libsdl.org/SDL_RenamePath)
func RenamePath(oldpath string, newpath string) error {
	cOldpath := C.CString(oldpath)
	defer C.free(unsafe.Pointer(cOldpath))
	cNewpath := C.CString(newpath)
	defer C.free(unsafe.Pointer(cNewpath))
	if !C.SDL_RenamePath(cOldpath, cNewpath) {
		return GetError()
	}
	return nil
}

// Copy a file.
// (https://wiki.libsdl.org/SDL_CopyFile)
func CopyFile(oldpath string, newpath string) error {
	cOldpath := C.CString(oldpath)
	defer C.free(unsafe.Pointer(cOldpath))
	cNewpath := C.CString(newpath)
	defer C.free(unsafe.Pointer(cNewpath))
	if !C.SDL_CopyFile(cOldpath, cNewpath) {
		return GetError()
	}
	return nil
}

// Get information about a filesystem path.
// (https://wiki.libsdl.org/SDL_GetPathInfo)
func GetPathInfo(path string) (info PathInfo, err error) {
	cPath := C.CString(path)
	defer C.free(unsafe.Pointer(cPath))
	if !C.SDL_GetPathInfo(cPath, info.cptr()) {
		return info, GetError()
	}
	return info, nil
}

// Enumerate a directory tree, filtered by pattern, and return a list.
// (https://wiki.libsdl.org/SDL_GlobDirectory)
func GlobDirectory(path string, pattern string, flags GlobFlags) (files []string, err error) {
	cPath := C.CString(path)
	defer C.free(unsafe.Pointer(cPath))
	cPattern := C.CString(pattern)
	defer C.free(unsafe.Pointer(cPattern))
	var count C.int
	cFiles := C.SDL_GlobDirectory(cPath, cPattern, C.SDL_GlobFlags(flags), &count)
	if cFiles == nil {
		return nil, GetError()
	}
	defer C.SDL_free(unsafe.Pointer(cFiles))
	cSlice := unsafe.Slice(cFiles, count)
	files = make([]string, len(cSlice))
	for i, cStr := range cSlice {
		files[i] = C.GoString(cStr)
	}
	return files, nil
}

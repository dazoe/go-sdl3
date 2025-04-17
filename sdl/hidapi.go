package sdl

/*
#include <SDL3/SDL_hidapi.h>
#include <stdlib.h>
*/
import "C"
import (
	"fmt"
	"unsafe"
)

// # CategoryHIDAPI
// (https://wiki.libsdl.org/SDL3/CategoryHIDAPI)

// An opaque handle representing an open HID device.
// (https://wiki.libsdl.org/SDL3/SDL_hid_device)
type HIDDevice C.SDL_hid_device

// HID underlying bus types.
// (https://wiki.libsdl.org/SDL3/SDL_hid_bus_type)
type HIDBusType C.SDL_hid_bus_type

const (
	/** Unknown bus type */
	HID_API_BUS_UNKNOWN HIDBusType = C.SDL_HID_API_BUS_UNKNOWN

	/** USB bus
	Specifications:
	https://usb.org/hid */
	HID_API_BUS_USB HIDBusType = C.SDL_HID_API_BUS_USB

	/** Bluetooth or Bluetooth LE bus
	Specifications:
	https://www.bluetooth.com/specifications/specs/human-interface-device-profile-1-1-1/
	https://www.bluetooth.com/specifications/specs/hid-service-1-0/
	https://www.bluetooth.com/specifications/specs/hid-over-gatt-profile-1-0/ */
	HID_API_BUS_BLUETOOTH HIDBusType = C.SDL_HID_API_BUS_BLUETOOTH

	/** I2C bus
	Specifications:
	https://docs.microsoft.com/previous-versions/windows/hardware/design/dn642101(v=vs.85) */
	HID_API_BUS_I2C HIDBusType = C.SDL_HID_API_BUS_I2C

	/** SPI bus
	Specifications:
	https://www.microsoft.com/download/details.aspx?id=103325 */
	HID_API_BUS_SPI HIDBusType = C.SDL_HID_API_BUS_SPI
)

// hidapi info structure

// TODO: will need special handling for this struct, it's a linked list...
// Information about a connected HID device
// (https://wiki.libsdl.org/SDL3/SDL_hid_device_info)
type HIDDeviceInfo struct {
	/** Platform-specific device path */
	Path string // char
	/** Device Vendor ID */
	VendorID uint16
	/** Device Product ID */
	ProductID uint16
	/** Serial Number */
	SerialNumber string //	wchar_t
	/** Device Release Number in binary-coded decimal,
	also known as Device Version Number */
	ReleaseNumber uint16
	/** Manufacturer String */
	ManufacturerString string //	wchar_t
	/** Product string */
	ProductString string //	wchar_t
	/** Usage Page for this Device/Interface
	(Windows/Mac/hidraw only) */
	UsagePage uint16
	/** Usage for this Device/Interface
	(Windows/Mac/hidraw only) */
	Usage uint16
	/** The USB interface which this logical device
	represents.

	Valid only if the device is a USB HID device.
	Set to -1 in all other cases.
	*/
	InterfaceNumber int32

	/** Additional information about the USB interface.
	Valid on libusb and Android implementations. */
	InterfaceClass    int32
	InterfaceSubclass int32
	InterfaceProtocol int32

	/** Underlying bus type */
	BusType HIDBusType

	/** Pointer to the next device */
	// Next *HIDDeviceInfo
}

type HIDError struct {
	Code    int
	Message string
}

func (e HIDError) Error() string {
	return fmt.Sprintf("HID Error: %d - %s", e.Code, e.Message)
}

// Initialize the HIDAPI library.
// (https://wiki.libsdl.org/SDL3/SDL_hid_init)
func HIDInit() error {
	ret := C.SDL_hid_init()
	if ret != 0 {
		return HIDError{Code: int(ret), Message: C.GoString(C.SDL_GetError())}
	}
	return nil
}

// Finalize the HIDAPI library.
// (https://wiki.libsdl.org/SDL3/SDL_hid_exit)
func HIDExit() error {
	ret := C.SDL_hid_exit()
	if ret != 0 {
		return HIDError{Code: int(ret), Message: C.GoString(C.SDL_GetError())}
	}
	return nil
}

// Check to see if devices may have been added or removed.
// (https://wiki.libsdl.org/SDL3/SDL_hid_device_change_count)
func HIDDeviceChangeCount() uint32 {
	return uint32(C.SDL_hid_device_change_count())
}

func wcharToString(wstr *C.wchar_t) string {
	if wstr == nil {
		return ""
	}
	var str []rune
	p := uintptr(unsafe.Pointer(wstr))
	sz := unsafe.Sizeof(*wstr)
	for {
		w := *((*C.wchar_t)(unsafe.Pointer(p)))
		if w == 0 {
			break
		}
		str = append(str, rune(w))
		p += sz
	}
	return string(str)
}
func stringToWchar(s string) *C.wchar_t {
	if s == "" {
		return nil
	}
	buf := make([]C.wchar_t, len(s)+1)
	for i, r := range s {
		buf[i] = C.wchar_t(r)
	}
	return &buf[0]
}

// Enumerate the HID Devices.
// (https://wiki.libsdl.org/SDL3/SDL_hid_enumerate)
func HIDEnumerate(vendorID uint16, productID uint16) ([]*HIDDeviceInfo, error) {
	ret := C.SDL_hid_enumerate(C.uint16_t(vendorID), C.uint16_t(productID))
	if ret == nil {
		return nil, GetError()
	}
	defer C.SDL_hid_free_enumeration(ret)
	var devices []*HIDDeviceInfo
	for p := ret; p != nil; p = p.next {
		devices = append(devices, &HIDDeviceInfo{
			Path:               C.GoString(p.path),
			VendorID:           uint16(p.vendor_id),
			ProductID:          uint16(p.product_id),
			SerialNumber:       wcharToString(p.serial_number),
			ReleaseNumber:      uint16(p.release_number),
			ManufacturerString: wcharToString(p.manufacturer_string),
			ProductString:      wcharToString(p.product_string),
			UsagePage:          uint16(p.usage_page),
			Usage:              uint16(p.usage),
			InterfaceNumber:    int32(p.interface_number),
			InterfaceClass:     int32(p.interface_class),
			InterfaceSubclass:  int32(p.interface_class),
			InterfaceProtocol:  int32(p.interface_protocol),
			BusType:            HIDBusType(p.bus_type),
		})
	}
	return devices, nil
}

// Free an enumeration linked list.
// (https://wiki.libsdl.org/SDL3/SDL_hid_free_enumeration)
// func HIDFreeEnumeration(devs *HIDDeviceInfo) {
// 	C.SDL_hid_free_enumeration(devs)
// }

// Open a HID device using a Vendor ID (VID), Product ID (PID) and optionally
// a serial number.
// (https://wiki.libsdl.org/SDL3/SDL_hid_open)
func HIDOpen(vendorID uint16, productID uint16, serialNumber string) (*HIDDevice, error) {
	var wstr *C.wchar_t
	if len(serialNumber) > 0 {
		wstr = stringToWchar(serialNumber)
	}
	ret := C.SDL_hid_open(C.uint16_t(vendorID), C.uint16_t(productID), wstr)
	if ret == nil {
		return nil, GetError()
	}
	return (*HIDDevice)(ret), nil
}

// Open a HID device by its path name.
// (https://wiki.libsdl.org/SDL3/SDL_hid_open_path)
func HIDOpenPath(path string) (*HIDDevice, error) {
	cstr := C.CString(path)
	defer C.free(unsafe.Pointer(cstr))
	ret := C.SDL_hid_open_path(cstr)
	if ret == nil {
		return nil, GetError()
	}
	return (*HIDDevice)(ret), nil
}

// Write an Output report to a HID device.
// (https://wiki.libsdl.org/SDL3/SDL_hid_write)
func (dev *HIDDevice) Write(data []byte) (int32, error) {
	ret := C.SDL_hid_write((*C.SDL_hid_device)(dev), (*C.uchar)(unsafe.Pointer(&data[0])), C.size_t(len(data)))
	if ret < 0 {
		return 0, GetError()
	}
	return int32(ret), nil
}

// Read an Input report from a HID device with timeout.
// (https://wiki.libsdl.org/SDL3/SDL_hid_read_timeout)
func (dev *HIDDevice) ReadTimeout(data []byte, milliseconds int32) (int32, error) {
	ret := C.SDL_hid_read_timeout((*C.SDL_hid_device)(dev), (*C.uchar)(unsafe.Pointer(&data[0])), C.size_t(len(data)), C.int(milliseconds))
	if ret < 0 {
		return 0, GetError()
	}
	return int32(ret), nil
}

// Read an Input report from a HID device.
// (https://wiki.libsdl.org/SDL3/SDL_hid_read)
func (dev *HIDDevice) Read(data []byte) (int32, error) {
	return dev.ReadTimeout(data, -1)
	// ret := C.SDL_hid_read((*C.SDL_hid_device)(dev), (*C.uchar)(unsafe.Pointer(&data[0])), C.size_t(len(data)))
	// if ret < 0 {
	// 	return 0, GetError()
	// }
	// return int32(ret), nil
}

// Set the device handle to be non-blocking.
// (https://wiki.libsdl.org/SDL3/SDL_hid_set_nonblocking)
func (dev *HIDDevice) SetNonBlocking(nonblock int32) error {
	ret := C.SDL_hid_set_nonblocking((*C.SDL_hid_device)(dev), C.int(nonblock))
	if ret != 0 {
		return HIDError{Code: int(ret), Message: C.GoString(C.SDL_GetError())}
	}
	return nil
}

// Send a Feature report to the device.
// (https://wiki.libsdl.org/SDL3/SDL_hid_send_feature_report)
func (dev *HIDDevice) SendFeatureReport(data []byte) (int32, error) {
	ret := C.SDL_hid_send_feature_report((*C.SDL_hid_device)(dev), (*C.uchar)(unsafe.Pointer(&data[0])), C.size_t(len(data)))
	if ret < 0 {
		return 0, GetError()
	}
	return int32(ret), nil
}

// Get a feature report from a HID device.
// (https://wiki.libsdl.org/SDL3/SDL_hid_get_feature_report)
func (dev *HIDDevice) GetFeatureReport(data []byte) (int32, error) {
	ret := C.SDL_hid_get_feature_report((*C.SDL_hid_device)(dev), (*C.uchar)(unsafe.Pointer(&data[0])), C.size_t(len(data)))
	if ret < 0 {
		return 0, GetError()
	}
	return int32(ret), nil
}

// Get an input report from a HID device.
// (https://wiki.libsdl.org/SDL3/SDL_hid_get_input_report)
func (dev *HIDDevice) GetInputReport(data []byte) (int32, error) {
	ret := C.SDL_hid_get_input_report((*C.SDL_hid_device)(dev), (*C.uchar)(unsafe.Pointer(&data[0])), C.size_t(len(data)))
	if ret < 0 {
		return 0, GetError()
	}
	return int32(ret), nil
}

// Close a HID device.
// (https://wiki.libsdl.org/SDL3/SDL_hid_close)
func (dev *HIDDevice) Close() error {
	ret := C.SDL_hid_close((*C.SDL_hid_device)(dev))
	if ret != 0 {
		return HIDError{Code: int(ret), Message: C.GoString(C.SDL_GetError())}
	}
	return nil
}

// Get The Manufacturer String from a HID device.
// (https://wiki.libsdl.org/SDL3/SDL_hid_get_manufacturer_string)
func (dev *HIDDevice) GetManufacturerString(maxlen int32) (string, error) {
	var wstr *C.wchar_t
	if maxlen > 0 {
		wstr = (*C.wchar_t)(C.malloc(C.size_t(maxlen) * C.size_t(unsafe.Sizeof(C.wchar_t(0)))))
		defer C.free(unsafe.Pointer(wstr))
	}
	ret := C.SDL_hid_get_manufacturer_string((*C.SDL_hid_device)(dev), wstr, C.size_t(maxlen))
	if ret < 0 {
		return "", GetError()
	}
	return wcharToString(wstr), nil
}

// Get The Product String from a HID device.
// (https://wiki.libsdl.org/SDL3/SDL_hid_get_product_string)
func (dev *HIDDevice) GetProductString(maxlen int32) (string, error) {
	var wstr *C.wchar_t
	if maxlen > 0 {
		wstr = (*C.wchar_t)(C.malloc(C.size_t(maxlen) * C.size_t(unsafe.Sizeof(C.wchar_t(0)))))
		defer C.free(unsafe.Pointer(wstr))
	}
	ret := C.SDL_hid_get_product_string((*C.SDL_hid_device)(dev), wstr, C.size_t(maxlen))
	if ret < 0 {
		return "", GetError()
	}
	return wcharToString(wstr), nil
}

// Get The Serial Number String from a HID device.
// (https://wiki.libsdl.org/SDL3/SDL_hid_get_serial_number_string)
func (dev *HIDDevice) GetSerialNumberString(maxlen int32) (string, error) {
	var wstr *C.wchar_t
	if maxlen > 0 {
		wstr = (*C.wchar_t)(C.malloc(C.size_t(maxlen) * C.size_t(unsafe.Sizeof(C.wchar_t(0)))))
		defer C.free(unsafe.Pointer(wstr))
	}
	ret := C.SDL_hid_get_serial_number_string((*C.SDL_hid_device)(dev), wstr, C.size_t(maxlen))
	if ret < 0 {
		return "", GetError()
	}
	return wcharToString(wstr), nil
}

// Get a string from a HID device, based on its string index.
// (https://wiki.libsdl.org/SDL3/SDL_hid_get_indexed_string)
func (dev *HIDDevice) GetIndexedString(stringIndex int32, maxlen int32) (string, error) {
	var wstr *C.wchar_t
	if maxlen > 0 {
		wstr = (*C.wchar_t)(C.malloc(C.size_t(maxlen) * C.size_t(unsafe.Sizeof(C.wchar_t(0)))))
		defer C.free(unsafe.Pointer(wstr))
	}
	ret := C.SDL_hid_get_indexed_string((*C.SDL_hid_device)(dev), C.int(stringIndex), wstr, C.size_t(maxlen))
	if ret < 0 {
		return "", GetError()
	}
	return wcharToString(wstr), nil
}

// Get the device info from a HID device.
// (https://wiki.libsdl.org/SDL3/SDL_hid_get_device_info)
func (dev *HIDDevice) GetDeviceInfo() (*HIDDeviceInfo, error) {
	ret := C.SDL_hid_get_device_info((*C.SDL_hid_device)(dev))
	if ret == nil {
		return nil, GetError()
	}
	return &HIDDeviceInfo{
		Path:               C.GoString(ret.path),
		VendorID:           uint16(ret.vendor_id),
		ProductID:          uint16(ret.product_id),
		SerialNumber:       wcharToString(ret.serial_number),
		ReleaseNumber:      uint16(ret.release_number),
		ManufacturerString: wcharToString(ret.manufacturer_string),
		ProductString:      wcharToString(ret.product_string),
		UsagePage:          uint16(ret.usage_page),
		Usage:              uint16(ret.usage),
		InterfaceNumber:    int32(ret.interface_number),
		InterfaceClass:     int32(ret.interface_class),
		InterfaceSubclass:  int32(ret.interface_subclass),
		InterfaceProtocol:  int32(ret.interface_protocol),
		BusType:            HIDBusType(ret.bus_type),
	}, nil
}

// Get a report descriptor from a HID device.
// (https://wiki.libsdl.org/SDL3/SDL_hid_get_report_descriptor)
func (dev *HIDDevice) GetReportDescriptor(data []byte) (int32, error) {
	ret := C.SDL_hid_get_report_descriptor((*C.SDL_hid_device)(dev), (*C.uchar)(unsafe.Pointer(&data[0])), C.size_t(len(data)))
	if ret < 0 {
		return 0, GetError()
	}
	return int32(ret), nil
}

// Start or stop a BLE scan on iOS and tvOS to pair Steam Controllers.
// (https://wiki.libsdl.org/SDL3/SDL_hid_ble_scan)
func BLEScan(active bool) {
	C.SDL_hid_ble_scan(C.bool(active))
}

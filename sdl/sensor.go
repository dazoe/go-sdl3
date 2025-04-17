package sdl

//#include <SDL3/SDL_sensor.h>
import "C"
import "unsafe"

// # CategorySensor
// (https://wiki.libsdl.org/SDL3/CategorySensor)

type Sensor C.SDL_Sensor

func (s *Sensor) cptr() *C.SDL_Sensor {
	return (*C.SDL_Sensor)(s)
}

// This is a unique ID for a sensor for the time it is connected to the
// system, and is never reused for the lifetime of the application.
// (https://wiki.libsdl.org/SDL3/SDL_SensorID)
type SensorID C.SDL_SensorID

// A constant to represent standard gravity for accelerometer sensors.
// (https://wiki.libsdl.org/SDL3/SDL_STANDARD_GRAVITY)
const STANDARD_GRAVITY = C.SDL_STANDARD_GRAVITY

// The different sensors defined by SDL.
// (https://wiki.libsdl.org/SDL3/SDL_SensorType)
type SensorType C.SDL_SensorType

func (t SensorType) c() C.SDL_SensorType {
	return C.SDL_SensorType(t)
}

const (
	SENSOR_INVALID = C.SDL_SENSOR_INVALID /**< Returned for an invalid sensor */
	SENSOR_UNKNOWN = C.SDL_SENSOR_UNKNOWN /**< Unknown sensor type */
	SENSOR_ACCEL   = C.SDL_SENSOR_ACCEL   /**< Accelerometer */
	SENSOR_GYRO    = C.SDL_SENSOR_GYRO    /**< Gyroscope */
	SENSOR_ACCEL_L = C.SDL_SENSOR_ACCEL_L /**< Accelerometer for left Joy-Con controller and Wii nunchuk */
	SENSOR_GYRO_L  = C.SDL_SENSOR_GYRO_L  /**< Gyroscope for left Joy-Con controller */
	SENSOR_ACCEL_R = C.SDL_SENSOR_ACCEL_R /**< Accelerometer for right Joy-Con controller */
	SENSOR_GYRO_R  = C.SDL_SENSOR_GYRO_R  /**< Gyroscope for right Joy-Con controller */
)

/* Function prototypes */

// Get a list of currently connected sensors.
// (https://wiki.libsdl.org/SDL3/SDL_GetSensors)
func GetSensors() (sensors []SensorID, err error) {
	var count C.int
	ret := C.SDL_GetSensors(&count)
	if ret == nil {
		return nil, GetError()
	}
	return unsafe.Slice((*SensorID)(ret), int(count)), nil
}

// Get the implementation dependent name of a sensor.
// (https://wiki.libsdl.org/SDL3/SDL_GetSensorNameForID)
func (instance_id SensorID) GetName() string {
	return C.GoString(C.SDL_GetSensorNameForID(C.SDL_SensorID(instance_id)))
}

// Get the type of a sensor.
// (https://wiki.libsdl.org/SDL3/SDL_GetSensorTypeForID)
func (instance_id SensorID) GetType() SensorType {
	return SensorType(C.SDL_GetSensorTypeForID(C.SDL_SensorID(instance_id)))
}

// Get the platform dependent type of a sensor.
// (https://wiki.libsdl.org/SDL3/SDL_GetSensorNonPortableTypeForID)
func (instance_id SensorID) GetNonPortableType() int32 {
	return int32(C.SDL_GetSensorNonPortableTypeForID(C.SDL_SensorID(instance_id)))
}

// Open a sensor for use.
// (https://wiki.libsdl.org/SDL3/SDL_OpenSensor)
func OpenSensor(instance_id SensorID) (sensor *Sensor, err error) {
	sensor = (*Sensor)(C.SDL_OpenSensor(C.SDL_SensorID(instance_id)))
	if sensor == nil {
		err = GetError()
	}
	return
}

// Return the SDL_Sensor associated with an instance ID.
// (https://wiki.libsdl.org/SDL3/SDL_GetSensorFromID)
func (instance_id SensorID) GetSensor() (sensor *Sensor, err error) {
	sensor = (*Sensor)(C.SDL_GetSensorFromID(C.SDL_SensorID(instance_id)))
	if sensor == nil {
		err = GetError()
	}
	return
}

// Get the properties associated with a sensor.
// (https://wiki.libsdl.org/SDL3/SDL_GetSensorProperties)
func (sensor *Sensor) GetProperties() (properties PropertiesID, err error) {
	properties = (PropertiesID(C.SDL_GetSensorProperties(sensor.cptr())))
	if properties == 0 {
		err = GetError()
	}
	return
}

// Get the implementation dependent name of a sensor.
// (https://wiki.libsdl.org/SDL3/SDL_GetSensorName)
func (sensor *Sensor) GetName() (name string, err error) {
	name = C.GoString(C.SDL_GetSensorName(sensor.cptr()))
	if len(name) == 0 {
		err = GetError()
	}
	return
}

// Get the type of a sensor.
// (https://wiki.libsdl.org/SDL3/SDL_GetSensorType)
func (sensor *Sensor) GetType() SensorType {
	return SensorType(C.SDL_GetSensorType(sensor.cptr()))
}

// Get the platform dependent type of a sensor.
// (https://wiki.libsdl.org/SDL3/SDL_GetSensorNonPortableType)
func (sensor *Sensor) GetNonPortableType() int32 {
	return int32(C.SDL_GetSensorNonPortableType(sensor.cptr()))
}

// Get the instance ID of a sensor.sensors
// (https://wiki.libsdl.org/SDL3/SDL_GetSensorID)
func (sensor *Sensor) GetID() (id SensorID, err error) {
	id = SensorID(C.SDL_GetSensorID(sensor.cptr()))
	if id == 0 {
		err = GetError()
	}
	return
}

// Get the current state of an opened sensor.
// (https://wiki.libsdl.org/SDL3/SDL_GetSensorData)
func (sensor *Sensor) GetData(data []float32) (err error) {
	dataPtr := (*C.float)(unsafe.SliceData(data))
	num_values := C.int(len(data))
	ret := C.SDL_GetSensorData(sensor.cptr(), dataPtr, num_values)
	if !ret {
		err = GetError()
	}
	return
}

// Close a sensor previously opened with SDL_OpenSensor().
// (https://wiki.libsdl.org/SDL3/SDL_CloseSensor)
func (sensor *Sensor) Close() {
	C.SDL_CloseSensor(sensor.cptr())
}

// Update the current state of the open sensors.
// (https://wiki.libsdl.org/SDL3/SDL_UpdateSensors)
func UpdateSensors() {
	C.SDL_UpdateSensors()
}

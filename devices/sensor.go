package devices

type Sensor interface {
	Detect() SensorStatus
}

type SensorStatus = int

const (
	MovementDetected SensorStatus = iota
)

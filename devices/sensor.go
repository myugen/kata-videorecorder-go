package devices

type Sensor interface {
	Detect() SensorStatus
}

type SensorStatus = int

const (
	NoMovementDetected SensorStatus = iota
	MovementDetected
)

package devices

//go:generate mockgen -destination=mocks/sensormock.go -package=mocks . Sensor
type Sensor interface {
	Detect() SensorStatus
}

type SensorStatus = int

const (
	NoMovementDetected SensorStatus = iota
	MovementDetected
)

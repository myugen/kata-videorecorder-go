package devices

import "time"

type Sensor interface {
	Subscribe(listener SensorListener)
}

type SensorListener interface {
	OnStatusChange(event SensorEvent)
}

type SensorEvent struct {
	status    SensorStatus
	timestamp time.Time
}

func (s SensorEvent) IsMovementDetected() bool {
	return s.status == MovementDetected
}

func (s SensorEvent) Timestamp() time.Time {
	return s.timestamp
}

func NewSensorEvent(status SensorStatus) SensorEvent {
	return SensorEvent{
		status:    status,
		timestamp: time.Now(),
	}
}

type SensorStatus = int

const (
	NoMovementDetected SensorStatus = iota
	MovementDetected
)

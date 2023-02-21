package mocksensor

import "github.com/myugen/kata-videorecorder-go/devices"

type MockSensor struct {
	listener devices.SensorListener
}

func (m *MockSensor) Subscribe(listener devices.SensorListener) {
	m.listener = listener
}

func (m *MockSensor) SimulateEventWith(simulatedEvent devices.SensorEvent) {
	m.listener.OnStatusChange(simulatedEvent)
}

func New() *MockSensor {
	return &MockSensor{}
}

var (
	SimulatedMovementDetectedEvent   = devices.NewSensorEvent(devices.MovementDetected)
	SimulatedNoMovementDetectedEvent = devices.NewSensorEvent(devices.NoMovementDetected)
)

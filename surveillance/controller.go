package surveillance

import (
	"github.com/myugen/kata-videorecorder-go/devices"
)

type Controller struct {
	sensors   []devices.Sensor
	recorders []devices.Recorder
}

func (c Controller) CheckSensors() {
	sensorsDetectSomething := false
	for _, sensor := range c.sensors {
		sensorStatus := sensor.Detect()
		if sensorStatus == devices.MovementDetected {
			sensorsDetectSomething = true
		}
	}

	if sensorsDetectSomething {
		for _, recorder := range c.recorders {
			recorder.StartRecording()
		}
	}
}

func NewController(sensors []devices.Sensor, recorders []devices.Recorder) *Controller {
	return &Controller{sensors: sensors, recorders: recorders}
}

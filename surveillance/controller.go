package surveillance

import (
	"github.com/myugen/kata-videorecorder-go/devices"
)

type Controller struct {
	sensors   []devices.Sensor
	recorders []devices.Recorder
}

func (c Controller) Scan() {
	someSensorsDetectedMovement := c.checkAllSensors()

	if someSensorsDetectedMovement {
		c.startRecordingAllCameras()
	} else {
		c.stopRecordingAllCameras()
	}
}

func (c Controller) checkAllSensors() bool {
	for _, sensor := range c.sensors {
		sensorStatus := sensor.Detect()
		if sensorStatus == devices.MovementDetected {
			return true
		}
	}
	return false
}

func (c Controller) startRecordingAllCameras() {
	for _, recorder := range c.recorders {
		recorder.StartRecording()
	}
}

func (c Controller) stopRecordingAllCameras() {
	for _, recorder := range c.recorders {
		recorder.StopRecording()
	}
}

func NewController(sensors []devices.Sensor, recorders []devices.Recorder) *Controller {
	return &Controller{sensors: sensors, recorders: recorders}
}

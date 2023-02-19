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

	recordCommandToRun := stopRecordCommandToRun
	if someSensorsDetectedMovement {
		recordCommandToRun = startRecordCommandToRun
	}

	c.executeInAllRecorders(recordCommandToRun)
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

func (c Controller) executeInAllRecorders(command recorderCommandToRun) {
	for _, recorder := range c.recorders {
		command(recorder)
	}
}

func NewController(sensors []devices.Sensor, recorders []devices.Recorder) *Controller {
	return &Controller{sensors: sensors, recorders: recorders}
}

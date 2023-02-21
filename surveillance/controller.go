package surveillance

import (
	"github.com/myugen/kata-videorecorder-go/devices"
)

type Controller struct {
	sensors   []devices.Sensor
	recorders []devices.Recorder
}

func (c *Controller) OnStatusChange(event devices.SensorEvent) {
	recorderComandToRun := stopRecordCommandToRun
	if event.IsMovementDetected() {
		recorderComandToRun = startRecordCommandToRun
	}
	c.executeInAllRecorders(recorderComandToRun)
}

func (c *Controller) executeInAllRecorders(command recorderCommandToRun) {
	for _, recorder := range c.recorders {
		command(recorder)
	}
}

func (c *Controller) withSensorsEventsSubscription() *Controller {
	for _, sensor := range c.sensors {
		sensor.Subscribe(c)
	}
	return c
}

func NewController(sensors []devices.Sensor, recorders []devices.Recorder) *Controller {
	instance := &Controller{sensors: sensors, recorders: recorders}
	return instance.withSensorsEventsSubscription()
}

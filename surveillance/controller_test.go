package surveillance_test

import (
	"testing"

	"github.com/myugen/kata-videorecorder-go/devices"
	"github.com/myugen/kata-videorecorder-go/surveillance"
	"github.com/stretchr/testify/assert"
)

func TestController_ShouldStartRecording_WhenSensorDetectsMovement(t *testing.T) {
	sensor := new(SensorEmittingMovementStub)
	camera := new(VideoCameraSpy)
	controller := surveillance.NewController([]devices.Sensor{sensor}, []devices.Recorder{camera})

	controller.CheckSensors()

	assert.Equal(t, 1, camera.TimesOfPlayRecordingCalled())
}

type SensorEmittingMovementStub struct {
}

func (s SensorEmittingMovementStub) Detect() devices.SensorStatus {
	return devices.MovementDetected
}

type VideoCameraSpy struct {
	timesOfPlayRecordingCalled int
	timesOfStopRecordingCalled int
}

func (v *VideoCameraSpy) TimesOfPlayRecordingCalled() int {
	return v.timesOfPlayRecordingCalled
}

func (v *VideoCameraSpy) TimesOfStopRecordingCalled() int {
	return v.timesOfStopRecordingCalled
}

func (v *VideoCameraSpy) StartRecording() {
	v.timesOfPlayRecordingCalled += 1
}

func (v *VideoCameraSpy) StopRecording() {
	v.timesOfStopRecordingCalled += 1
}

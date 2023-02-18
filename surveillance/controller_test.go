package surveillance_test

import (
	"testing"

	"github.com/myugen/kata-videorecorder-go/devices"
	"github.com/myugen/kata-videorecorder-go/surveillance"
	"github.com/myugen/kata-videorecorder-go/testutils/recordermock"
	"github.com/myugen/kata-videorecorder-go/testutils/sensorstub"
	"github.com/stretchr/testify/assert"
)

func TestController_ShouldStartRecording_WhenSomeSensorDetectsMovement(t *testing.T) {
	sensor := sensorstub.NewAlwaysDetectingMovement()
	camera := recordermock.NewVideoCamera()
	controller := surveillance.NewController([]devices.Sensor{sensor}, []devices.Recorder{camera})

	controller.Scan()

	assert.Equal(t, 1, camera.TimesOfPlayRecordingCalled())
}

func TestController_ShouldStopRecording_WhenAllSensorsDoNotDetectMovement(t *testing.T) {
	sensor := sensorstub.NewNeverDetectingMovement()
	camera := recordermock.NewVideoCamera()
	controller := surveillance.NewController([]devices.Sensor{sensor}, []devices.Recorder{camera})

	controller.Scan()

	assert.Equal(t, 1, camera.TimesOfStopRecordingCalled())
}

package surveillance_test

import (
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/myugen/kata-videorecorder-go/devices"
	"github.com/myugen/kata-videorecorder-go/devices/mocks"
	"github.com/myugen/kata-videorecorder-go/surveillance"
)

func TestController_ShouldStartRecording_WhenSomeSensorDetectsMovement(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockedSensor := mocks.NewMockSensor(mockCtrl)
	mockedSensor.EXPECT().Detect().Return(devices.MovementDetected).AnyTimes()
	mockedRecorder := mocks.NewMockRecorder(mockCtrl)
	mockedRecorder.EXPECT().StartRecording().Times(1)
	controller := surveillance.NewController([]devices.Sensor{mockedSensor}, []devices.Recorder{mockedRecorder})

	controller.Scan()
}

func TestController_ShouldStopRecording_WhenAllSensorsDoNotDetectMovement(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockedSensor := mocks.NewMockSensor(mockCtrl)
	mockedSensor.EXPECT().Detect().Return(devices.NoMovementDetected).AnyTimes()
	mockedRecorder := mocks.NewMockRecorder(mockCtrl)
	mockedRecorder.EXPECT().StopRecording().Times(1)
	controller := surveillance.NewController([]devices.Sensor{mockedSensor}, []devices.Recorder{mockedRecorder})

	controller.Scan()
}

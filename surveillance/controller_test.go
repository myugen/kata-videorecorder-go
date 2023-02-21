package surveillance_test

import (
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/myugen/kata-videorecorder-go/devices"
	"github.com/myugen/kata-videorecorder-go/devices/mocks"
	"github.com/myugen/kata-videorecorder-go/surveillance"
	"github.com/myugen/kata-videorecorder-go/testutils/mocksensor"
)

func TestController_ShouldStartRecording_WhenSomeSensorDetectsMovement(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockedSensor := mocksensor.New()
	mockedRecorder := mocks.NewMockRecorder(mockCtrl)
	mockedRecorder.EXPECT().StartRecording().Times(1)
	_ = surveillance.NewController([]devices.Sensor{mockedSensor}, []devices.Recorder{mockedRecorder})

	mockedSensor.SimulateEventWith(mocksensor.SimulatedMovementDetectedEvent)
}

func TestController_ShouldStopRecording_WhenAllSensorsDoNotDetectMovement(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockedSensor := mocksensor.New()
	mockedRecorder := mocks.NewMockRecorder(mockCtrl)
	mockedRecorder.EXPECT().StopRecording().Times(1)
	_ = surveillance.NewController([]devices.Sensor{mockedSensor}, []devices.Recorder{mockedRecorder})

	mockedSensor.SimulateEventWith(mocksensor.SimulatedNoMovementDetectedEvent)
}

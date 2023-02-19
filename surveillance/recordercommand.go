package surveillance

import "github.com/myugen/kata-videorecorder-go/devices"

type recorderCommandToRun = func(devices.Recorder)

var (
	startRecordCommandToRun = func(recorder devices.Recorder) {
		recorder.StartRecording()
	}

	stopRecordCommandToRun = func(recorder devices.Recorder) {
		recorder.StopRecording()
	}
)

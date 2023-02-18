package recordermock

type VideoCamera struct {
	timesOfPlayRecordingCalled int
	timesOfStopRecordingCalled int
}

func (v *VideoCamera) TimesOfPlayRecordingCalled() int {
	return v.timesOfPlayRecordingCalled
}

func (v *VideoCamera) TimesOfStopRecordingCalled() int {
	return v.timesOfStopRecordingCalled
}

func (v *VideoCamera) StartRecording() {
	v.timesOfPlayRecordingCalled += 1
}

func (v *VideoCamera) StopRecording() {
	v.timesOfStopRecordingCalled += 1
}

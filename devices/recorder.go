package devices

//go:generate mockgen -destination=mocks/recordermock.go -package=mocks . Recorder
type Recorder interface {
	StartRecording()
	StopRecording()
}

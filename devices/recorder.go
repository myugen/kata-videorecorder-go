package devices

//go:generate mockgen -destination=mocks/mockrecord.go -package=mocks . Recorder
type Recorder interface {
	StartRecording()
	StopRecording()
}

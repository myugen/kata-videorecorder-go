package sensorstub

import "github.com/myugen/kata-videorecorder-go/devices"

type AlwaysDetectingMovement struct {
}

func (s AlwaysDetectingMovement) Detect() devices.SensorStatus {
	return devices.MovementDetected
}

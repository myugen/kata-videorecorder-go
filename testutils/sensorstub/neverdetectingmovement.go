package sensorstub

import "github.com/myugen/kata-videorecorder-go/devices"

type NeverDetectingMovement struct {
}

func (s NeverDetectingMovement) Detect() devices.SensorStatus {
	return devices.NoMovementDetected
}

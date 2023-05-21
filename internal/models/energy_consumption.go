package models

import "time"

type EnergyConsumption struct {
	ID         int       `json:"id"`
	Timestamp  time.Time `json:"timestamp"`
	Value      float64   `json:"value"`
	DeviceID   int       `json:"device_id"`
	DeviceName string    `json:"device_name"`
}

func NewEnergyConsumption(timestamp time.Time, value float64, deviceID int, deviceName string) *EnergyConsumption {
	return &EnergyConsumption{
		Timestamp:  timestamp,
		Value:      value,
		DeviceID:   deviceID,
		DeviceName: deviceName,
	}
}

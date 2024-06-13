package domain

import "time"

type Device struct {
	Id          uint64
	HouseId     uint64
	UserId      uint64
	Name        string
	Model       string
	Type        DeviceType
	Description *string
	Units       string
	UUID        string
	CreatedDate time.Time
	UpdatedDate time.Time
	DeletedDate *time.Time
}

type DeviceType string

const (
	TempSensor  DeviceType = "TEMPERATURE_SENSOR"
	LockSensor  DeviceType = "LOCK_SENSOR"
	MoveSensor  DeviceType = "MOVE_SENSOR"
	LightSensor DeviceType = "LIGHT_SENSOR"
)

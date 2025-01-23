package domain

import "time"

type Measurement struct {
	Id          uint64
	Value       float64
	DeviceId    uint64
	CreatedDate time.Time
}

type Measurements struct {
	Items []Measurement
	Pages uint64
	Total uint64
}

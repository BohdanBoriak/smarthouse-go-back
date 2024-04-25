package domain

import "time"

type House struct {
	Id          uint64
	UserId      uint64
	Name        string
	Address     string
	Lat         float64
	Lon         float64
	CreatedDate time.Time
	UpdatedDate time.Time
	DeletedDate *time.Time
}

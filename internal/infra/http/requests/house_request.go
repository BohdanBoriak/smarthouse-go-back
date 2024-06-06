package requests

import "github.com/BohdanBoriak/boilerplate-go-back/internal/domain"

type HouseRequest struct {
	Name    string  `json:"name" validate:"required"`
	Address string  `json:"address" validate:"required"`
	Lat     float64 `json:"lat" validate:"required"`
	Lon     float64 `json:"lon" validate:"required"`
}

func (r HouseRequest) ToDomainModel() (interface{}, error) {
	return domain.House{
		Name:    r.Name,
		Address: r.Address,
		Lat:     r.Lat,
		Lon:     r.Lon,
	}, nil
}

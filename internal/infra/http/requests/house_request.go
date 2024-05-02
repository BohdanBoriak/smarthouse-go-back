package requests

import "github.com/BohdanBoriak/boilerplate-go-back/internal/domain"

type HouseRequest struct {
	Name    string  `json:"name"`
	Address string  `json:"address"`
	Lat     float64 `json:"lat"`
	Lon     float64 `json:"lon"`
}

func (r HouseRequest) ToDomainModel() (interface{}, error) {
	return domain.House{
		Name:    r.Name,
		Address: r.Address,
		Lat:     r.Lat,
		Lon:     r.Lon,
	}, nil
}

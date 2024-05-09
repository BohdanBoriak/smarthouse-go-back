package resources

import "github.com/BohdanBoriak/boilerplate-go-back/internal/domain"

type HouseDto struct {
	Id      uint64  `json:"id"`
	UserId  uint64  `json:"userId"`
	Name    string  `json:"name"`
	Address string  `json:"address"`
	Lat     float64 `json:"lat"`
	Lon     float64 `json:"lon"`
}

func (d HouseDto) DomainToDto(house domain.House) HouseDto {
	return HouseDto{
		Id:      house.Id,
		UserId:  house.UserId,
		Name:    house.Name,
		Address: house.Address,
		Lat:     house.Lat,
		Lon:     house.Lon,
	}
}

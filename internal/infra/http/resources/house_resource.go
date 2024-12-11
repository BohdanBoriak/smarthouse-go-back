package resources

import "github.com/BohdanBoriak/boilerplate-go-back/internal/domain"

type HouseDto struct {
	Id      uint64     `json:"id"`
	UserId  uint64     `json:"userId"`
	Name    string     `json:"name"`
	Address string     `json:"address"`
	Lat     float64    `json:"lat"`
	Lon     float64    `json:"lon"`
	Devices DevicesDto `json:"devices"`
}

type HousesDto struct {
	Houses []HouseDto `json:"houses"`
}

func (d HouseDto) DomainToDto(house domain.House) HouseDto {
	var devices DevicesDto
	devices = DeviceDto{}.DomainToDtoCollection(house.Devices)
	return HouseDto{
		Id:      house.Id,
		UserId:  house.UserId,
		Name:    house.Name,
		Address: house.Address,
		Lat:     house.Lat,
		Lon:     house.Lon,
		Devices: devices,
	}
}

func (d HouseDto) DomainToDtoCollection(hs []domain.House) HousesDto {
	var hsDto []HouseDto
	for _, h := range hs {
		hDto := d.DomainToDto(h)
		hsDto = append(hsDto, hDto)
	}

	return HousesDto{
		Houses: hsDto,
	}
}

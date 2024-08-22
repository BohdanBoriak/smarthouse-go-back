package resources

import "github.com/BohdanBoriak/boilerplate-go-back/internal/domain"

type DeviceDto struct {
	Id         uint64            `json:"id"`
	HouseId    uint64            `json:"houseId"`
	UserId     uint64            `json:"userId"`
	Name       string            `json:"name"`
	Model      string            `json:"model"`
	Type       domain.DeviceType `json:"type"`
	Desription *string           `json:"description"`
	Units      string            `json:"units"`
	UUID       string            `json:"uuid"`
}

type DevicesDto struct {
	Devices []DeviceDto `json:"devices"`
}

func (d DeviceDto) DomainToDto(device domain.Device) DeviceDto {
	return DeviceDto{
		Id:         device.Id,
		HouseId:    device.HouseId,
		UserId:     device.UserId,
		Name:       device.Name,
		Model:      device.Model,
		Type:       device.Type,
		Desription: device.Description,
		Units:      device.Units,
		UUID:       device.UUID,
	}
}

func (d DeviceDto) DomainToDtoCollection(dc []domain.Device) DevicesDto {
	var dcDto []DeviceDto
	for _, h := range dc {
		dDto := d.DomainToDto(h)
		dcDto = append(dcDto, dDto)
	}

	return DevicesDto{
		Devices: dcDto,
	}
}

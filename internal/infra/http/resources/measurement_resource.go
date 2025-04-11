package resources

import (
	"time"

	"github.com/BohdanBoriak/boilerplate-go-back/internal/domain"
)

type MeasurementDto struct {
	Id          uint64    `json:"id"`
	Value       float64   `json:"value"`
	DeviceId    uint64    `json:"deviceId"`
	CreatedDate time.Time `json:"time"`
}

type MeasurementsDto struct {
	Measurements []MeasurementDto `json:"measurements"`
	Pages        uint64           `json:"pages"`
	Total        uint64           `json:"total"`
}

func (d MeasurementDto) DomainToDto(m domain.Measurement) MeasurementDto {
	return MeasurementDto{
		Id:          m.Id,
		Value:       m.Value,
		DeviceId:    m.DeviceId,
		CreatedDate: m.CreatedDate,
	}
}

func (d MeasurementsDto) DomainToDtoCollection(ms domain.Measurements) MeasurementsDto {
	var msDto []MeasurementDto
	for _, m := range ms.Items {
		msDto = append(msDto, MeasurementDto{}.DomainToDto(m))
	}

	return MeasurementsDto{
		Measurements: msDto,
		Pages:        ms.Pages,
		Total:        ms.Total,
	}
}

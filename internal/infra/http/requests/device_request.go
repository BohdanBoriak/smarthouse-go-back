package requests

import "github.com/BohdanBoriak/boilerplate-go-back/internal/domain"

type DeviceRequest struct {
	Name       string            `json:"name" validate:"required"`
	Model      string            `json:"model" validate:"required"`
	Type       domain.DeviceType `json:"type" validate:"required,oneof=TEMPERATURE_SENSOR,LOCK_SENSOR,MOVE_SENSOR,LIGHT_SENSOR"`
	Desription *string           `json:"description"`
	Units      string            `json:"units" validate:"required"`
}

func (r DeviceRequest) ToDomainModel() (interface{}, error) {
	return domain.Device{
		Name:       r.Name,
		Model:      r.Model,
		Type:       r.Type,
		Desription: r.Desription,
		Units:      r.Units,
	}, nil
}

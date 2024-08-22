package controllers

import (
	"log"
	"net/http"

	"github.com/BohdanBoriak/boilerplate-go-back/internal/app"
	"github.com/BohdanBoriak/boilerplate-go-back/internal/domain"
	"github.com/BohdanBoriak/boilerplate-go-back/internal/infra/http/requests"
	"github.com/BohdanBoriak/boilerplate-go-back/internal/infra/http/resources"
)

type DeviceController struct {
	deviceService app.DeviceService
}

func NewDeviceController(ds app.DeviceService) DeviceController {
	return DeviceController{
		deviceService: ds,
	}
}

func (c DeviceController) Save() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		house := r.Context().Value(HouseKey).(domain.House)
		user := r.Context().Value(UserKey).(domain.User)
		device, err := requests.Bind(r, requests.DeviceRequest{}, domain.Device{})
		if err != nil {
			log.Printf("HouseController: %s", err)
			BadRequest(w, err)
			return
		}

		device.HouseId = house.Id
		device.UserId = user.Id
		device, err = c.deviceService.Save(device)
		if err != nil {
			log.Printf("HouseController: %s", err)
			InternalServerError(w, err)
			return
		}

		var response resources.DeviceDto
		response = response.DomainToDto(device)
		Created(w, response)
	}
}

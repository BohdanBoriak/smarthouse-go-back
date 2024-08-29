package controllers

import (
	"errors"
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

func (c DeviceController) FindById() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		user := r.Context().Value(UserKey).(domain.User)
		device := r.Context().Value(DeviceKey).(domain.Device)
		if user.Id != device.UserId {
			err := errors.New("access denied")
			Forbidden(w, err)
			return
		}

		var response resources.DeviceDto
		response = response.DomainToDto(device)
		Success(w, response)
	}
}

func (c DeviceController) Update() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		user := r.Context().Value(UserKey).(domain.User)
		device := r.Context().Value(DeviceKey).(domain.Device)
		if user.Id != device.UserId {
			err := errors.New("access denied")
			Forbidden(w, err)
			return
		}

		deviceNew, err := requests.Bind(r, requests.DeviceRequest{}, domain.Device{})
		if err != nil {
			log.Printf("DeviceController: %s", err)
			BadRequest(w, err)
			return
		}

		device.Name = deviceNew.Name
		device.Model = deviceNew.Model
		device.Description = deviceNew.Description
		device.Type = deviceNew.Type
		device.Units = deviceNew.Units
		device, err = c.deviceService.Update(device)
		if err != nil {
			log.Printf("DeviceController: %s", err)
			InternalServerError(w, err)
			return
		}

		var response resources.DeviceDto
		response = response.DomainToDto(device)
		Success(w, response)
	}
}

func (c DeviceController) Delete() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		user := r.Context().Value(UserKey).(domain.User)
		device := r.Context().Value(DeviceKey).(domain.Device)
		if user.Id != device.UserId {
			err := errors.New("access denied")
			Forbidden(w, err)
			return
		}

		err := c.deviceService.Delete(device.Id)
		if err != nil {
			log.Printf("DeviceController: %s", err)
			InternalServerError(w, err)
			return
		}

		Ok(w)
	}
}

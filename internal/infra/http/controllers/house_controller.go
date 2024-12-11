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

type HouseController struct {
	houseService  app.HouseService
	deviceService app.DeviceService
}

func NewHouseController(hs app.HouseService, ds app.DeviceService) HouseController {
	return HouseController{
		houseService:  hs,
		deviceService: ds,
	}
}

func (c HouseController) Save() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		user := r.Context().Value(UserKey).(domain.User)
		house, err := requests.Bind(r, requests.HouseRequest{}, domain.House{})
		if err != nil {
			log.Printf("HouseController: %s", err)
			BadRequest(w, err)
			return
		}

		house.UserId = user.Id
		house, err = c.houseService.Save(house)
		if err != nil {
			log.Printf("HouseController: %s", err)
			InternalServerError(w, err)
			return
		}

		var response resources.HouseDto
		response = response.DomainToDto(house)
		Created(w, response)
	}
}

func (c HouseController) HousesList() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		user := r.Context().Value(UserKey).(domain.User)
		houses, err := c.houseService.FindByUserId(user.Id)
		if err != nil {
			log.Printf("HouseController: %s", err)
			InternalServerError(w, err)
			return
		}

		var hDto resources.HouseDto
		response := hDto.DomainToDtoCollection(houses)
		Success(w, response)
	}
}

func (c HouseController) FindById() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		user := r.Context().Value(UserKey).(domain.User)
		house := r.Context().Value(HouseKey).(domain.House)
		if user.Id != house.UserId {
			err := errors.New("access denied")
			Forbidden(w, err)
			return
		}
		var err error
		house.Devices, err = c.deviceService.FindByHouseId(house.Id)
		if err != nil {
			log.Printf("HouseController: %s", err)
			BadRequest(w, err)
			return
		}

		var response resources.HouseDto
		response = response.DomainToDto(house)
		Success(w, response)
	}
}

func (c HouseController) Update() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		user := r.Context().Value(UserKey).(domain.User)
		house := r.Context().Value(HouseKey).(domain.House)
		if user.Id != house.UserId {
			err := errors.New("access denied")
			Forbidden(w, err)
			return
		}

		houseNew, err := requests.Bind(r, requests.HouseRequest{}, domain.House{})
		if err != nil {
			log.Printf("HouseController: %s", err)
			BadRequest(w, err)
			return
		}

		house.Name = houseNew.Name
		house.Address = houseNew.Address
		house.Lat = houseNew.Lat
		house.Lon = houseNew.Lon
		house, err = c.houseService.Update(house)
		if err != nil {
			log.Printf("HouseController: %s", err)
			InternalServerError(w, err)
			return
		}

		var response resources.HouseDto
		response = response.DomainToDto(houseNew)
		Success(w, response)
	}
}

func (c HouseController) Delete() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		user := r.Context().Value(UserKey).(domain.User)
		house := r.Context().Value(HouseKey).(domain.House)
		if user.Id != house.UserId {
			err := errors.New("access denied")
			Forbidden(w, err)
			return
		}

		err := c.houseService.Delete(house.Id)
		if err != nil {
			log.Printf("HouseController: %s", err)
			InternalServerError(w, err)
			return
		}

		Ok(w)
	}
}

package controllers

import (
	"log"
	"net/http"

	"github.com/BohdanBoriak/boilerplate-go-back/internal/app"
	"github.com/BohdanBoriak/boilerplate-go-back/internal/domain"
	"github.com/BohdanBoriak/boilerplate-go-back/internal/infra/http/requests"
	"github.com/BohdanBoriak/boilerplate-go-back/internal/infra/http/resources"
)

type HouseController struct {
	houseService app.HouseService
}

func NewHouseController(hs app.HouseService) HouseController {
	return HouseController{
		houseService: hs,
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

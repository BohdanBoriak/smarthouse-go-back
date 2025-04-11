package controllers

import (
	"github.com/BohdanBoriak/boilerplate-go-back/internal/domain"
	"github.com/BohdanBoriak/boilerplate-go-back/internal/infra/http/requests"
	"github.com/BohdanBoriak/boilerplate-go-back/internal/infra/http/resources"
	"log"
	"net/http"
	"time"

	"github.com/BohdanBoriak/boilerplate-go-back/internal/app"
	"github.com/BohdanBoriak/boilerplate-go-back/internal/infra/database"
)

type MeasurementController struct {
	measurementService app.MeasurementService
}

func NewMeasurementController(ms app.MeasurementService) MeasurementController {
	return MeasurementController{
		measurementService: ms,
	}
}

func (mc MeasurementController) FindList() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		p, err := requests.DecodePaginationQuery(r)
		if err != nil {
			log.Printf("MeasurementController: %s", err)
			BadRequest(w, err)
			return
		}

		var date time.Time
		dateString := r.URL.Query().Get("date")
		if dateString != "" {
			date, err = time.Parse(time.DateTime, dateString)
			if err != nil {
				log.Printf("MeasurementController: %s", err)
				BadRequest(w, err)
				return
			}
		}

		f := database.MeasurementSearchParams{
			DeviceId:   r.Context().Value(DeviceKey).(domain.Device).Id,
			Date:       &date,
			Pagination: p,
		}

		measurements, err := mc.measurementService.MeasurementsList(f)
		if err != nil {
			log.Printf("MeasurementController: %s", err)
			InternalServerError(w, err)
			return
		}

		measurementsDto := resources.MeasurementsDto{}
		measurementsDto = measurementsDto.DomainToDtoCollection(measurements)
		Success(w, measurementsDto)
	}
}

func (mc MeasurementController) Save() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		measurement, err := requests.Bind(r, requests.MeasurementRequest{}, domain.Measurement{})
		if err != nil {
			log.Printf("MeasurementController: %s", err)
			BadRequest(w, err)
			return
		}

		device := r.Context().Value(DeviceKey).(domain.Device)
		measurement.DeviceId = device.Id
		_, err = mc.measurementService.Save(measurement)
		if err != nil {
			log.Printf("MeasurementController: %s", err)
			InternalServerError(w, err)
			return
		}

		noContent(w)
	}
}

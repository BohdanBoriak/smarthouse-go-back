package controllers

import (
	"fmt"
	"github.com/BohdanBoriak/boilerplate-go-back/internal/domain"
	"github.com/BohdanBoriak/boilerplate-go-back/internal/infra/http/requests"
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

		trs, err := mc.measurementService.MeasurementsList(f)
		if err != nil {
			log.Printf("MeasurementController: %s", err)
			InternalServerError(w, err)
			return
		}

		fmt.Println(trs)

		//todo: finish with resources
		//var psDto resources.TranslationRequestsDto
		//Success(w, psDto.DomainToDtoPagination(trs))
	}
}

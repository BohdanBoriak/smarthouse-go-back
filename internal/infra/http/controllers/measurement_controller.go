package controllers

import (
	"log"
	"net/http"

	"github.com/BohdanBoriak/boilerplate-go-back/internal/app"
	"github.com/BohdanBoriak/boilerplate-go-back/internal/infra/database"
)

type MeasurementController struct {
	measurementService app.MeasurementService
}

func NewMeasurementController(ds app.MeasurementService) MeasurementController {
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

		f := database.MeasurementSearchParams{
			From:   domain.LangCode(r.URL.Query().Get("from")),
			To:     domain.LangCode(r.URL.Query().Get("to")),
			Search: r.URL.Query().Get("search"),
			Status: domain.TranslationRequestStatus(r.URL.Query().Get("status")),
		}

		trs, err := mc.measurementService.MeasurementsList()
		if err != nil {
			log.Printf("MeasurementController: %s", err)
			InternalServerError(w, err)
			return
		}

		var psDto resources.TranslationRequestsDto
		Success(w, psDto.DomainToDtoPagination(trs))
	}
}

package middlewares

import (
	"context"
	"github.com/BohdanBoriak/boilerplate-go-back/internal/app"
	"github.com/BohdanBoriak/boilerplate-go-back/internal/infra/http/controllers"
	"github.com/go-chi/chi/v5"
	"log"
	"net/http"
)

func PathUUIDDevice(service app.DeviceService) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		hfn := func(w http.ResponseWriter, r *http.Request) {
			uuid := chi.URLParam(r, "uuid")

			device, err := service.FindByUUId(uuid)
			if err != nil {
				log.Print(err)
				controllers.BadRequest(w, err)
				return
			}

			ctx := context.WithValue(r.Context(), controllers.DeviceKey, device)

			next.ServeHTTP(w, r.WithContext(ctx))
		}
		return http.HandlerFunc(hfn)
	}
}

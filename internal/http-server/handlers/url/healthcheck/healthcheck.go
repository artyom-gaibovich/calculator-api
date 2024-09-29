package healthcheck

import (
	resp "calculator-api/internal/lib/api/response"
	"calculator-api/internal/lib/logger/sl"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/render"
	"github.com/go-playground/validator/v10"
	"golang.org/x/exp/slog"
	"net/http"
)

type Request struct {
}

type Response struct {
	resp.Response
	Message string `json:"message,omitempty"`
}

type HealthChecker interface {
	HealthCheck() (string, error)
}

func New(log *slog.Logger, healthChecker HealthChecker) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		const fn = "handlers.url.check.New"
		log = log.With(
			slog.String("fn", fn),
			slog.String("request_id", middleware.GetReqID(r.Context())),
		)
		var req Request

		if err := validator.New().Struct(req); err != nil {
			validateErr := err.(validator.ValidationErrors)
			log.Error("invalid request", sl.Err(err))
			render.JSON(w, r, resp.ValidationError(validateErr))
			return
		}
		log.Info("service is working...")
		msg, _ := healthChecker.HealthCheck()
		render.JSON(w, r, Response{
			Response: resp.OK(),
			Message:  msg,
		})
	}

}

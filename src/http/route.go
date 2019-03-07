package http

import (
	"github.com/go-chi/chi"
	chiMiddleware "github.com/go-chi/chi/middleware"
	"github.com/jinzhu/gorm"
	"github.com/sirupsen/logrus"
	"io"
	"net/http"
	"time"
)

func NewRoute(logger *logrus.Entry, db *gorm.DB) http.Handler {
	h := NewWebHandler(db)
	r := chi.NewRouter()
	r.Use(chiMiddleware.RequestID)
	r.Use(chiMiddleware.Timeout(10 * time.Second))
	r.Use(Logger(logger))
	r.Use(Recoverer)

	r.Route("/api/v1", func(r chi.Router) {
		r.Get("/search", h.WrapFunc(Search))
		r.Get("/autocomplete", h.WrapFunc(Autocomplete))
	})
	r.NotFound(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusNotFound)
		io.WriteString(w, `{"status": 404, "data": "URL_NOT_FOUND"}`) // nolint:errcheck
	})

	return r
}

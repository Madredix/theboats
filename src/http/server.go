package http

import (
	"context"
	"net/http"
	"time"

	"github.com/sirupsen/logrus"

	"github.com/jinzhu/gorm"
	"strconv"
)

type Server interface {
	Start()
	Stop()
}

type server struct {
	srv    *http.Server
	logger *logrus.Entry
	db     *gorm.DB
}

func NewHTTPServer(port int, logger *logrus.Logger, db *gorm.DB) Server {
	l := logger.WithField(`module`, `web`)
	return server{
		&http.Server{Addr: ":" + strconv.Itoa(port), Handler: NewRoute(l, db)},
		l,
		db,
	}
}

func (s server) Start() {
	go func() {
		if err := s.srv.ListenAndServe(); err != nil {
			s.logger.WithError(err).WithField(`action`, `listen and serve`).Error()
		}
	}()

	s.logger.WithField(`action`, `start`).Info("server gracefully started")
}

func (s server) Stop() {
	s.logger.WithField(`action`, `stop`).Info("server is shutting down")

	// shut down gracefully, but wait no longer than 5 seconds before halting
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := s.srv.Shutdown(ctx); err != nil {
		s.logger.WithError(err).WithField(`action`, `stop`).Error()
	} else {
		s.logger.WithField(`action`, `stop`).Info("server gracefully stopped")
	}
}

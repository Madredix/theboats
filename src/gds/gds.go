package gds

import (
	"github.com/Madredix/theboats/src/models"
	"time"
)

type GDS interface {
	GetName() string
	GetUpdateInterval() time.Duration
	GetBuilders() (models.Builders, error)
	GetModels() (models.Models, error)
	GetCompanies() (models.Companies, error)
	GetYachts() (models.Yachts, error)
	GetReservations() (models.Reservations, error)
}

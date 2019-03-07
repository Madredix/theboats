package models

import (
	"github.com/Madredix/theboats/src/libs/db"
	"github.com/jinzhu/gorm"
	"time"
)

const sqlConflictReservation = `ON CONFLICT (id) DO UPDATE SET period_from = EXCLUDED.period_from, period_to = EXCLUDED.period_to`

type ReservationRepo struct {
	db *gorm.DB
}

type (
	Reservation struct {
		ID         uint `json:"id" gorm:"primary_key"`
		YachtID    uint
		PeriodFrom time.Time
		PeriodTo   time.Time
	}
	Reservations []Reservation
)

func NewReservationRepo(db *gorm.DB) ReservationRepo {
	return ReservationRepo{db: db}
}

func (r ReservationRepo) CreateOrUpdate(reservation Reservation) error {
	return r.db.Set("gorm:insert_option", sqlConflictReservation).Create(&reservation).Error
}

func (r ReservationRepo) CreateOrUpdateAll(reservations Reservations) (err error) {
	butch := make([]interface{}, len(reservations))
	for i := range reservations {
		butch[i] = reservations[i]
	}
	return db.BatchInsert(r.db, butch, sqlConflictReservation)
}

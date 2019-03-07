package gds

import (
	"github.com/Madredix/theboats/src/models"
	"github.com/jinzhu/gorm"
	"github.com/sirupsen/logrus"
	"time"
)

type GDSUpdater interface {
	Update(gds GDS, manual bool) bool
}

type gdsUpdater struct {
	db     *gorm.DB
	logger *logrus.Entry
}

func NewGDSUpdater(db *gorm.DB, logger *logrus.Logger) GDSUpdater {
	return &gdsUpdater{db: db, logger: logger.WithField(`module`, `updater`)}
}

func (g gdsUpdater) Update(gds GDS, manual bool) bool {
	l := g.logger.WithField(`gds`, gds.GetName())
	l.Debug(`start update`)

	// check last update
	gdsRepo := models.NewGDSRepo(g.db)
	gdsModel, err := gdsRepo.GetByName(gds.GetName())
	if err != nil {
		l.Error(`gds not found in db`)
		return false
	}

	lastUpdate := gdsModel.LastUpdate()
	if lastUpdate != nil && (lastUpdate.Add(gds.GetUpdateInterval()).After(time.Now()) || !manual) {
		l.Error(`too short period or run first with exists data`)
		return false
	}

	if err = gdsModel.StartUpdate(); err != nil {
		l.WithError(err).Error(`start update`)
		return false
	}
	defer gdsModel.StopUpdate(err == nil)

	tx := g.db.Begin()

	builderRepo := models.NewBuilderRepo(tx)
	builders, err := gds.GetBuilders()
	if err != nil {
		tx.Rollback()
		return false
	}
	err = builderRepo.CreateAll(builders)
	if err != nil {
		tx.Rollback()
		l.WithError(err).Error(`update builders`)
		return false
	}

	modelRepo := models.NewModelRepo(tx)
	m, err := gds.GetModels()
	if err != nil {
		tx.Rollback()
		return false
	}
	err = modelRepo.CreateAll(m)
	if err != nil {
		tx.Rollback()
		l.WithError(err).Error(`update models`)
		return false
	}

	companyRepo := models.NewCompanyRepo(tx)
	companies, err := gds.GetCompanies()
	if err != nil {
		tx.Rollback()
		return false
	}
	err = companyRepo.CreateAll(companies)
	if err != nil {
		tx.Rollback()
		l.WithError(err).Error(`update companies`)
		return false
	}

	yachtRepo := models.NewYachtRepo(tx)
	yachts, err := gds.GetYachts()
	if err != nil {
		tx.Rollback()
		return false
	}
	err = yachtRepo.CreateOrUpdateAll(yachts)
	if err != nil {
		tx.Rollback()
		l.WithError(err).Error(`update yachts`)
		return false
	}

	reservationRepo := models.NewReservationRepo(tx)
	reservations, err := gds.GetReservations()
	if err != nil {
		tx.Rollback()
		return false
	}
	err = reservationRepo.CreateOrUpdateAll(reservations)
	if err != nil {
		tx.Rollback()
		l.WithError(err).Error(`update reservations`)
		return false
	}

	if err = tx.Commit().Error; err != nil {
		l.WithError(err).Error(`transaction commit`)
		return false
	}

	l.Info(`data update seccussfully`)

	return true
}

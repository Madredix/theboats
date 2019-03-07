package models

import (
	"github.com/jinzhu/gorm"
	"github.com/pkg/errors"
	"time"
)

const errGDSUpdateNotStart = `update is not started`

type GDSRepo struct {
	db *gorm.DB
}

type GDS struct {
	ID      uint `json:"id" gorm:"primary_key"`
	Name    string
	db      *gorm.DB  `sql:"-"`
	Updates []Updater `sql:"-" gorm:"foreignkey:gds_id"`
}

func NewGDSRepo(db *gorm.DB) GDSRepo {
	return GDSRepo{db: db}
}

func (g GDSRepo) GetByName(name string) (*GDS, error) {
	gds := &GDS{db: g.db}
	err := g.db.First(gds, `name = ?`, name).Error
	return gds, err
}

func (g GDSRepo) GetByID(id uint) (gds *GDS, err error) {
	err = g.db.First(&gds, `id = ?`, id).Error
	return
}

func (g *GDS) LastUpdate() *time.Time {
	g.db.Model(g).Order(`GREATEST(start_at, stop_at) desc`).Limit(1).Related(&g.Updates)
	switch {
	case len(g.Updates) > 0 && g.Updates[0].StopAt != nil:
		return g.Updates[0].StopAt
	case len(g.Updates) > 0:
		return &g.Updates[0].StartAt
	default:
		return nil
	}
}

func (g *GDS) StartUpdate() error {
	u := Updater{GDSID: g.ID}
	err := g.db.Create(&u).Error
	if err == nil {
		g.Updates = append(g.Updates, u)
	}
	return err
}

func (g *GDS) StopUpdate(status bool) error {
	cur := len(g.Updates) - 1
	if cur == -1 || g.Updates[cur].StopAt != nil {
		return errors.New(errGDSUpdateNotStart)
	}
	u := &g.Updates[cur]
	err := g.db.Model(u).Where(`id = ?`, u.ID).Updates(map[string]interface{}{`stop_at`: gorm.Expr("NOW()"), `status`: status}).Error
	if err == nil {
		g.db.First(&u)
	}
	return err
}

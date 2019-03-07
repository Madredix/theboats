package models

import (
	"github.com/Madredix/theboats/src/libs/db"
	"github.com/jinzhu/gorm"
	"time"
)

const sqlConflictYacht = `ON CONFLICT (id) DO UPDATE SET deleted_at = NULL, updated_at = EXCLUDED.updated_at, name = EXCLUDED.name, company_id = EXCLUDED.company_id, model_id = EXCLUDED.model_id`

type YachtRepo struct {
	db *gorm.DB
}

type (
	Yacht struct {
		gorm.Model
		CompanyID uint `json:"companyId"`
		ModelID   uint `json:"yachtModelId"`
		Name      string
	}
	Yachts []Yacht
)

func NewYachtRepo(db *gorm.DB) YachtRepo {
	return YachtRepo{db: db}
}

func (y YachtRepo) CreateOrUpdate(yacht Yacht) error {
	return y.db.Set("gorm:insert_option", sqlConflictYacht).Create(&yacht).Error
}

func (y YachtRepo) CreateOrUpdateAll(yachts Yachts) (err error) {
	t := time.Now()
	butch := make([]interface{}, len(yachts))
	for i := range yachts {
		yachts[i].CreatedAt = t
		yachts[i].UpdatedAt = t
		butch[i] = yachts[i]
	}
	err = db.BatchInsert(y.db, butch, sqlConflictYacht)
	if err == nil {
		y.db.Delete(&Yachts{}, `updated_at < ?`, t)
	}
	return err
}

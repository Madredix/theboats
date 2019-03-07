package models

import (
	"github.com/Madredix/theboats/src/libs/db"
	"github.com/jinzhu/gorm"
)

type ModelRepo struct {
	db *gorm.DB
}

type (
	Model struct {
		ID        uint `json:"id" gorm:"primary_key"`
		BuilderID uint `json:"yachtBuilderId"`
		Name      string
	}
	Models []Model
)

func NewModelRepo(db *gorm.DB) ModelRepo {
	return ModelRepo{db: db}
}

func (m ModelRepo) Create(model Model) error {
	return m.db.FirstOrCreate(&model).Error
}

func (m ModelRepo) CreateAll(models Models) (err error) {
	butch := make([]interface{}, len(models))
	for i := range models {
		butch[i] = models[i]
	}
	return db.BatchInsert(m.db, butch, `ON CONFLICT (id) DO NOTHING`)
}

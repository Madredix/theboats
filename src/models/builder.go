package models

import (
	"github.com/Madredix/theboats/src/libs/db"
	"github.com/jinzhu/gorm"
)

type BuilderRepo struct {
	db *gorm.DB
}

type (
	Builder struct {
		ID   uint `json:"id" gorm:"primary_key"`
		Name string
	}
	Builders []Builder
)

func NewBuilderRepo(db *gorm.DB) BuilderRepo {
	return BuilderRepo{db: db}
}

func (b BuilderRepo) Create(builder Builder) error {
	return b.db.FirstOrCreate(&builder).Error
}

func (b BuilderRepo) CreateAll(builders Builders) (err error) {
	butch := make([]interface{}, len(builders))
	for i := range builders {
		butch[i] = builders[i]
	}
	return db.BatchInsert(b.db, butch, `ON CONFLICT (id) DO NOTHING`)
}

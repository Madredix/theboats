package models

import (
	"github.com/Madredix/theboats/src/libs/db"
	"github.com/jinzhu/gorm"
)

type CompanyRepo struct {
	db *gorm.DB
}

type (
	Company struct {
		ID   uint `json:"id" gorm:"primary_key"`
		Name string
	}
	Companies []Company
)

func NewCompanyRepo(db *gorm.DB) CompanyRepo {
	return CompanyRepo{db: db}
}

func (c CompanyRepo) Create(company Company) error {
	return c.db.FirstOrCreate(&company).Error
}

func (c CompanyRepo) CreateAll(companies Companies) (err error) {
	butch := make([]interface{}, len(companies))
	for i := range companies {
		butch[i] = companies[i]
	}
	return db.BatchInsert(c.db, butch, `ON CONFLICT (id) DO NOTHING`)
}

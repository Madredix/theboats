package def

import (
	"github.com/jinzhu/gorm"
	"github.com/sarulabs/di"
	"github.com/sirupsen/logrus"
)

func GetDbConnection(di di.Container) *gorm.DB {
	var db *gorm.DB

	err := di.Fill(DatabaseDef, &db)
	if err != nil {
		panic(err)
	}

	return db
}

func GetLogger(di di.Container) *logrus.Logger {
	var l *logrus.Logger

	err := di.Fill(LoggerDef, &l)
	if err != nil {
		panic(err)
	}

	return l
}

func GetConfig(di di.Container) *Config {
	var c Config

	err := di.Fill(CfgDef, &c)
	if err != nil {
		panic(err)
	}

	return &c
}

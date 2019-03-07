package nausys

import (
	"github.com/Madredix/theboats/src/gds"
	"github.com/Madredix/theboats/src/libs/web"
	"github.com/sirupsen/logrus"
	"time"
)

const Name = `nausys`
const intervalRequest = time.Second
const intervalUpdate = time.Minute

type nausys struct {
	url      string
	login    string
	password string
	logger   *logrus.Entry
	web      web.IntervalRequest
}

func NewNausys(
	url string,
	login string,
	password string,
	logger *logrus.Logger,
) gds.GDS {
	return &nausys{
		url:      url,
		login:    login,
		password: password,
		logger:   logger.WithField(`gds`, Name),
		web:      web.NewIntervalRequest(intervalRequest),
	}
}

package nausys

import (
	"github.com/Madredix/theboats/src/domain"
	"github.com/Madredix/theboats/src/models"
	"strconv"
	"time"
)

func (n nausys) GetName() string {
	return Name
}

func (n nausys) GetUpdateInterval() time.Duration {
	return intervalUpdate
}

func (n nausys) GetBuilders() (models.Builders, error) {
	var responce struct {
		Status   string
		Builders models.Builders
	}
	params := map[string]string{`username`: n.login, `password`: n.password}
	if err := n.web.PostJsonData(n.url+`catalogue/v6/yachtBuilders`, params, &responce, map[string]string{`Content-Type`: `application/json`}); err != nil {
		n.logger.Error(err)
		return nil, err
	}
	return responce.Builders, nil
}

func (n nausys) GetModels() (models.Models, error) {
	var responce struct {
		Status string
		Models models.Models
	}
	params := map[string]string{`username`: n.login, `password`: n.password}
	if err := n.web.PostJsonData(n.url+`catalogue/v6/yachtModels`, params, &responce, map[string]string{`Content-Type`: `application/json`}); err != nil {
		n.logger.Error(err)
		return nil, err
	}
	return responce.Models, nil
}

func (n nausys) GetCompanies() (models.Companies, error) {
	var responce struct {
		Status    string
		Companies models.Companies
	}
	params := map[string]string{`username`: n.login, `password`: n.password}
	if err := n.web.PostJsonData(n.url+`catalogue/v6/charterCompanies`, params, &responce, map[string]string{`Content-Type`: `application/json`}); err != nil {
		n.logger.Error(err)
		return nil, err
	}
	return responce.Companies, nil
}

func (n nausys) GetYachts() (models.Yachts, error) {
	var responce struct {
		Status string
		Yachts models.Yachts
	}
	yachts := make(models.Yachts, 0)
	companies, err := n.GetCompanies()
	if err != nil {
		return nil, err
	}
	for _, company := range companies {
		params := map[string]string{`username`: n.login, `password`: n.password}
		if err := n.web.PostJsonData(n.url+`catalogue/v6/yachts/`+strconv.Itoa(int(company.ID)), params, &responce, map[string]string{`Content-Type`: `application/json`}); err != nil {
			n.logger.Error(err)
			return nil, err
		}
		yachts = append(yachts, responce.Yachts...)
	}

	return responce.Yachts, nil
}

func (n nausys) GetReservations() (models.Reservations, error) {
	var responce struct {
		Status       string
		Reservations []struct {
			ID         uint
			YachtID    uint        `json:"yachtId"`
			PeriodFrom domain.Date `json:"periodFrom"`
			PeriodTo   domain.Date `json:"periodTo"`
		}
	}
	params := map[string]interface{}{
		`credentials`: map[string]string{
			`username`: n.login,
			`password`: n.password,
		},
		`periodFrom`: time.Now().AddDate(0, -6, 0).Format(`02.01.2006`),
		`periodTo`:   time.Now().AddDate(1, 6, 0).Format(`02.01.2006`),
	}
	if err := n.web.PostJsonData(n.url+`yachtReservation/v6/reservations`, params, &responce, map[string]string{`Content-Type`: `application/json`}); err != nil {
		n.logger.Error(err)
		return nil, err
	}
	reservations := make(models.Reservations, len(responce.Reservations))
	for i, reservation := range responce.Reservations {
		reservations[i] = models.Reservation{ID: reservation.ID, YachtID: reservation.YachtID, PeriodFrom: reservation.PeriodFrom.Time, PeriodTo: reservation.PeriodTo.Time}
	}
	return reservations, nil
}

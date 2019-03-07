package nausys

import (
	"github.com/Madredix/theboats/src/libs/web"
	"github.com/Madredix/theboats/src/models"
	"github.com/Madredix/theboats/src/test"
	"github.com/google/go-cmp/cmp"
	"github.com/sirupsen/logrus"
	"net/http"
	"testing"
	"time"
)

const testDataPathYachtReservation = `../../../testing/mock/gds/yachtReservation.json`

// Пример теста метода, остальные не стал покрывать
func TestNausys_GetReservations(t *testing.T) {
	test.MockHTTPResponse(http.StatusOK, map[string]string{"Content-Type": "application/json"}, test.GetTestData(testDataPathYachtReservation, t))
	nausys := getNausys()
	reservations, err := nausys.GetReservations()
	if err != nil {
		t.Errorf("\nUnexpected error: %s", err.Error())
	}

	expected := []models.Reservation{
		{ID: 265052527, YachtID: 479288, PeriodFrom: time.Date(2019, 4, 6, 9, 0, 0, 0, time.UTC), PeriodTo: time.Date(2019, 4, 13, 19, 0, 0, 0, time.UTC)},
		{ID: 265075450, YachtID: 3355085, PeriodFrom: time.Date(2019, 10, 19, 17, 0, 0, 0, time.UTC), PeriodTo: time.Date(2019, 10, 26, 10, 0, 0, 0, time.UTC)},
	}

	if len(expected) != len(reservations) {
		t.Errorf("\nExpected count: %d\nReceived count: %d\n", len(expected), len(reservations))
	}
	for i := range reservations {
		if !cmp.Equal(reservations[i], expected[i]) {
			t.Errorf("\nIteration: %d\nExpected: %+v\nReceived: %+v\n", i, expected[i], reservations[i])
		}
	}
}

func getNausys() nausys {
	return nausys{
		url:      `http://localhost`,
		login:    ``,
		password: ``,
		logger:   logrus.New().WithField(`gds`, Name),
		web:      web.NewIntervalRequest(time.Nanosecond),
	}
}

package domain

import (
	"encoding/json"
	"github.com/google/go-cmp/cmp"
	"testing"
	"time"
)

func TestDate_UnmarshalJSON(t *testing.T) {
	source := []byte(`["01.01.2018 12:16", "15.07.2011 05:00", null]`)
	expected := []Date{
		{time.Date(2018, time.January, 1, 12, 16, 0, 0, time.UTC)},
		{time.Date(2011, time.July, 15, 5, 0, 0, 0, time.UTC)},
		{time.Time{}},
	}
	var test []Date
	err := json.Unmarshal(source, &test)
	if !cmp.Equal(test, expected) || err != nil {
		t.Errorf("\nExpected: %+v\nReceived: %+v", expected, test)
	}
}

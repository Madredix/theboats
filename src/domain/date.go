package domain

import (
	"strings"
	"time"
)

type Date struct {
	time.Time
}

const dateLayout = "02.01.2006 15:04"

func (s *Date) UnmarshalJSON(b []byte) (err error) {
	t := strings.Trim(string(b), "\"")
	if t == "null" {
		s.Time = time.Time{}
		return
	}
	s.Time, err = time.Parse(dateLayout, t)
	return
}

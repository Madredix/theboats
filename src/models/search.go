package models

import (
	"github.com/jinzhu/gorm"
	"time"
)

type SearchRepo struct {
	db *gorm.DB
}

type (
	Autocomplite struct {
		Name string
	}
	SearchYacht struct {
		Name             string
		Builder          string
		Model            string
		Company          string
		Availability     bool
		AvailabilityFrom time.Time
		AvailabilityTo   *time.Time
	}
)

func NewSearchRepo(db *gorm.DB) SearchRepo {
	return SearchRepo{db: db}
}

// Автокомплит названия производителя или модели
// Может выводить названия производителя/модели у которых нет доступных яхт
func (s SearchRepo) Autocomplite(q string) []Autocomplite {
	result := make([]Autocomplite, 0)
	if len(q) >= 3 {
		// префикс имени производителя яхты или префикс названия модели
		s.db.Raw(`
		SELECT s.name FROM(
			SELECT name FROM builders WHERE name ILIKE $1
			UNION
			SELECT name FROM models WHERE name ILIKE $1
		) AS s ORDER BY name ASC LIMIT 100
	`, q+`%`).Find(&result)
	}
	return result
}

func (s SearchRepo) Search(q string) []SearchYacht {
	result := make([]SearchYacht, 0)
	if len(q) >= 3 {
		// поиск по префиксу имени производителя яхты или префиксу названия модели
		/*
			Полное имя производителя;
			Полное название модели;
			Имя владельца яхты (флота);
			Доступность яхты на текущую дату (доступна/в резерве);
			Ближайшие даты доступности яхты (доступна с/по).
		*/
		s.db.Raw(`
			SELECT distinct on (y.name)
  				y.name                  as name,
  				b.name                  as builder,
  				m.name                  as model,
  				c.name                  as company,
  				COALESCE(now() BETWEEN r1.period_from AND r1.period_to, false) = false as availability,
  				CASE WHEN r1.period_from > now() or r1.period_from IS NULL THEN now() ELSE r1.period_to END as availability_from,
  				CASE WHEN r1.period_from > now() THEN r2.period_from ELSE r1.period_from END as availability_to
			FROM yachts AS y
				JOIN companies AS c ON c.id = y.company_id
				JOIN models AS m ON m.id = y.model_id
				JOIN builders AS b ON b.id = m.builder_id
				LEFT JOIN reservations AS r1 ON r1.yacht_id = y.id and period_to > now()
				LEFT JOIN reservations AS r2 ON r2.yacht_id = y.id and r2.period_from > r1.period_from
				WHERE b.name ILIKE $1 OR m.name ILIKE $1
				ORDER BY y.name ASC
			LIMIT 100;
		`, q+`%`).Find(&result)
	}

	return result
}

func (a Autocomplite) MarshalJSON() ([]byte, error) {
	return []byte(`"` + a.Name + `"`), nil
}

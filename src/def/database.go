package def

import (
	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
	"github.com/sarulabs/di"
)

const DatabaseDef = "database"

func init() {
	Register(func(builder *Builder, params map[string]interface{}) error {
		builder.Add(di.Def{
			Name: DatabaseDef,
			Build: func(ctx di.Container) (_ interface{}, err error) {
				cfg := GetConfig(ctx)
				logger := GetLogger(ctx)
				connStr := `postgres://` + cfg.Database.User + `:` + cfg.Database.Password + `@` + cfg.Database.Host + `:` + cfg.Database.Port + `/` + cfg.Database.Name + `?sslmode=disable&binary_parameters=yes`

				db, err := gorm.Open(`postgres`, connStr)
				if err != nil {
					logger.WithError(err).WithField(`str`, connStr).Fatal(`init db`)
				}

				if db.Exec("SELECT 1").Error != nil {
					logger.WithError(err).Fatal(`try select`)
				}

				return db, nil
			},
		})

		return nil
	})
}

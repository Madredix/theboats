package def

import (
	"github.com/Madredix/theboats/src/gds"
	"github.com/sarulabs/di"
)

const GDSUpdaterDef = "gds_updater"

func init() {
	Register(func(builder *Builder, params map[string]interface{}) error {
		builder.Add(di.Def{
			Name: GDSUpdaterDef,
			Build: func(ctx di.Container) (_ interface{}, err error) {
				return gds.NewGDSUpdater(GetDbConnection(ctx), GetLogger(ctx)), nil
			},
		})

		return nil
	})
}

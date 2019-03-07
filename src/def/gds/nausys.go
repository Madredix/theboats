package gdsDef

import (
	"github.com/Madredix/theboats/src/def"
	"github.com/Madredix/theboats/src/gds/nausys"
)

const NausysDef = nausys.Name

func init() {
	def.Register(func(builder *def.Builder, params map[string]interface{}) error {
		return builder.Add(def.Definition{
			Name: NausysDef,
			Build: func(ctx def.Context) (_ interface{}, err error) {
				cfg := ctx.Get(def.CfgDef).(def.Config)
				logger := def.GetLogger(ctx)

				cfgNausys, ok := cfg.Gds[NausysDef]
				if !ok {
					logger.WithField(`gds`, NausysDef).Fatal(`read gds config`)
				}
				return nausys.NewNausys(cfgNausys.Url, cfgNausys.Login, cfgNausys.Password, logger), nil
			},
		})
	})
}

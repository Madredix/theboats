package def

import (
	"github.com/Madredix/theboats/src/http"
	_ "github.com/lib/pq"
	"github.com/sarulabs/di"
)

const HttpDef = "http"

func init() {
	Register(func(builder *Builder, params map[string]interface{}) error {
		builder.Add(di.Def{
			Name: HttpDef,
			Build: func(ctx di.Container) (_ interface{}, err error) {
				cfg := GetConfig(ctx)
				logger := GetLogger(ctx)
				db := GetDbConnection(ctx)
				server := http.NewHTTPServer(cfg.Http.Port, logger, db)
				return server, nil
			},
		})

		return nil
	})
}

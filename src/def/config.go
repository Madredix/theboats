package def

import (
	"errors"
	"fmt"
	"path/filepath"
	"time"

	"github.com/spf13/viper"
)

const DateTimeFormat = time.RFC3339Nano
const TimeZone = "Europe/Moscow"

type (
	Config struct {
		Environment string `mapstructure:"environment"`

		Log struct {
			Level  string `mapstructure:"level"`
			Format string `mapstructure:"format"`
			Output string `mapstructure:"output"`
		} `mapstructure:"log"`

		Database struct {
			Host     string `mapstructure:"host"`
			Password string `mapstructure:"password"`
			Port     string `mapstructure:"port"`
			Name     string `mapstructure:"name"`
			User     string `mapstructure:"user"`
		} `mapstructure:"database"`

		Migration struct {
			Path string `mapstructure:"path"`
		} `mapstructure:"migration"`

		Http struct {
			Port int `mapstructure:"port"`
		} `mapstructure:"http"`

		Gds map[string]struct {
			Url      string `mapstructure:"url"`
			Password string `mapstructure:"password"`
			Login    string `mapstructure:"login"`
		} `mapstructure:"gds"`

		Testing struct {
			SeedPath string `mapstructure:"seed_path"`
			MockPath string `mapstructure:"mock_path"`
		} `mapstructure:"testing"`

		baseDir string
	}
)

const (
	CfgDef    = "config"
	CfgType   = "yml"
	importOpt = "imports"
)

func init() {
	Register(func(builder *Builder, params map[string]interface{}) error {
		var ok bool
		if _, ok = params["configFile"]; !ok {
			return errors.New("can't get required parameter config path")
		}

		var path string
		if path, ok = params["configFile"].(string); !ok {
			return errors.New(`parameter "configFile" should be string`)
		}

		builder.Add(Definition{
			Name: CfgDef,
			Build: func(ctx Context) (interface{}, error) {
				var config Config
				var err error

				config.baseDir, err = filepath.Abs(filepath.Dir(path))
				if err != nil {
					panic(err.Error())
				}

				err = config.importFile(path, true)
				if err != nil {
					panic(err.Error())
				}

				return config, nil
			},
		})

		return nil
	})
}

func (c *Config) importFile(path string, isMain bool) error {
	viperInst := viper.New()
	viperInst.SetConfigType(CfgType)

	var cfgPath = path
	var err error

	if !filepath.IsAbs(path) {
		if isMain {
			cfgPath, err = filepath.Abs(path)
			if err != nil {
				return err
			}
		} else {
			cfgPath, err = filepath.Abs(c.baseDir + string(filepath.Separator) + filepath.Clean(path))
			if err != nil {
				return err
			}
		}
	}
	viperInst.SetConfigFile(cfgPath)

	if err := viperInst.ReadInConfig(); err != nil {
		return fmt.Errorf("Fatal error config file: '%s' \n", err)
	}

	imports := viperInst.GetStringSlice(importOpt)
	for _, filePath := range imports {
		if err := c.importFile(filePath, false); err != nil {
			return err
		}
	}

	if err := viperInst.Unmarshal(c); err != nil {
		return err
	}

	return nil
}

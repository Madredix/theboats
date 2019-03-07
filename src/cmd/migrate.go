package cmd

import (
	"errors"
	"os"
	"time"

	"github.com/Madredix/theboats/src/def"
	"github.com/golang-migrate/migrate"
	"github.com/golang-migrate/migrate/database/postgres"
	_ "github.com/golang-migrate/migrate/database/postgres"
	_ "github.com/golang-migrate/migrate/source/file"
	"github.com/jinzhu/gorm"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

const (
	migrateActionUp     = `up`
	migrateActionDown   = `down`
	migrateActionCreate = `create`
)

var (
	command string
	name    string

	MigrateCmd = &cobra.Command{
		Use:   `migrate [command]`,
		Short: `Database migration tool`,
	}
	MigrateCmdUp = &cobra.Command{
		Use:   `up`,
		Short: `Migration up`,
		Run:   runMigrate(migrateActionUp),
	}
	MigrateCmdDown = &cobra.Command{
		Use:   `down`,
		Short: `Migration down`,
		Run:   runMigrate(migrateActionDown),
	}
	MigrateCmdCreate = &cobra.Command{
		Use:   `create`,
		Short: `Migration create`,
		Run:   runMigrate(migrateActionCreate),
	}
)

func init() {
	MigrateCmdCreate.Flags().StringVarP(&name, `name`, `n`, ``, `Name of the migration (required)`)
	MigrateCmdCreate.MarkFlagRequired(`name`)
	MigrateCmd.AddCommand(MigrateCmdUp)
	MigrateCmd.AddCommand(MigrateCmdDown)
	MigrateCmd.AddCommand(MigrateCmdCreate)
	RootCmd.AddCommand(MigrateCmd)
}

func runMigrate(action string) func(cmd *cobra.Command, args []string) {
	return func(cmd *cobra.Command, args []string) {
		cfg := diContext.Get(def.CfgDef).(def.Config)
		db := diContext.Get(def.DatabaseDef).(*gorm.DB)
		logger := diContext.Get(def.LoggerDef).(*logrus.Logger)

		driver, err := postgres.WithInstance(db.DB(), &postgres.Config{})
		if err != nil {
			logger.WithError(err).Fatal(`init driver`)
		}

		m, err := migrate.NewWithDatabaseInstance(
			`file://`+cfg.Migration.Path,
			`postgres`,
			driver,
		)
		if err != nil {
			logger.WithError(err).Fatal(`init migrate`)
		}

		switch action {
		case migrateActionUp:
			err = m.Up()
		case migrateActionDown:
			err = m.Down()
		case migrateActionCreate:
			fileName := cfg.Migration.Path + `/` + time.Now().Format(`20060102150405`) + `_` + name
			if err = createMigration(fileName + `.up.sql`); err == nil {
				err = createMigration(fileName + `.down.sql`)
			}
		default:
			err = errors.New(`unknown command ` + command)
		}

		if err != nil && err != migrate.ErrNoChange {
			logger.WithError(err).Fatal(`execute command`)
		}
		logger.Info(`migrate ` + action + ` successfully`)
	}
}

func createMigration(fileName string) error {
	file, err := os.Create(fileName)
	if err != nil {
		return err
	}
	defer file.Close()

	return nil
}

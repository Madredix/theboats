package cmd

import (
	"github.com/Madredix/theboats/src/def"
	"github.com/Madredix/theboats/src/libs/db"
	"github.com/jinzhu/gorm"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var SeedCmd = &cobra.Command{
	Use:   "seed",
	Short: "Seed test data",
	Run:   runSeed,
}

func init() {
	RootCmd.AddCommand(SeedCmd)
}

func runSeed(cmd *cobra.Command, args []string) {
	MigrateCmdDown.Run(cmd, args)
	MigrateCmdUp.Run(cmd, args)

	logger := diContext.Get(def.LoggerDef).(*logrus.Logger)
	cfg := diContext.Get(def.CfgDef).(def.Config)
	database := diContext.Get(def.DatabaseDef).(*gorm.DB).DB()
	err := db.SeedFile(cfg.Testing.SeedPath, database)
	if err != nil {
		logger.WithError(err).Fatal(`seed data`)
	} else {
		logger.Info(`seed successfully`)
	}
}

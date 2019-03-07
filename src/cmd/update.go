package cmd

import (
	"github.com/Madredix/theboats/src/def"
	"github.com/Madredix/theboats/src/def/gds"
	"github.com/Madredix/theboats/src/gds"
	"github.com/spf13/cobra"
)

var (
	UpdateCmd = &cobra.Command{
		Use:   "update",
		Short: "Update all GDS",
		Run:   runUpdate(true),
	}
	UpdateCmdFirst = &cobra.Command{
		Run:    runUpdate(false),
		Hidden: true,
	}
)

func init() {
	RootCmd.AddCommand(UpdateCmd)
	RootCmd.AddCommand(UpdateCmdFirst)
}

func runUpdate(manual bool) func(cmd *cobra.Command, args []string) {
	return func(cmd *cobra.Command, args []string) {
		updater := diContext.Get(def.GDSUpdaterDef).(gds.GDSUpdater)
		updater.Update(diContext.Get(gdsDef.NausysDef).(gds.GDS), manual)
	}
}

package cmd

import (
	"github.com/Madredix/theboats/src/def"
	"github.com/spf13/cobra"
)

var (
	// Config file path.
	configFile string

	// DI Container.
	diContext def.Context

	RootCmd = &cobra.Command{
		Use:   `theboats`,
		Short: `Test work for http://theboats.com`,
		Long:  ``,
		Run:   runApp,
		PersistentPreRunE: func(cmd *cobra.Command, args []string) (err error) {
			diContext, err = def.Instance(map[string]interface{}{
				"configFile": configFile,
			})
			return err
		},
	}
)

func init() {
	RootCmd.PersistentFlags().StringVarP(&configFile, "config", "c", "./config/config.json", "config file")
}

func runApp(cmd *cobra.Command, args []string) {
	MigrateCmdUp.Run(cmd, args)
	UpdateCmdFirst.Run(cmd, args)
	HttpCmd.Run(cmd, args)
}

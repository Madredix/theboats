package cmd

import (
	"github.com/Madredix/theboats/src/def"
	"github.com/Madredix/theboats/src/http"
	"github.com/spf13/cobra"
	"os"
	"os/signal"
	"syscall"
)

var HttpCmd = &cobra.Command{
	Use:   "http",
	Short: "HTTP Rest api server",
	Run:   runHttp,
}

func init() {
	RootCmd.AddCommand(HttpCmd)
}

func runHttp(cmd *cobra.Command, args []string) {
	// subscribe to SIGINT signals
	stopChan := make(chan os.Signal, 1)
	signal.Notify(stopChan, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	// Init & run webserver
	webserver := diContext.Get(def.HttpDef).(http.Server)
	webserver.Start()

	// wait for SIGINT
	<-stopChan
	webserver.Stop()
}

package cmd

import (
	"time"

	rest "github.com/yogabagas/jatis-BE/transport/rest"

	"github.com/spf13/cobra"
)

var serverCommand = &cobra.Command{
	Use: "serve",
	PreRun: func(cmd *cobra.Command, args []string) {
		initModule()
	},
	Run: func(cmd *cobra.Command, args []string) {
		rest := rest.NewRest(&rest.Options{
			Port:         ":9000",
			Db:           Db,
			WriteTimeout: 30 * time.Second,
			ReadTimeout:  30 * time.Second,
		})

		rest.Serve()
	},
}

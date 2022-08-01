package server

import (
	"memos/server/config"
	"memos/server/db"
	"memos/server/logger"
	"memos/server/router"

	"github.com/spf13/cobra"
)

var (
	configYml string
	apiCheck  bool
	StartCmd  = &cobra.Command{
		Use:          "server",
		Short:        "Start API server",
		Example:      "memos server -c config/settings.yml",
		SilenceUsage: true,
		PreRun: func(cmd *cobra.Command, args []string) {
			setup()
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			return run()
		},
	}
)

func init() {
	StartCmd.PersistentFlags().StringVarP(&configYml, "config", "c", "config/settings.yml", "Start server with provided configuration file")
	StartCmd.PersistentFlags().BoolVarP(&apiCheck, "api", "a", false, "Start server with check api data")
}

func setup() {
}

func run() error {
	config.InitConf()

	if err := logger.Init(config.Conf.Logger, config.Conf.Application.Mode); err != nil {
		return err
	}

	logger.Infof("init database...")
	db.InitDB()

	logger.Infof("init routers...")
	router.InitRouter()

	logger.Infof("run server...")
	router.RunServer()
	return nil
}

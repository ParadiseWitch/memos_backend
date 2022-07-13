package cmd

import (
	"errors"
	"fmt"
	"memos/cmd/server"
	"memos/server/logger"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:          "memos",
	Short:        "memos",
	SilenceUsage: true,
	Long:         `memos`,
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) < 1 {
			tip()
			return errors.New(logger.Red("requires at least one arg"))
		}
		return nil
	},
	PersistentPreRunE: func(*cobra.Command, []string) error { return nil },
	Run: func(cmd *cobra.Command, args []string) {
		tip()
	},
}

func tip() {
	usageStr := `Welecome to ` + logger.Green(`memos.`) + `Use ` + logger.Red(`-h`) + ` for help`
	fmt.Printf("%s\n", usageStr)
}

func init() {
	rootCmd.AddCommand(server.StartCmd)
}

//Execute : apply commands
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(-1)
	}
}

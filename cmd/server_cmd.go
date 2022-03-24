package cmd

import (
	"course-teacher/config"
	"course-teacher/db"
	"course-teacher/server"

	"github.com/spf13/cobra"
)

var serveCmd = &cobra.Command{
	Use:     "serve",
	Short:   "Starts application server",
	Run:     runServeCmd,
	Aliases: []string{"s", "server"},
}

func init() {
	rootCmd.AddCommand(serveCmd)
}

func runServeCmd(cmd *cobra.Command, args []string) {

	config.Init()
	db.Init()
	server.Init()

}

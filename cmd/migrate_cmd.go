package cmd

import (
	"course-teacher/config"
	"database/sql"
	"fmt"
	"log"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/lib/pq"
	"github.com/spf13/cobra"
)

var migrateCmd = &cobra.Command{
	Use:     "migrate",
	Short:   "Migrate cmd is used for database migration",
	Run:     runMigrateUp,
	Aliases: []string{"m", "migrate"},
}

func init() {
	rootCmd.AddCommand(migrateCmd)
}

func runMigrateUp(cmd *cobra.Command, args []string) {
	config.Init()
	db, err := sql.Open("postgres", config.GetDatabaseURL())
	fmt.Println(config.GetDatabaseURL())
	if err != nil {
		log.Panicln(err)
	}
	driver, err := postgres.WithInstance(db, &postgres.Config{})
	m, err := migrate.NewWithDatabaseInstance(
		"file://migrations",
		"postgres", driver)
	if err := m.Up(); err != nil {
		log.Fatal(err)
	}
}

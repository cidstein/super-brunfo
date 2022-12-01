package cmd

import (
	"database/sql"
	"fmt"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	"github.com/golang-migrate/migrate/v4/source/file"
	"github.com/spf13/cobra"
)

var migrateUpCmd *cobra.Command

func init() {
	migrateUpCmd = &cobra.Command{
		Use:   "up",
		Short: "Migrate to v1 command",
		Long:  `Migrate to install version 1 of the application`,
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("Running migrate up command")

			db, err := sql.Open("postgres", "postgres://postgres:postgres@db:5432/postgres?sslmode=disable")
			if err != nil {
				fmt.Printf("sql.Open: %v \n", err)
				panic(err)
			}

			driver, err := postgres.WithInstance(db, &postgres.Config{})
			if err != nil {
				fmt.Printf("instance error: %v \n", err)
			}

			fileSource, err := (&file.File{}).Open("file://migrations")
			if err != nil {
				fmt.Printf("fileSource error: %v \n", err)
				panic(err)
			}

			m, err := migrate.NewWithInstance("file", fileSource, "postgres", driver)
			if err != nil {
				fmt.Printf("NewWithInstance: %v \n", err)
				panic(err)
			}

			if err = m.Up(); err != nil {
				fmt.Printf("migrate Up error: %v \n", err)
				panic(err)
			}

			fmt.Println("Migrated up done with sucess")
		},
	}

	migrateCmd.AddCommand(migrateUpCmd)
}

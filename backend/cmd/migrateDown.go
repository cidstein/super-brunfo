package cmd

import (
	"database/sql"
	"fmt"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	"github.com/golang-migrate/migrate/v4/source/file"
	"github.com/spf13/cobra"
)

var migrateDownCmd *cobra.Command

func init() {
	migrateDownCmd = &cobra.Command{
		Use:   "down",
		Short: "Migrate from v2 to v1",
		Long:  `Command to downgrade the database from v2 to v1`,
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("Running migrate down command")

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

			if err = m.Down(); err != nil {
				fmt.Printf("migrate Down error: %v \n", err)
				panic(err)
			}

			fmt.Println("Migrated down done with sucess")
		},
	}

	migrateCmd.AddCommand(migrateDownCmd)
}

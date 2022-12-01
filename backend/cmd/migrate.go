package cmd

import "github.com/spf13/cobra"

var migrateCmd = &cobra.Command{
	Use:   "migrate",
	Short: "Migrate the database",
	Long:  `Migrate the database`,
}

func init() {
	rootCmd.AddCommand(migrateCmd)
}

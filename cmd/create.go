package cmd

import (
	"fmt"

	"github.com/intentional_mitsake/db_shit/pkg/config"
	"github.com/intentional_mitsake/db_shit/pkg/db"
	"github.com/spf13/cobra"
)

var createCmd = &cobra.Command{
	Use:   "create",
	Short: "Open a connection to a database server and create a new database",
	Long: `This command opens a connection to a PostgreSQL database server and creates a new database.
	Format: 
	dbcli create --username username --password password --host localhost -- port 5432 --database new_db
	`,
	RunE: func(cmd *cobra.Command, args []string) error {
		cfgFile := config.LoadDatabaseConfig()
		client := db.NewPGClient(cfgFile)
		//call create, connection is opened inside create and closed as well
		if err := client.Create(); err != nil {
			return fmt.Errorf("failed to create database: %w", err)
		}
		fmt.Println("Database created successfully.")
		return nil
	},
}

func init() {
	//adds openCmd to the parent cmd i.e. root
	rootCmd.AddCommand(createCmd)
}

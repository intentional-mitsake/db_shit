package cmd

import (
	"fmt"

	"github.com/intentional_mitsake/db_shit/pkg/config"
	"github.com/intentional_mitsake/db_shit/pkg/db"
	"github.com/spf13/cobra"
)

var openCmd = &cobra.Command{
	Use:   "open",
	Short: "Open a connection to a database server",
	RunE: func(cmd *cobra.Command, args []string) error {
		cfgFile := config.LoadDatabaseConfig()
		client := db.NewPGClient(cfgFile)
		if err := client.Connect(); err != nil {
			return fmt.Errorf("connection failure: %w", err)
		}
		defer client.Close()
		fmt.Println("Connection opened successfully.")
		return nil
	},
}

func init() {
	//adds openCmd to the parent cmd i.e. root
	rootCmd.AddCommand(openCmd)
}

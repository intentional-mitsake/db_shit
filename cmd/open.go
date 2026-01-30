package cmd

import (
	"fmt"

	"github.com/intentional_mitsake/db_shit/pkg/config"
	"github.com/intentional_mitsake/db_shit/pkg/db"
	"github.com/spf13/cobra"
)

// have to keep in mind this is for testing the connection only
// the connection is closed immediately after this command is done executing
// so this funcitonaltiiy needs to be added to other functions like cresate, backup and restore
// FOR TEST ONLY
var openCmd = &cobra.Command{
	Use:   "open",
	Short: "Open a connection to a database server",
	RunE: func(cmd *cobra.Command, args []string) error {
		cfgFile := config.LoadDatabaseConfig()
		client := db.NewPGClient(cfgFile)
		if err := client.Connect(true); err != nil {
			return fmt.Errorf("connection failure: %w", err)
		}
		//defer executes when the function returns
		//so the connection is closed after the function returns
		defer client.Close()
		fmt.Println("Connection opened successfully.")
		return nil
	},
}

func init() {
	//adds openCmd to the parent cmd i.e. root
	rootCmd.AddCommand(openCmd)
}

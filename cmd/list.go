package cmd

import (
	"fmt"

	"github.com/intentional_mitsake/db_shit/pkg/config"
	"github.com/intentional_mitsake/db_shit/pkg/db"
	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "Open a connection to a database server and list existing databases",
	Long: `This command opens a connection to a PostgreSQL database server and lists existing databases.
	Format: 
	dbcli list --username username --password password --host localhost -- port 5432
	`,
	RunE: func(cmd *cobra.Command, args []string) error {
		cfgFile := config.LoadDatabaseConfig()
		client := db.NewPGClient(cfgFile)
		//call create, connection is opened inside create and closed as well
		list, err := client.List()
		if err != nil {
			return fmt.Errorf("failed to list databases: %w", err)
		}
		//looping over the list
		fmt.Println("List of databases:")
		//'_' can be ignored but if u use 'index' u have to use the var
		//here '_' is the index of the list and 'db' is the value at that index
		//this pretty much loops over the entire lsit like for i := 0; i < len(list); i++
		for _, db := range list {
			fmt.Println(db)
		}
		return nil
	},
}

func init() {
	//adds openCmd to the parent cmd i.e. root
	rootCmd.AddCommand(listCmd)
}

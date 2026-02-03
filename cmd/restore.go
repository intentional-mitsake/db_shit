package cmd

import (
	"github.com/intentional_mitsake/db_shit/pkg/config"
	"github.com/intentional_mitsake/db_shit/pkg/db"
	"github.com/intentional_mitsake/db_shit/pkg/utils"
	"github.com/spf13/cobra"
)

var restoreCmd = &cobra.Command{
	Use:   "restore",
	Short: "Open a connection to a database server and lrestore a database",
	Long: `This command opens a connection to a PostgreSQL database server and lists existing databases.
	Format: 
	dbcli restore --username username --password password --host localhost --database dbname -- port 5432 --destination file
	`,
	RunE: func(cmd *cobra.Command, args []string) error {
		logger := utils.CreateLogger()
		cfgFile := config.LoadDatabaseConfig()
		client := db.NewPGClient(cfgFile)
		if err := client.Restore(); err != nil {
			logger.Error(err.Error())
			return nil
		}
		logger.Info("Database Restored.")
		return nil
	},
}

func init() {
	//adds openCmd to the parent cmd i.e. root
	rootCmd.AddCommand(restoreCmd)
}

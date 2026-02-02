package cmd

import (
	"github.com/intentional_mitsake/db_shit/pkg/config"
	"github.com/intentional_mitsake/db_shit/pkg/db"
	"github.com/intentional_mitsake/db_shit/pkg/utils"
	"github.com/spf13/cobra"
)

var backupCmd = &cobra.Command{
	Use:   "backup",
	Short: "Open a connection to a database server and list existing databases",
	Long: `This command opens a connection to a PostgreSQL database server and lists existing databases.
	Format: 
	dbcli backup --username username --password password --host localhost --database dbname -- port 5432 --destination file
	`,
	RunE: func(cmd *cobra.Command, args []string) error {
		logger := utils.CreateLogger()
		cfgFile := config.LoadDatabaseConfig()
		client := db.NewPGClient(cfgFile)
		if err := client.Backup(); err != nil {
			logger.Error(err.Error())
			return nil
		}
		logger.Info("Backup Created.")
		return nil
	},
}

func init() {
	//adds openCmd to the parent cmd i.e. root
	rootCmd.AddCommand(backupCmd)
}

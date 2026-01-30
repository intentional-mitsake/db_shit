/*
Copyright © 2026 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var cfgFile string

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "dbcli",
	Short: "A CLI tool for DB management.",
	Long:  `CLI based utility tool for manual DB backups, restores, creation and basic CRUD.`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	// Run: func(cmd *cobra.Command, args []string) { },
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	// Here you will define your flags and configuration settings.
	cobra.OnInitialize(initConfig)
	// Cobra supports persistent flags, which, if defined here,
	//these flags can be used by all subcommands. no need to redefine in each command file.
	//cuz rootcmd itself has no run func, these flags will be available to all subcommands but the rootcmd itself won't use them
	rootCmd.PersistentFlags().String("host", "localhost", "Database Host")
	rootCmd.PersistentFlags().String("username", "admin", "Database User")
	rootCmd.PersistentFlags().String("password", "", "Database Password")
	rootCmd.PersistentFlags().String("db", "testdb", "Database Name")
	rootCmd.PersistentFlags().Int("port", 5432, "Database Port")
	rootCmd.PersistentFlags().String("type", "mysql", "Database Type (mysql, postgres)")
	rootCmd.PersistentFlags().String("destination", "", "Destination File")
	// will be global for your application.
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.db_shit.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

}

func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := os.UserHomeDir()
		cobra.CheckErr(err)

		// Search config in home directory with name ".db_shit" (without extension).
		viper.AddConfigPath(home)
		viper.AddConfigPath(".")
		viper.SetConfigType("yaml")
		viper.SetConfigName(".db_shit")
	}

	viper.AutomaticEnv() // Read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Fprintln(os.Stderr, "Using config file:", viper.ConfigFileUsed())
	}

	viper.BindPFlag("host", rootCmd.PersistentFlags().Lookup("host"))
	viper.BindPFlag("username", rootCmd.PersistentFlags().Lookup("username"))
	viper.BindPFlag("password", rootCmd.PersistentFlags().Lookup("password"))
	viper.BindPFlag("db", rootCmd.PersistentFlags().Lookup("db"))
	viper.BindPFlag("port", rootCmd.PersistentFlags().Lookup("port"))
	viper.BindPFlag("type", rootCmd.PersistentFlags().Lookup("type"))
	viper.BindPFlag("destination", rootCmd.PersistentFlags().Lookup("destination"))
}

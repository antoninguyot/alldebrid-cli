package cmd

import (
	"context"
	"github.com/adrg/xdg"
	"github.com/antoninguyot/alldebrid-cli/cmd/auth"
	"github.com/antoninguyot/alldebrid-cli/cmd/links"
	"github.com/antoninguyot/alldebrid-cli/cmd/magnets"
	"github.com/antoninguyot/alldebrid-cli/pkg/http"
	"github.com/jedib0t/go-pretty/v6/table"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var cfgFile string

// RootCmd represents the base command when called without any subcommands
var RootCmd = &cobra.Command{
	Use:   "alldebrid",
	Short: "A CLI application for Alldebrid",
	Long: `alldebrid-cli is a full-featured CLI application for Alldebrid.
It can manage your magnets and links, and more to come.`,
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the RootCmd.
func Execute() {
	err := RootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	//allDebridClient = http.NewClient()

	cobra.OnInitialize(initConfig)

	// Global flags
	RootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.alldebrid-cli.yaml)")
	RootCmd.PersistentFlags().StringP("output", "o", "table", "output format (table, csv)")
	RootCmd.PersistentFlags().String("token", "", "alldebrid token")
	viper.BindPFlag("auth.token", RootCmd.PersistentFlags().Lookup("token"))

	// Object instantiation
	RootCmd.PersistentPreRun = func(cmd *cobra.Command, args []string) {
		cmd.SetContext(initHttpClient(cmd.Context(), viper.GetString("auth.token")))
		cmd.SetContext(initTableWriter(cmd.Context()))
	}

	// Subcommands
	RootCmd.AddCommand(auth.AuthCmd)
	RootCmd.AddCommand(magnets.MagnetsCmd)
	RootCmd.AddCommand(links.LinksCmd)

}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		configFilePath, err := xdg.ConfigFile("alldebrid-cli/config.yaml")
		cobra.CheckErr(err)
		viper.SetConfigFile(configFilePath)
	}

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		//
	}
}

func initHttpClient(ctx context.Context, token string) context.Context {
	return context.WithValue(ctx, "client", http.NewClient(token))
}

func initTableWriter(ctx context.Context) context.Context {
	writer := table.NewWriter()
	writer.SetOutputMirror(os.Stdout)

	return context.WithValue(ctx, "writer", writer)
}

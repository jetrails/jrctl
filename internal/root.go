package internal

import (
	"fmt"
	"os"
	"github.com/jetrails/jrctl/sdk/color"
	"github.com/jetrails/jrctl/sdk/version"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var config string

var rootCmd = &cobra.Command {
	Use:    "jrctl",
	Version: version.VersionString,
	Short:  "Command line tool to help interact with the " + color.GetLogo () + " API",
	Long:   "Command line tool to help interact with the " + color.GetLogo () + " API\n" +
			"Hosted on Github, " + color.GreenString ("https://github.com/jetrails/jrctl") + ".\n" +
			"For issues/requests, please open an issue in our Github repository.",
	DisableAutoGenTag: true,
}

func GetRootCommand () * cobra.Command {
	return rootCmd
}

func Execute () {
	if error := rootCmd.Execute (); error != nil {
		fmt.Println ( error )
		os.Exit ( 1 )
	}
}

func init () {
	cobra.OnInitialize ( initConfig )
	version.CheckVersion ( viper.GetBool ("debug") )
}

func initConfig () {
	if config != "" {
		viper.SetConfigFile ( config )
	} else {
		viper.AddConfigPath ("$HOME")
		viper.SetConfigName (".jrctl")
		viper.SetConfigType ("yaml")
		viper.SetEnvPrefix ("JR")
		viper.SetDefault ( "public_api_endpoint", "api-public.jetrails.cloud" )
		viper.SetDefault ( "secret_endpoint", "secret.jetrails.cloud" )
		viper.SetDefault ( "daemon_endpoint", "localhost:27482" )
		viper.SetDefault ( "daemon_token", "" )
		viper.SafeWriteConfig ()
	}
	viper.AutomaticEnv ()
	if error := viper.ReadInConfig (); error == nil {
		if ( viper.GetBool ("debug") ) {
			fmt.Println ( color.CyanString ( "config:" ), viper.ConfigFileUsed () )
			fmt.Println ( color.CyanString ( "debug:" ), viper.GetString ("debug") )
			fmt.Println ( color.CyanString ( "color:" ), viper.GetString ("color") )
			fmt.Println ( color.CyanString ( "public_api_endpoint:" ), viper.GetString ("public_api_endpoint") )
			fmt.Println ( color.CyanString ( "secret_endpoint:" ), viper.GetString ("secret_endpoint") )
			fmt.Println ( color.CyanString ( "daemon_endpoint:" ), viper.GetString ("daemon_endpoint") )
			fmt.Println ( color.CyanString ( "daemon_token:" ), viper.GetString ("daemon_token") )
			fmt.Println ()
		}
	}
}

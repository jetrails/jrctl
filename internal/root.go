package internal

import "fmt"
import "os"
import "github.com/spf13/cobra"
import "github.com/spf13/viper"
import "github.com/fatih/color"

var configFile string
var jetrails = color.GreenString (">") + "jetrails" + color.GreenString ("_")

var rootCmd = &cobra.Command {
	Use:   "jrctl",
	Short: "Command line tool to help interact with the " + jetrails + " API",
	Version: "1.0.1",
}

func Execute () {
	if error := rootCmd.Execute (); error != nil {
		fmt.Println ( error )
		os.Exit ( 1 )
	}
}

func init () {
	cobra.OnInitialize ( initConfig )
}

func initConfig () {
	if configFile != "" {
		viper.SetConfigFile ( configFile )
	} else {
		viper.AddConfigPath ("$HOME")
		viper.SetConfigName (".jrctl")
		viper.SetConfigType ("yaml")
		viper.SetEnvPrefix ("jr")
		viper.SetDefault ( "debug", false )
		viper.SetDefault ( "endpoint_postfix", "" )
		viper.SafeWriteConfig ()
	}
	viper.AutomaticEnv ()
	if error := viper.ReadInConfig (); error == nil {
		if ( viper.GetBool ("debug") ) {
			fmt.Println ( "Debug:", viper.GetString ("debug") )
			fmt.Println ( "Endpoint:", viper.GetString ("endpoint") )
			fmt.Println ( "Config File:", viper.ConfigFileUsed () )
		}
	}
}

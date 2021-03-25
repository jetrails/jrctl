package internal

import (
	"fmt"
	"os"
	"os/user"
	"path"
	"github.com/jetrails/jrctl/sdk/color"
	"github.com/jetrails/jrctl/sdk/version"
	"github.com/jetrails/jrctl/sdk/utils"
	"github.com/jetrails/jrctl/sdk/server"
	"github.com/jetrails/jrctl/sdk/env"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var config string

var rootCmd = &cobra.Command {
	Use: "jrctl",
	Version: version.VersionString,
	Short: "Command line tool to help interact with the " + color.GetLogo () + " API.",
	Long: utils.Paragraph ( [] string {
		"Command line tool to help interact with the " + color.GetLogo () + " API.",
		"Hosted on Github, " + color.GreenString ("https://github.com/jetrails/jrctl") + ".",
		"For issues/requests, please open an issue in our Github repository.",
	}),
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
	version.CheckVersion ( env.GetBool ( "debug", false ) )
}

func initConfig () {
	if config != "" {
		viper.SetConfigFile ( config )
	} else {
		viper.AddConfigPath ("$HOME/.jrctl")
		viper.SetConfigName ("config")
		viper.SetConfigType ("yaml")
		viper.SetDefault ( "servers", [] server.Entry {
			server.Entry {
				Endpoint: "127.0.0.1:27482",
				Token: "REPLACE_WITH_AUTH_TOKEN",
				Types: [] string { "localhost" },
			},
		})
		if usr, error := user.Current (); error == nil {
			os.MkdirAll ( path.Join ( usr.HomeDir, ".jrctl" ), 0755 )
			viper.SafeWriteConfig ()
		}
	}
	viper.ReadInConfig ()
	if ( env.GetBool ( "debug", false ) ) {
		fmt.Println ("---")
		fmt.Printf ( "%s: %v\n", color.CyanString ("config"), viper.ConfigFileUsed () )
		fmt.Printf ( "%s: %v\n", color.CyanString ("debug"), env.GetBool ( "debug", false ) )
		fmt.Printf ( "%s: %v\n", color.CyanString ("insecure"), env.GetBool ( "insecure", false ) )
		fmt.Printf ( "%s: %v\n", color.CyanString ("color"), env.GetBool ( "color", true ) )
		fmt.Printf ( "%s: %v\n", color.CyanString ("public_api_endpoint"), env.GetString ( "public_api_endpoint", "api-public.jetrails.cloud" ) )
		fmt.Printf ( "%s: %v\n", color.CyanString ("secret_endpoint"), env.GetString ( "secret_endpoint", "secret.jetrails.cloud" ) )
		fmt.Println ("---")
	}
}

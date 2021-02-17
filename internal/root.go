package internal

import (
	"fmt"
	"os"
	"encoding/json"
	"github.com/jetrails/jrctl/sdk"
	"github.com/jetrails/jrctl/sdk/utils"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/fatih/color"
	"github.com/parnurzeal/gorequest"
	"github.com/hashicorp/go-version"
)

const ReleasesUrl = "https://api.github.com/repos/jetrails/jrctl/releases"
const TagUrlTemplate = "https://github.com/jetrails/jrctl/releases/tag/%s"

var configFile string
var jetrails = color.GreenString (">") + "jetrails" + color.GreenString ("_")

type Release struct {
	TagName string `json:"tag_name"`
}

type ReleaseResponse struct {
	Collection [] Release
}

var rootCmd = &cobra.Command {
	Use:    "jrctl",
	Short:  "Command line tool to help interact with the " + jetrails + " API",
	Long:   "Command line tool to help interact with the " + jetrails + " API\n" +
			"Hosted on Github, " + color.GreenString ("https://github.com/jetrails/jrctl") + ".\n" +
			"For issues/requests, please open an issue in our Github repository.",
	Version: sdk.Version,
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
	checkVersion ()
}

func checkVersion () {
	var cacheWindow int64 = 60 * 60
	cachedVersion, hit := utils.GetCache ( "version:" + sdk.Version, cacheWindow )
	versionObj, _ := version.NewVersion ( sdk.Version )
	if hit {
		cachedVersionObj, _ := version.NewVersion ( cachedVersion )
		if versionObj.LessThan ( cachedVersionObj ) {
			fmt.Printf (
				"Software is out-of-date. Update to the latest version: %s.\n%s\n\n",
				cachedVersion,
				color.RedString ( fmt.Sprintf ( TagUrlTemplate, cachedVersion ) ),
			)
		}
		return
	}
	var request = gorequest.New ()
	var debug = viper.GetBool ("debug")
	request.SetDebug ( debug )
	response, body, _ := request.Get ( ReleasesUrl ).End ()
	if response.StatusCode == 200 {
		releases := make ( [] Release, 0 )
		json.Unmarshal ( [] byte ( body ), &releases )
		newest := releases [ 0 ]
		utils.SetCache ( "version:" + sdk.Version, newest.TagName, cacheWindow )
		targetVersionObj, _ := version.NewVersion ( newest.TagName )
		if versionObj.LessThan ( targetVersionObj ) {
			fmt.Printf (
				"Software is out-of-date. Update to the latest version: %s.\n%s\n\n",
				newest.TagName,
				color.RedString ( fmt.Sprintf ( TagUrlTemplate, newest.TagName ) ),
			)
		}
	}
}

func initConfig () {
	if configFile != "" {
		viper.SetConfigFile ( configFile )
	} else {
		viper.AddConfigPath ("$HOME")
		viper.SetConfigName (".jrctl")
		viper.SetConfigType ("yaml")
		viper.SetEnvPrefix ("JR")
		viper.SetDefault ( "debug", false )
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
			fmt.Println ( color.CyanString ( "public_api_endpoint:" ), viper.GetString ("public_api_endpoint") )
			fmt.Println ( color.CyanString ( "secret_endpoint:" ), viper.GetString ("secret_endpoint") )
			fmt.Println ( color.CyanString ( "daemon_endpoint:" ), viper.GetString ("daemon_endpoint") )
			fmt.Println ( color.CyanString ( "daemon_token:" ), viper.GetString ("daemon_token") )
			fmt.Println ()
		}
	}
}

package internal

import "fmt"
import "os"
import "encoding/json"
import "github.com/jetrails/jrctl/sdk/utils"
import "github.com/spf13/cobra"
import "github.com/spf13/viper"
import "github.com/fatih/color"
import "github.com/parnurzeal/gorequest"
import "github.com/hashicorp/go-version"

const Version = "1.0.4"
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
	Version: Version,
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
	cachedVersion, hit := utils.GetCache ( "version:" + Version, cacheWindow )
	versionObj, _ := version.NewVersion ( Version )
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
		utils.SetCache ( "version:" + Version, newest.TagName, cacheWindow )
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
		viper.SetEnvPrefix ("jr")
		viper.SetDefault ( "debug", false )
		viper.SetDefault ( "endpoint_postfix", "" )
		viper.SafeWriteConfig ()
	}
	viper.AutomaticEnv ()
	if error := viper.ReadInConfig (); error == nil {
		if ( viper.GetBool ("debug") ) {
			fmt.Println ( color.CyanString ( "Debug:" ), viper.GetString ("debug") )
			fmt.Println ( color.CyanString ( "Endpoint Postfix:" ), viper.GetString ("endpoint") )
			fmt.Println ( color.CyanString ( "Config File:" ), viper.ConfigFileUsed () )
			fmt.Println ()
		}
	}
}

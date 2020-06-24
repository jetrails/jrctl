package internal

import "fmt"
import "os"
import "encoding/json"
import "github.com/jetrails/jrctl/sdk/utils"
import "github.com/spf13/cobra"
import "github.com/spf13/viper"
import "github.com/fatih/color"
import "github.com/parnurzeal/gorequest"

const Version = "1.0.2"
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
	Use:   "jrctl",
	Short: "Command line tool to help interact with the " + jetrails + " API",
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
	if hit {
		if Version != cachedVersion {
			fmt.Printf (
				"Software is out-of-date. Update to the latest version: %s.\n%s\n\n",
				cachedVersion,
				color.GreenString ( fmt.Sprintf ( TagUrlTemplate, cachedVersion ) ),
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
		if Version != newest.TagName {
			fmt.Printf (
				"Software is out-of-date. Update to the latest version: %s.\n%s\n\n",
				newest.TagName,
				color.GreenString ( fmt.Sprintf ( TagUrlTemplate, newest.TagName ) ),
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
			fmt.Println ( "Debug:", viper.GetString ("debug") )
			fmt.Println ( "Endpoint:", viper.GetString ("endpoint") )
			fmt.Println ( "Config File:", viper.ConfigFileUsed () )
		}
	}
}

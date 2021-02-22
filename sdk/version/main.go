package version

import (
	"fmt"
	"encoding/json"
	"github.com/jetrails/jrctl/sdk/utils"
	"github.com/jetrails/jrctl/sdk/color"
	"github.com/parnurzeal/gorequest"
	vercmp "github.com/hashicorp/go-version"
)

const VersionString string = "1.0.4"
const ReleasesUrl = "https://api.github.com/repos/jetrails/jrctl/releases"
const TagUrlTemplate = "https://github.com/jetrails/jrctl/releases/tag/%s"

type ReleaseEntry struct {
	TagName string `json:"tag_name"`
}

type ReleaseResponse struct {
	Collection [] ReleaseEntry
}

func CheckVersion ( debug bool ) {
	var cacheWindow int64 = 60 * 60
	cachedVersion, hit := utils.GetCache ( "version:" + VersionString, cacheWindow )
	versionObj, _ := vercmp.NewVersion ( VersionString )
	if hit {
		cachedVersionObj, _ := vercmp.NewVersion ( cachedVersion )
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
	request.SetDebug ( debug )
	response, body, _ := request.Get ( ReleasesUrl ).End ()
	if response.StatusCode == 200 {
		releases := make ( [] ReleaseEntry, 0 )
		json.Unmarshal ( [] byte ( body ), &releases )
		newest := releases [ 0 ]
		utils.SetCache ( "version:" + VersionString, newest.TagName, cacheWindow )
		targetVersionObj, _ := vercmp.NewVersion ( newest.TagName )
		if versionObj.LessThan ( targetVersionObj ) {
			fmt.Printf (
				"Software is out-of-date. Update to the latest version: %s.\n%s\n\n",
				newest.TagName,
				color.RedString ( fmt.Sprintf ( TagUrlTemplate, newest.TagName ) ),
			)
		}
	}
}

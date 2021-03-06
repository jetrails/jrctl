package version

import (
	"encoding/json"
	"errors"
	"fmt"
	"strings"

	vercmp "github.com/hashicorp/go-version"
	"github.com/jetrails/jrctl/pkg/cache"
	"github.com/jetrails/jrctl/pkg/color"
	"github.com/parnurzeal/gorequest"
)

const VersionString string = "2.1.0"
const ReleasesUrl = "https://api.github.com/repos/jetrails/jrctl/releases"
const TagUrlTemplate = "https://github.com/jetrails/jrctl/releases/tag/%s"

type ReleaseEntry struct {
	TagName    string `json:"tag_name"`
	PreRelease bool   `json:"prerelease"`
	Draft      bool   `json:"draft"`
}

func FindStable(releases []ReleaseEntry) (ReleaseEntry, error) {
	for _, release := range releases {
		isAlpha := strings.HasSuffix(release.TagName, "-alpha")
		isBeta := strings.HasSuffix(release.TagName, "-beta")
		if !release.PreRelease && !release.Draft && !isAlpha && !isBeta {
			return release, nil
		}
	}
	return ReleaseEntry{}, errors.New("No stable version found")
}

func CheckVersion(debug bool) {
	cachedVersion, _ := cache.Get("latest-version-" + VersionString)
	versionObj, _ := vercmp.NewVersion(VersionString)
	if len(string(cachedVersion)) > 0 {
		cachedVersionObj, _ := vercmp.NewVersion(string(cachedVersion))
		if versionObj.LessThan(cachedVersionObj) {
			fmt.Printf(
				"Software is out-of-date. Update to the latest version: %s.\n%s\n\n",
				string(cachedVersion),
				color.RedString(fmt.Sprintf(TagUrlTemplate, string(cachedVersion))),
			)
		}
		return
	}
	var request = gorequest.New()
	request.SetDebug(debug)
	response, body, _ := request.
		Get(ReleasesUrl).
		Query("page=1&per_page=100").
		End()
	if response != nil && response.StatusCode == 200 {
		var releases []ReleaseEntry
		if error := json.Unmarshal([]byte(body), &releases); error == nil {
			if newest, error := FindStable(releases); error == nil {
				cache.Set("latest-version-"+VersionString, []byte(newest.TagName), 60*60)
				targetVersionObj, _ := vercmp.NewVersion(newest.TagName)
				if versionObj.LessThan(targetVersionObj) {
					fmt.Printf(
						"Software is out-of-date. Update to the latest version: %s.\n%s\n\n",
						newest.TagName,
						color.RedString(fmt.Sprintf(TagUrlTemplate, newest.TagName)),
					)
				}
			}
		}
	}
}

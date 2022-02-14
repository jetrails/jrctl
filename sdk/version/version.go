package version

import (
	"encoding/json"
	"errors"
	"strings"
	"time"

	vercmp "github.com/hashicorp/go-version"
	"github.com/jetrails/jrctl/pkg/cache"
	"github.com/parnurzeal/gorequest"
)

const VersionString string = "2.3.1"
const ReleasesUrl = "https://github.com/jetrails/jrctl/releases"
const ReleasesApiUrl = "https://api.github.com/repos/jetrails/jrctl/releases"
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
		isLatest := release.TagName == "latest"
		if !release.PreRelease && !release.Draft && !isAlpha && !isBeta && !isLatest {
			return release, nil
		}
	}
	return ReleaseEntry{}, errors.New("no stable version found")
}

func CheckVersion(debug bool) (bool, string) {
	versionObj, _ := vercmp.NewVersion(VersionString)
	var request = gorequest.New()
	request.SetDebug(debug)
	response, body, _ := request.
		Timeout(5 * time.Second).
		Get(ReleasesApiUrl).
		Query("page=1&per_page=100").
		End()
	if response != nil && response.StatusCode == 200 {
		var releases []ReleaseEntry
		if error := json.Unmarshal([]byte(body), &releases); error == nil {
			if newest, error := FindStable(releases); error == nil {
				cache.Set("latest-version-"+VersionString, []byte(newest.TagName), 60*60)
				targetVersionObj, _ := vercmp.NewVersion(newest.TagName)
				return !versionObj.LessThan(targetVersionObj), newest.TagName
			}
		}
	}
	return false, ""
}

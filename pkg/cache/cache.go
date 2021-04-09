package cache

import (
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"regexp"
	"sort"
	"strconv"
	"time"
)

var UserCacheDir string
var ProgramName string

func Clean() error {
	if UserCacheDir != "" {
		if info, error := ioutil.ReadDir(UserCacheDir); error == nil {
			now := time.Now().Unix()
			for _, file := range info {
				filename := file.Name()
				if re := regexp.MustCompile(`-(\d+)$`); re.MatchString(filename) {
					matches := re.FindStringSubmatch(filename)
					expires, _ := strconv.Atoi(matches[1])
					if now >= int64(expires) {
						os.Remove(path.Join(UserCacheDir, filename))
					}
				}
			}
		} else {
			return error
		}
	}
	return nil
}

func Set(key string, value []byte, ttl int64) (string, error) {
	expires := time.Now().Unix() + ttl
	filename := fmt.Sprintf("%s-%d", key, expires)
	filepath := path.Join(UserCacheDir, filename)
	if file, error := os.OpenFile(filepath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0600); error == nil {
		defer file.Close()
		file.Write(value)
		return filepath, nil
	} else {
		return "", error
	}
}

func Get(key string) ([]byte, error) {
	if UserCacheDir != "" {
		if error := Clean(); error == nil {
			if info, error := ioutil.ReadDir(UserCacheDir); error == nil {
				files := []string{}
				for _, file := range info {
					files = append(files, file.Name())
				}
				sort.Sort(sort.Reverse(sort.StringSlice(files)))
				for _, filename := range files {
					pattern := regexp.QuoteMeta(key) + `-\d+$`
					if re := regexp.MustCompile(pattern); re.MatchString(filename) {
						return ioutil.ReadFile(path.Join(UserCacheDir, filename))
					}
				}
				return []byte{}, fmt.Errorf("Could not find cache entry with key %q.", key)
			} else {
				return []byte{}, error
			}
		} else {
			return []byte{}, error
		}
	}
	return []byte{}, fmt.Errorf("Could not find cache entry with key %q.", key)
}

func init() {
	ProgramName := path.Base(os.Args[0])
	if dir, error := os.UserCacheDir(); error == nil {
		UserCacheDir = path.Join(dir, ProgramName)
		if _, error := os.Stat(UserCacheDir); os.IsNotExist(error) {
			os.Mkdir(UserCacheDir, 0700)
		}
	}
}

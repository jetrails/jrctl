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
		if info, err := ioutil.ReadDir(UserCacheDir); err == nil {
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
			return err
		}
	}
	return nil
}

func Set(key string, value []byte, ttl int64) (string, error) {
	expires := time.Now().Unix() + ttl
	filename := fmt.Sprintf("%s-%d", key, expires)
	filepath := path.Join(UserCacheDir, filename)
	if file, err := os.OpenFile(filepath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0600); err == nil {
		defer file.Close()
		file.Write(value)
		return filepath, nil
	} else {
		return "", err
	}
}

func Get(key string) ([]byte, error) {
	if UserCacheDir != "" {
		if err := Clean(); err == nil {
			if info, err := ioutil.ReadDir(UserCacheDir); err == nil {
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
				return []byte{}, err
			}
		} else {
			return []byte{}, err
		}
	}
	return []byte{}, fmt.Errorf("Could not find cache entry with key %q.", key)
}

func init() {
	ProgramName := path.Base(os.Args[0])
	if dir, err := os.UserCacheDir(); err == nil {
		UserCacheDir = path.Join(dir, ProgramName)
		if _, err := os.Stat(UserCacheDir); os.IsNotExist(err) {
			os.Mkdir(UserCacheDir, 0700)
		}
	}
}

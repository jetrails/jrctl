package cache

import (
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"time"
)

func getCacheFileName(key string, window int64) string {
	now := time.Now().Unix()
	postfix := now / window
	return fmt.Sprintf("jrctl-%s-%d", key, postfix)
}

func SetCache(key string, value string, window int64) string {
	name := getCacheFileName(key, window)
	filepath := path.Join(os.TempDir(), name)
	file, _ := os.Create(filepath)
	_, _ = file.WriteString(value)
	file.Close()
	return filepath
}

func GetCache(key string, window int64) (string, bool) {
	name := getCacheFileName(key, window)
	filepath := path.Join(os.TempDir(), name)
	info, error := os.Stat(filepath)
	if os.IsNotExist(error) || info.IsDir() {
		return "", false
	}
	contents, _ := ioutil.ReadFile(filepath)
	return string(contents), true
}

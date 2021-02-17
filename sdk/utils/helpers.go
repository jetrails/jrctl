package utils

import (
	"regexp"
	"strings"
	"os/user"
)

func Examples ( examples [] string ) string {
	for i, e := range examples {
		examples [ i ] = "  " + e
	}
	return strings.Join ( examples, "\n" )
}

func GetUser () string {
	user, error := user.Current ()
	if error != nil {
		return "unknown"
	}
	return SafeString ( user.Username )
}

func SafeString ( input string ) string {
	return regexp.MustCompile ("[^a-zA-Z0-9]+").ReplaceAllString ( input, "_" )
}

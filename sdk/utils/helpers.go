package utils

import (
	"regexp"
	"strings"
	"os"
	"os/user"
)

func Includes ( a string, list [] string ) bool {
	for _, b := range list {
		if b == a {
			return true
		}
	}
	return false
}

func Examples ( examples [] string ) string {
	space := "  "
	if os.Getenv ("JR_DOCS") == "true" {
		space = ""
	}
	for i, e := range examples {
		examples [ i ] = space + e
	}
	return strings.Join ( examples, "\n" )
}

func Combine ( lines [] string ) string {
	return strings.Join ( lines, "\n\n" )
}

func Paragraph ( lines [] string ) string {
	width := 80
	result := ""
	line := ""
	delim := ""
	lineDelim := ""
	combined := strings.Join ( lines, " " )
	for _, word := range strings.Split ( combined, " " ) {
		temp := line + delim + word
		if ( len ( temp ) <= width ) {
			line = temp
		} else {
			result = result + lineDelim + line
			line = word
			lineDelim = "\n"
		}
		delim = " "
	}
	return result + lineDelim + line
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

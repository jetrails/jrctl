package utils

import "strings"
import "os"
import "fmt"
import "bufio"
import "io/ioutil"
import "golang.org/x/crypto/ssh/terminal"

func PromptPassword ( prompt string, value string ) string {
	if value == "" {
		fmt.Printf ( prompt )
		input, _ := terminal.ReadPassword ( 0 )
		fmt.Printf ("\r\033[K")
		return string ( input )
	}
	return value
}

func ReadFile ( path string ) string {
	content, error := ioutil.ReadFile ( path )
	if error != nil {
		HandleErrorResponse ( InternalError ( error ) )
	}
	return string ( content )
}

func PromptContent ( prompt string ) string {
	var breaks = 0
	var input = ""
	reader := bufio.NewReader ( os.Stdin )
	fmt.Printf ("Hint: Entering 3 empty lines in a row will halt input.\n")
	fmt.Printf ( "%s:\n\n", prompt )
	for breaks < 2 {
		line, _ := reader.ReadString ('\n')
		input += line
		if line == "\n" {
			breaks++
		} else {
			breaks = 0
		}
	}
	return strings.TrimSpace ( input )
}

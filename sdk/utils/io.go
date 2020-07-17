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
	var input = ""
	reader := bufio.NewReader ( os.Stdin )
	fmt.Printf ( "%s (Ctrl-D to end input):\n\n", prompt )
	for true {
		byte, _ := reader.ReadByte ()
		if byte == 0 {
			break
		}
		input += string ( byte )
	}
	return strings.TrimSpace ( input )
}

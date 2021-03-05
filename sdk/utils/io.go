package utils

import (
	"strings"
	"os"
	"fmt"
	"bufio"
	"io/ioutil"
	"golang.org/x/crypto/ssh/terminal"
)

func PromptPassword ( prompt string, value string ) string {
	if value == "" {
		fmt.Printf ( prompt )
		input, _ := terminal.ReadPassword ( 0 )
		fmt.Printf ("\r\033[K")
		return string ( input )
	}
	return value
}

func ReadFile ( path string ) ( string, error ) {
	content, error := ioutil.ReadFile ( path )
	if error != nil {
		return "", error
	}
	return string ( content ), nil
}

func SaveScreen () {
	fmt.Printf ("\033[?1049h\033[H")
}

func RestoreScreen () {
	fmt.Printf ("\033[?1049l")
	fmt.Printf ("\033[34h\033[?25h")
}

func PromptContent ( prompt string ) string {
	SaveScreen ()
	defer RestoreScreen ()
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

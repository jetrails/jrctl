package color

import (
	"os"
	"fmt"
)

const Format string = "\x1b[%dm%s\x1b[0m"
const Black int = 30
const Red int = 31
const Green int = 32
const Yellow int = 33
const Blue int = 34
const Magenta int = 35
const Cyan int = 36
const White int = 37

func hasColor () bool {
	return os.Getenv ("JR_COLOR") == "" || os.Getenv ("JR_COLOR") == "true"
}

func GetLogo () string {
	return GreenString (">") + "jetrails" + GreenString ("_")
}

func GreenString ( data string ) string {
	if hasColor () {
		return fmt.Sprintf ( Format, Green, data )
	}
	return data
}

func CyanString ( data string ) string {
	if hasColor () {
		return fmt.Sprintf ( Format, Cyan, data )
	}
	return data
}

func RedString ( data string ) string {
	if hasColor () {
		return fmt.Sprintf ( Format, Red, data )
	}
	return data
}

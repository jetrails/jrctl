package color

import (
	"fmt"
	"os"
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

var hasColorCached *bool

func hasColor() bool {
	if hasColorCached != nil {
		return *hasColorCached
	}
	_, result := os.LookupEnv("NO_COLOR")
	result = !result
	hasColorCached = &result
	return *hasColorCached
}

func BlackString(data string) string {
	if hasColor() {
		return fmt.Sprintf(Format, Black, data)
	}
	return data
}

func RedString(data string) string {
	if hasColor() {
		return fmt.Sprintf(Format, Red, data)
	}
	return data
}

func GreenString(data string) string {
	if hasColor() {
		return fmt.Sprintf(Format, Green, data)
	}
	return data
}

func YellowString(data string) string {
	if hasColor() {
		return fmt.Sprintf(Format, Yellow, data)
	}
	return data
}

func BlueString(data string) string {
	if hasColor() {
		return fmt.Sprintf(Format, Blue, data)
	}
	return data
}

func MagentaString(data string) string {
	if hasColor() {
		return fmt.Sprintf(Format, Magenta, data)
	}
	return data
}

func CyanString(data string) string {
	if hasColor() {
		return fmt.Sprintf(Format, Cyan, data)
	}
	return data
}

func WhiteString(data string) string {
	if hasColor() {
		return fmt.Sprintf(Format, White, data)
	}
	return data
}

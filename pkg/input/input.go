package input

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func ReadFile(path string) (string, error) {
	content, error := ioutil.ReadFile(path)
	if error != nil {
		return "", error
	}
	return string(content), nil
}

func PromptYesNo(prompt string) bool {
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Printf("%s [y/n]: ", prompt)
		response, error := reader.ReadString('\n')
		if error != nil {
			return false
		}
		response = strings.ToLower(strings.TrimSpace(response))
		if response == "y" || response == "yes" {
			return true
		} else if response == "n" || response == "no" {
			return false
		}
	}
}

func SaveScreen() {
	fmt.Printf("\033[?1049h\033[H")
}

func RestoreScreen() {
	fmt.Printf("\033[?1049l")
	fmt.Printf("\033[34h\033[?25h")
}

func PromptContent(prompt string) string {
	SaveScreen()
	defer RestoreScreen()
	var input = ""
	reader := bufio.NewReader(os.Stdin)
	fmt.Printf("%s (Ctrl-D to end input):\n\n", prompt)
	for true {
		byte, _ := reader.ReadByte()
		if byte == 0 {
			break
		}
		input += string(byte)
	}
	return strings.TrimSpace(input)
}

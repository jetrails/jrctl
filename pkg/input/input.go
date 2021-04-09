package input

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func saveScreen() {
	fmt.Printf("\033[?1049h\033[H")
}

func restoreScreen() {
	fmt.Printf("\033[?1049l")
	fmt.Printf("\033[34h\033[?25h")
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

func PromptContent(prompt string) string {
	saveScreen()
	defer restoreScreen()
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

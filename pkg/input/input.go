package input

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
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
	if value, error := getInputFromEditor(); error == nil {
		return value
	}
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

func openFileInEditor(filename string) error {
	if editor := os.Getenv("EDITOR"); editor != "" {
		if executable, error := exec.LookPath(editor); error == nil {
			cmd := exec.Command(executable, filename)
			cmd.Stdin = os.Stdin
			cmd.Stdout = os.Stdout
			cmd.Stderr = os.Stderr
			return cmd.Run()
		} else {
			return error
		}
	} else {
		return fmt.Errorf("EDITOR env variable not set to anything")
	}
}

func getInputFromEditor() (string, error) {
	file, error := ioutil.TempFile(os.TempDir(), "*")
	if error != nil {
		return "", error
	}
	filename := file.Name()
	defer os.Remove(filename)
	if error = file.Close(); error != nil {
		return "", error
	}
	if error = openFileInEditor(filename); error != nil {
		return "", error
	}
	bytes, error := ioutil.ReadFile(filename)
	if error != nil {
		return "", error
	}
	return string(bytes), nil
}

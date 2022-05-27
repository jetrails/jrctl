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
		response, err := reader.ReadString('\n')
		if err != nil {
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
	if value, err := getInputFromEditor(); err == nil {
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
		if executable, err := exec.LookPath(editor); err == nil {
			cmd := exec.Command(executable, filename)
			cmd.Stdin = os.Stdin
			cmd.Stdout = os.Stdout
			cmd.Stderr = os.Stderr
			return cmd.Run()
		} else {
			return err
		}
	} else {
		return fmt.Errorf("EDITOR env variable not set to anything")
	}
}

func getInputFromEditor() (string, error) {
	file, err := ioutil.TempFile(os.TempDir(), "*")
	if err != nil {
		return "", err
	}
	filename := file.Name()
	defer os.Remove(filename)
	if err = file.Close(); err != nil {
		return "", err
	}
	if err = openFileInEditor(filename); err != nil {
		return "", err
	}
	bytes, err := ioutil.ReadFile(filename)
	if err != nil {
		return "", err
	}
	return string(bytes), nil
}

func HasDataInPipe() bool {
	if stat, _ := os.Stdin.Stat(); stat.Mode()&os.ModeCharDevice == 0 {
		return true
	}
	return false
}

func GetPipeData() string {
	if bytes, err := ioutil.ReadAll(os.Stdin); err == nil {
		return strings.TrimSpace(string(bytes))
	}
	return ""
}

func GetFirstArgumentOrPipe(args []string) string {
	if len(args) == 0 {
		return GetPipeData()
	} else {
		return args[0]
	}
	return ""
}

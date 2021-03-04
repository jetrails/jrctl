package utils

import (
	"fmt"
)

func CollectErrors ( errors [] error ) [] string {
	var messages [] string
	for _, error := range errors {
		messages = append ( messages, error.Error () )
	}
	return messages
}

func PrintErrors ( code int, status string ) {
	if code != 200 {
		if status == "" {
			status = "Client Error"
		}
		fmt.Printf ( "Error %d: %s\n", code, status )
	}
}

func PrintMessages ( messages [] string ) {
	for _, message := range messages {
		fmt.Printf ( "Message: %s\n", message )
	}
}

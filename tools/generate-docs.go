package main

import (
	"log"
	"github.com/jetrails/jrctl/internal"
	"github.com/spf13/cobra/doc"
)

func main () {
	if error := doc.GenMarkdownTree ( internal.GetRootCommand (), "./docs" ); error != nil {
		log.Fatal ( error )
	}
}

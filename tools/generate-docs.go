package main

import (
	"fmt"
	"log"
	"time"

	"github.com/jetrails/jrctl/internal"
	"github.com/spf13/cobra/doc"
)

func main() {
	now := time.Now()
	header := &doc.GenManHeader{
		Title:   "JetRails CLI",
		Section: "1",
		Source:  fmt.Sprintf("Copyright %4d ADF, Inc. All Rights Reserved ", now.Year()),
		Date:    &now,
	}
	if err := doc.GenMarkdownTree(internal.RootCmd, "./docs"); err != nil {
		log.Fatal(err)
	}
	if err := doc.GenManTree(internal.RootCmd, header, "./man"); err != nil {
		log.Fatal(err)
	}
}

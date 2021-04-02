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
	if error := doc.GenMarkdownTree(internal.GetRootCommand(), "./docs"); error != nil {
		log.Fatal(error)
	}
	if error := doc.GenManTree(internal.GetRootCommand(), header, "./man"); error != nil {
		log.Fatal(error)
	}
}

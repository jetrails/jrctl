package main

import (
	"github.com/jetrails/jrctl/internal"
	"github.com/jetrails/jrctl/pkg/env"
)

func main() {
	internal.RootCmd.Execute()
}

func init() {
	env.EnvPrefix = "JR_"
}

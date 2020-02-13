package main

import (
	"os"

	"github.com/wxytjustb/devicebeat/cmd"

	_ "github.com/wxytjustb/devicebeat/include"
)

func main() {
	if err := cmd.RootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}

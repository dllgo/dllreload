package main

import (
	"flag"
	"fmt"
	"github.com/dllgo/dllreload/runner"
	"os"
)

func main() {
	configPath := flag.String("c", "", "config file path")
	flag.Parse()

	if *configPath != "" {
		if _, err := os.Stat(*configPath); err != nil {
			fmt.Printf("Can't find config file `%s`\n", *configPath)
			os.Exit(1)
		} else {
			os.Setenv("RUNNER_CONFIG_PATH", *configPath)
		}
	}

	runner.Start()
}

package main

import (
	"order-service/internal/cli"
	"os"
)

func main() {
	//os.Setenv("KV_VIPER_FILE", "config.yaml")
	if !cli.Run(os.Args) {
		os.Exit(1)
	}
}

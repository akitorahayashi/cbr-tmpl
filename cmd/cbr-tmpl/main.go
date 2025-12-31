package main

import (
	"os"

	"github.com/akitorahayashi/cbr-tmpl/internal"
	"github.com/akitorahayashi/cbr-tmpl/internal/cmd"
)

func main() {
	storage := internal.NewFilesystemStorage("")
	rootCmd := cmd.NewRootCmd(storage)

	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}

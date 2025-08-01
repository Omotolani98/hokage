package main

import (
	"context"
	"os"

	"github.com/Omotolani98/hokage/pkg/commands"

	"github.com/charmbracelet/fang"
	"github.com/spf13/cobra"
)

// FEATURE: A file arranger - moves files to a better folder based on the file extensions
func main() {
	rootCmd := &cobra.Command{
		Use:   "hokage",
		Short: "A simple file arranger - moves files to a better folder based on the file extensions",
	}

	rootCmd.AddCommand(
		commands.Apply(),
	)

	if err := fang.Execute(context.Background(), rootCmd); err != nil {
		os.Exit(1)
	}
}

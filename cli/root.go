package cli

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCommand = &cobra.Command{
	Use:   "lfrd",
	Short: "LFRD is a simple static website generator",
	Long:  "LFRD (Lightweight Framework for Rapid Design) is a tool for quickly creating personal and commercial static websites",
}

// Execute runs the root command
func Execute() {
	if err := rootCommand.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	rootCommand.AddCommand(initCliCommand)
}

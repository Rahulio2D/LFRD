package cli

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/Rahulio2D/LFRD/helpers"
	"github.com/spf13/cobra"
)

var installCliCommand = &cobra.Command{
	Use:   "install",
	Short: "Install the default templates into your home directory",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			fmt.Println("Error: Please provide the full path to the LFRD source directory as an argument.")
			return
		}

		lfrdSourceDirectory := args[0]
		templatesSourceDirectory := filepath.Join(lfrdSourceDirectory, "templates")

		homeDirectory, err := os.UserHomeDir()
		if err != nil {
			fmt.Println("Error finding home directory:", err)
			return
		}
		targetDirectory := filepath.Join(homeDirectory, ".lfrd", "templates")

		// Check if directory exists and delete if it does
		if _, err := os.Stat(targetDirectory); err == nil {
			fmt.Printf("Directory %s already exists, deleting...\n", targetDirectory)
			if err := os.RemoveAll(targetDirectory); err != nil {
				fmt.Println("Error deleting existing directory:", err)
				return
			}
		}

		if err := helpers.CopyDir(templatesSourceDirectory, targetDirectory); err != nil {
			fmt.Println("Error installing templates:", err)
			return
		}

		fmt.Printf("Templates installed to %s\n", targetDirectory)
	},
}

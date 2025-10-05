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

		homeDirectory, err := os.UserHomeDir()
		if err != nil {
			fmt.Println("Error finding home directory:", err)
			return
		}

		templateDirectory := filepath.Join(homeDirectory, ".lfrd", "templates")
		if err := helpers.CopyDir(lfrdSourceDirectory, templateDirectory); err != nil {
			fmt.Println("Error installing templates:", err)
			return
		}

		fmt.Printf("Templates installed to %s\n", templateDirectory)
	},
}

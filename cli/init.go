package cli

import (
	"fmt"
	"os"

	"github.com/Rahulio2D/LFRD/helpers"
	"github.com/spf13/cobra"
)

var initCliCommand = &cobra.Command{
	Use:   "init [name]",
	Short: "Initialise a new LFRD website project",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		projectName := args[0]

		// Create project folder and assign permissions (rwxr-xr-x)
		if err := os.Mkdir(projectName, 0755); err != nil {
			fmt.Println("Error creating project:", err)
			return
		}

		templateDirectory := "templates"
		if err := helpers.CopyDir(templateDirectory, projectName); err != nil {
			fmt.Println("Error copying LFRD templates files:", err)
			return
		}

		fmt.Printf("✅ Project '%s' created successfully!\n", projectName)
	},
}

package cli

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/Rahulio2D/LFRD/helpers"
	"github.com/spf13/cobra"
)

var buildCliCommand = &cobra.Command{
	Use:   "build",
	Short: "Build the website into the docs directory ready for deployment.",
	Run: func(cmd *cobra.Command, args []string) {
		workingDirectory, err := os.Getwd()
		if err != nil {
			fmt.Println("Error getting working directory:", err)
			return
		}
		destinationDirectory := filepath.Join(workingDirectory, "docs")

		fmt.Println("Building site...")
		if err := helpers.CopyDir(workingDirectory, destinationDirectory); err != nil {
			fmt.Println("Error building site:", err)
			return
		}

		fmt.Println("✅ Site built successfully!")
	},
}

// TODO: Right now all this does is copy the current directory to docs.
// This needs to be expanded to actually process templates, markdown, etc.
// Ideally it should first fill out templates in the HTML files
// Then another update to convert markdown files to HTML (i.e. blog posts)

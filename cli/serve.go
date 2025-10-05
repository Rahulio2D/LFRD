package cli

import (
	"fmt"
	"net/http"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
)

var serveCliCommand = &cobra.Command{
	Use:   "serve",
	Short: "Serve the site locally for testing",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		port := args[0]

		workingDirectory, _ := os.Getwd()
		siteDirectory := filepath.Join(workingDirectory, "docs")

		fileServer := http.FileServer(http.Dir(siteDirectory))
		http.Handle("/", fileServer)

		fmt.Println("Serving site at http://localhost:" + port)
		err := http.ListenAndServe(":"+port, nil)
		if err != nil {
			fmt.Println("Error starting server:", err)
		}
	},
}

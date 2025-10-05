package helpers

import (
	"io"
	"io/fs"
	"os"
	"path/filepath"
)

// CopyDir copies all files and subdirectories from sourceLocation to destinationLocation
func CopyDir(sourceLocation string, destinationLocation string) error {
	return filepath.WalkDir(sourceLocation, func(fileName string, fileEntry fs.DirEntry, err error) error {
		if err != nil {
			return err
		}

		relPath, err := filepath.Rel(sourceLocation, fileName)
		if err != nil {
			return err
		}

		targetPath := filepath.Join(destinationLocation, relPath)

		if fileEntry.IsDir() {
			if fileEntry.Name() == "docs" {
				return filepath.SkipDir
			}
			return os.MkdirAll(targetPath, 0755)
		}

		return CopyFile(fileName, targetPath)
	})
}

func CopyFile(sourceLocation string, destinationLocation string) error {
	in, err := os.Open(sourceLocation)
	if err != nil {
		return err
	}
	defer in.Close()

	out, err := os.Create(destinationLocation)
	if err != nil {
		return err
	}
	defer out.Close()

	_, err = io.Copy(out, in)
	if err != nil {
		return err
	}

	return out.Close()
}

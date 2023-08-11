package tools

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
)

func CopyFolder(source, destination string) error {
	sourceInfo, err := os.Stat(source)
	if err != nil {
		return err
	}

	if !sourceInfo.IsDir() {
		return fmt.Errorf("el origen debe ser una carpeta")
	}

	// Crea la carpeta de destino si no existe
	if err := os.MkdirAll(destination, sourceInfo.Mode()); err != nil {
		return err
	}

	entries, err := os.ReadDir(source)
	if err != nil {
		return err
	}

	for _, entry := range entries {
		sourcePath := filepath.Join(source, entry.Name())
		destinationPath := filepath.Join(destination, entry.Name())

		if entry.IsDir() {
			if err := CopyFolder(sourcePath, destinationPath); err != nil {
				return err
			}
		} else {
			if err := CopyFile(sourcePath, destinationPath); err != nil {
				return err
			}
		}
	}

	return nil
}

func CopyFile(source, destination string) error {
	sourceFile, err := os.Open(source)
	if err != nil {
		return err
	}
	defer sourceFile.Close()

	destinationFile, err := os.Create(destination)
	if err != nil {
		return err
	}
	defer destinationFile.Close()

	_, err = io.Copy(destinationFile, sourceFile)
	return err
}

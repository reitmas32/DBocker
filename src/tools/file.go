package tools

import (
	"os"
	"strings"
)

func ReplaceTextInFile(filePath, oldText, newText string) error {
	content, err := os.ReadFile(filePath)
	if err != nil {
		return err
	}

	newContent := strings.ReplaceAll(string(content), oldText, newText)

	err = os.WriteFile(filePath, []byte(newContent), os.ModePerm)
	if err != nil {
		return err
	}

	return nil
}

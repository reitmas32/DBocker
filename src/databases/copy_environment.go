package databases

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/UniHacksOrg/DBocker/src/models"
	"github.com/UniHacksOrg/DBocker/src/tools"
)

func CopyDB_Project(base string, config models.DataBase) string {
	// Obtener la ruta del directorio actual
	currentDir, err := os.Getwd()
	if err != nil {
		fmt.Println("Error al obtener la ruta del directorio actual:", err)
		return "error"
	}
	sourceFolder := filepath.Join(currentDir, "..", "assets", base)
	destinationFolder := filepath.Join(currentDir, "sandbox", config.Name)

	err = tools.CopyFolder(sourceFolder, destinationFolder)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Carpeta copiada exitosamente")
	}
	return destinationFolder
}

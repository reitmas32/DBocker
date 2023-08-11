package databases

import (
	"fmt"
	"strconv"

	"github.com/UniHacksOrg/DBocker/src/models"
	"github.com/UniHacksOrg/DBocker/src/tools"
)

func ConfigureDB_Project(destinationFolder string, config models.DataBase) {
	placeholders := map[string]string{
		"__container_name__": config.ContainerName,
		"__port__":           strconv.Itoa(config.Port),
		"__name_db__":        config.Name,
		"__password__":       config.Password,
		"__root_password__":  config.RootPassword,
		"__user__":           config.User,
		"__volumen_route__":  fmt.Sprintf("%s/volumen", destinationFolder),
	}

	composeFilePath := fmt.Sprintf("%s/docker-compose.yml", destinationFolder)
	envFilePath := fmt.Sprintf("%s/.env", destinationFolder)

	for placeholder, value := range placeholders {
		err := tools.ReplaceTextInFile(composeFilePath, placeholder, value)
		if err != nil {
			fmt.Printf("Error reemplazando \"%s\" en %s : %s\n", placeholder, composeFilePath, err)
		} else {
			fmt.Printf("Texto \"%s\" reemplazado exitosamente en %s\n", placeholder, composeFilePath)
		}

		err = tools.ReplaceTextInFile(envFilePath, placeholder, value)
		if err != nil {
			fmt.Printf("Error reemplazando \"%s\" en %s : %s\n", placeholder, envFilePath, err)
		} else {
			fmt.Printf("Texto \"%s\" reemplazado exitosamente en %s\n", placeholder, envFilePath)
		}
	}

}

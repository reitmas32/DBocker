package main

import (
	"fmt"
	"log"

	"github.com/UniHacksOrg/DBocker/src/api/routes"
	"github.com/UniHacksOrg/DBocker/src/config"
)

func main() {
	log.Println("Run Server on ENVIRONMENT:", config.ENVIRONMENT)
	log.Println("Run Server on PORT:", config.PORT)
	log.Println("Run Server on DEBUG:", config.DEBUG)
	log.Println("Run Server on API_VERSION:", config.API_VERSION)

	config.LoadEnvs()

	config.SetupDB()

	config.MigrateDB()

	// Configuración de la conexión a la base de datos MySQL
	config.SetupRouter()

	routes.Index()

	routes.DataBases()

	config.Router.Run(fmt.Sprintf(":%d", config.PORT))
}

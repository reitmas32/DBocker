package views

import (
	"fmt"
	"strings"

	"github.com/UniHacksOrg/DBocker/src/api/schemas"
	"github.com/UniHacksOrg/DBocker/src/api/services"
	"github.com/UniHacksOrg/DBocker/src/databases"
	"github.com/UniHacksOrg/DBocker/src/models"
	"github.com/UniHacksOrg/DBocker/src/tools"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

// @Summary add a new item of the computers
// @ID create-computer
// @Tags Computers
// @Produce json
// @Param data body schemas.ComputerCreateSchema true "Schema by Create New Computer"
// @Success 200 {object} models.Response
// @Failure 400 {object} models.Response
// @Router /api/v1/computer [post]
func DataBase_POST(c *gin.Context) {

	responseCreateDataBase := models.Response{
		Message: "No Create DataBase",
		Success: false,
		Data:    "{}",
	}

	// Decodificar el objeto JSON recibido
	var dataBaseCreateSchema schemas.DataBaseCreateSchema
	if err := c.ShouldBindWith(&dataBaseCreateSchema, binding.JSON); err != nil {
		responseCreateLoan := models.Response{
			Message: "Error to Get Content JSON",
			Success: false,
			Data:    strings.Split(err.Error(), "\n"),
		}
		c.Header("Content-Type", "application/json")
		c.JSON(200, responseCreateLoan)
		return
	}

	database := models.DataBase{
		Name:          dataBaseCreateSchema.Name,
		User:          dataBaseCreateSchema.User,
		Password:      dataBaseCreateSchema.Password,
		Port:          2000,
		ContainerName: dataBaseCreateSchema.Name + "_" + tools.GenerateRandomString(5),
		Engine:        dataBaseCreateSchema.Engine,
		RootPassword:  dataBaseCreateSchema.Password,
	}

	// Registrar la DB
	result, message, new_database := services.CreateDataBase(database)

	if result {

		responseCreateDataBase = models.Response{
			Message: message,
			Success: result,
			Data:    new_database,
		}
	} else {
		responseCreateDataBase = models.Response{
			Message: message,
			Success: responseCreateDataBase.Success,
			Data:    nil,
		}
	}

	destinationFolder := databases.CopyDB_Project(dataBaseCreateSchema.Engine, new_database)

	databases.ConfigureDB_Project(destinationFolder, new_database)

	// Darle permisos al dir Volumen
	_, message = tools.GrantVolumePermissions(destinationFolder)

	fmt.Println(message)

	// docker-compose up -d

	_, message = tools.ComposeUp(destinationFolder)

	fmt.Println(message)

	c.Header("Content-Type", "application/json")
	c.JSON(200, responseCreateDataBase)
}

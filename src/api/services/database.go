package services

import (
	"github.com/UniHacksOrg/DBocker/src/config"
	"github.com/UniHacksOrg/DBocker/src/models"
)

func CreateDataBase(database models.DataBase) (bool, string, models.DataBase) {
	result := config.DB.Create(&database)
	if result.Error != nil {
		return false, result.Error.Error(), database
	}

	// Update Port after inserting the record with the assigned ID
	database.Port = int(database.ID) + config.DB_INIT_PORT
	config.DB.Save(&database)

	return true, "Create DataBase Successful", database
}

package routes

import (
	"fmt"

	"github.com/UniHacksOrg/DBocker/src/api/views"
	"github.com/UniHacksOrg/DBocker/src/config"
)

func DataBases() {
	config.Router.POST(fmt.Sprintf("/api/%s/database", config.API_VERSION), views.DataBase_POST)
}

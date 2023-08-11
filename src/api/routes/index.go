package routes

import (
	"github.com/UniHacksOrg/DBocker/src/api/views"
	"github.com/UniHacksOrg/DBocker/src/config"
)

func Index() {
	config.Router.GET("/", views.Index)
}

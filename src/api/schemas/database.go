package schemas

type DataBaseCreateSchema struct {
	Name     string `json:"name" binding:"required"`
	Engine   string `json:"engine" binding:"required"`
	User     string `json:"user" binding:"required"`
	Password string `json:"password" binding:"required"`
}

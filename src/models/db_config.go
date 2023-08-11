package models

type DB_Config struct {
	RootPassword  string
	Name          string
	User          string
	Password      string
	Port          int
	ContainerName string
	Engine        string
}

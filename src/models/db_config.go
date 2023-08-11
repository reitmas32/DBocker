package models

import "gorm.io/gorm"

type DataBase struct {
	gorm.Model
	RootPassword  string
	Name          string `gorm:"unique"`
	User          string
	Password      string
	Port          int
	ContainerName string
	Engine        string
}

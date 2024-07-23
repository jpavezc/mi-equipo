package main

import (
	"mi-equipo.cl/mi-equipo/configs"
	"mi-equipo.cl/mi-equipo/models"
)

func init() {
	configs.ConnectToDB()
}

func main() {
	configs.DB.AutoMigrate(&models.Player{})
}

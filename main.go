package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"mi-equipo.cl/mi-equipo/configs"
	"mi-equipo.cl/mi-equipo/routes"
	"mi-equipo.cl/mi-equipo/services"
)

func init() {
	configs.ConnectToDB()
}

func main() {

	services.ElegirRandom()

	r := gin.Default()

	routes.PlayerRouter(r)

	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Hello world from server Go.",
		})
	})

	r.Run()

}

package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"mi-equipo.cl/mi-equipo/configs"
	"mi-equipo.cl/mi-equipo/routes"
)

func init() {
	configs.ConnectToDB()
}

func main() {

	r := gin.Default()

	routes.PlayerRouter(r)

	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Hello world from server Go.",
		})
	})
	r.Run()
}

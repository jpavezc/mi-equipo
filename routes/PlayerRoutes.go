package routes

import (
	"github.com/gin-gonic/gin"
	"mi-equipo.cl/mi-equipo/controllers"
)

func PlayerRouter(router *gin.Engine) {

	routes := router.Group("api/v1/players")
	routes.POST("", controllers.PlayerCreate)
	routes.GET("", controllers.PlayerGet)
	routes.GET("/:id", controllers.PlayerGetById)
	routes.PUT("/:id", controllers.PlayerUpdate)
	routes.DELETE("/:id", controllers.PlayerDelete)

}

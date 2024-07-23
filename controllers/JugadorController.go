package controllers

import (
	"github.com/gin-gonic/gin"
	"mi-equipo.cl/mi-equipo/configs"
	"mi-equipo.cl/mi-equipo/models"
)

type PlayerRequestBody struct {
	Name     string `json:"name"`
	Position string `json:"position"`
	Shoot    int    `json:"shoot"`
}

func PlayerCreate(c *gin.Context) {

	body := PlayerRequestBody{}

	c.BindJSON(&body)

	player := &models.Player{Name: body.Name, Position: body.Position, Shoot: body.Shoot}

	result := configs.DB.Create(&player)

	if result.Error != nil {
		c.JSON(500, gin.H{"Error": "Failed to insert"})
		return
	}

	c.JSON(200, &player)
}

func PlayerGet(c *gin.Context) {
	var persons []models.Player
	configs.DB.Find(&persons)
	c.JSON(200, &persons)
	return
}

func PlayerGetById(c *gin.Context) {
	id := c.Param("id")
	var player models.Player
	configs.DB.First(&player, id)
	c.JSON(200, &player)
	return
}

func PlayerUpdate(c *gin.Context) {

	id := c.Param("id")
	var player models.Player
	configs.DB.First(&player, id)

	body := PlayerRequestBody{}
	c.BindJSON(&body)
	data := &models.Player{Name: body.Name, Position: body.Position, Shoot: body.Shoot}

	result := configs.DB.Model(&player).Updates(data)

	if result.Error != nil {
		c.JSON(500, gin.H{"Error": true, "message": "Failed to update"})
		return
	}

	c.JSON(200, &player)
}

func PlayerDelete(c *gin.Context) {
	id := c.Param("id")
	var player models.Player
	configs.DB.Delete(&player, id)
	c.JSON(200, gin.H{"deleted": true})
	return
}

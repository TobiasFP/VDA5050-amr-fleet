package node

import (
	"TobiasFP/BotNana/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Create(ctx *gin.Context) {
	var node models.Node
	ctx.BindJSON(&node)

	models.DB.Create(&node)
	ctx.JSON(http.StatusOK, node)
}

func All(ctx *gin.Context) {
	var nodes []models.Node
	models.DB.Find(&nodes)
	ctx.JSON(http.StatusOK, gin.H{"data": nodes})
}

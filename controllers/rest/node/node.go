package node

import (
	"TobiasFP/BotNana/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func Create(ctx *gin.Context) {
	var node models.Node
	ctx.BindJSON(&node)
	node.NodeID = uuid.New().String()
	models.DB.Create(&node)
	ctx.JSON(http.StatusOK, node)
}

func All(ctx *gin.Context) {
	var nodes []models.Node
	models.DB.Preload("NodePosition").Find(&nodes)
	ctx.JSON(http.StatusOK, gin.H{"data": nodes})
}

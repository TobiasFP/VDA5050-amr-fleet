package node

import (
	"TobiasFP/BotNana/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// @BasePath /api/v1

func Create(ctx *gin.Context) {
	var node models.NodeMeta
	ctx.BindJSON(&node)
	node.Node.NodeID = uuid.New().String()
	models.DB.Create(&node)
	ctx.JSON(http.StatusOK, node)
}

// @Summary Get all nodes
// @Schemes
// @Accept json
// @Produce json
// @Success 200 {slice} []models.NodeMeta data "ok"
// @Router /nodes/all [get]
func All(ctx *gin.Context) {
	var nodes []models.NodeMeta
	models.DB.Preload("Node").Preload("Node.NodePosition").Find(&nodes)
	ctx.JSON(http.StatusOK, gin.H{"data": nodes})
}

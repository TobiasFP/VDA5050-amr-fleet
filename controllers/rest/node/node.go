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
	models.SqlDB.Create(&node)
	ctx.JSON(http.StatusOK, node)
}

// @Summary Get all nodes
// @Schemes
// @Accept json
// @Produce json
// @Success 200 {slice} []models.NodeMeta data "ok"
// @Router /nodes/all [get]
func All(ctx *gin.Context) {
	var unfilteredNodes []models.NodeMeta

	models.SqlDB.Preload("Node").Preload("Node.NodePosition").Find(&unfilteredNodes)

	// This could be optimised with a better db query, but the gains are
	// most likely  too small
	mapID := ctx.Param("mapid")
	if mapID != "" {
		var filteredNodes []models.NodeMeta
		for _, node := range unfilteredNodes {
			if node.Node.NodePosition.MapID == mapID {
				filteredNodes = append(filteredNodes, node)
			}
		}

		ctx.JSON(http.StatusOK, gin.H{"data": filteredNodes})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"data": unfilteredNodes})
}

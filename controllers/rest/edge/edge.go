package edge

import (
	"TobiasFP/BotNana/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// @BasePath /api/v1

func Create(ctx *gin.Context) {
	var edge models.Edge
	ctx.BindJSON(&edge)
	edge.EdgeID = uuid.New().String()

	models.DB.Create(&edge)
	ctx.JSON(http.StatusOK, edge)
}

// @Summary Get all edges
// @Schemes
// @Accept json
// @Produce json
// @Success 200 {slice} []models.Edge data "ok"
// @Router /edge/all [get]
func All(ctx *gin.Context) {
	var edges []models.Edge
	models.DB.Find(&edges)
	ctx.JSON(http.StatusOK, gin.H{"data": edges})
}

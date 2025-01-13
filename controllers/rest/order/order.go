package order

import (
	"TobiasFP/BotNana/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// @BasePath /api/v1

func Create(ctx *gin.Context) {
	var orderDetails models.OrderTemplateDetails
	// err = json.Unmarshal(jsonData, &orderDetails)
	err := ctx.ShouldBindJSON(&orderDetails)
	if err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
	}
	orderDetails.Order.OrderID = uuid.New().String()

	var nodes []models.Node
	var node models.Node
	for _, id := range orderDetails.NodeIds {

		result := models.DB.Where("nodeId = ?", id).First(&node)
		if result.RowsAffected == 1 && result.Error == nil {
			nodes = append(nodes, node)
		}
	}
	orderDetails.Order.Nodes = nodes
	models.DB.Create(&orderDetails)
	ctx.JSON(http.StatusOK, orderDetails)
}

// @Summary Get all orders
// @Schemes
// @Accept json
// @Produce json
// @Success 200 {slice} []models.Order data "ok"
// @Router /orders/all [get]
func All(ctx *gin.Context) {
	var nodes []models.OrderTemplateDetails
	models.DB.Preload("Order").Find(&nodes)
	ctx.JSON(http.StatusOK, gin.H{"data": nodes})
}

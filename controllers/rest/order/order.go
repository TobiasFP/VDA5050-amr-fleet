package order

import (
	"TobiasFP/BotNana/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type OrderDetails struct {
	Order   models.Order `json:"order"`
	NodeIds []string     `json:"nodeIds"`
}

func Create(ctx *gin.Context) {
	var orderDetails OrderDetails
	ctx.BindJSON(&orderDetails)
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
	models.DB.Create(&orderDetails.Order)
	ctx.JSON(http.StatusOK, orderDetails)
}

func All(ctx *gin.Context) {
	var nodes []models.Order
	models.DB.Find(&nodes)
	ctx.JSON(http.StatusOK, gin.H{"data": nodes})
}

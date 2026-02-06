package order

import (
	mqttstate "TobiasFP/BotNana/controllers/mqtt"
	"TobiasFP/BotNana/models"
	"net/http"
	"time"

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
		return
	}
	orderDetails.Order.OrderID = uuid.New().String()

	var nodes []models.Node
	var node models.Node
	for _, id := range orderDetails.NodeIds {

		result := models.SqlDB.Where("node_id = ?", id).First(&node)
		if result.RowsAffected == 1 && result.Error == nil {
			nodes = append(nodes, node)
		}
	}
	orderDetails.Order.Nodes = nodes
	models.SqlDB.Create(&orderDetails)
	ctx.JSON(http.StatusOK, orderDetails)
}

// @Summary Get all orders
// @Schemes
// @Accept json
// @Produce json
// @Success 200 {slice} []models.Order data "ok"
// @Router /orders/all [get]
func All(ctx *gin.Context) {
	var orderTemplate []models.OrderTemplateDetails
	models.SqlDB.Preload("Order").Find(&orderTemplate)
	ctx.JSON(http.StatusOK, gin.H{"data": orderTemplate})
}

func AssignAnonymous(ctx *gin.Context) {
	var assignOrder models.AssignOrder
	err := ctx.ShouldBindJSON(&assignOrder)
	if err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}

	var orderDetails models.OrderTemplateDetails

	orderDetailResults := models.SqlDB.Model(&models.OrderTemplateDetails{}).Preload("Order").Preload("Order.Nodes").Preload("Order.Edges").Where("ID = ?", assignOrder.ID).First(&orderDetails)
	if orderDetailResults.RowsAffected == 0 {
		ctx.JSON(400, gin.H{"error": "No Order Template with the given ID found."})
		return
	}

	if orderDetailResults.Error != nil {
		ctx.JSON(400, gin.H{"error": orderDetailResults.Error.Error()})
		return
	}

	// We sadly need to create a new order, as we cannot simply do:
	// order := orderDetails.Order as this would copy info related to gorm as well.
	order := models.Order{
		HeaderID:      0,
		Version:       "2.1.0",
		Timestamp:     time.Now().Format(time.RFC3339),
		Manufacturer:  orderDetails.Order.Manufacturer,
		SerialNumber:  "",
		OrderID:       uuid.New().String(),
		OrderUpdateID: 0,
		Nodes:         orderDetails.Order.Nodes,
		Edges:         orderDetails.Order.Edges,
		ZoneSetID:     "",
	}

	orderCreateRes := models.SqlDB.Create(&order)

	if orderCreateRes.Error != nil {
		ctx.JSON(400, gin.H{"error": orderCreateRes.Error.Error()})
		return
	}
	if mqttErr := mqttstate.AssignOrder(mqttstate.Client, order); mqttErr != nil {
		ctx.JSON(400, gin.H{"error": mqttErr.Error()})
		return
	}
	ctx.JSON(http.StatusOK, assignOrder)
}

package reststates

import (
	"TobiasFP/BotNana/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func AllStates(ctx *gin.Context) {
	var states []models.State
	models.DB.Preload("BatteryState").Preload("Maps").Preload("SafetyState").Preload("AgvPosition").Find(&states)

	ctx.JSON(http.StatusOK, gin.H{"data": states})
}

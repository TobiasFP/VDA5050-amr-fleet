package reststates

import (
	"TobiasFP/BotNana/models"
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

func AllStatesOnlyPositionData(ctx *gin.Context) {
	var states []models.State
	res := models.DB.Preload("Maps").Preload("AgvPosition").Select("serial_number").Find(&states)
	if res.Error != nil {
		ctx.Error(errors.New("no AMRs foundr"))
	}
	ctx.JSON(http.StatusOK, gin.H{"data": states})
}
func AllStates(ctx *gin.Context) {
	var states []models.State
	res := models.DB.Preload("BatteryState").Preload("Maps").Preload("SafetyState").Preload("AgvPosition").Find(&states)
	if res.Error != nil {
		ctx.Error(errors.New("no AMRs foundr"))
	}
	ctx.JSON(http.StatusOK, gin.H{"data": states})
}

func State(ctx *gin.Context) {
	SN := ctx.Query("SN")
	if SN == "" {
		ctx.Error(errors.New("state did not match"))
		return
	}
	var state models.State
	res := models.DB.Where("serial_number = ?", SN).Preload("BatteryState").Preload("Maps").Preload("SafetyState").Preload("AgvPosition").First(&state)
	if res.Error != nil {
		ctx.Error(errors.New("no AMR found with given serial number"))
	}
	ctx.JSON(http.StatusOK, gin.H{"data": state})
}

package action

import (
	"TobiasFP/BotNana/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// @BasePath /api/v1

func Create(ctx *gin.Context) {
	var action models.Action
	ctx.BindJSON(&action)
	action.ActionID = uuid.New().String()

	models.DB.Create(&action)
	ctx.JSON(http.StatusOK, action)
}

// @Summary Get all actions
// @Schemes
// @Accept json
// @Produce json
// @Success 200 {slice} []models.Action data "ok"
// @Router /edge/all [get]
func All(ctx *gin.Context) {
	var actions []models.Action
	models.DB.Find(&actions)
	ctx.JSON(http.StatusOK, gin.H{"data": actions})
}

// @Summary Get all actionParameters
// @Schemes
// @Accept json
// @Produce json
// @Success 200 {slice} []models.ActionParameter data "ok"
// @Router /edge/all [get]
func AllActionParams(ctx *gin.Context) {
	var actionParams []models.ActionParameter
	models.DB.Find(&actionParams)
	ctx.JSON(http.StatusOK, gin.H{"data": actionParams})
}

func CreateActionParameters(ctx *gin.Context) {
	var actionParam models.ActionParameter
	ctx.BindJSON(&actionParam)

	models.DB.Create(&actionParam)
	ctx.JSON(http.StatusOK, actionParam)
}

// @Summary Get all instantActions
// @Schemes
// @Accept json
// @Produce json
// @Success 200 {slice} []models.InstantAction data "ok"
// @Router /edge/all [get]
func AllInstantActions(ctx *gin.Context) {
	var instantActions []models.InstantAction
	models.DB.Find(&instantActions)
	ctx.JSON(http.StatusOK, gin.H{"data": instantActions})
}

func CreateInstantAction(ctx *gin.Context) {
	var instantAction models.InstantAction
	ctx.BindJSON(&instantAction)

	models.DB.Create(&instantAction)
	ctx.JSON(http.StatusOK, instantAction)
}

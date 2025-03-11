package action

import (
	"TobiasFP/BotNana/models"
	"context"
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// @BasePath /api/v1

func Create(ctx *gin.Context) {
	var action models.Action
	ctx.BindJSON(&action)
	action.ActionID = uuid.New().String()

	models.SqlDB.Create(&action)
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
	models.SqlDB.Find(&actions)
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

	allActionParametersRes, err := models.NoSqlDB.Search().Index("actionparameters").Do(context.Background())
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	for _, hit := range allActionParametersRes.Hits.Hits {
		var actionParam models.ActionParameter
		json.Unmarshal(hit.Source_, &actionParam)
		actionParams = append(actionParams, actionParam)
	}

	ctx.JSON(http.StatusOK, gin.H{"data": actionParams})
}

func CreateActionParameters(ctx *gin.Context) {
	var actionParam models.ActionParameter
	ctx.BindJSON(&actionParam)

	_, err := models.NoSqlDB.Index("actionparameters").Request(actionParam).Do(context.Background())

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

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
	models.SqlDB.Find(&instantActions)
	ctx.JSON(http.StatusOK, gin.H{"data": instantActions})
}

func CreateInstantAction(ctx *gin.Context) {
	var instantAction models.InstantAction
	ctx.BindJSON(&instantAction)

	models.SqlDB.Create(&instantAction)
	ctx.JSON(http.StatusOK, instantAction)
}

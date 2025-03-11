package restmap

import (
	"TobiasFP/BotNana/models"
	"encoding/base64"
	"io"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// @Summary Get all maps
// @Schemes
// @Accept json
// @Produce json
// @Success 200 {slice} []models.AmrMap data "ok"
// @Router /maps/all [get]
func AllMaps(ctx *gin.Context) {
	var maps []models.AmrMap
	models.SqlDB.Find(&maps)
	ctx.JSON(http.StatusOK, gin.H{"data": maps})
}

// @Summary Get the pgm map as b64
// @Schemes
// @Accept json
// @Produce json
// @Success 200 {string} data
// @Router /maps/:mapID [get]

// Deferring to use an actual NoSql DB as it just adds complexity.
// In the future, if we are bottlenecked or if this is problematic with scalability,
// we should simply instantiate a mongodb or an elasticsearch instance.
func Map(ctx *gin.Context) {
	mapID := ctx.Param("mapID")

	var amrMap models.AmrMap

	mapResult := models.SqlDB.Preload("MapData").Where("map_id = ?", mapID).First(&amrMap)
	if mapResult.Error != nil {
		ctx.Error(mapResult.Error)
	}
	mapDataAsB64 := base64.StdEncoding.EncodeToString(amrMap.MapData.Data)
	ctx.JSON(http.StatusOK, gin.H{"data": mapDataAsB64})
}

func Create(ctx *gin.Context) {
	var amrMap models.AmrMap
	// We are simply receiving an AMR Map.
	// This could be optimised by sending binary data from the frontend
	// instead of a b64 encoded string of the map data, but currently,
	// this is overkill.
	mapFile, mapsErr := ctx.FormFile("map")
	if mapsErr != nil {
		ctx.JSON(400, gin.H{"error": mapsErr.Error()})
		return
	}

	mapDescription, _ := ctx.GetPostForm("mapDescription")

	amrMap.MapDescription = mapDescription
	fileData, _ := mapFile.Open()

	byteContainer, err := io.ReadAll(fileData)
	if err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}
	amrMap.MapID = uuid.New()
	amrMap.MapData.Data = byteContainer
	createErr := models.SqlDB.Create(&amrMap)
	if createErr.Error != nil {
		ctx.JSON(400, gin.H{"error": createErr.Error.Error()})
		return
	}
	ctx.JSON(http.StatusOK, amrMap)
}

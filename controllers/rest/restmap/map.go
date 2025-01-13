package restmap

import (
	"TobiasFP/BotNana/models"
	"bufio"
	"encoding/base64"
	"io"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

// @Summary Get all maps
// @Schemes
// @Accept json
// @Produce json
// @Success 200 {slice} []models.AmrMap data "ok"
// @Router /maps/all [get]
func AllMaps(ctx *gin.Context) {
	var maps []models.AmrMap
	models.DB.Find(&maps)
	ctx.JSON(http.StatusOK, gin.H{"data": maps})
}

// @Summary Get the pgm map as b64
// @Schemes
// @Accept json
// @Produce json
// @Success 200 {string} data
// @Router /maps/map [get]
func Map(ctx *gin.Context) {
	mapPgmFile, err := os.Open("assets/maps/99187cd1-8b4b-4f5a-ac11-e455928409de.pgm")
	if err != nil {
		panic(err)
	}
	defer mapPgmFile.Close()
	PgmFileReader := bufio.NewReader(mapPgmFile)
	Pgmcontent, _ := io.ReadAll(PgmFileReader)
	PgmBase64 := base64.StdEncoding.EncodeToString(Pgmcontent)
	ctx.JSON(http.StatusOK, gin.H{"data": PgmBase64})
}

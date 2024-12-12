package map

import (
	"TobiasFP/BotNana/models"
	"bufio"
	"encoding/base64"
	"io"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func AllMaps(ctx *gin.Context) {
	var maps []models.AmrMap
	models.DB.Find(&maps)
	ctx.JSON(http.StatusOK, gin.H{"data": maps})
}

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


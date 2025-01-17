package restroutes

import (
	"TobiasFP/BotNana/config"
	"TobiasFP/BotNana/controllers/auth"
	"TobiasFP/BotNana/controllers/rest/edge"
	"TobiasFP/BotNana/controllers/rest/node"
	"TobiasFP/BotNana/controllers/rest/order"
	"TobiasFP/BotNana/controllers/rest/restmap"
	"TobiasFP/BotNana/controllers/rest/reststate"
	"TobiasFP/BotNana/models" // swagger embed files
	"log"
	"net/http"
	"os"
	"time"

	docs "TobiasFP/BotNana/docs"

	swaggerfiles "github.com/swaggo/files"

	ginSwagger "github.com/swaggo/gin-swagger" // gin-swagger middleware

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	ginglog "github.com/szuecs/gin-glog"
	"github.com/tbaehler/gin-keycloak/pkg/ginkeycloak"
	"golang.org/x/oauth2"
)

var (
	clientID        = ""
	clientIDDev     = "botnana"
	clientSecret    = ""
	clientSecretDev = "sJGdGuHDvP44SpwKqlVbcqNOu6u17V4K"
)

// StartGin function
func StartGin() {
	models.ConnectDatabase()

	conf := config.GetConfig()
	production := conf.GetBool("production")
	if !production {
		clientID = clientIDDev
		clientSecret = clientSecretDev
	}

	router := gin.Default()
	docs.SwaggerInfo.BasePath = "/api/v1"
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "OPTIONS", "DELETE"},
		AllowHeaders:     []string{"Origin", "content-type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		AllowOriginFunc: func(origin string) bool {
			if origin == "http://localhost:8100" {
				return true
			}
			// Allow your frontend through CORS
			return origin == conf.GetString("appUrl")
		},
		MaxAge: 12 * time.Hour,
	}))

	router.Use(ginglog.Logger(3 * time.Second))
	router.Use(gin.Recovery())
	authController := auth.Auth{
		Config: oauth2.Config{
			RedirectURL:  conf.GetString("apiUrl") + "/auth/callback",
			ClientID:     clientID,
			ClientSecret: clientSecret,
			Scopes:       []string{"profile", "email", "roles"},
			Endpoint: oauth2.Endpoint{
				AuthURL:  conf.GetString("OAuthUrl") + "/auth",
				TokenURL: conf.GetString("OAuthUrl") + "/token",
			},
		},
	}

	router.GET("/", authController.Login)
	router.GET("/auth/callback", authController.Callback)

	var keycloakconfig = ginkeycloak.KeycloakConfig{
		Url:   conf.GetString("keycloakUrl"),
		Realm: "botnana",
	}
	api := router.Group("/api")
	v1 := api.Group("/v1")
	v1.Use(ginkeycloak.Auth(ginkeycloak.AuthCheck(), keycloakconfig))

	amrGroup := v1.Group("/amrs")
	amrGroup.GET("/all", reststate.AllStates)
	amrGroup.GET("/positiondata", reststate.AllStatesOnlyPositionData)
	amrGroup.GET("/info", reststate.State)

	mapsGroup := v1.Group("/maps")
	mapsGroup.GET("/all", restmap.AllMaps)
	mapsGroup.GET("/:mapID", restmap.Map)
	mapsGroup.POST("/", restmap.Create)

	edgeGroup := v1.Group("/edges")
	edgeGroup.GET("/all", edge.All)
	edgeGroup.POST("/", edge.Create)

	nodeGroup := v1.Group("/nodes")
	nodeGroup.GET("/all", node.All)
	nodeGroup.POST("/", node.Create)

	orderGroup := v1.Group("/orders")
	orderGroup.GET("/all", order.All)
	orderGroup.POST("/", order.Create)
	orderGroup.POST("/assign", order.AssignAnonymous)

	helloWorldGroup := v1.Group("/helloworld")
	helloWorldGroup.GET("/", helloworld)

	router.NoRoute(func(ctx *gin.Context) {
		ctx.AbortWithStatus(http.StatusNotFound)
	})

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	port := os.Getenv("PORT")
	if port == "" {
		port = conf.GetString("apiPort")
	}
	err := router.Run(":" + port)
	if err != nil {
		log.Println(err.Error())
		return
	}
}

func helloworld(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{"data": "hello world"})
}

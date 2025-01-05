package restroutes

import (
	"TobiasFP/BotNana/controllers/auth"
	"TobiasFP/BotNana/controllers/rest/edge"
	"TobiasFP/BotNana/controllers/rest/node"
	"TobiasFP/BotNana/controllers/rest/order"
	"TobiasFP/BotNana/controllers/rest/restmap"
	"TobiasFP/BotNana/controllers/rest/reststate"
	"TobiasFP/BotNana/models"
	"log"
	"net/http"
	"os"
	"time"

	"TobiasFP/BotNana/config"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	ginglog "github.com/szuecs/gin-glog"
	"github.com/tbaehler/gin-keycloak/pkg/ginkeycloak"
	"golang.org/x/oauth2"
)

var (
	clientID        = "6f5944858fca4f20b1799a40647ff8c8"
	clientIDDev     = "botnana"
	clientSecret    = "3ab14fa856b24ff38b915a5ba2235a9b"
	clientSecretDev = "2DAb8HZ13h2zhg206vjKZVDmjnDIZQrb"
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
	api.Use(ginkeycloak.Auth(ginkeycloak.AuthCheck(), keycloakconfig))

	amrGroup := api.Group("/amrs")
	amrGroup.GET("/all", reststate.AllStates)
	amrGroup.GET("/positiondata", reststate.AllStatesOnlyPositionData)
	amrGroup.GET("/info", reststate.State)

	mapsGroup := api.Group("/maps")
	mapsGroup.GET("/all", restmap.AllMaps)
	mapsGroup.GET("/map", restmap.Map)

	edgeGroup := api.Group("/edges")
	edgeGroup.GET("/all", edge.All)
	edgeGroup.POST("/", edge.Create)

	nodeGroup := api.Group("/nodes")
	nodeGroup.GET("/all", node.All)
	nodeGroup.POST("/", node.Create)

	orderGroup := api.Group("/orders")
	orderGroup.GET("/all", order.All)
	orderGroup.POST("/", order.Create)

	helloWorldGroup := api.Group("/helloworld")
	helloWorldGroup.GET("/", helloworld)

	router.NoRoute(func(ctx *gin.Context) {
		ctx.AbortWithStatus(http.StatusNotFound)
	})

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

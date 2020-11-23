package main

import (
	"os"
	"strconv"
	"time"

	// innkeeper_handler "github.com/kubeinn/schutterij/internal/api/innkeeper"
	auth_handler "github.com/kubeinn/schutterij/internal/api/auth"
	pilgrim_handler "github.com/kubeinn/schutterij/internal/api/pilgrim"
	dbcontroller "github.com/kubeinn/schutterij/internal/controllers/dbcontroller"
	kubecontroller "github.com/kubeinn/schutterij/internal/controllers/kubecontroller"
	global "github.com/kubeinn/schutterij/internal/global"
	middleware "github.com/kubeinn/schutterij/internal/middleware"
	test "github.com/kubeinn/schutterij/test"
	go_cache "github.com/patrickmn/go-cache"

	cors "github.com/gin-contrib/cors"
	gin "github.com/gin-gonic/gin"
)

func main() {
	// Testing
	// test.TestInitEnvironmentVars()

	// Initialize variables
	initialize()

	// Testing
	test.TestCreateDefaultInnkeeper()
	// test.TestCreateDefaultReeve()

	// Start web server
	// Set the router as the default one shipped with Gin
	router := gin.Default()
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "OPTIONS"},
		AllowHeaders:     []string{"*"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	}))

	// // Setup route group for the innkeeper API endpoint
	// innkeeperAPI := router.Group(global.INNKEEPER_ROUTE_PREFIX)
	// innkeeperAPI.Use(middleware.TokenAuthMiddleware())
	// {
	// 	innkeeperAPI.GET("/test", innkeeper_handler.GetTestValidation)
	// }

	// Setup route group for the pilgrim API endpoint
	pilgrimAPI := router.Group(global.PILGRIM_ROUTE_PREFIX)
	pilgrimAPI.Use(middleware.TokenAuthMiddleware())
	{
		pilgrimAPI.POST("/create-project", pilgrim_handler.PostCreateProject)
		pilgrimAPI.POST("/delete-project", pilgrim_handler.PostDeleteProject)
	}

	// Setup route group for the authentication API endpoint
	authAPI := router.Group(global.AUTHENTICATION_ROUTE_PREFIX)
	{
		authAPI.POST("/login", auth_handler.PostValidateCredentialsHandler)
		authAPI.POST("/register/pilgrim", auth_handler.PostRegisterPilgrimHandler)
		authAPI.POST("/register/village", auth_handler.PostRegisterVillageHandler)
		authAPI.POST("/validate-regcode", auth_handler.PostValidateRegcodeHandler)
		authAPI.POST("/check-auth", auth_handler.PostCheckAuthHandler)
	}

	// Start and run the server
	router.Run(":8080")
}

func initialize() {
	// Instantiate global variables
	// Create a cache with a default expiration time of 5 minutes, and which
	// purges expired items every 10 minutes
	global.SESSION_CACHE = go_cache.New(15*time.Minute, 5*time.Minute)
	// Import signing key
	global.JWT_SIGNING_KEY = []byte(os.Getenv("JWT_SIGNING_KEY"))
	dbName := os.Getenv("PGDATABASE")
	dbHost := os.Getenv("PGHOST")
	dbPort, _ := strconv.Atoi(os.Getenv("PGPORT"))
	dbUser := os.Getenv("PGUSER")
	dbPassword := os.Getenv("POSTGRES_PASSWORD")
	global.PG_CONTROLLER = *dbcontroller.NewPostgresController(dbName, dbHost, dbPort, dbUser, dbPassword)

	// Create KUBE_CONTROLLER
	global.KUBE_CONTROLLER = *kubecontroller.NewKubeController(global.KUBE_CONFIG_ABSOLUTE_PATH)
}

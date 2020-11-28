package main

import (
	"os"
	"strconv"
	"time"

	auth_handler "github.com/kubeinn/src/backend/internal/api/auth"
	postgrest_handler "github.com/kubeinn/src/backend/internal/api/postgrest"

	dbcontroller "github.com/kubeinn/src/backend/internal/controllers/dbcontroller"
	kubecontroller "github.com/kubeinn/src/backend/internal/controllers/kubecontroller"
	global "github.com/kubeinn/src/backend/internal/global"
	middleware "github.com/kubeinn/src/backend/internal/middleware"
	test "github.com/kubeinn/src/backend/test"

	cors "github.com/gin-contrib/cors"
	gin "github.com/gin-gonic/gin"
	go_cache "github.com/patrickmn/go-cache"
)

func main() {
	// Testing (comment for production)
	test.TestInitEnvironmentVars()

	// Initialize variables
	initialize()

	// Testing
	test.TestCreateDefaultInnkeeper()

	// Start web server
	// Set the router as the default one shipped with Gin
	router := gin.Default()
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "OPTIONS", "PUT", "DELETE"},
		AllowHeaders:     []string{"*"},
		ExposeHeaders:    []string{"Content-Length", "Content-Range"},
		AllowCredentials: true,
	}))

	// Setup route group for the authentication API endpoint
	authAPI := router.Group(global.AUTHENTICATION_ROUTE_PREFIX)
	{
		authAPI.POST("/login", auth_handler.PostValidateCredentialsHandler)
		authAPI.POST("/register/pilgrim", auth_handler.PostRegisterPilgrimHandler)
		authAPI.POST("/check-auth", auth_handler.PostCheckAuthHandler)
	}

	// Setup route group for innkeeperAPI endpoint
	innkeeperAPI := router.Group(global.INNKEEPER_ROUTE_PREFIX + global.POSTGREST_ROUTE_PREFIX)
	innkeeperAPI.Use(middleware.TokenAuthMiddleware())
	{
		innkeeperAPI.Any("/innkeepers", postgrest_handler.ReverseProxy)
		innkeeperAPI.Any("/pilgrims", postgrest_handler.ReverseProxy)
		innkeeperAPI.Any("/projects", postgrest_handler.ReverseProxy)
		innkeeperAPI.Any("/tickets", postgrest_handler.ReverseProxy)
	}

	// Setup route group for pilgrimAPI endpoint
	pilgrimAPI := router.Group(global.PILGRIM_ROUTE_PREFIX + global.POSTGREST_ROUTE_PREFIX)
	pilgrimAPI.Use(middleware.TokenAuthMiddleware())
	{
		pilgrimAPI.Any("/innkeepers", postgrest_handler.ReverseProxy)
		pilgrimAPI.Any("/pilgrims", postgrest_handler.ReverseProxy)
		pilgrimAPI.Any("/projects", postgrest_handler.ReverseProxy)
		pilgrimAPI.Any("/tickets", postgrest_handler.ReverseProxy)
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

	// Import postgrest url
	global.POSTGREST_URL = os.Getenv("POSTGREST_URL")

	// Create PG_CONTROLLER
	dbName := os.Getenv("PGDATABASE")
	dbHost := os.Getenv("PGHOST")
	dbPort, _ := strconv.Atoi(os.Getenv("PGPORT"))
	dbUser := os.Getenv("PGUSER")
	dbPassword := os.Getenv("POSTGRES_PASSWORD")
	global.PG_CONTROLLER = *dbcontroller.NewPostgresController(dbName, dbHost, dbPort, dbUser, dbPassword)

	// Create KUBE_CONTROLLER
	global.KUBE_CONTROLLER = *kubecontroller.NewKubeController(global.KUBE_CONFIG_ABSOLUTE_PATH)
}

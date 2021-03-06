package main

import (
	"fmt"
	"os"
	"strconv"
	"time"

	auth_handler "github.com/kubeinn/kubeinn/src/internal/api/auth"
	postgrest_handler "github.com/kubeinn/kubeinn/src/internal/api/rproxy/postgrest_handler"
	prometheus_handler "github.com/kubeinn/kubeinn/src/internal/api/rproxy/prometheus_handler"

	dbcontroller "github.com/kubeinn/kubeinn/src/internal/controllers/dbcontroller"
	kubecontroller "github.com/kubeinn/kubeinn/src/internal/controllers/kubecontroller"
	global "github.com/kubeinn/kubeinn/src/internal/global"
	middleware "github.com/kubeinn/kubeinn/src/internal/middleware"

	cors "github.com/gin-contrib/cors"
	gin_static "github.com/gin-gonic/contrib/static"
	gin "github.com/gin-gonic/gin"
	go_cache "github.com/patrickmn/go-cache"
)

func main() {
	// Initialize variables
	initialize()

	// Create default admin account
	auth_handler.RegisterInnkeeper("admin", "admin", "admin")

	// Set the router as the default one shipped with Gin
	router := gin.Default()
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "OPTIONS", "PUT", "DELETE"},
		AllowHeaders:     []string{"*"},
		ExposeHeaders:    []string{"Content-Length", "Content-Range"},
		AllowCredentials: true,
	}))

	// Serve frontend static files
	router.Use(gin_static.Serve("/", gin_static.LocalFile("./client/crossroads/build", true)))
	router.Use(gin_static.Serve("/innkeeper", gin_static.LocalFile("./client/innkeeper/build", true)))
	router.Use(gin_static.Serve("/pilgrim", gin_static.LocalFile("./client/pilgrim/build", true)))

	// Serve 404 page when page not found
	router.NoRoute(func(c *gin.Context) {
		c.JSON(404, gin.H{"code": "PAGE_NOT_FOUND", "message": "Page not found"})
	})

	// Setup route group for the authentication API endpoint
	authAPI := router.Group(global.API_ROUTE_PREFIX + global.AUTHENTICATION_ROUTE_PREFIX)
	{
		authAPI.POST("/login", auth_handler.PostValidateCredentialsHandler)
		authAPI.POST("/register/pilgrim", auth_handler.PostRegisterPilgrimHandler)
		authAPI.POST("/check-auth", auth_handler.PostCheckAuthHandler)
	}

	// Setup route group for innkeeperAPI endpoint
	innkeeperAPI := router.Group(global.API_ROUTE_PREFIX + global.INNKEEPER_ROUTE_PREFIX)
	innkeeperAPI.Use(middleware.TokenAuthMiddleware())
	{
		innkeeperAPI.Any("/innkeepers", postgrest_handler.PostgrestHandler)
		innkeeperAPI.Any("/pilgrims", postgrest_handler.PostgrestHandler)
		innkeeperAPI.Any("/projects", postgrest_handler.PostgrestHandler)
		innkeeperAPI.Any("/tickets", postgrest_handler.PostgrestHandler)
		innkeeperAPI.Any("/pods", prometheus_handler.PrometheusHandler)
	}

	// Setup route group for pilgrimAPI endpoint
	pilgrimAPI := router.Group(global.API_ROUTE_PREFIX + global.PILGRIM_ROUTE_PREFIX)
	pilgrimAPI.Use(middleware.TokenAuthMiddleware())
	{
		pilgrimAPI.Any("/projects", postgrest_handler.PostgrestHandler)
		pilgrimAPI.Any("/tickets", postgrest_handler.PostgrestHandler)
		pilgrimAPI.Any("/pods", prometheus_handler.PrometheusHandler)
	}

	// Start and run the server
	router.Run(":8080")
}

// initialize instantiates global variables
func initialize() {
	// Wait 3 minutes for database to start
	fmt.Println("Waiting 3 minutes for database to initialize...")
	time.Sleep(3 * time.Minute)
	fmt.Println("Wait completed.")

	// Cache with a default expiration time of 5 minutes, and which purges expired items every 10 minutes
	global.SESSION_CACHE = go_cache.New(15*time.Minute, 5*time.Minute)

	// Import signing key
	global.JWT_SIGNING_KEY = []byte(os.Getenv("JWT_SIGNING_KEY"))

	// Import postgrest url
	postgrestUrl := os.Getenv("PGTURL")
	postgrestPort := os.Getenv("PGTPORT")
	global.POSTGREST_URL = postgrestUrl + ":" + postgrestPort

	// Import Prometheus url
	prometheusUrl := os.Getenv("PROMETHEUS_URL")
	prometheusPort := os.Getenv("PROMETHEUS_PORT")
	global.PROMETHEUS_URL = prometheusUrl + ":" + prometheusPort

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

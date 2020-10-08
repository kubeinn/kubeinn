package main

import (
	"log"
	"os"
	"path/filepath"
	"strconv"
	"time"

	innkeeper_handler "github.com/kubeinn/schutterij/internal/api/innkeeper"
	// pilgrim_handler "github.com/kubeinn/schutterij/internal/api/pilgrim"
	auth_handler "github.com/kubeinn/schutterij/internal/api/auth"
	db_controller "github.com/kubeinn/schutterij/internal/controllers/dbcontroller"
	global "github.com/kubeinn/schutterij/internal/global"
	middleware "github.com/kubeinn/schutterij/internal/middleware"
	test "github.com/kubeinn/schutterij/test"
	go_cache "github.com/patrickmn/go-cache"

	cors "github.com/gin-contrib/cors"
	gin "github.com/gin-gonic/gin"
	urfavecli "github.com/urfave/cli/v2"
	// clientcmd "k8s.io/client-go/tools/clientcmd"
	homedir "k8s.io/client-go/util/homedir"
)

func main() {
	// Testing
	test.TestInitEnvironmentVars()

	// Initialize variables
	initialize()

	// Testing
	test.TestDatabaseConnection()

	var kubecfg string

	// Run the application
	app := &urfavecli.App{
		Name:  "Schutterij",
		Usage: "API endpoint for KubeInn multi-tenancy manager for Kubernetes.",
		Flags: []urfavecli.Flag{
			&urfavecli.StringFlag{
				Name:        "kubecfg",
				Value:       filepath.Join(homedir.HomeDir(), ".kube", "config"),
				Usage:       "Specify the filepath of kubeconfig.",
				Destination: &kubecfg,
				Required:    true,
			},
		},
		Action: func(c *urfavecli.Context) error {
			// for {
			// 	fmt.Println("Waiting for kubeconfig to be uploaded...")
			// 	if _, err := os.Stat(c.String("kubecfg")); !os.IsNotExist(err) {
			// 		break
			// 	}
			// 	time.Sleep(5 * time.Second)
			// }

			// // Read in kube config
			// var err error
			// global.KUBE_CONFIG, err = clientcmd.BuildConfigFromFlags("", c.String("kubecfg"))
			// if err != nil {
			// 	panic(err)
			// }

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

			// Setup route group for the innkeeper API endpoint
			innkeeperAPI := router.Group(global.INNKEEPER_ROUTE_PREFIX)
			innkeeperAPI.Use(middleware.TokenAuthMiddleware())
			{
				innkeeperAPI.GET("/test", innkeeper_handler.GetTestValidation)
			}

			// // Setup route group for the pilgrim API endpoint
			// innkeeperAPI := router.Group(global.PILGRIM_ROUTE_PREFIX)
			// innkeeperAPI.Use(middleware.TokenAuthMiddleware())
			// {
			// 	innkeeperAPI.POST("/", innkeeper_handler.PostExtraResourcesHandler)
			// }

			// Setup route group for the authentication API endpoint
			authAPI := router.Group(global.AUTHENTICATION_ROUTE_PREFIX)
			{
				authAPI.POST("/login", auth_handler.PostValidateCredentialsHandler)
				authAPI.POST("/register", auth_handler.PostRegisterPilgrim)
				authAPI.POST("/check-auth", auth_handler.PostCheckAuthHandler)
			}

			// Start and run the server
			router.Run(":8080")
			return nil
		},
	}
	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
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
	global.PG_CONTROLLER = *db_controller.NewPostgresController(dbName, dbHost, dbPort, dbUser, dbPassword)
}

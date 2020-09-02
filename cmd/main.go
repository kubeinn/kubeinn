package main

import (
	// "context"
	// "net/http"
	"fmt"
	"log"
	"math/rand"
	"os"
	"path/filepath"
	"strconv"
	"time"

	// innkeeper_handler "github.com/kubeinn/schutterij/internal/api/innkeeper"
	// pilgrim_handler "github.com/kubeinn/schutterij/internal/api/pilgrim"
	db_controller "github.com/kubeinn/schutterij/internal/controllers/DBController"
	global "github.com/kubeinn/schutterij/internal/global"
	// middleware "github.com/kubeinn/schutterij/internal/middleware"

	cors "github.com/gin-contrib/cors"
	gin "github.com/gin-gonic/gin"
	urfavecli "github.com/urfave/cli/v2"
	clientcmd "k8s.io/client-go/tools/clientcmd"
	homedir "k8s.io/client-go/util/homedir"
)

func main() {
	initialize()

	var kubecfg string

	// Run the application
	app := &urfavecli.App{
		Name:  "Schutterij",
		Usage: "API endpoint for kubeinn multi-tenancy manager for Kubernetes.",
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
			for {
				fmt.Println("Waiting for kubeconfig to be uploaded...")
				if _, err := os.Stat(c.String("kubecfg")); !os.IsNotExist(err) {
					break
				}
				time.Sleep(5 * time.Second)
			}

			// Read in kube config
			var err error
			global.KUBE_CONFIG, err = clientcmd.BuildConfigFromFlags("", c.String("kubecfg"))
			if err != nil {
				panic(err)
			}

			// Start web server
			// Set the router as the default one shipped with Gin
			router := gin.Default()
			router.Use(cors.Default())

			// Setup route group for the innkeeper API endpoint
			// innkeeperAPI := router.Group(global.INNKEEPER_API_ENDPOINT_PREFIX)
			// innkeeperAPI.Use(middleware.TokenAuthMiddleware())
			// {
			// 	// innkeeperAPI.POST("/resources/extra", innkeeper_handler.PostExtraResourcesHandler)
			// }

			// Setup route group for the pilgrim API endpoint
			// pilgrimAPI := router.Group(global.PILGRIM_API_ENDPOINT_PREFIX)

			// Setup route group for the authentication API endpoint
			// authAPI := router.Group(global.AUTH_API_ENDPOINT_PREFIX)
			// {
			// 	// authAPI.POST()
			// }

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
	global.JWT_SIGNING_KEY = make([]byte, 32)
	rand.Seed(time.Now().UnixNano())
	rand.Read(global.JWT_SIGNING_KEY)
	fmt.Print("global.JWT_SIGNING_KEY: ", string(global.JWT_SIGNING_KEY))

	// Create Postgres Controller
	dbName := os.Getenv("PGDATABASE")
	dbHost := os.Getenv("PGHOST")
	dbPort, _ := strconv.Atoi(os.Getenv("PGPORT"))
	dbUser := os.Getenv("PGUSER")
	dbPassword := os.Getenv("PGDATABASE")
	global.PG_CONTROLLER = *db_controller.NewPostgresController(dbName, dbHost, dbPort, dbUser, dbPassword)
}

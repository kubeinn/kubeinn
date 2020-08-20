package main

import (
	// "context"
	"fmt"
	"log"
	// "net/http"
	"os"
	"path/filepath"
	"time"

	global "github.com/kubeinn/pilgrim-go/internal/global"
	handlers "github.com/kubeinn/pilgrim-go/internal/handlers"
	postgres "github.com/kubeinn/pilgrim-go/internal/postgres"

	cors "github.com/gin-contrib/cors"
	static "github.com/gin-gonic/contrib/static"
	gin "github.com/gin-gonic/gin"
	urfavecli "github.com/urfave/cli/v2"
	clientcmd "k8s.io/client-go/tools/clientcmd"
	homedir "k8s.io/client-go/util/homedir"
)

func main() {
	var port string
	var kubecfg string

	// Instantiate global variables
	global.PostgresController = pg.NewPostgresController()

	app := &urfavecli.App{
		Name:  "Administration Portal",
		Usage: "Starts the Asi@Connect Administration Portal.",
		Flags: []urfavecli.Flag{
			&urfavecli.StringFlag{
				Name:        "port",
				Value:       "30000",
				Usage:       "Specify the container port to listen on.",
				Destination: &port,
				Required:    true,
			},
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
				fmt.Print("Waiting for kubeconfig to be uploaded...\n")
				if _, err := os.Stat(c.String("kubecfg")); !os.IsNotExist(err) {
					break
				}
				time.Sleep(5 * time.Second)
			}

			// Read in kube config
			var err error
			global.Kubeconfig, err = clientcmd.BuildConfigFromFlags("", c.String("kubecfg"))
			if err != nil {
				panic(err)
			}

			// Start web server
			// Set the router as the default one shipped with Gin
			router := gin.Default()
			router.Use(cors.Default())

			// Serve frontend static files
			router.Use(static.Serve("/", static.LocalFile("./web/build", true)))

			// Setup route group for the API
			api := router.Group("/api")
			{
				api.GET("/users", handlers.GetUsersHandler)
				api.GET("/resources/summary", handlers.GetResourceSummaryHandler)
				api.POST("/resources/register", handlers.PostRegisterUserHandler)
				api.POST("/resources/extra", handlers.PostExtraResourcesHandler)
			}
			// Start and run the server
			router.Run(":5000")
			return nil
		},
	}
	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}

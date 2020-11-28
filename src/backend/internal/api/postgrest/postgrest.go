package postgrest

import (
	"bytes"

	gin "github.com/gin-gonic/gin"
	hooks "github.com/kubeinn/src/backend/internal/api/hooks"
	global "github.com/kubeinn/src/backend/internal/global"
	"io/ioutil"
	"log"
	http "net/http"
	"net/http/httputil"
	"strings"
)

func ReverseProxy(c *gin.Context) {
	// Parse context request
	method := c.Request.Method
	path := strings.TrimPrefix(c.Request.URL.Path, "/innkeeper/postgrest")
	path = strings.TrimPrefix(path, "/pilgrim/postgrest")
	url := global.POSTGREST_URL + path
	body, err := ioutil.ReadAll(c.Request.Body)
	c.Request.Body.Close()
	c.Request.Body = ioutil.NopCloser(bytes.NewBuffer(body))

	// log.Println("===============================")
	// log.Println("method: " + method)
	// log.Println("path: " + path)
	// log.Println("url: " + url)
	// log.Println("===============================")

	// Create proxy request
	log.Println("Creating proxy request...")
	proxyReq, err := http.NewRequest(method, url, nil)
	if err != nil {
		log.Println(err)
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "failed to create new request"})
		return
	}

	// Set proxy request query params
	log.Println("Set proxy request query params...")
	proxyReq.URL.RawQuery = c.Request.URL.RawQuery

	// Set proxy request headers
	log.Println("Set proxy request headers")
	proxyReq.Header.Add("Authorization", c.Request.Header.Get("Authorization"))

	// PreHooks
	switch method {
	case "POST":
		if path == "/projects" {
			body, err = hooks.PreCreateProjectHook(c)
			if err != nil {
				log.Println(err)
				c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "Error creating project"})
				return
			}
		}
		if path == "/innkeepers" {
			body, err = hooks.PreCreateInnkeeperHook(c)
			if err != nil {
				log.Println(err)
				c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "Error creating innkeeper"})
				return
			}
		}
		if path == "/pilgrims" {
			body, err = hooks.PreCreatePilgrimHook(c)
			if err != nil {
				log.Println(err)
				c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "Error creating pilgrim"})
				return
			}
		}
	case "PUT":
		if path == "/innkeepers" {
			body, err = hooks.PreEditInnkeeperHook(c)
			if err != nil {
				log.Println(err)
				c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "Error editing innkeeper"})
				return
			}
		}
		if path == "/pilgrims" {
			body, err = hooks.PreEditPilgrimHook(c)
			if err != nil {
				log.Println(err)
				c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "Error editing pilgrim"})
				return
			}
		}
	case "DELETE":
		if path == "/projects" {
			err := hooks.PreDeleteProjectHook(c)
			if err != nil {
				c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "Error deleting project"})
				return
			}
		}
	case "GET":
		proxyReq.Header.Add("Prefer", c.Request.Header.Get("Prefer"))
	default:

	}

	// Create proxy request body
	proxyReq.Body = ioutil.NopCloser(bytes.NewReader(body))

	log.Println("Dumping proxy request...")
	data, err := httputil.DumpRequest(proxyReq, true)
	if err != nil {
		log.Fatal("Error")
	}
	log.Println(string(data))

	log.Println("Sending proxy request...")
	// Send proxy request
	client := &http.Client{}
	proxyRes, err := client.Do(proxyReq)
	if err != nil {
		log.Println(err)
		return
	}
	defer proxyRes.Body.Close()

	// PostHooks
	// switch method {
	// case "POST":
	//
	// 	}
	// default:

	// }

	// log.Println("Dumping proxy response...")
	// data, err = httputil.DumpResponse(proxyRes, true)
	// if err != nil {
	// 	log.Fatal("Error")
	// }
	// log.Println(string(data))

	body, err = ioutil.ReadAll(proxyRes.Body)
	//  Write response headers
	for header, values := range proxyRes.Header {
		for _, value := range values {
			log.Println(header + ": " + value)
			c.Writer.Header().Set(header, value)
		}
	}
	c.Data(proxyRes.StatusCode, "application/json", body)
}

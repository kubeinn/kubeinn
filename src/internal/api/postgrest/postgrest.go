package postgrest

import (
	"bytes"

	gin "github.com/gin-gonic/gin"
	hooks "github.com/kubeinn/kubeinn/src/internal/api/hooks"
	global "github.com/kubeinn/kubeinn/src/internal/global"
	"io/ioutil"
	"log"
	http "net/http"
	"net/http/httputil"
	"strings"
)

func ReverseProxy(c *gin.Context) {
	// Parse context request
	var audience string
	var path string
	var subject string
	method := c.Request.Method
	// Identify source
	if strings.HasPrefix(c.Request.URL.Path, global.API_ROUTE_PREFIX+global.INNKEEPER_ROUTE_PREFIX) {
		audience = global.JWT_AUDIENCE_INNKEEPER
		path = strings.TrimPrefix(c.Request.URL.Path, global.API_ROUTE_PREFIX+global.INNKEEPER_ROUTE_PREFIX+global.POSTGREST_ROUTE_PREFIX)
	} else if strings.HasPrefix(c.Request.URL.Path, global.API_ROUTE_PREFIX+global.PILGRIM_ROUTE_PREFIX) {
		audience = global.JWT_AUDIENCE_PILGRIM
		path = strings.TrimPrefix(c.Request.URL.Path, global.API_ROUTE_PREFIX+global.PILGRIM_ROUTE_PREFIX+global.POSTGREST_ROUTE_PREFIX)
	}
	subject = c.Request.Header.Get("subject")
	url := "http://" + global.POSTGREST_URL + path
	body, err := ioutil.ReadAll(c.Request.Body)
	c.Request.Body.Close()
	c.Request.Body = ioutil.NopCloser(bytes.NewBuffer(body))

	log.Println("===============================")
	log.Println("method: " + method)
	log.Println("path: " + path)
	log.Println("url: " + url)
	log.Println("===============================")

	// Create proxy request
	log.Println("Creating proxy request...")
	proxyReq, err := http.NewRequest(method, url, nil)
	if err != nil {
		log.Println(err)
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "failed to create new request"})
		return
	}

	// Set proxy request headers
	log.Println("Set proxy request headers")
	proxyReq.Header.Add("Authorization", c.Request.Header.Get("Authorization"))

	// Set proxy request query params
	log.Println("Set proxy request query params...")
	proxyReq.URL.RawQuery = c.Request.URL.RawQuery

	// PreHooks
	switch method {
	case "POST":
		if path == "/projects" {
			body, err = hooks.PreCreateProjectHook(c, audience)
			if err != nil {
				log.Println(err)
				c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "Error creating project"})
				return
			}
			proxyReq.Body = ioutil.NopCloser(bytes.NewReader(body))
		}
		if path == "/innkeepers" {
			body, err = hooks.PreCreateInnkeeperHook(c)
			if err != nil {
				log.Println(err)
				c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "Error creating innkeeper"})
				return
			}
			proxyReq.URL.Path = "/rpc/create_innkeeper"
			proxyReq.URL.RawQuery = ""
			proxyReq.Method = "POST"
			proxyReq.Body = ioutil.NopCloser(bytes.NewReader(body))
		}
		if path == "/pilgrims" {
			body, err = hooks.PreCreatePilgrimHook(c)
			if err != nil {
				log.Println(err)
				c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "Error creating pilgrim"})
				return
			}
			proxyReq.URL.Path = "/rpc/create_pilgrim"
			proxyReq.URL.RawQuery = ""
			proxyReq.Method = "POST"
		}
	case "PUT":
		if path == "/innkeepers" {
			body, err = hooks.PreEditInnkeeperHook(c)
			if err != nil {
				log.Println(err)
				c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "Error editing innkeeper"})
				return
			}
			proxyReq.URL.Path = "/rpc/update_innkeeper"
			proxyReq.URL.RawQuery = ""
			proxyReq.Method = "POST"
		}
		if path == "/pilgrims" {
			body, err = hooks.PreEditPilgrimHook(c)
			if err != nil {
				log.Println(err)
				c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "Error editing pilgrim"})
				return
			}
			proxyReq.URL.Path = "/rpc/update_pilgrim"
			proxyReq.URL.RawQuery = ""
			proxyReq.Method = "POST"
		}
	case "DELETE":
		if path == "/projects" {
			err := hooks.PreDeleteProjectHook(c, audience, subject)
			if err != nil {
				c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "Error deleting project"})
				return
			}
		}
	case "GET":
		proxyReq.Header.Add("Prefer", "count=exact")
		proxyReq.Header.Add("Content-Type", "application/json")
		if path == "/innkeepers" {
			proxyReq.URL.RawQuery = c.Request.URL.RawQuery + "&select=id,username,email"
		}
		if path == "/pilgrims" {
			proxyReq.URL.RawQuery = c.Request.URL.RawQuery + "&select=id,organization,description,username,email,status"
		}
	default:

	}

	// Create proxy request body
	log.Println("Creating proxy request...")
	proxyReq.Body = ioutil.NopCloser(bytes.NewReader(body))

	// Dump proxy request
	log.Println("Dumping proxy request...")
	data, err := httputil.DumpRequest(proxyReq, true)
	if err != nil {
		log.Fatal("Error")
	}
	log.Println(string(data))

	// Send proxy request
	log.Println("Sending proxy request...")
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

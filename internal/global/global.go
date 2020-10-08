package global

import (
	dbcontroller "github.com/kubeinn/schutterij/internal/controllers/dbcontroller"
	go_cache "github.com/patrickmn/go-cache"
	rest "k8s.io/client-go/rest"
)

var JWT_SIGNING_KEY []byte
var PG_CONTROLLER dbcontroller.PostgresController
var KUBE_CONFIG *rest.Config
var SESSION_CACHE *go_cache.Cache

const JWT_AUDIENCE_INNKEEPER string = "innkeeper"
const JWT_AUDIENCE_PILGRIM string = "pilgrim"

const AUTHENTICATION_ROUTE_PREFIX string = "/auth"
const INNKEEPER_ROUTE_PREFIX string = "/innkeeper"
const PILGRIM_ROUTE_PREFIX string = "/pilgrim"

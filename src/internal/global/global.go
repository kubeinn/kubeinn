package global

import (
	dbcontroller "github.com/kubeinn/src/backend/internal/controllers/dbcontroller"
	kubecontroller "github.com/kubeinn/src/backend/internal/controllers/kubecontroller"
	go_cache "github.com/patrickmn/go-cache"
)

var JWT_SIGNING_KEY []byte
var POSTGREST_URL string

var PG_CONTROLLER dbcontroller.PostgresController
var KUBE_CONTROLLER kubecontroller.KubeController
var SESSION_CACHE *go_cache.Cache

const KUBE_CONFIG_ABSOLUTE_PATH string = "/root/.kube/admin-config"

const JWT_AUDIENCE_INNKEEPER string = "innkeeper"
const JWT_AUDIENCE_PILGRIM string = "pilgrim"

const API_ROUTE_PREFIX string = "/api"
const AUTHENTICATION_ROUTE_PREFIX string = "/auth"
const POSTGREST_ROUTE_PREFIX string = "/postgrest"
const INNKEEPER_ROUTE_PREFIX string = "/innkeeper"
const PILGRIM_ROUTE_PREFIX string = "/pilgrim"

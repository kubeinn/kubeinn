package global

import (
	dbcontroller "github.com/kubeinn/schutterij/internal/controllers/dbcontroller"
	kubecontroller "github.com/kubeinn/schutterij/internal/controllers/kubecontroller"
	go_cache "github.com/patrickmn/go-cache"
)

var JWT_SIGNING_KEY []byte
var PG_CONTROLLER dbcontroller.PostgresController
var KUBE_CONTROLLER kubecontroller.KubeController
var SESSION_CACHE *go_cache.Cache

const KUBE_CONFIG_ABSOLUTE_PATH string = "/root/.kube/admin-config"

const JWT_AUDIENCE_INNKEEPER string = "Innkeeper"
const JWT_AUDIENCE_PILGRIM string = "Pilgrim"
const JWT_AUDIENCE_REEVE string = "Reeve"

const AUTHENTICATION_ROUTE_PREFIX string = "/auth"
const INNKEEPER_ROUTE_PREFIX string = "/innkeeper"
const PILGRIM_ROUTE_PREFIX string = "/pilgrim"
const REEVE_ROUTE_PREFIX string = "/reeve"

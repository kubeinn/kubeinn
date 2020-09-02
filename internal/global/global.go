package global

import (
	dbcontroller "github.com/kubeinn/schutterij/internal/controllers/DBController"
	rest "k8s.io/client-go/rest"
)

var JWT_SIGNING_KEY []byte

var PG_CONTROLLER dbcontroller.PostgresController

var KUBE_CONFIG *rest.Config

const JWT_SUBJECT_INNKEEPER string = "innkeeper"
const JWT_SUBJECT_PILGRIM string = "pilgrim"
const JWT_SUBJECT_AUTH string = "auth"

const INNKEEPER_API_ENDPOINT_PREFIX string = "/api/innkeeper"
const PILGRIM_API_ENDPOINT_PREFIX string = "/api/pilgrim"
const AUTH_API_ENDPOINT_PREFIX string = "/api/auth"

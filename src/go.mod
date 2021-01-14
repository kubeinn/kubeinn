module github.com/kubeinn/kubeinn/src

go 1.14

require (
	github.com/dgrijalva/jwt-go v3.2.0+incompatible
	github.com/gin-contrib/cors v1.3.1
	github.com/gin-gonic/contrib v0.0.0-20201101042839-6a891bf89f19
	github.com/gin-gonic/gin v1.6.3
	github.com/imdario/mergo v0.3.11 // indirect
	github.com/jackc/pgx/v4 v4.10.1
	github.com/patrickmn/go-cache v2.1.0+incompatible
	golang.org/x/crypto v0.0.0-20201221181555-eec23a3978ad
	golang.org/x/time v0.0.0-20201208040808-7e3f01d25324 // indirect
	k8s.io/api v0.19.0
	k8s.io/apimachinery v0.19.0
	k8s.io/client-go v0.19.0
	k8s.io/utils v0.0.0-20210111153108-fddb29f9d009 // indirect
)

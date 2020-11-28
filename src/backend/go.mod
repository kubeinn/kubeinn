module github.com/kubeinn/src/backend

go 1.14

require (
	github.com/dgrijalva/jwt-go v3.2.0+incompatible
	github.com/gin-contrib/cors v1.3.1
	github.com/gin-gonic/gin v1.6.3
	github.com/imdario/mergo v0.3.11 // indirect
	github.com/jackc/pgx/v4 v4.9.2
	github.com/patrickmn/go-cache v2.1.0+incompatible
	golang.org/x/crypto v0.0.0-20201124201722-c8d3bf9c5392
	golang.org/x/time v0.0.0-20200630173020-3af7569d3a1e // indirect
	k8s.io/api v0.19.4
	k8s.io/apimachinery v0.19.4
	k8s.io/client-go v0.19.0
	k8s.io/klog v1.0.0 // indirect
	k8s.io/utils v0.0.0-20201110183641-67b214c5f920 // indirect
)

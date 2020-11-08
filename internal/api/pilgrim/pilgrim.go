package pilgrim

import (
	gin "github.com/gin-gonic/gin"
	"net/http"
	"strconv"

	global "github.com/kubeinn/schutterij/internal/global"
)

// PostCreateProject is ...
func PostCreateProject(c *gin.Context) {

	namespace := c.Query("namespace")
	cpu, _ := strconv.ParseInt(c.Query("cpu"), 10, 64)
	memory, _ := strconv.ParseInt(c.Query("cpu"), 10, 64)
	storage, _ := strconv.ParseInt(c.Query("cpu"), 10, 64)

	err := global.KUBE_CONTROLLER.CreateNamespace(namespace)
	if err != nil {
		// Registration failed
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "Error creating namespace"})
		return
	}

	err = global.KUBE_CONTROLLER.CreateResourceQuota(namespace, cpu, memory, storage)
	if err != nil {
		// Registration failed
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "Error creating resource quota"})
		return
	}

	err = global.KUBE_CONTROLLER.CreateServiceAccount(namespace)
	if err != nil {
		// Registration failed
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "Error creating service account"})
		return
	}

	err = global.KUBE_CONTROLLER.CreateRole(namespace)
	if err != nil {
		// Registration failed
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "Error creating role"})
		return
	}

	err = global.KUBE_CONTROLLER.CreateRoleBinding(namespace)
	if err != nil {
		// Registration failed
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "Error creating role binding"})
		return
	}

	c.String(http.StatusOK, "Project created!")
}

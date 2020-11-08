package kubecontroller

import (
	"context"
	// "flag"
	// "fmt"
	// "path/filepath"
	// "time"
	"os"

	// "k8s.io/apimachinery/pkg/api/errors"
	resource "k8s.io/apimachinery/pkg/api/resource"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	// "k8s.io/client-go/tools/clientcmd"
	// "k8s.io/client-go/util/homedir"
	corev1 "k8s.io/api/core/v1"
	rbac "k8s.io/api/rbac/v1"
	kubernetes "k8s.io/client-go/kubernetes"
	rest "k8s.io/client-go/rest"
)

type KubeController struct {
	clientset *kubernetes.Clientset
}

func NewKubeController(config *rest.Config) *KubeController {
	kc := KubeController{}
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		os.Exit(1)
	}
	kc.clientset = clientset
	return &kc
}

func (kc *KubeController) CreateNamespace(namespace string) error {
	nsSpec := &corev1.Namespace{ObjectMeta: metav1.ObjectMeta{Name: namespace}}
	_, err := kc.clientset.CoreV1().Namespaces().Create(context.TODO(), nsSpec, metav1.CreateOptions{})
	return err
}

func (kc *KubeController) CreateResourceQuota(namespace string, cpu int64, memory int64, storage int64) error {
	resourceList := make(corev1.ResourceList)
	resourceList[corev1.ResourceCPU] = *resource.NewQuantity(cpu, resource.DecimalSI)
	resourceList[corev1.ResourceMemory] = *resource.NewQuantity(memory, resource.DecimalSI)
	resourceList[corev1.ResourceStorage] = *resource.NewQuantity(storage, resource.DecimalSI)

	rqSpec := &corev1.ResourceQuota{
		ObjectMeta: metav1.ObjectMeta{
			Name:      "compute-resources",
			Namespace: namespace,
		},
		Spec: corev1.ResourceQuotaSpec{
			Hard: resourceList,
		},
	}
	_, err := kc.clientset.CoreV1().ResourceQuotas(namespace).Create(context.TODO(), rqSpec, metav1.CreateOptions{})
	return err
}

func (kc *KubeController) CreateServiceAccount(namespace string) error {

	saSpec := &corev1.ServiceAccount{
		ObjectMeta: metav1.ObjectMeta{
			Name:      namespace,
			Namespace: namespace,
		},
	}
	_, err := kc.clientset.CoreV1().ServiceAccounts(namespace).Create(context.TODO(), saSpec, metav1.CreateOptions{})
	return err
}

func (kc *KubeController) CreateRole(namespace string) error {

	rules := []rbac.PolicyRule{
		rbac.PolicyRule{
			Verbs:     []string{"*"},
			Resources: []string{"*"},
			APIGroups: []string{"", "batch", "extensions", "apps"},
		},
	}

	rSpec := &rbac.Role{
		ObjectMeta: metav1.ObjectMeta{
			Name:      namespace,
			Namespace: namespace,
		},
		Rules: rules,
	}

	_, err := kc.clientset.RbacV1().Roles(namespace).Create(context.TODO(), rSpec, metav1.CreateOptions{})
	return err
}

func (kc *KubeController) CreateRoleBinding(namespace string) error {

	subjects := []rbac.Subject{
		rbac.Subject{
			Kind:     "ServiceAccount",
			Name:     namespace,
			APIGroup: "",
		},
	}

	roleref := rbac.RoleRef{
		Kind:     "Role",
		Name:     namespace,
		APIGroup: "rbac.authorization.k8s.io",
	}

	rbSpec := &rbac.RoleBinding{
		ObjectMeta: metav1.ObjectMeta{
			Name:      namespace,
			Namespace: namespace,
		},
		Subjects: subjects,
		RoleRef:  roleref,
	}

	_, err := kc.clientset.RbacV1().RoleBindings(namespace).Create(context.TODO(), rbSpec, metav1.CreateOptions{})
	return err
}

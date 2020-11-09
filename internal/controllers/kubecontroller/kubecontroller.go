package kubecontroller

import (
	"context"
	"fmt"
	"os"
	"time"

	corev1 "k8s.io/api/core/v1"
	rbac "k8s.io/api/rbac/v1"
	resource "k8s.io/apimachinery/pkg/api/resource"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	kubernetes "k8s.io/client-go/kubernetes"
	rest "k8s.io/client-go/rest"
	clientcmd "k8s.io/client-go/tools/clientcmd"
	clientcmdapi "k8s.io/client-go/tools/clientcmd/api"
)

type KubeController struct {
	clientset  *kubernetes.Clientset
	restconfig *rest.Config
}

func NewKubeController(path string) *KubeController {
	kc := KubeController{}
	fmt.Println("Loading kubeconfig from: " + path)
	config, err := clientcmd.BuildConfigFromFlags("", path)
	if err != nil {
		fmt.Println("Error loading kubeconfig.")
		fmt.Println(err)
		os.Exit(1)
	}
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		os.Exit(1)
	}
	kc.clientset = clientset
	kc.restconfig = config
	return &kc
}

func (kc *KubeController) CreateNamespace(namespace string) error {
	fmt.Println("Creating namespace...")

	nsSpec := &corev1.Namespace{ObjectMeta: metav1.ObjectMeta{Name: namespace}}
	_, err := kc.clientset.CoreV1().Namespaces().Create(context.TODO(), nsSpec, metav1.CreateOptions{})
	if err != nil {
		fmt.Println(err)
		return err
	}
	for i := 0; i < 10; i++ {
		_, err := kc.clientset.CoreV1().Namespaces().Get(context.TODO(), namespace, metav1.GetOptions{})
		if err != nil {
			time.Sleep(5 * time.Second)
		} else {
			return nil
		}
	}
	return err
}

func (kc *KubeController) CreateResourceQuota(namespace string, cpu int64, memory int64, storage int64) error {
	fmt.Println("Creating resource quota...")

	resourceList := make(map[corev1.ResourceName]resource.Quantity)
	resourceList[corev1.ResourceLimitsCPU] = *resource.NewQuantity(cpu, resource.DecimalSI)
	resourceList[corev1.ResourceLimitsMemory] = *resource.NewQuantity(memory, resource.DecimalSI)
	resourceList[corev1.ResourceRequestsStorage] = *resource.NewQuantity(storage, resource.DecimalSI)

	rqSpec := &corev1.ResourceQuota{
		ObjectMeta: metav1.ObjectMeta{
			Name:      namespace,
			Namespace: namespace,
		},
		Spec: corev1.ResourceQuotaSpec{
			Hard: resourceList,
		},
	}
	_, err := kc.clientset.CoreV1().ResourceQuotas(namespace).Create(context.TODO(), rqSpec, metav1.CreateOptions{})

	if err != nil {
		fmt.Println(err)
		return err
	}
	for i := 0; i < 10; i++ {
		_, err := kc.clientset.CoreV1().ResourceQuotas(namespace).Get(context.TODO(), namespace, metav1.GetOptions{})
		if err != nil {
			time.Sleep(5 * time.Second)
		} else {
			return nil
		}
	}
	return err
}

func (kc *KubeController) CreateServiceAccount(namespace string) error {
	fmt.Println("Creating service account...")

	saSpec := &corev1.ServiceAccount{
		ObjectMeta: metav1.ObjectMeta{
			Name:      namespace,
			Namespace: namespace,
		},
	}
	_, err := kc.clientset.CoreV1().ServiceAccounts(namespace).Create(context.TODO(), saSpec, metav1.CreateOptions{})

	if err != nil {
		fmt.Println(err)
		return err
	}
	for i := 0; i < 10; i++ {
		_, err := kc.clientset.CoreV1().ServiceAccounts(namespace).Get(context.TODO(), namespace, metav1.GetOptions{})
		if err != nil {
			time.Sleep(5 * time.Second)
		} else {
			return nil
		}
	}
	return err
}

func (kc *KubeController) CreateRole(namespace string) error {
	fmt.Println("Creating role...")

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

	if err != nil {
		fmt.Println(err)
		return err
	}
	for i := 0; i < 10; i++ {
		_, err := kc.clientset.RbacV1().Roles(namespace).Get(context.TODO(), namespace, metav1.GetOptions{})
		if err != nil {
			time.Sleep(5 * time.Second)
		} else {
			return nil
		}
	}
	return err
}

func (kc *KubeController) CreateRoleBinding(namespace string) error {
	fmt.Println("Creating role binding...")

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

	if err != nil {
		fmt.Println(err)
		return err
	}
	for i := 0; i < 10; i++ {
		_, err := kc.clientset.RbacV1().RoleBindings(namespace).Get(context.TODO(), namespace, metav1.GetOptions{})
		if err != nil {
			time.Sleep(5 * time.Second)
		} else {
			return nil
		}
	}
	return err
}

func (kc *KubeController) GenerateKubeConfiguration(namespace string) error {
	fmt.Println("Generating kube config...")

	// Get secret list for namespace
	secretList, err := kc.clientset.CoreV1().Secrets(namespace).List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		return err
	}

	// Get secret containing authentication data
	secret := secretList.Items[0]

	clusters := make(map[string]*clientcmdapi.Cluster)
	clusters["default-cluster"] = &clientcmdapi.Cluster{
		Server:                   kc.restconfig.Host,
		CertificateAuthorityData: secret.Data["ca.crt"],
	}

	contexts := make(map[string]*clientcmdapi.Context)
	contexts["default-context"] = &clientcmdapi.Context{
		Cluster:   "default-cluster",
		Namespace: namespace,
		AuthInfo:  namespace,
	}

	authinfos := make(map[string]*clientcmdapi.AuthInfo)
	authinfos[namespace] = &clientcmdapi.AuthInfo{
		Token: string(secret.Data["token"]),
	}

	clientConfig := clientcmdapi.Config{
		Kind:           "Config",
		APIVersion:     "v1",
		Clusters:       clusters,
		Contexts:       contexts,
		CurrentContext: "default-context",
		AuthInfos:      authinfos,
	}

	clientcmd.WriteToFile(clientConfig, namespace+"-config")

	return nil
}

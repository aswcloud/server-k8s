package k8s

import (
	"fmt"
	"path/filepath"

	// appsv1 "k8s.io/api/apps/v1"
	// corev1 "k8s.io/api/core/v1"
	// metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/util/homedir"
)

type Client struct {
	app        *kubernetes.Clientset
	kubeconfig string
}

func New() *Client {
	kubeconfig := filepath.Join(homedir.HomeDir(), ".kube", "config")
	// kubeconfig := flag.String("kubeconfig", filepath.Join(homedir.HomeDir(), ".kube", "config"), "(optional) absolute path to the kubeconfig file")
	// config, err := rest.InClusterConfig()
	// if err != nil {
	// 	fmt.Printf("unable to load in-cluster config: %v", err)
	// }

	// config, _ := clientcmd.BuildConfigFromFlags("kubernetes.default.svc", "")
	config, _ := clientcmd.BuildConfigFromFlags("", kubeconfig)
	clientset, err2 := kubernetes.NewForConfig(config)
	if err2 != nil {
		fmt.Printf("unable to create a client: %v", err2)
	}

	return &Client{
		app:        clientset,
		kubeconfig: "~!~!~!",
	}
}

func (self *Client) Namespace() *Namespace {
	// list, _ := .List(context.TODO(), metav1.ListOptions{})
	return &Namespace{
		app: self.app.CoreV1().Namespaces(),
	}
}

func (self *Client) Deployment(namespace string) *Deployment {
	return &Deployment{
		namespace: namespace,
		app:       self.app.AppsV1().Deployments(namespace),
	}
}

func (self *Client) Service(namespace string) *Service {
	// list, _ := self.app.CoreV1().Services(namespace).List(context.TODO(), metav1.ListOptions{})

	return &Service{
		app:       self.app.CoreV1().Services(namespace),
		namespace: namespace,
	}
}

func (self *Client) Storage(namespace string) *Storage {
	return &Storage{
		namespace: namespace,
		pvc:       self.app.CoreV1().PersistentVolumeClaims(namespace),
		pv:        self.app.CoreV1().PersistentVolumes(),
	}

}

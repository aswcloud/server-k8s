package k8s

import (
	"context"

	template "github.com/aswcloud/server-k8s/k8s/template"
	appsv1 "k8s.io/api/apps/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	k8sv1 "k8s.io/client-go/kubernetes/typed/apps/v1"
)

type Deployment struct {
	namespace string
	app       k8sv1.DeploymentInterface
}

func (self *Deployment) List() (*appsv1.DeploymentList, error) {
	return self.app.List(context.TODO(), metav1.ListOptions{})
}

func (self *Deployment) Create(data template.DeploymentTemplate) (*appsv1.Deployment, error) {
	deploy := template.ParseDeployment(data)
	return self.app.Create(context.TODO(), deploy, metav1.CreateOptions{})
}

func (self *Deployment) Remove(name string) error {
	return self.app.Delete(context.TODO(), name, metav1.DeleteOptions{})
}

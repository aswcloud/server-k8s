package k8s

import (
	"context"

	template "github.com/aswcloud/server-k8s/k8s/template"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	k8sns "k8s.io/client-go/kubernetes/typed/core/v1"
)

type Service struct {
	app       k8sns.ServiceInterface
	namespace string
}

func (self *Service) List() (*corev1.ServiceList, error) {
	return self.app.List(context.TODO(), metav1.ListOptions{})
}

func (self *Service) Create(data template.ServiceTemplate) (*corev1.Service, error) {
	service := template.ParseService(data)

	return self.app.Create(context.TODO(), service, metav1.CreateOptions{})
}

func (self *Service) Remove(name string) error {
	return self.app.Delete(context.TODO(), name, metav1.DeleteOptions{})
}

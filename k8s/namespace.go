package k8s

import (
	"context"

	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	k8sns "k8s.io/client-go/kubernetes/typed/core/v1"
)

type Namespace struct {
	app k8sns.NamespaceInterface
}

func (self *Namespace) List() (*corev1.NamespaceList, error) {
	return self.app.List(context.TODO(), metav1.ListOptions{})
}

func (self *Namespace) Create(name string) (*corev1.Namespace, error) {
	nsName := &corev1.Namespace{
		ObjectMeta: metav1.ObjectMeta{
			Name: name,
		},
	}

	return self.app.Create(context.TODO(), nsName, metav1.CreateOptions{})
}

func (self *Namespace) Remove(name string) error {
	return self.app.Delete(context.TODO(), name, metav1.DeleteOptions{})
}

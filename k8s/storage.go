package k8s

import (
	"context"

	template "github.com/aswcloud/server-k8s/k8s/template"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	k8sns "k8s.io/client-go/kubernetes/typed/core/v1"
)

type Storage struct {
	namespace string
	pvc       k8sns.PersistentVolumeClaimInterface
	pv        k8sns.PersistentVolumeInterface
}

func (self *Storage) PersistentVolume() *PersistentVolume {
	return &PersistentVolume{
		namespace: self.namespace,
		app:       self.pv,
	}
}

func (self *Storage) PersistentVolumeClaim() *PersistentVolumeClaim {
	return &PersistentVolumeClaim{
		namespace: self.namespace,
		app:       self.pvc,
	}
}

type PersistentVolume struct {
	namespace string
	app       k8sns.PersistentVolumeInterface
}

type PersistentVolumeClaim struct {
	namespace string
	app       k8sns.PersistentVolumeClaimInterface
}

func (self *PersistentVolume) List() (*corev1.PersistentVolumeList, error) {
	return self.app.List(context.TODO(), metav1.ListOptions{})
}

func (self *PersistentVolume) Create(data template.PersistentVolumeTemplate) (*corev1.PersistentVolume, error) {
	pv := template.ParsePersistentVolume(data)

	return self.app.Create(context.TODO(), pv, metav1.CreateOptions{})
}

func (self *PersistentVolume) Remove(name string) error {
	return self.app.Delete(context.TODO(), name, metav1.DeleteOptions{})
}

func (self *PersistentVolumeClaim) List() (*corev1.PersistentVolumeClaimList, error) {
	return self.app.List(context.TODO(), metav1.ListOptions{})
}

func (self *PersistentVolumeClaim) Create(data template.PersistentVolumeClaimTemplate) (*corev1.PersistentVolumeClaim, error) {
	pvc := template.ParsePersistentVolumeClaim(data)
	return self.app.Create(context.TODO(), pvc, metav1.CreateOptions{})
}

func (self *PersistentVolumeClaim) Remove(name string) error {
	return self.app.Delete(context.TODO(), name, metav1.DeleteOptions{})
}

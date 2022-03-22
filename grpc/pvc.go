package grpc

import (
	"context"
	"encoding/json"
	"log"

	pb "github.com/aswcloud/idl/v1/kubernetes"
	"github.com/aswcloud/server-k8s/k8s"
	"github.com/aswcloud/server-k8s/k8s/template"
	funk "github.com/thoas/go-funk"
	"google.golang.org/protobuf/encoding/protojson"
	v1 "k8s.io/api/core/v1"
)

func (self *KubernetesServer) CreatePersistentVolumeClaim(ctx context.Context, name *pb.Pvc) (*pb.Result, error) {
	k8s := k8s.New()

	log.Println(name.String())
	temp := template.PersistentVolumeClaimTemplate{}
	data, _ := protojson.Marshal(name)
	json.Unmarshal(data, &temp)

	_, err := k8s.Storage(name.Namespace).PersistentVolumeClaim().Create(temp)
	result := true
	errText := ""
	if err != nil {
		result = false
		errText = err.Error()
	}

	return &pb.Result{
		Result: result,
		Error:  &errText,
	}, nil
}

func (self *KubernetesServer) DeletePersistentVolumeClaim(ctx context.Context, name *pb.DeletePvc) (*pb.Result, error) {
	k8s := k8s.New()

	err := k8s.Storage(name.Namespace).PersistentVolumeClaim().Remove(name.Name)
	result := true
	errText := ""
	if err != nil {
		result = false
		errText = err.Error()
	}

	return &pb.Result{
		Result: result,
		Error:  &errText,
	}, nil
}

func (self *KubernetesServer) ListPersistentVolumeClaim(ctx context.Context, name *pb.Namespace) (*pb.ListPvc, error) {
	log.Println("receive")

	k8s := k8s.New()
	list, e := k8s.Storage(name.Namespace).PersistentVolumeClaim().List()
	if e != nil {
		log.Println(e)
		return nil, e
	}

	data := funk.Map(list.Items, func(item v1.PersistentVolumeClaim) *pb.Pvc {
		return &pb.Pvc{
			Namespace: item.ObjectMeta.Namespace,
			Name:      item.ObjectMeta.Name,
			Capacity:  item.Spec.Resources.Requests.Storage().String(),
			AccessMode: funk.Map(item.Spec.AccessModes, func(data v1.PersistentVolumeAccessMode) string {
				return string(data)
			}).([]string),
			StorageClassName: *item.Spec.StorageClassName,
		}
	}).([]*pb.Pvc)
	log.Println(list)
	return &pb.ListPvc{
		Name: data,
	}, nil
}

package grpc

import (
	"context"
	"encoding/json"
	"log"

	pb "github.com/aswcloud/idl/v1/kubernetes"
	"github.com/aswcloud/server-k8s/k8s"
	"github.com/aswcloud/server-k8s/k8s/template"
	"google.golang.org/protobuf/encoding/protojson"
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
	_, e := k8s.Storage(name.Namespace).PersistentVolumeClaim().List()
	if e != nil {
		log.Println(e)
		return nil, e
	}

	// data := funk.Map(list.Items, func(item v1.PersistentVolumeClaim) pb.Pvc {
	// 	// return &pb.Pvc{
	// 	// 	Namespace: item.Namespace,
	// 	// 	Name:      item.ObjectMeta.Name,
	// 	// 	// Capacity:  item.Spec.Capacity,
	// 	// }
	// })
	// log.Println(data)
	return &pb.ListPvc{
		Name: []*pb.Pvc{},
	}, nil
}

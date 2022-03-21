package grpc

import (
	"context"
	"encoding/json"
	"log"

	pb "github.com/aswcloud/idl/v1/kubernetes"
	"github.com/aswcloud/server-k8s/k8s"
	"github.com/aswcloud/server-k8s/k8s/template"
	"google.golang.org/protobuf/encoding/protojson"
	v1 "k8s.io/api/core/v1"

	funk "github.com/thoas/go-funk"
)

func (self *KubernetesServer) CreateService(ctx context.Context, name *pb.Service) (*pb.Service, error) {
	k8s := k8s.New()

	log.Println(name.String())
	temp := template.ServiceTemplate{}
	data, _ := protojson.Marshal(name)
	json.Unmarshal(data, &temp)

	svc, err := k8s.Service(name.Namespace).Create(temp)

	if err != nil {
		log.Println(err)
	}
	ports := funk.Map(svc.Spec.Ports, func(data v1.ServicePort) *pb.ServicePort {
		return &pb.ServicePort{
			Name:          data.Name,
			TargetPort:    data.TargetPort.IntVal,
			ContainerPort: data.Port,
			NodePort:      data.NodePort,
		}
	}).([]*pb.ServicePort)
	return &pb.Service{
		Namespace:    svc.Namespace,
		Name:         svc.Name,
		Type:         string(svc.Spec.Type),
		TemplateName: svc.Spec.Selector["app"],
		Ports:        ports,
	}, nil
}

func (self *KubernetesServer) DeleteService(ctx context.Context, name *pb.DeleteService) (*pb.Result, error) {
	k8s := k8s.New()

	err := k8s.Service(name.Namespace).Remove(name.Name)

	result := true
	errText := ""
	if err != nil {
		result = false
		errText = err.Error()
	}

	return &pb.Result{
		Result: result,
		Error:  &errText,
	}, err
}

func (self *KubernetesServer) ListService(ctx context.Context, name *pb.Namespace) (*pb.ListService, error) {
	k8s := k8s.New()

	list, err := k8s.Service(name.Namespace).List()
	if err != nil {
		return &pb.ListService{}, err
	}

	result := funk.Map(list.Items, func(data v1.Service) string {
		return data.Name
	}).([]string)

	return &pb.ListService{
		Name: result,
	}, nil
}

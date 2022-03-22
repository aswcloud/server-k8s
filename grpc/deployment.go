package grpc

import (
	"context"
	"encoding/json"
	"log"

	"github.com/aswcloud/server-k8s/k8s"
	"github.com/aswcloud/server-k8s/k8s/template"
	funk "github.com/thoas/go-funk"
	"google.golang.org/protobuf/encoding/protojson"
	v1 "k8s.io/api/apps/v1"

	pb "github.com/aswcloud/idl/v1/kubernetes"
)

func (self *KubernetesServer) CreateDeployment(ctx context.Context, name *pb.Deployment) (*pb.Result, error) {
	k8s := k8s.New()

	log.Println(name.String())
	temp := template.DeploymentTemplate{}
	data, _ := protojson.Marshal(name)
	json.Unmarshal(data, &temp)
	log.Println(temp)
	_, err := k8s.Deployment(name.Namespace).Create(temp)

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

func (self *KubernetesServer) DeleteDeployment(ctx context.Context, name *pb.DeleteDeployment) (*pb.Result, error) {
	k8s := k8s.New()

	err := k8s.Deployment(name.Namespace).Remove(name.Name)

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

func (self *KubernetesServer) ListDeployment(ctx context.Context, name *pb.Namespace) (*pb.ListDeployment, error) {
	k8s := k8s.New()

	list, err := k8s.Deployment(name.Namespace).List()
	if err != nil {
		return &pb.ListDeployment{}, err
	}

	result := funk.Map(list.Items, func(data v1.Deployment) string {
		return data.Name
	}).([]string)

	return &pb.ListDeployment{
		Name: result,
	}, nil
}

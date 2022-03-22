package grpc

import (
	"context"
	"fmt"
	"log"

	pb "github.com/aswcloud/idl/v1/kubernetes"
	"github.com/aswcloud/server-k8s/k8s"
	funk "github.com/thoas/go-funk"
	v1 "k8s.io/api/core/v1"
)

func (self *KubernetesServer) CreateNamespace(ctx context.Context, name *pb.Namespace) (*pb.Result, error) {
	k8s := k8s.New()
	_, e := k8s.Namespace().Create(name.Namespace)
	if e != nil {
		fmt.Println(e)
		return nil, e
	}
	return &pb.Result{
		Result: true,
		Error:  nil,
	}, nil
}

func (self *KubernetesServer) DeleteNamespace(ctx context.Context, name *pb.Namespace) (*pb.Result, error) {
	k8s := k8s.New()
	e := k8s.Namespace().Remove(name.Namespace)
	if e != nil {
		fmt.Println(e)
		return nil, e
	}
	return &pb.Result{
		Result: true,
		Error:  nil,
	}, nil
}

func (self *KubernetesServer) ListNamespace(ctx context.Context, name *pb.Empty) (*pb.ListNamespace, error) {
	log.Println("receive")

	k8s := k8s.New()
	list, e := k8s.Namespace().List()
	if e != nil {
		log.Println(e)
		return nil, e
	}

	data := funk.Map(list.Items, func(item v1.Namespace) string {
		return item.ObjectMeta.Name
	})
	log.Println(data)
	return &pb.ListNamespace{
		Name: data.([]string),
	}, nil
}

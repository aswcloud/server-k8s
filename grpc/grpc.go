package grpc

import (
	pb "github.com/aswcloud/idl/v1/kubernetes"
)

type KubernetesServer struct {
	pb.UnimplementedKubernetesServer
}

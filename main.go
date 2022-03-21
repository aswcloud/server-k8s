package main

import (
	"fmt"
	"log"
	"net"

	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	grpc_logrus "github.com/grpc-ecosystem/go-grpc-middleware/logging/logrus"
	grpc_recovery "github.com/grpc-ecosystem/go-grpc-middleware/recovery"
	"github.com/sirupsen/logrus"
	"github.com/subosito/gotenv"

	pb "github.com/aswcloud/idl/v1/kubernetes"
	og "github.com/aswcloud/server-k8s/grpc"
	"google.golang.org/grpc"
)

func main() {
	gotenv.Load()

	log.Println("loaded")
	lis, err := net.Listen("tcp", ":8000")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer(
		grpc.UnaryInterceptor(grpc_middleware.ChainUnaryServer(
			// grpc_prometheus.UnaryServerInterceptor,
			grpc_logrus.UnaryServerInterceptor(logrus.NewEntry(logrus.StandardLogger())),
			grpc_recovery.UnaryServerInterceptor(),
		)),
	)

	pb.RegisterKubernetesServer(s, &og.KubernetesServer{})

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}

	fmt.Println("ERROR!333")
}

// // 님은 방구 뿡뿡이~ 님은 방구 뿡뿡잉~~

package main

import (
	"context"
	"flag"
	"fmt"
	"path/filepath"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/util/homedir"
)

func main() {
	kubeconfig := flag.String("kubeconfig", filepath.Join(homedir.HomeDir(), ".kube", "config"), "(optional) absolute path to the kubeconfig file")

	fmt.Println(*kubeconfig)
	config, _ := clientcmd.BuildConfigFromFlags("", *kubeconfig)
	// clientset을 생성한다
	clientset, _ := kubernetes.NewForConfig(config)
	// 파드를 나열하기 위해 API에 접근한다
	pods, _ := clientset.CoreV1().Pods("default").List(context.TODO(), v1.ListOptions{})
	fmt.Printf("There are %d pods in the cluster\n", len(pods.Items))
}

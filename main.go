package main

import (
	"fmt"

	"github.com/aswcloud/server-k8s/k8s"
	"github.com/aswcloud/server-k8s/k8s/template"
)

type Template struct {
	Name           string
	TemplateName   string
	ReplicaCount   int
	ContainerImage string
	Ports          []int
}

func main() {
	k8s := k8s.New()
	list, err := k8s.Namespace().List()
	if err != nil {
		fmt.Println(err)
	}

	for idx, item := range list.Items {
		fmt.Println(idx, " : ", item.Name)
	}
	_, err = k8s.Namespace().Create("tttt")
	if err != nil {
		fmt.Println("List Error!")
	}

	k8s.Deployment("tttt").Create(template.DeploymentTemplate{
		Name:           "nginx-deployment",
		TemplateName:   "nginx-template",
		ReplicaCount:   8,
		ContainerImage: "nginx:latest",
		Ports:          []int{80},
	})
	k8s.Service("tttt").Create(template.ServiceTemplate{
		Name:         "nginx-service",
		Type:         "NodePort",
		TemplateName: "nginx-template",
		Ports: []template.ServicePortTemplate{
			{
				Name:          "http",
				TargetPort:    80,
				ContainerPort: 80,
				NodePort:      30001,
			},
		},
	})

	// k8s.Deployment("tttt").Remove("nginx-deployment")
	// k8s.Service("tttt").Remove("nginx-service")
	// k8s.Namespace().Remove("tttt")
	// // 님은 방구 뿡뿡이~ 님은 방구 뿡뿡잉~~
}

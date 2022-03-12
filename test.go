package main

import (
	"fmt"

	"github.com/aswcloud/server-k8s/k8s"
	"github.com/aswcloud/server-k8s/k8s/template"
)

func CreatePersistent() {
	k8s := k8s.New()

	// k8s.Storage("tttt").PersistentVolume().Create(template.PersistentVolumeTemplate{
	// 	Name:     "pv-uuid-test",
	// 	Capacity: "3Gi",
	// 	HostPath: "/mnt/test",
	// })
	k8s.Storage("tttt").PersistentVolumeClaim().Create(template.PersistentVolumeClaimTemplate{
		Name:     "pvc-tttt-nginx",
		Capacity: "3Gi",
	})
}

func CreateNamespace() {
	k8s := k8s.New()
	k8s.Namespace().Create("tttt")
}

func CreateService() {
	k8s := k8s.New()

	value, _ := k8s.Service("tttt").Create(template.ServiceTemplate{
		Name:         "nginx-service",
		Type:         "NodePort",
		TemplateName: "nginx-template",
		Ports: []template.ServicePortTemplate{
			{
				Name:          "http",
				TargetPort:    80,
				ContainerPort: 80,
			},
			{
				Name:          "sftp",
				TargetPort:    22,
				ContainerPort: 22,
				NodePort:      30002,
			},
		},
	})
	for idx, item := range value.Spec.Ports {
		fmt.Println(idx, " : ", item.Name, " / ", item.NodePort)
	}

}

func CreateDeployment() {
	k8s := k8s.New()

	k8s.Deployment("tttt").Create(template.DeploymentTemplate{
		Name:         "nginx-deployment",
		TemplateName: "nginx-template",
		ReplicaCount: 1,
		Volume: []template.DeploymentVolumeTemplate{
			{
				Name:      "pvc-volume",
				ClaimName: "pvc-tttt-nginx",
			},
		},
		Containers: []template.DeploymentContainerTemplate{
			{
				Name:  "nginx",
				Image: "aoikazto/php-apache:7.4",
				Ports: []int{80},
				Env:   []template.KeyValue{},
				VolumeMount: []template.DeploymentVolumeMountTemplate{
					{
						Name:      "pvc-volume",
						MountPath: "/var/www/html",
					},
				},
			},
			{
				Name:  "ftp",
				Image: "amimof/sftp",
				Ports: []int{22},
				Env: []template.KeyValue{
					{
						Key:   "SSH_USERNAME",
						Value: "\"sftpuser\"",
					}, {
						Key:   "SSH_PASSWORD",
						Value: "\"sftpuser\"",
					}, {
						Key:   "SSH_USERID",
						Value: "\"1000\"",
					},
				},
				VolumeMount: []template.DeploymentVolumeMountTemplate{
					{
						Name:      "pvc-volume",
						MountPath: "/home/sftpuser/data",
					},
				},
			},
		},
	})
}

func DeleteAll() {
	k8s := k8s.New()
	k8s.Service("tttt").Remove("nginx-service")
	k8s.Deployment("tttt").Remove("nginx-deployment")
	k8s.Storage("tttt").PersistentVolumeClaim().Remove("pvc-tttt-nginx")
	k8s.Namespace().Remove("tttt")
}

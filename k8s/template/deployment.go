package template

import (
	"bytes"
	"fmt"
	"os"
	"strings"

	temp "text/template"

	appsv1 "k8s.io/api/apps/v1"
	k8syaml "k8s.io/apimachinery/pkg/util/yaml"
)

func ParseDeployment(template DeploymentTemplate) *appsv1.Deployment {
	a := temp.New("0301929")
	data, _ := os.ReadFile("./template/deployment-template.yaml")
	a, err := a.Parse(string(data))
	if err != nil {
		fmt.Println("ERROR : ", err)
	}

	var sb strings.Builder
	a.Execute(&sb, template)
	fmt.Println(sb.String())

	deploy := &appsv1.Deployment{}
	dec := k8syaml.NewYAMLToJSONDecoder(bytes.NewReader([]byte(sb.String())))
	dec.Decode(deploy)
	return deploy
}

type DeploymentContainerTemplate struct {
	Name        string                          `json: "name"`
	Image       string                          `json: "image"`
	Ports       []int                           `json: "ports"`
	Env         []KeyValue                      `json: "env"`
	VolumeMount []DeploymentVolumeMountTemplate `json: "volumeMount"`
}

type KeyValue struct {
	Key   string `json: "key"`
	Value string `json: "value"`
}

type DeploymentVolumeTemplate struct {
	Name      string `json: "name"`
	ClaimName string `json: "claimName"`
}

type DeploymentVolumeMountTemplate struct {
	Name      string `json: "name"`
	MountPath string `json: "mountPath"`
}

type DeploymentTemplate struct {
	Name         string                        `json: "name"`
	TemplateName string                        `json: "templateName"`
	ReplicaCount int                           `json: "replicaCount"`
	Volume       []DeploymentVolumeTemplate    `json: "volume"`
	Containers   []DeploymentContainerTemplate `json: "containers"`
}

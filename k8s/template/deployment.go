package template

import (
	"bytes"
	"os"
	"strings"

	temp "text/template"

	appsv1 "k8s.io/api/apps/v1"
	k8syaml "k8s.io/apimachinery/pkg/util/yaml"
)

func ParseDeployment(template DeploymentTemplate) *appsv1.Deployment {
	a := temp.New("test")
	data, _ := os.ReadFile("./template/deployment-template.yaml")
	a, _ = a.Parse(string(data))

	var sb strings.Builder
	a.Execute(&sb, template)
	deploy := &appsv1.Deployment{}
	dec := k8syaml.NewYAMLToJSONDecoder(bytes.NewReader([]byte(sb.String())))
	dec.Decode(deploy)
	return deploy
}

type DeploymentTemplate struct {
	Name           string
	TemplateName   string
	ReplicaCount   int
	ContainerImage string
	Ports          []int
}

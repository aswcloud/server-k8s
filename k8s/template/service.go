package template

import (
	"bytes"
	"fmt"
	"os"
	"strings"

	temp "text/template"

	corev1 "k8s.io/api/core/v1"
	k8syaml "k8s.io/apimachinery/pkg/util/yaml"
	// applyv1 "k8s.io/client-go/applyconfigurations/core/v1"
)

func ParseService(template ServiceTemplate) *corev1.Service {
	a := temp.New("test")
	data, _ := os.ReadFile("./template/service-template.yaml")
	a, _ = a.Parse(string(data))

	var sb strings.Builder
	a.Execute(&sb, template)
	fmt.Println(sb.String())
	service := &corev1.Service{}
	dec := k8syaml.NewYAMLToJSONDecoder(bytes.NewReader([]byte(sb.String())))
	dec.Decode(service)
	return service
}

type ServicePortTemplate struct {
	Name          string `json: "name"`
	TargetPort    int32  `json: "targetPort"`
	NodePort      int32  `json: "nodePort"`
	ContainerPort int32  `json: "containerPort"`
}

type ServiceTemplate struct {
	Name         string                `json: "name"`
	Type         string                `json: "type"`
	TemplateName string                `json: "templateName"`
	Ports        []ServicePortTemplate `json: "ports"`
}

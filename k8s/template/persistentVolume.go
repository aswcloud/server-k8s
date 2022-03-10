package template

import (
	"bytes"
	"fmt"
	"os"
	"strings"

	temp "text/template"

	corev1 "k8s.io/api/core/v1"
	k8syaml "k8s.io/apimachinery/pkg/util/yaml"
)

func ParsePersistentVolume(template PersistentVolumeTemplate) *corev1.PersistentVolume {
	a := temp.New("test")
	data, _ := os.ReadFile("./template/persistent-volume.yaml")
	a, _ = a.Parse(string(data))

	var sb strings.Builder
	a.Execute(&sb, template)
	fmt.Println(sb.String())
	pv := &corev1.PersistentVolume{}
	dec := k8syaml.NewYAMLToJSONDecoder(bytes.NewReader([]byte(sb.String())))
	dec.Decode(pv)
	return pv
}

type PersistentVolumeTemplate struct {
	Name     string
	Capacity string
	HostPath string
}

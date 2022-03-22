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

func ParsePersistentVolumeClaim(template PersistentVolumeClaimTemplate) *corev1.PersistentVolumeClaim {
	a := temp.New("test")
	data, _ := os.ReadFile("./template/persistent-volume-claim.yaml")
	a, _ = a.Parse(string(data))

	var sb strings.Builder
	a.Execute(&sb, template)
	fmt.Println(sb.String())
	pvc := &corev1.PersistentVolumeClaim{}
	dec := k8syaml.NewYAMLToJSONDecoder(bytes.NewReader([]byte(sb.String())))
	dec.Decode(pvc)
	return pvc
}

type PersistentVolumeClaimTemplate struct {
	Name             string   `json: "name"`
	Capacity         string   `json: "capacity"`
	AccessMode       []string `json: "accessMode"`
	StorageClassName string   `json: "storageClassName"`
}

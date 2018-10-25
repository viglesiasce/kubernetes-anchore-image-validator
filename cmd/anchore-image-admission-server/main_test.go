package main

import (
	"encoding/json"
	"fmt"
	"github.com/golang/glog"
	"k8s.io/api/admission/v1beta1"
	"k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"testing"
)


func TestValidate(t *testing.T) {
	tpod := &v1.Pod{
		Spec: v1.PodSpec{Containers: []v1.Container{
			{
				Name:    "Container1",
				Image:   "ubuntu",
				Command: []string{"bin/bash", "bin"},
			},
		},
		},
	}

	marshalledPod, err := json.Marshal(tpod)

	if (err != nil ) {
		glog.Error("Failed marshalling pod spec")
		t.Error("Failed marshalling")
	}

	admSpec := v1beta1.AdmissionRequest {
		UID: "abc123",
		Kind: metav1.GroupVersionKind{Group: v1beta1.GroupName, Version: v1beta1.SchemeGroupVersion.Version, Kind: "Pod"},
		Resource: metav1.GroupVersionResource{Group: metav1.GroupName, Version: "v1", Resource: "pods"},
		SubResource: "someresource",
		Name: "somename",
		Namespace: "default",
		Operation: "CREATE",
		Object: runtime.RawExtension{Raw: marshalledPod},
	}

	adm := admissionHook{}
	resp := adm.Validate(&admSpec)
	obj, err := json.Marshal(resp)

	if(err == nil) {
		fmt.Println(string(obj[:]))
	}

}

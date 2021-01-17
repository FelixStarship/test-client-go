
package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"k8s.io/apimachinery/pkg/api/meta"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/serializer/yaml"
	yamlutil "k8s.io/apimachinery/pkg/util/yaml"
	"k8s.io/client-go/dynamic"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/restmapper"
	"k8s.io/client-go/tools/clientcmd"
	"log"
	"os"

	//
	// Uncomment to load all auth plugins
	// _ "k8s.io/client-go/plugin/pkg/client/auth"
	//
	// Or uncomment to load specific auth plugins
	// _ "k8s.io/client-go/plugin/pkg/client/auth/azure"
	// _ "k8s.io/client-go/plugin/pkg/client/auth/gcp"
	// _ "k8s.io/client-go/plugin/pkg/client/auth/oidc"
	// _ "k8s.io/client-go/plugin/pkg/client/auth/openstack"
)

func main() {

	nameSpace := "demo01"
	k8sConfig, err := clientcmd.BuildConfigFromFlags("", "config/config-bak.yaml")  // 使用 kubectl 默认配置 ~/.kube/config
	if err != nil {
		fmt.Printf("%v",err)
		return
	}

	// 创建一个k8s客户端
	clientSet, err := kubernetes.NewForConfig(k8sConfig)
	if err != nil {
		fmt.Printf("%v",err)
		return
	}
	dd, err := dynamic.NewForConfig(k8sConfig)
	if err != nil {
		log.Fatal(err)
	}

	filebytes, err := ioutil.ReadFile("config/kube_config_cluster.yml")
	if err != nil {
		fmt.Printf("%v\n", err)
	}

	decoder := yamlutil.NewYAMLOrJSONDecoder(bytes.NewReader(filebytes), 100)
	for {
		var rawObj runtime.RawExtension
		if err = decoder.Decode(&rawObj); err != nil {
			break
		}

		obj, gvk, err := yaml.NewDecodingSerializer(unstructured.UnstructuredJSONScheme).Decode(rawObj.Raw, nil, nil)
		unstructuredMap, err := runtime.DefaultUnstructuredConverter.ToUnstructured(obj)
		if err != nil {
			log.Fatal(err)
		}

		unstructuredObj := &unstructured.Unstructured{Object: unstructuredMap}

		gr, err := restmapper.GetAPIGroupResources(clientSet.Discovery())
		if err != nil {
			log.Fatal(err)
		}

		mapper := restmapper.NewDiscoveryRESTMapper(gr)
		mapping, err := mapper.RESTMapping(gvk.GroupKind(), gvk.Version)
		if err != nil {
			log.Fatal(err)
		}

		var dri dynamic.ResourceInterface
		if mapping.Scope.Name() == meta.RESTScopeNameNamespace {
			if unstructuredObj.GetNamespace() == "" {
				unstructuredObj.SetNamespace(nameSpace)
			}
			dri = dd.Resource(mapping.Resource).Namespace(unstructuredObj.GetNamespace())
		} else {
			dri = dd.Resource(mapping.Resource)
		}

		obj2, err := dri.Create(unstructuredObj, metav1.CreateOptions{})
		if  err != nil {
			log.Fatal(err)
		}

		fmt.Printf("%s/%s created", obj2.GetKind(), obj2.GetName())
	}
}

func homeDir() string {
	if h := os.Getenv("HOME"); h != "" {
		return h
	}
	return os.Getenv("USERPROFILE") // windows
}

//func newDeployment() *appsv1.Deployment {
//	labels := map[string]string{
//		"app":        "nginx",
//		"controller": foo.Name,
//	}
//	return &appsv1.Deployment{
//		ObjectMeta: metav1.ObjectMeta{
//			Name:      foo.Spec.DeploymentName,
//			Namespace: foo.Namespace,
//			OwnerReferences: []metav1.OwnerReference{
//				*metav1.NewControllerRef(foo, samplev1alpha1.SchemeGroupVersion.WithKind("Foo")),
//			},
//		},
//		Spec: appsv1.DeploymentSpec{
//			Replicas: foo.Spec.Replicas,
//			Selector: &metav1.LabelSelector{
//				MatchLabels: labels,
//			},
//			Template: corev1.PodTemplateSpec{
//				ObjectMeta: metav1.ObjectMeta{
//					Labels: labels,
//				},
//				Spec: corev1.PodSpec{
//					Containers: []corev1.Container{
//						{
//							Name:  "nginx",
//							Image: "nginx:latest",
//						},
//					},
//				},
//			},
//		},
//	}
//}
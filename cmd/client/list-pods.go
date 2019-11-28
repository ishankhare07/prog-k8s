package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	_ "k8s.io/client-go/plugin/pkg/client/auth/gcp"
	"k8s.io/client-go/tools/clientcmd"
)

func main() {
	kubeconfig := flag.String("kubeconfig", filepath.Join(os.Getenv("HOME"), ".kube/config"), "kubeconfig file")
	flag.Parse()
	config, err := clientcmd.BuildConfigFromFlags("", *kubeconfig)
	if err != nil {
		panic(err)
	}

	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic(err)
	}

	pods, err := clientset.CoreV1().Pods("default").List(metav1.ListOptions{})
	if err != nil {
		panic(err)
	}

	fmt.Printf("Found %d pods running\n", len(pods.Items))

	fmt.Printf("%-50s%-15s%-40s%s\n", "NAME", "NAMESPACE", "UID", "STATUS")
	for _, pod := range pods.Items {
		fmt.Printf("%-50s%-15s%-40s%s\n", pod.Name, pod.Namespace, pod.UID, pod.Status.Phase)
	}
}

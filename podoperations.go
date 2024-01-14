package main

import (
    "context"
    "fmt"
    "k8s.io/client-go/kubernetes"
    metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func GetPods(clientset *kubernetes.Clientset) {
    pods, err := clientset.CoreV1().Pods("").List(context.TODO(), metav1.ListOptions{})
    if err != nil {
        panic(err.Error())
    }
    for _, pod := range pods.Items {
        fmt.Printf("Name: %s, Namespace: %s\n", pod.Name, pod.Namespace)
    }
}

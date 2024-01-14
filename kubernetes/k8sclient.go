package main

import (
    "k8s.io/client-go/kubernetes"
    "k8s.io/client-go/rest"
)

func CreateK8sClient() (*kubernetes.Clientset, error) {
    config, err := rest.InClusterConfig()
    if err != nil {
        return nil, err
    }
    clientset, err := kubernetes.NewForConfig(config)
    if err != nil {
        return nil, err
    }
    return clientset, nil
}

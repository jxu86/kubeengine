package main

import (
	"context"
	"fmt"

	"github.com/jxu86/kubeengine/core/common/log"
	"github.com/jxu86/kubeengine/core/kubeclient"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var logger = log.GetLogger("main", log.INFO)

func main() {
	// service.Server()
	fmt.Println("kube start")
	kc := kubeclient.NewClients("./config/kubeconfig.yaml")

	ns, err := kc.KubeClient.CoreV1().Nodes().List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		logger.Errorf(err.Error())
	}
	for _, n := range ns.Items {
		logger.Infof("Node：", n.Name, n.Status.Addresses)
	}

	pods, err := kc.KubeClient.CoreV1().Pods("").List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		panic(err.Error())
	}
	fmt.Printf("There are %d pods in the cluster\n", len(pods.Items))

}

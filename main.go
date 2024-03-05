package main

import (
	"context"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"log"
	ctrl "sigs.k8s.io/controller-runtime"
	"time"
)

const NameSpace = "rbd-system"

func main() {
	clientset, err := kubernetes.NewForConfig(ctrl.GetConfigOrDie())
	if err != nil {
		panic(err.Error())
	}

	pods, err := clientset.CoreV1().Pods(NameSpace).List(context.TODO(), metav1.ListOptions{
		LabelSelector: "name=rbd-gateway",
	})

	if err != nil || len(pods.Items) == 0 {
		log.Fatalf("get pod error %v", err)
	}

	err = WriteHosts("/newetc/hosts", pods.Items[0].Status.PodIP)
	if err != nil {
		log.Fatalf("write host error %v", err)
	}

	err = SyncDockerCertFromSecret(clientset, NameSpace, "hub-image-repository")
	if err != nil {
		log.Fatalf("sync docker cert from secret error %v", err)
	}
	log.Println("all job success")

	time.Sleep(time.Hour * 1)
}

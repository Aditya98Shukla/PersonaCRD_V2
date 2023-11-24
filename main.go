package main

import (
	"context"
	"flag"
	"fmt"
	"os"
	/*"os/signal"
	"syscall"
	"time"

	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/cache"*/
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
  "k8s.io/client-go/util/homedir"
	//"k8s.io/client-go/util/wait"
	//"k8s.io/client-go/util/workqueue"
	"k8s.io/client-go/tools/clientcmd"
	/*"k8s.io/client-go/informers"
	"k8s.io/client-go/kubernetes/scheme"

	myv1alpha1 "path-to-your-custom-resource-package/apis/mygroup/v1alpha1" // Import your custom resource package
  */
  klient "github.com/Aditya98Shukla/PersonaCRD_V2/pkg/generated/clientset/versioned"
  //compv2 "github.com/Aditya98Shukla/PersonaCRD_V2/pkg/api/compute/v2"
)

const(
  DefaultNamespace string = "default"
)

var kubeconfig string

func main() {
	// Set up signals so we handle the first shutdown signal gracefully
	//stopCh := setupSignalHandler()
  fmt.Println("Start Program")
  if home := homedir.HomeDir(); home != "" {
    flag.StringVar(&kubeconfig, "kubeconfig", home+"/.kube/config", "(optional) absolute path to the kubeconfig file")
  } else {
    flag.StringVar(&kubeconfig, "kubeconfig", "", "absolute path to the kubeconfig file")
  }
  flag.Parse()
	
  fmt.Println("Create a Kubernetes client")
	config, err := clientcmd.BuildConfigFromFlags("", kubeconfig)
	if err != nil {
		fmt.Printf("Error building kubeconfig: %s\n", err.Error())
		os.Exit(1)
	}

  fmt.Println("Create a Kubernetes clientset")
	klientset, err := klient.NewForConfig(config)
	if err != nil {
		fmt.Printf("Error creating Kubernetes client: %s\n", err.Error())
		os.Exit(1)
  }

  personas, err := klientset.ComputeV2().Personas(DefaultNamespace).List(context.TODO(),metav1.ListOptions{})
  if err != nil {
    fmt.Printf("Error getting all personas: %s\n", err.Error())
    os.Exit(1)
  }
  fmt.Println(personas.Items)
  fmt.Println("Exit Program")
}

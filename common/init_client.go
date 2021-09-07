package common

import (
	"log"

	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
)

// Init the kubernetes client instance
func InitClient() (clientset *kubernetes.Clientset, err error) {
	var (
		restconf *rest.Config
	)

	restconf, err = GetRestConf()
	if err != nil {
		log.Fatal(err)
	}

	// Init the clientset configurations
	clientset, err = kubernetes.NewForConfig(restconf)
	if err != nil {
		log.Fatal(err)
	}

	return clientset, nil
}

// Get the Kubernetes Restful Client configurations
func GetRestConf() (restconf *rest.Config, err error) {

	// read the local kubeconfig
	// kubeconfig := filepath.Join(
	// 	os.Getenv("HOME"), ".kube", "config",
	// )

	// get the restful configuration
	restconf, err = clientcmd.BuildConfigFromFlags("", "./admin.conf")
	if err != nil {
		log.Fatal(err)
	}

	return
}

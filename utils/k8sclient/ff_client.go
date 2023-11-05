package k8s

import (
	"context"
	"ffapi/config"
	configmodels "ffapi/config/config-models"
	"ffapi/pkg/apis/deployment/v1alpha1"
	"ffapi/pkg/generated/clientset/versioned"
	"fmt"
	"log"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
)

var (
	FFClient *K8sClient
)

func init() {
	cfg, ok := config.Cfg.Configs["k8sclient"].(*configmodels.K8sClientConfig)
	if !ok {
		log.Fatal(fmt.Errorf("unable to cast config %s to IViperConfig", cfg.GetName()))
	}
	FFClient = newK8sFFClient(*cfg)
}

// This package contains modules for interacting with the Kubernetes API.

type K8sClient struct {
	// the k8s client itself
	Client           *versioned.Clientset
	DefaultNamespace string
}

// create a new k8s client using loaded config
func newK8sFFClient(cfg configmodels.K8sClientConfig) *K8sClient {
	var restCfg *rest.Config
	var err error
	// create a new k8s clientset
	if restCfg, err = rest.InClusterConfig(); err != nil {
		log.Println("Unable to load in cluster config, trying to load from kubeconfig")
		if restCfg, err = clientcmd.BuildConfigFromFlags("", cfg.KubeConfigPath); err != nil {
			panic(fmt.Errorf("unable to load kubeconfig: %w", err))
		}
	}

	if client, err := versioned.NewForConfig(restCfg); err == nil {
		return &K8sClient{
			Client:           client,
			DefaultNamespace: cfg.Namespace,
		}
	}

	panic(fmt.Errorf("unable to create k8s client: %w", err))
}

// Get the featureflag resource
func (kc *K8sClient) Get(namespace string, name string) (*v1alpha1.FeatureFlag, error) {
	if namespace == "" {
		namespace = kc.DefaultNamespace
	}
	return kc.Client.DeploymentV1alpha1().FeatureFlags(namespace).Get(context.Background(), name, metav1.GetOptions{})
}

// List the featureflag resources
func (kc *K8sClient) List(namespace string) (*v1alpha1.FeatureFlagList, error) {
	if namespace == "" {
		namespace = kc.DefaultNamespace
	}
	return kc.Client.DeploymentV1alpha1().FeatureFlags(namespace).List(context.Background(), metav1.ListOptions{})
}

// Create the featureflag resource
func (kc *K8sClient) Create(ff *v1alpha1.FeatureFlag) (*v1alpha1.FeatureFlag, error) {
	if ff.Namespace == "" {
		ff.Namespace = kc.DefaultNamespace
	}
	return kc.Client.DeploymentV1alpha1().FeatureFlags(ff.Namespace).Create(context.Background(), ff, metav1.CreateOptions{})
}

// Update the featureflag resource
func (kc *K8sClient) Update(namespace string, ff *v1alpha1.FeatureFlag) (*v1alpha1.FeatureFlag, error) {
	if namespace == "" {
		namespace = kc.DefaultNamespace
	}
	return kc.Client.DeploymentV1alpha1().FeatureFlags(namespace).Update(context.Background(), ff, metav1.UpdateOptions{})
}

// Update the status of the featureflag resource
func (kc *K8sClient) UpdateStatus(namespace string, ff *v1alpha1.FeatureFlag) (*v1alpha1.FeatureFlag, error) {
	if namespace == "" {
		namespace = kc.DefaultNamespace
	}
	return kc.Client.DeploymentV1alpha1().FeatureFlags(namespace).UpdateStatus(context.Background(), ff, metav1.UpdateOptions{})
}

// Delete the featureflag resource
func (kc *K8sClient) Delete(namespace string, name string) error {
	if namespace == "" {
		namespace = kc.DefaultNamespace
	}
	return kc.Client.DeploymentV1alpha1().FeatureFlags(namespace).Delete(context.Background(), name, metav1.DeleteOptions{})
}

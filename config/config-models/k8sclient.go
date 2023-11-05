package configmodels

type K8sClientConfig struct {
	// The path to the kubeconfig file
	KubeConfigPath string `mapstructure:"kubeConfigPath"`
	// The namespace to use for the k8s client
	Namespace string `mapstructure:"namespace"`
}

func (K8sClientConfig) Default() interface{} {
	return &K8sClientConfig{
		KubeConfigPath: "",
		Namespace:      "default",
	}
}

func (K8sClientConfig) GetName() string {
	return "k8sclient"
}

package k8s_client

import (
	"context"
	"fmt"
	"strings"

	"k8s.io/client-go/kubernetes"

	_ "k8s.io/client-go/plugin/pkg/client/auth"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/tools/clientcmd/api"
)

type Client struct {
	services map[string]K8sServices
	contexts []string
	paths    map[string]struct{}

	Context string
}

func (c *Client) K8sServices() K8sServices {
	return c.services[c.Context]
}

func (c Client) CopyWithContext(k8sContext string) *Client {
	return &Client{
		services: c.services,
		contexts: c.contexts,
		paths:    c.paths,
		Context:  k8sContext,
	}
}

func (c *Client) SetServices(s map[string]K8sServices) {
	c.services = s
	contexts := make([]string, 0, len(s))
	for k := range s {
		contexts = append(contexts, k)
	}
	c.contexts = contexts
}

func NewClient(ctx context.Context, k8sConfigFile string, k8sContextSlice []string) (*Client, error) {
	loadingRules := clientcmd.NewDefaultClientConfigLoadingRules()
	loadingRules.DefaultClientConfig = &clientcmd.DefaultClientConfig
	loadingRules.ExplicitPath = k8sConfigFile
	configOverrides := &clientcmd.ConfigOverrides{
		//ClusterDefaults: clientcmd.ClusterDefaults,
		//CurrentContext:  "", // or "current-context name"
	}
	kubeConfig := clientcmd.NewNonInteractiveDeferredLoadingClientConfig(
		loadingRules,
		configOverrides,
	)
	rawKubeConfig, err := kubeConfig.RawConfig()
	if err != nil {
		return nil, err
	}

	var contexts []string
	switch len(k8sContextSlice) {
	case 0:
		contexts = []string{rawKubeConfig.CurrentContext}
	case 1:
		if k8sContextSlice[0] == "*" {
			for cName := range rawKubeConfig.Contexts {
				contexts = append(contexts, cName)
			}
		} else {
			if _, ok := rawKubeConfig.Contexts[k8sContextSlice[0]]; !ok {
				return nil, fmt.Errorf("context %q doesn't exist in kube configuration", k8sContextSlice[0])
			}
			contexts = []string{k8sContextSlice[0]}
		}
	default:
		for _, cName := range k8sContextSlice {
			if _, ok := rawKubeConfig.Contexts[cName]; !ok {
				return nil, fmt.Errorf("context %q doesn't exist in kube configuration", cName)
			}
			contexts = append(contexts, cName)
		}
	}

	if len(contexts) == 0 {
		return nil, fmt.Errorf("could not find any context. Try to add context, https://kubernetes.io/docs/reference/kubectl/cheatsheet/#kubectl-context-and-configuration")
	}

	c := Client{
		services: make(map[string]K8sServices),
		contexts: contexts,
		Context:  contexts[0],
		paths:    make(map[string]struct{}),
	}

	for _, ctxName := range contexts {
		kClient, err := buildKubeClient(rawKubeConfig, ctxName)
		if err != nil {
			return nil, fmt.Errorf("failed to build k8s client for context %q: %w", ctxName, err)
		}
		c.paths, err = getAPIsMap(kClient)
		if err != nil {
		}
		c.services[ctxName] = initServices(kClient)
	}

	return &c, nil
}

func buildKubeClient(kubeConfig api.Config, ctx string) (*kubernetes.Clientset, error) {
	override := &clientcmd.ConfigOverrides{CurrentContext: ctx}
	clientConfig := clientcmd.NewNonInteractiveClientConfig(
		kubeConfig,
		override.CurrentContext,
		override,
		&clientcmd.ClientConfigLoadingRules{},
	)
	restConfig, err := clientConfig.ClientConfig()
	if err != nil {
		return nil, err
	}
	return kubernetes.NewForConfig(restConfig)
}

func getAPIsMap(client *kubernetes.Clientset) (map[string]struct{}, error) {
	doc, err := client.OpenAPISchema()
	if err != nil {
		return nil, err
	}
	paths := make(map[string]struct{})
	for _, p := range doc.Paths.Path {
		path := p.Name
		if strings.HasPrefix(path, "/apis/") {
			paths[path] = struct{}{}
		}
	}
	return paths, nil
}

func initServices(client *kubernetes.Clientset) K8sServices {
	return K8sServices{
		Client:          client,
		CronJobs:        client.BatchV1().CronJobs(""),
		DaemonSets:      client.AppsV1().DaemonSets(""),
		Deployments:     client.AppsV1().Deployments(""),
		Endpoints:       client.CoreV1().Endpoints(""),
		Jobs:            client.BatchV1().Jobs(""),
		LimitRanges:     client.CoreV1().LimitRanges(""),
		Namespaces:      client.CoreV1().Namespaces(),
		NetworkPolicies: client.NetworkingV1().NetworkPolicies(""),
		Nodes:           client.CoreV1().Nodes(),
		Pods:            client.CoreV1().Pods(""),
		ReplicaSets:     client.AppsV1().ReplicaSets(""),
		ResourceQuotas:  client.CoreV1().ResourceQuotas(""),
		RoleBindings:    client.RbacV1().RoleBindings(""),
		Roles:           client.RbacV1().Roles(""),
		ServiceAccounts: client.CoreV1().ServiceAccounts(""),
		K8sServices:     client.CoreV1().Services(""),
		StatefulSets:    client.AppsV1().StatefulSets(""),
	}
}

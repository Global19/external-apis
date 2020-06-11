// Code generated by skv2. DO NOT EDIT.

package v1

import (
	v1 "github.com/solo-io/external-apis/pkg/api/k8s/core/v1"

	"k8s.io/client-go/rest"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

/*
  The intention of these providers are to be used for Mocking.
  They expose the Clients as interfaces, as well as factories to provide mocked versions
  of the clients when they require building within a component.

  See package `github.com/solo-io/skv2/pkg/multicluster/register` for example
*/

// Provider for Secret Client from Client
func SecretClientProvider(client client.Client) v1.SecretClient {
	return v1.NewSecretClient(client)
}

type SecretClientFactory func(client client.Client) v1.SecretClient

func SecretClientFactoryProvider() SecretClientFactory {
	return SecretClientProvider
}

type SecretClientFromConfigFactory func(cfg *rest.Config) (v1.SecretClient, error)

func SecretClientFromConfigFactoryProvider() SecretClientFromConfigFactory {
	return func(cfg *rest.Config) (v1.SecretClient, error) {
		clients, err := v1.NewClientsetFromConfig(cfg)
		if err != nil {
			return nil, err
		}
		return clients.Secrets(), nil
	}
}

// Provider for ServiceAccount Client from Client
func ServiceAccountClientProvider(client client.Client) v1.ServiceAccountClient {
	return v1.NewServiceAccountClient(client)
}

type ServiceAccountClientFactory func(client client.Client) v1.ServiceAccountClient

func ServiceAccountClientFactoryProvider() ServiceAccountClientFactory {
	return ServiceAccountClientProvider
}

type ServiceAccountClientFromConfigFactory func(cfg *rest.Config) (v1.ServiceAccountClient, error)

func ServiceAccountClientFromConfigFactoryProvider() ServiceAccountClientFromConfigFactory {
	return func(cfg *rest.Config) (v1.ServiceAccountClient, error) {
		clients, err := v1.NewClientsetFromConfig(cfg)
		if err != nil {
			return nil, err
		}
		return clients.ServiceAccounts(), nil
	}
}

// Provider for ConfigMap Client from Client
func ConfigMapClientProvider(client client.Client) v1.ConfigMapClient {
	return v1.NewConfigMapClient(client)
}

type ConfigMapClientFactory func(client client.Client) v1.ConfigMapClient

func ConfigMapClientFactoryProvider() ConfigMapClientFactory {
	return ConfigMapClientProvider
}

type ConfigMapClientFromConfigFactory func(cfg *rest.Config) (v1.ConfigMapClient, error)

func ConfigMapClientFromConfigFactoryProvider() ConfigMapClientFromConfigFactory {
	return func(cfg *rest.Config) (v1.ConfigMapClient, error) {
		clients, err := v1.NewClientsetFromConfig(cfg)
		if err != nil {
			return nil, err
		}
		return clients.ConfigMaps(), nil
	}
}

// Provider for Service Client from Client
func ServiceClientProvider(client client.Client) v1.ServiceClient {
	return v1.NewServiceClient(client)
}

type ServiceClientFactory func(client client.Client) v1.ServiceClient

func ServiceClientFactoryProvider() ServiceClientFactory {
	return ServiceClientProvider
}

type ServiceClientFromConfigFactory func(cfg *rest.Config) (v1.ServiceClient, error)

func ServiceClientFromConfigFactoryProvider() ServiceClientFromConfigFactory {
	return func(cfg *rest.Config) (v1.ServiceClient, error) {
		clients, err := v1.NewClientsetFromConfig(cfg)
		if err != nil {
			return nil, err
		}
		return clients.Services(), nil
	}
}

// Provider for Pod Client from Client
func PodClientProvider(client client.Client) v1.PodClient {
	return v1.NewPodClient(client)
}

type PodClientFactory func(client client.Client) v1.PodClient

func PodClientFactoryProvider() PodClientFactory {
	return PodClientProvider
}

type PodClientFromConfigFactory func(cfg *rest.Config) (v1.PodClient, error)

func PodClientFromConfigFactoryProvider() PodClientFromConfigFactory {
	return func(cfg *rest.Config) (v1.PodClient, error) {
		clients, err := v1.NewClientsetFromConfig(cfg)
		if err != nil {
			return nil, err
		}
		return clients.Pods(), nil
	}
}

// Provider for Namespace Client from Client
func NamespaceClientProvider(client client.Client) v1.NamespaceClient {
	return v1.NewNamespaceClient(client)
}

type NamespaceClientFactory func(client client.Client) v1.NamespaceClient

func NamespaceClientFactoryProvider() NamespaceClientFactory {
	return NamespaceClientProvider
}

type NamespaceClientFromConfigFactory func(cfg *rest.Config) (v1.NamespaceClient, error)

func NamespaceClientFromConfigFactoryProvider() NamespaceClientFromConfigFactory {
	return func(cfg *rest.Config) (v1.NamespaceClient, error) {
		clients, err := v1.NewClientsetFromConfig(cfg)
		if err != nil {
			return nil, err
		}
		return clients.Namespaces(), nil
	}
}

// Provider for Node Client from Client
func NodeClientProvider(client client.Client) v1.NodeClient {
	return v1.NewNodeClient(client)
}

type NodeClientFactory func(client client.Client) v1.NodeClient

func NodeClientFactoryProvider() NodeClientFactory {
	return NodeClientProvider
}

type NodeClientFromConfigFactory func(cfg *rest.Config) (v1.NodeClient, error)

func NodeClientFromConfigFactoryProvider() NodeClientFromConfigFactory {
	return func(cfg *rest.Config) (v1.NodeClient, error) {
		clients, err := v1.NewClientsetFromConfig(cfg)
		if err != nil {
			return nil, err
		}
		return clients.Nodes(), nil
	}
}

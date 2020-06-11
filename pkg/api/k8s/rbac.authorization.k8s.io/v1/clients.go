// Code generated by skv2. DO NOT EDIT.

//go:generate mockgen -source ./clients.go -destination mocks/clients.go

package v1

import (
	"context"

	"github.com/solo-io/skv2/pkg/controllerutils"
	"github.com/solo-io/skv2/pkg/multicluster"
	rbac_authorization_k8s_io_v1 "k8s.io/api/rbac/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/client-go/kubernetes/scheme"
	"k8s.io/client-go/rest"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

// MulticlusterClientset for the rbac.authorization.k8s.io/v1 APIs
type MulticlusterClientset interface {
	// Cluster returns a Clientset for the given cluster
	Cluster(cluster string) (Clientset, error)
}

type multiclusterClientset struct {
	client multicluster.Client
}

func NewMulticlusterClientset(client multicluster.Client) MulticlusterClientset {
	return &multiclusterClientset{client: client}
}

func (m *multiclusterClientset) Cluster(cluster string) (Clientset, error) {
	client, err := m.client.Cluster(cluster)
	if err != nil {
		return nil, err
	}
	return NewClientset(client), nil
}

// clienset for the rbac.authorization.k8s.io/v1 APIs
type Clientset interface {
	// clienset for the rbac.authorization.k8s.io/v1/v1 APIs
	Roles() RoleClient
	// clienset for the rbac.authorization.k8s.io/v1/v1 APIs
	RoleBindings() RoleBindingClient
	// clienset for the rbac.authorization.k8s.io/v1/v1 APIs
	ClusterRoles() ClusterRoleClient
	// clienset for the rbac.authorization.k8s.io/v1/v1 APIs
	ClusterRoleBindings() ClusterRoleBindingClient
}

type clientSet struct {
	client client.Client
}

func NewClientsetFromConfig(cfg *rest.Config) (Clientset, error) {
	scheme := scheme.Scheme
	if err := AddToScheme(scheme); err != nil {
		return nil, err
	}
	client, err := client.New(cfg, client.Options{
		Scheme: scheme,
	})
	if err != nil {
		return nil, err
	}
	return NewClientset(client), nil
}

func NewClientset(client client.Client) Clientset {
	return &clientSet{client: client}
}

// clienset for the rbac.authorization.k8s.io/v1/v1 APIs
func (c *clientSet) Roles() RoleClient {
	return NewRoleClient(c.client)
}

// clienset for the rbac.authorization.k8s.io/v1/v1 APIs
func (c *clientSet) RoleBindings() RoleBindingClient {
	return NewRoleBindingClient(c.client)
}

// clienset for the rbac.authorization.k8s.io/v1/v1 APIs
func (c *clientSet) ClusterRoles() ClusterRoleClient {
	return NewClusterRoleClient(c.client)
}

// clienset for the rbac.authorization.k8s.io/v1/v1 APIs
func (c *clientSet) ClusterRoleBindings() ClusterRoleBindingClient {
	return NewClusterRoleBindingClient(c.client)
}

// Reader knows how to read and list Roles.
type RoleReader interface {
	// Get retrieves a Role for the given object key
	GetRole(ctx context.Context, key client.ObjectKey) (*rbac_authorization_k8s_io_v1.Role, error)

	// List retrieves list of Roles for a given namespace and list options.
	ListRole(ctx context.Context, opts ...client.ListOption) (*rbac_authorization_k8s_io_v1.RoleList, error)
}

// RoleTransitionFunction instructs the RoleWriter how to transition between an existing
// Role object and a desired on an Upsert
type RoleTransitionFunction func(existing, desired *rbac_authorization_k8s_io_v1.Role) error

// Writer knows how to create, delete, and update Roles.
type RoleWriter interface {
	// Create saves the Role object.
	CreateRole(ctx context.Context, obj *rbac_authorization_k8s_io_v1.Role, opts ...client.CreateOption) error

	// Delete deletes the Role object.
	DeleteRole(ctx context.Context, key client.ObjectKey, opts ...client.DeleteOption) error

	// Update updates the given Role object.
	UpdateRole(ctx context.Context, obj *rbac_authorization_k8s_io_v1.Role, opts ...client.UpdateOption) error

	// Patch patches the given Role object.
	PatchRole(ctx context.Context, obj *rbac_authorization_k8s_io_v1.Role, patch client.Patch, opts ...client.PatchOption) error

	// DeleteAllOf deletes all Role objects matching the given options.
	DeleteAllOfRole(ctx context.Context, opts ...client.DeleteAllOfOption) error

	// Create or Update the Role object.
	UpsertRole(ctx context.Context, obj *rbac_authorization_k8s_io_v1.Role, transitionFuncs ...RoleTransitionFunction) error
}

// StatusWriter knows how to update status subresource of a Role object.
type RoleStatusWriter interface {
	// Update updates the fields corresponding to the status subresource for the
	// given Role object.
	UpdateRoleStatus(ctx context.Context, obj *rbac_authorization_k8s_io_v1.Role, opts ...client.UpdateOption) error

	// Patch patches the given Role object's subresource.
	PatchRoleStatus(ctx context.Context, obj *rbac_authorization_k8s_io_v1.Role, patch client.Patch, opts ...client.PatchOption) error
}

// Client knows how to perform CRUD operations on Roles.
type RoleClient interface {
	RoleReader
	RoleWriter
	RoleStatusWriter
}

type roleClient struct {
	client client.Client
}

func NewRoleClient(client client.Client) *roleClient {
	return &roleClient{client: client}
}

func (c *roleClient) GetRole(ctx context.Context, key client.ObjectKey) (*rbac_authorization_k8s_io_v1.Role, error) {
	obj := &Role{}
	if err := c.client.Get(ctx, key, obj); err != nil {
		return nil, err
	}
	return obj, nil
}

func (c *roleClient) ListRole(ctx context.Context, opts ...client.ListOption) (*rbac_authorization_k8s_io_v1.RoleList, error) {
	list := &RoleList{}
	if err := c.client.List(ctx, list, opts...); err != nil {
		return nil, err
	}
	return list, nil
}

func (c *roleClient) CreateRole(ctx context.Context, obj *rbac_authorization_k8s_io_v1.Role, opts ...client.CreateOption) error {
	return c.client.Create(ctx, obj, opts...)
}

func (c *roleClient) DeleteRole(ctx context.Context, key client.ObjectKey, opts ...client.DeleteOption) error {
	obj := &Role{}
	obj.SetName(key.Name)
	obj.SetNamespace(key.Namespace)
	return c.client.Delete(ctx, obj, opts...)
}

func (c *roleClient) UpdateRole(ctx context.Context, obj *rbac_authorization_k8s_io_v1.Role, opts ...client.UpdateOption) error {
	return c.client.Update(ctx, obj, opts...)
}

func (c *roleClient) PatchRole(ctx context.Context, obj *rbac_authorization_k8s_io_v1.Role, patch client.Patch, opts ...client.PatchOption) error {
	return c.client.Patch(ctx, obj, patch, opts...)
}

func (c *roleClient) DeleteAllOfRole(ctx context.Context, opts ...client.DeleteAllOfOption) error {
	obj := &Role{}
	return c.client.DeleteAllOf(ctx, obj, opts...)
}

func (c *roleClient) UpsertRole(ctx context.Context, obj *rbac_authorization_k8s_io_v1.Role, transitionFuncs ...RoleTransitionFunction) error {
	genericTxFunc := func(existing, desired runtime.Object) error {
		for _, txFunc := range transitionFuncs {
			if err := txFunc(existing.(*rbac_authorization_k8s_io_v1.Role), desired.(*rbac_authorization_k8s_io_v1.Role)); err != nil {
				return err
			}
		}
		return nil
	}
	_, err := controllerutils.Upsert(ctx, c.client, obj, genericTxFunc)
	return err
}

func (c *roleClient) UpdateRoleStatus(ctx context.Context, obj *rbac_authorization_k8s_io_v1.Role, opts ...client.UpdateOption) error {
	return c.client.Status().Update(ctx, obj, opts...)
}

func (c *roleClient) PatchRoleStatus(ctx context.Context, obj *rbac_authorization_k8s_io_v1.Role, patch client.Patch, opts ...client.PatchOption) error {
	return c.client.Status().Patch(ctx, obj, patch, opts...)
}

// Reader knows how to read and list RoleBindings.
type RoleBindingReader interface {
	// Get retrieves a RoleBinding for the given object key
	GetRoleBinding(ctx context.Context, key client.ObjectKey) (*rbac_authorization_k8s_io_v1.RoleBinding, error)

	// List retrieves list of RoleBindings for a given namespace and list options.
	ListRoleBinding(ctx context.Context, opts ...client.ListOption) (*rbac_authorization_k8s_io_v1.RoleBindingList, error)
}

// RoleBindingTransitionFunction instructs the RoleBindingWriter how to transition between an existing
// RoleBinding object and a desired on an Upsert
type RoleBindingTransitionFunction func(existing, desired *rbac_authorization_k8s_io_v1.RoleBinding) error

// Writer knows how to create, delete, and update RoleBindings.
type RoleBindingWriter interface {
	// Create saves the RoleBinding object.
	CreateRoleBinding(ctx context.Context, obj *rbac_authorization_k8s_io_v1.RoleBinding, opts ...client.CreateOption) error

	// Delete deletes the RoleBinding object.
	DeleteRoleBinding(ctx context.Context, key client.ObjectKey, opts ...client.DeleteOption) error

	// Update updates the given RoleBinding object.
	UpdateRoleBinding(ctx context.Context, obj *rbac_authorization_k8s_io_v1.RoleBinding, opts ...client.UpdateOption) error

	// Patch patches the given RoleBinding object.
	PatchRoleBinding(ctx context.Context, obj *rbac_authorization_k8s_io_v1.RoleBinding, patch client.Patch, opts ...client.PatchOption) error

	// DeleteAllOf deletes all RoleBinding objects matching the given options.
	DeleteAllOfRoleBinding(ctx context.Context, opts ...client.DeleteAllOfOption) error

	// Create or Update the RoleBinding object.
	UpsertRoleBinding(ctx context.Context, obj *rbac_authorization_k8s_io_v1.RoleBinding, transitionFuncs ...RoleBindingTransitionFunction) error
}

// StatusWriter knows how to update status subresource of a RoleBinding object.
type RoleBindingStatusWriter interface {
	// Update updates the fields corresponding to the status subresource for the
	// given RoleBinding object.
	UpdateRoleBindingStatus(ctx context.Context, obj *rbac_authorization_k8s_io_v1.RoleBinding, opts ...client.UpdateOption) error

	// Patch patches the given RoleBinding object's subresource.
	PatchRoleBindingStatus(ctx context.Context, obj *rbac_authorization_k8s_io_v1.RoleBinding, patch client.Patch, opts ...client.PatchOption) error
}

// Client knows how to perform CRUD operations on RoleBindings.
type RoleBindingClient interface {
	RoleBindingReader
	RoleBindingWriter
	RoleBindingStatusWriter
}

type roleBindingClient struct {
	client client.Client
}

func NewRoleBindingClient(client client.Client) *roleBindingClient {
	return &roleBindingClient{client: client}
}

func (c *roleBindingClient) GetRoleBinding(ctx context.Context, key client.ObjectKey) (*rbac_authorization_k8s_io_v1.RoleBinding, error) {
	obj := &RoleBinding{}
	if err := c.client.Get(ctx, key, obj); err != nil {
		return nil, err
	}
	return obj, nil
}

func (c *roleBindingClient) ListRoleBinding(ctx context.Context, opts ...client.ListOption) (*rbac_authorization_k8s_io_v1.RoleBindingList, error) {
	list := &RoleBindingList{}
	if err := c.client.List(ctx, list, opts...); err != nil {
		return nil, err
	}
	return list, nil
}

func (c *roleBindingClient) CreateRoleBinding(ctx context.Context, obj *rbac_authorization_k8s_io_v1.RoleBinding, opts ...client.CreateOption) error {
	return c.client.Create(ctx, obj, opts...)
}

func (c *roleBindingClient) DeleteRoleBinding(ctx context.Context, key client.ObjectKey, opts ...client.DeleteOption) error {
	obj := &RoleBinding{}
	obj.SetName(key.Name)
	obj.SetNamespace(key.Namespace)
	return c.client.Delete(ctx, obj, opts...)
}

func (c *roleBindingClient) UpdateRoleBinding(ctx context.Context, obj *rbac_authorization_k8s_io_v1.RoleBinding, opts ...client.UpdateOption) error {
	return c.client.Update(ctx, obj, opts...)
}

func (c *roleBindingClient) PatchRoleBinding(ctx context.Context, obj *rbac_authorization_k8s_io_v1.RoleBinding, patch client.Patch, opts ...client.PatchOption) error {
	return c.client.Patch(ctx, obj, patch, opts...)
}

func (c *roleBindingClient) DeleteAllOfRoleBinding(ctx context.Context, opts ...client.DeleteAllOfOption) error {
	obj := &RoleBinding{}
	return c.client.DeleteAllOf(ctx, obj, opts...)
}

func (c *roleBindingClient) UpsertRoleBinding(ctx context.Context, obj *rbac_authorization_k8s_io_v1.RoleBinding, transitionFuncs ...RoleBindingTransitionFunction) error {
	genericTxFunc := func(existing, desired runtime.Object) error {
		for _, txFunc := range transitionFuncs {
			if err := txFunc(existing.(*rbac_authorization_k8s_io_v1.RoleBinding), desired.(*rbac_authorization_k8s_io_v1.RoleBinding)); err != nil {
				return err
			}
		}
		return nil
	}
	_, err := controllerutils.Upsert(ctx, c.client, obj, genericTxFunc)
	return err
}

func (c *roleBindingClient) UpdateRoleBindingStatus(ctx context.Context, obj *rbac_authorization_k8s_io_v1.RoleBinding, opts ...client.UpdateOption) error {
	return c.client.Status().Update(ctx, obj, opts...)
}

func (c *roleBindingClient) PatchRoleBindingStatus(ctx context.Context, obj *rbac_authorization_k8s_io_v1.RoleBinding, patch client.Patch, opts ...client.PatchOption) error {
	return c.client.Status().Patch(ctx, obj, patch, opts...)
}

// Reader knows how to read and list ClusterRoles.
type ClusterRoleReader interface {
	// Get retrieves a ClusterRole for the given object key
	GetClusterRole(ctx context.Context, name string) (*rbac_authorization_k8s_io_v1.ClusterRole, error)

	// List retrieves list of ClusterRoles for a given namespace and list options.
	ListClusterRole(ctx context.Context, opts ...client.ListOption) (*rbac_authorization_k8s_io_v1.ClusterRoleList, error)
}

// ClusterRoleTransitionFunction instructs the ClusterRoleWriter how to transition between an existing
// ClusterRole object and a desired on an Upsert
type ClusterRoleTransitionFunction func(existing, desired *rbac_authorization_k8s_io_v1.ClusterRole) error

// Writer knows how to create, delete, and update ClusterRoles.
type ClusterRoleWriter interface {
	// Create saves the ClusterRole object.
	CreateClusterRole(ctx context.Context, obj *rbac_authorization_k8s_io_v1.ClusterRole, opts ...client.CreateOption) error

	// Delete deletes the ClusterRole object.
	DeleteClusterRole(ctx context.Context, name string, opts ...client.DeleteOption) error

	// Update updates the given ClusterRole object.
	UpdateClusterRole(ctx context.Context, obj *rbac_authorization_k8s_io_v1.ClusterRole, opts ...client.UpdateOption) error

	// Patch patches the given ClusterRole object.
	PatchClusterRole(ctx context.Context, obj *rbac_authorization_k8s_io_v1.ClusterRole, patch client.Patch, opts ...client.PatchOption) error

	// DeleteAllOf deletes all ClusterRole objects matching the given options.
	DeleteAllOfClusterRole(ctx context.Context, opts ...client.DeleteAllOfOption) error

	// Create or Update the ClusterRole object.
	UpsertClusterRole(ctx context.Context, obj *rbac_authorization_k8s_io_v1.ClusterRole, transitionFuncs ...ClusterRoleTransitionFunction) error
}

// StatusWriter knows how to update status subresource of a ClusterRole object.
type ClusterRoleStatusWriter interface {
	// Update updates the fields corresponding to the status subresource for the
	// given ClusterRole object.
	UpdateClusterRoleStatus(ctx context.Context, obj *rbac_authorization_k8s_io_v1.ClusterRole, opts ...client.UpdateOption) error

	// Patch patches the given ClusterRole object's subresource.
	PatchClusterRoleStatus(ctx context.Context, obj *rbac_authorization_k8s_io_v1.ClusterRole, patch client.Patch, opts ...client.PatchOption) error
}

// Client knows how to perform CRUD operations on ClusterRoles.
type ClusterRoleClient interface {
	ClusterRoleReader
	ClusterRoleWriter
	ClusterRoleStatusWriter
}

type clusterRoleClient struct {
	client client.Client
}

func NewClusterRoleClient(client client.Client) *clusterRoleClient {
	return &clusterRoleClient{client: client}
}

func (c *clusterRoleClient) GetClusterRole(ctx context.Context, name string) (*rbac_authorization_k8s_io_v1.ClusterRole, error) {
	obj := &ClusterRole{}
	key := client.ObjectKey{
		Name: name,
	}
	if err := c.client.Get(ctx, key, obj); err != nil {
		return nil, err
	}
	return obj, nil
}

func (c *clusterRoleClient) ListClusterRole(ctx context.Context, opts ...client.ListOption) (*rbac_authorization_k8s_io_v1.ClusterRoleList, error) {
	list := &ClusterRoleList{}
	if err := c.client.List(ctx, list, opts...); err != nil {
		return nil, err
	}
	return list, nil
}

func (c *clusterRoleClient) CreateClusterRole(ctx context.Context, obj *rbac_authorization_k8s_io_v1.ClusterRole, opts ...client.CreateOption) error {
	return c.client.Create(ctx, obj, opts...)
}

func (c *clusterRoleClient) DeleteClusterRole(ctx context.Context, name string, opts ...client.DeleteOption) error {
	obj := &ClusterRole{}
	obj.SetName(name)
	return c.client.Delete(ctx, obj, opts...)
}

func (c *clusterRoleClient) UpdateClusterRole(ctx context.Context, obj *rbac_authorization_k8s_io_v1.ClusterRole, opts ...client.UpdateOption) error {
	return c.client.Update(ctx, obj, opts...)
}

func (c *clusterRoleClient) PatchClusterRole(ctx context.Context, obj *rbac_authorization_k8s_io_v1.ClusterRole, patch client.Patch, opts ...client.PatchOption) error {
	return c.client.Patch(ctx, obj, patch, opts...)
}

func (c *clusterRoleClient) DeleteAllOfClusterRole(ctx context.Context, opts ...client.DeleteAllOfOption) error {
	obj := &ClusterRole{}
	return c.client.DeleteAllOf(ctx, obj, opts...)
}

func (c *clusterRoleClient) UpsertClusterRole(ctx context.Context, obj *rbac_authorization_k8s_io_v1.ClusterRole, transitionFuncs ...ClusterRoleTransitionFunction) error {
	genericTxFunc := func(existing, desired runtime.Object) error {
		for _, txFunc := range transitionFuncs {
			if err := txFunc(existing.(*rbac_authorization_k8s_io_v1.ClusterRole), desired.(*rbac_authorization_k8s_io_v1.ClusterRole)); err != nil {
				return err
			}
		}
		return nil
	}
	_, err := controllerutils.Upsert(ctx, c.client, obj, genericTxFunc)
	return err
}

func (c *clusterRoleClient) UpdateClusterRoleStatus(ctx context.Context, obj *rbac_authorization_k8s_io_v1.ClusterRole, opts ...client.UpdateOption) error {
	return c.client.Status().Update(ctx, obj, opts...)
}

func (c *clusterRoleClient) PatchClusterRoleStatus(ctx context.Context, obj *rbac_authorization_k8s_io_v1.ClusterRole, patch client.Patch, opts ...client.PatchOption) error {
	return c.client.Status().Patch(ctx, obj, patch, opts...)
}

// Reader knows how to read and list ClusterRoleBindings.
type ClusterRoleBindingReader interface {
	// Get retrieves a ClusterRoleBinding for the given object key
	GetClusterRoleBinding(ctx context.Context, name string) (*rbac_authorization_k8s_io_v1.ClusterRoleBinding, error)

	// List retrieves list of ClusterRoleBindings for a given namespace and list options.
	ListClusterRoleBinding(ctx context.Context, opts ...client.ListOption) (*rbac_authorization_k8s_io_v1.ClusterRoleBindingList, error)
}

// ClusterRoleBindingTransitionFunction instructs the ClusterRoleBindingWriter how to transition between an existing
// ClusterRoleBinding object and a desired on an Upsert
type ClusterRoleBindingTransitionFunction func(existing, desired *rbac_authorization_k8s_io_v1.ClusterRoleBinding) error

// Writer knows how to create, delete, and update ClusterRoleBindings.
type ClusterRoleBindingWriter interface {
	// Create saves the ClusterRoleBinding object.
	CreateClusterRoleBinding(ctx context.Context, obj *rbac_authorization_k8s_io_v1.ClusterRoleBinding, opts ...client.CreateOption) error

	// Delete deletes the ClusterRoleBinding object.
	DeleteClusterRoleBinding(ctx context.Context, name string, opts ...client.DeleteOption) error

	// Update updates the given ClusterRoleBinding object.
	UpdateClusterRoleBinding(ctx context.Context, obj *rbac_authorization_k8s_io_v1.ClusterRoleBinding, opts ...client.UpdateOption) error

	// Patch patches the given ClusterRoleBinding object.
	PatchClusterRoleBinding(ctx context.Context, obj *rbac_authorization_k8s_io_v1.ClusterRoleBinding, patch client.Patch, opts ...client.PatchOption) error

	// DeleteAllOf deletes all ClusterRoleBinding objects matching the given options.
	DeleteAllOfClusterRoleBinding(ctx context.Context, opts ...client.DeleteAllOfOption) error

	// Create or Update the ClusterRoleBinding object.
	UpsertClusterRoleBinding(ctx context.Context, obj *rbac_authorization_k8s_io_v1.ClusterRoleBinding, transitionFuncs ...ClusterRoleBindingTransitionFunction) error
}

// StatusWriter knows how to update status subresource of a ClusterRoleBinding object.
type ClusterRoleBindingStatusWriter interface {
	// Update updates the fields corresponding to the status subresource for the
	// given ClusterRoleBinding object.
	UpdateClusterRoleBindingStatus(ctx context.Context, obj *rbac_authorization_k8s_io_v1.ClusterRoleBinding, opts ...client.UpdateOption) error

	// Patch patches the given ClusterRoleBinding object's subresource.
	PatchClusterRoleBindingStatus(ctx context.Context, obj *rbac_authorization_k8s_io_v1.ClusterRoleBinding, patch client.Patch, opts ...client.PatchOption) error
}

// Client knows how to perform CRUD operations on ClusterRoleBindings.
type ClusterRoleBindingClient interface {
	ClusterRoleBindingReader
	ClusterRoleBindingWriter
	ClusterRoleBindingStatusWriter
}

type clusterRoleBindingClient struct {
	client client.Client
}

func NewClusterRoleBindingClient(client client.Client) *clusterRoleBindingClient {
	return &clusterRoleBindingClient{client: client}
}

func (c *clusterRoleBindingClient) GetClusterRoleBinding(ctx context.Context, name string) (*rbac_authorization_k8s_io_v1.ClusterRoleBinding, error) {
	obj := &ClusterRoleBinding{}
	key := client.ObjectKey{
		Name: name,
	}
	if err := c.client.Get(ctx, key, obj); err != nil {
		return nil, err
	}
	return obj, nil
}

func (c *clusterRoleBindingClient) ListClusterRoleBinding(ctx context.Context, opts ...client.ListOption) (*rbac_authorization_k8s_io_v1.ClusterRoleBindingList, error) {
	list := &ClusterRoleBindingList{}
	if err := c.client.List(ctx, list, opts...); err != nil {
		return nil, err
	}
	return list, nil
}

func (c *clusterRoleBindingClient) CreateClusterRoleBinding(ctx context.Context, obj *rbac_authorization_k8s_io_v1.ClusterRoleBinding, opts ...client.CreateOption) error {
	return c.client.Create(ctx, obj, opts...)
}

func (c *clusterRoleBindingClient) DeleteClusterRoleBinding(ctx context.Context, name string, opts ...client.DeleteOption) error {
	obj := &ClusterRoleBinding{}
	obj.SetName(name)
	return c.client.Delete(ctx, obj, opts...)
}

func (c *clusterRoleBindingClient) UpdateClusterRoleBinding(ctx context.Context, obj *rbac_authorization_k8s_io_v1.ClusterRoleBinding, opts ...client.UpdateOption) error {
	return c.client.Update(ctx, obj, opts...)
}

func (c *clusterRoleBindingClient) PatchClusterRoleBinding(ctx context.Context, obj *rbac_authorization_k8s_io_v1.ClusterRoleBinding, patch client.Patch, opts ...client.PatchOption) error {
	return c.client.Patch(ctx, obj, patch, opts...)
}

func (c *clusterRoleBindingClient) DeleteAllOfClusterRoleBinding(ctx context.Context, opts ...client.DeleteAllOfOption) error {
	obj := &ClusterRoleBinding{}
	return c.client.DeleteAllOf(ctx, obj, opts...)
}

func (c *clusterRoleBindingClient) UpsertClusterRoleBinding(ctx context.Context, obj *rbac_authorization_k8s_io_v1.ClusterRoleBinding, transitionFuncs ...ClusterRoleBindingTransitionFunction) error {
	genericTxFunc := func(existing, desired runtime.Object) error {
		for _, txFunc := range transitionFuncs {
			if err := txFunc(existing.(*rbac_authorization_k8s_io_v1.ClusterRoleBinding), desired.(*rbac_authorization_k8s_io_v1.ClusterRoleBinding)); err != nil {
				return err
			}
		}
		return nil
	}
	_, err := controllerutils.Upsert(ctx, c.client, obj, genericTxFunc)
	return err
}

func (c *clusterRoleBindingClient) UpdateClusterRoleBindingStatus(ctx context.Context, obj *rbac_authorization_k8s_io_v1.ClusterRoleBinding, opts ...client.UpdateOption) error {
	return c.client.Status().Update(ctx, obj, opts...)
}

func (c *clusterRoleBindingClient) PatchClusterRoleBindingStatus(ctx context.Context, obj *rbac_authorization_k8s_io_v1.ClusterRoleBinding, patch client.Patch, opts ...client.PatchOption) error {
	return c.client.Status().Patch(ctx, obj, patch, opts...)
}

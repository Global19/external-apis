// Code generated by skv2. DO NOT EDIT.

//go:generate mockgen -source ./clients.go -destination mocks/clients.go

package v1

import (
	"context"

	"github.com/solo-io/skv2/pkg/controllerutils"
	"github.com/solo-io/skv2/pkg/multicluster"
	batch_v1 "k8s.io/api/batch/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/client-go/kubernetes/scheme"
	"k8s.io/client-go/rest"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

// MulticlusterClientset for the batch/v1 APIs
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

// clienset for the batch/v1 APIs
type Clientset interface {
	// clienset for the batch/v1/v1 APIs
	Jobs() JobClient
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

// clienset for the batch/v1/v1 APIs
func (c *clientSet) Jobs() JobClient {
	return NewJobClient(c.client)
}

// Reader knows how to read and list Jobs.
type JobReader interface {
	// Get retrieves a Job for the given object key
	GetJob(ctx context.Context, key client.ObjectKey) (*batch_v1.Job, error)

	// List retrieves list of Jobs for a given namespace and list options.
	ListJob(ctx context.Context, opts ...client.ListOption) (*batch_v1.JobList, error)
}

// JobTransitionFunction instructs the JobWriter how to transition between an existing
// Job object and a desired on an Upsert
type JobTransitionFunction func(existing, desired *batch_v1.Job) error

// Writer knows how to create, delete, and update Jobs.
type JobWriter interface {
	// Create saves the Job object.
	CreateJob(ctx context.Context, obj *batch_v1.Job, opts ...client.CreateOption) error

	// Delete deletes the Job object.
	DeleteJob(ctx context.Context, key client.ObjectKey, opts ...client.DeleteOption) error

	// Update updates the given Job object.
	UpdateJob(ctx context.Context, obj *batch_v1.Job, opts ...client.UpdateOption) error

	// Patch patches the given Job object.
	PatchJob(ctx context.Context, obj *batch_v1.Job, patch client.Patch, opts ...client.PatchOption) error

	// DeleteAllOf deletes all Job objects matching the given options.
	DeleteAllOfJob(ctx context.Context, opts ...client.DeleteAllOfOption) error

	// Create or Update the Job object.
	UpsertJob(ctx context.Context, obj *batch_v1.Job, transitionFuncs ...JobTransitionFunction) error
}

// StatusWriter knows how to update status subresource of a Job object.
type JobStatusWriter interface {
	// Update updates the fields corresponding to the status subresource for the
	// given Job object.
	UpdateJobStatus(ctx context.Context, obj *batch_v1.Job, opts ...client.UpdateOption) error

	// Patch patches the given Job object's subresource.
	PatchJobStatus(ctx context.Context, obj *batch_v1.Job, patch client.Patch, opts ...client.PatchOption) error
}

// Client knows how to perform CRUD operations on Jobs.
type JobClient interface {
	JobReader
	JobWriter
	JobStatusWriter
}

type jobClient struct {
	client client.Client
}

func NewJobClient(client client.Client) *jobClient {
	return &jobClient{client: client}
}

func (c *jobClient) GetJob(ctx context.Context, key client.ObjectKey) (*batch_v1.Job, error) {
	obj := &Job{}
	if err := c.client.Get(ctx, key, obj); err != nil {
		return nil, err
	}
	return obj, nil
}

func (c *jobClient) ListJob(ctx context.Context, opts ...client.ListOption) (*batch_v1.JobList, error) {
	list := &JobList{}
	if err := c.client.List(ctx, list, opts...); err != nil {
		return nil, err
	}
	return list, nil
}

func (c *jobClient) CreateJob(ctx context.Context, obj *batch_v1.Job, opts ...client.CreateOption) error {
	return c.client.Create(ctx, obj, opts...)
}

func (c *jobClient) DeleteJob(ctx context.Context, key client.ObjectKey, opts ...client.DeleteOption) error {
	obj := &Job{}
	obj.SetName(key.Name)
	obj.SetNamespace(key.Namespace)
	return c.client.Delete(ctx, obj, opts...)
}

func (c *jobClient) UpdateJob(ctx context.Context, obj *batch_v1.Job, opts ...client.UpdateOption) error {
	return c.client.Update(ctx, obj, opts...)
}

func (c *jobClient) PatchJob(ctx context.Context, obj *batch_v1.Job, patch client.Patch, opts ...client.PatchOption) error {
	return c.client.Patch(ctx, obj, patch, opts...)
}

func (c *jobClient) DeleteAllOfJob(ctx context.Context, opts ...client.DeleteAllOfOption) error {
	obj := &Job{}
	return c.client.DeleteAllOf(ctx, obj, opts...)
}

func (c *jobClient) UpsertJob(ctx context.Context, obj *batch_v1.Job, transitionFuncs ...JobTransitionFunction) error {
	genericTxFunc := func(existing, desired runtime.Object) error {
		for _, txFunc := range transitionFuncs {
			if err := txFunc(existing.(*batch_v1.Job), desired.(*batch_v1.Job)); err != nil {
				return err
			}
		}
		return nil
	}
	_, err := controllerutils.Upsert(ctx, c.client, obj, genericTxFunc)
	return err
}

func (c *jobClient) UpdateJobStatus(ctx context.Context, obj *batch_v1.Job, opts ...client.UpdateOption) error {
	return c.client.Status().Update(ctx, obj, opts...)
}

func (c *jobClient) PatchJobStatus(ctx context.Context, obj *batch_v1.Job, patch client.Patch, opts ...client.PatchOption) error {
	return c.client.Status().Patch(ctx, obj, patch, opts...)
}

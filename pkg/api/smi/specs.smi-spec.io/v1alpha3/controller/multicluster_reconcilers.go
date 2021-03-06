// Code generated by skv2. DO NOT EDIT.

//go:generate mockgen -source ./multicluster_reconcilers.go -destination mocks/multicluster_reconcilers.go

// Definitions for the multicluster Kubernetes Controllers
package controller

import (
	"context"

	specs_smi_spec_io_v1alpha3 "github.com/servicemeshinterface/smi-sdk-go/pkg/apis/specs/v1alpha3"

	"github.com/pkg/errors"
	"github.com/solo-io/skv2/pkg/ezkube"
	"github.com/solo-io/skv2/pkg/multicluster"
	mc_reconcile "github.com/solo-io/skv2/pkg/multicluster/reconcile"
	"github.com/solo-io/skv2/pkg/reconcile"
	"sigs.k8s.io/controller-runtime/pkg/predicate"
)

// Reconcile Upsert events for the HTTPRouteGroup Resource across clusters.
// implemented by the user
type MulticlusterHTTPRouteGroupReconciler interface {
	ReconcileHTTPRouteGroup(clusterName string, obj *specs_smi_spec_io_v1alpha3.HTTPRouteGroup) (reconcile.Result, error)
}

// Reconcile deletion events for the HTTPRouteGroup Resource across clusters.
// Deletion receives a reconcile.Request as we cannot guarantee the last state of the object
// before being deleted.
// implemented by the user
type MulticlusterHTTPRouteGroupDeletionReconciler interface {
	ReconcileHTTPRouteGroupDeletion(clusterName string, req reconcile.Request) error
}

type MulticlusterHTTPRouteGroupReconcilerFuncs struct {
	OnReconcileHTTPRouteGroup         func(clusterName string, obj *specs_smi_spec_io_v1alpha3.HTTPRouteGroup) (reconcile.Result, error)
	OnReconcileHTTPRouteGroupDeletion func(clusterName string, req reconcile.Request) error
}

func (f *MulticlusterHTTPRouteGroupReconcilerFuncs) ReconcileHTTPRouteGroup(clusterName string, obj *specs_smi_spec_io_v1alpha3.HTTPRouteGroup) (reconcile.Result, error) {
	if f.OnReconcileHTTPRouteGroup == nil {
		return reconcile.Result{}, nil
	}
	return f.OnReconcileHTTPRouteGroup(clusterName, obj)
}

func (f *MulticlusterHTTPRouteGroupReconcilerFuncs) ReconcileHTTPRouteGroupDeletion(clusterName string, req reconcile.Request) error {
	if f.OnReconcileHTTPRouteGroupDeletion == nil {
		return nil
	}
	return f.OnReconcileHTTPRouteGroupDeletion(clusterName, req)
}

type MulticlusterHTTPRouteGroupReconcileLoop interface {
	// AddMulticlusterHTTPRouteGroupReconciler adds a MulticlusterHTTPRouteGroupReconciler to the MulticlusterHTTPRouteGroupReconcileLoop.
	AddMulticlusterHTTPRouteGroupReconciler(ctx context.Context, rec MulticlusterHTTPRouteGroupReconciler, predicates ...predicate.Predicate)
}

type multiclusterHTTPRouteGroupReconcileLoop struct {
	loop multicluster.Loop
}

func (m *multiclusterHTTPRouteGroupReconcileLoop) AddMulticlusterHTTPRouteGroupReconciler(ctx context.Context, rec MulticlusterHTTPRouteGroupReconciler, predicates ...predicate.Predicate) {
	genericReconciler := genericHTTPRouteGroupMulticlusterReconciler{reconciler: rec}

	m.loop.AddReconciler(ctx, genericReconciler, predicates...)
}

func NewMulticlusterHTTPRouteGroupReconcileLoop(name string, cw multicluster.ClusterWatcher, options reconcile.Options) MulticlusterHTTPRouteGroupReconcileLoop {
	return &multiclusterHTTPRouteGroupReconcileLoop{loop: mc_reconcile.NewLoop(name, cw, &specs_smi_spec_io_v1alpha3.HTTPRouteGroup{}, options)}
}

type genericHTTPRouteGroupMulticlusterReconciler struct {
	reconciler MulticlusterHTTPRouteGroupReconciler
}

func (g genericHTTPRouteGroupMulticlusterReconciler) ReconcileDeletion(cluster string, req reconcile.Request) error {
	if deletionReconciler, ok := g.reconciler.(MulticlusterHTTPRouteGroupDeletionReconciler); ok {
		return deletionReconciler.ReconcileHTTPRouteGroupDeletion(cluster, req)
	}
	return nil
}

func (g genericHTTPRouteGroupMulticlusterReconciler) Reconcile(cluster string, object ezkube.Object) (reconcile.Result, error) {
	obj, ok := object.(*specs_smi_spec_io_v1alpha3.HTTPRouteGroup)
	if !ok {
		return reconcile.Result{}, errors.Errorf("internal error: HTTPRouteGroup handler received event for %T", object)
	}
	return g.reconciler.ReconcileHTTPRouteGroup(cluster, obj)
}

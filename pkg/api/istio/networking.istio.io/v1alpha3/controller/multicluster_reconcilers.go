// Code generated by skv2. DO NOT EDIT.

//go:generate mockgen -source ./multicluster_reconcilers.go -destination mocks/multicluster_reconcilers.go

// Definitions for the multicluster Kubernetes Controllers
package controller

import (
	"context"

	networking_istio_io_v1alpha3 "istio.io/client-go/pkg/apis/networking/v1alpha3"

	"github.com/pkg/errors"
	"github.com/solo-io/skv2/pkg/ezkube"
	"github.com/solo-io/skv2/pkg/multicluster"
	mc_reconcile "github.com/solo-io/skv2/pkg/multicluster/reconcile"
	"github.com/solo-io/skv2/pkg/reconcile"
	"sigs.k8s.io/controller-runtime/pkg/predicate"
)

// Reconcile Upsert events for the DestinationRule Resource across clusters.
// implemented by the user
type MulticlusterDestinationRuleReconciler interface {
	ReconcileDestinationRule(clusterName string, obj *networking_istio_io_v1alpha3.DestinationRule) (reconcile.Result, error)
}

// Reconcile deletion events for the DestinationRule Resource across clusters.
// Deletion receives a reconcile.Request as we cannot guarantee the last state of the object
// before being deleted.
// implemented by the user
type MulticlusterDestinationRuleDeletionReconciler interface {
	ReconcileDestinationRuleDeletion(clusterName string, req reconcile.Request)
}

type MulticlusterDestinationRuleReconcilerFuncs struct {
	OnReconcileDestinationRule         func(clusterName string, obj *networking_istio_io_v1alpha3.DestinationRule) (reconcile.Result, error)
	OnReconcileDestinationRuleDeletion func(clusterName string, req reconcile.Request)
}

func (f *MulticlusterDestinationRuleReconcilerFuncs) ReconcileDestinationRule(clusterName string, obj *networking_istio_io_v1alpha3.DestinationRule) (reconcile.Result, error) {
	if f.OnReconcileDestinationRule == nil {
		return reconcile.Result{}, nil
	}
	return f.OnReconcileDestinationRule(clusterName, obj)
}

func (f *MulticlusterDestinationRuleReconcilerFuncs) ReconcileDestinationRuleDeletion(clusterName string, req reconcile.Request) {
	if f.OnReconcileDestinationRuleDeletion == nil {
		return
	}
	f.OnReconcileDestinationRuleDeletion(clusterName, req)
}

type MulticlusterDestinationRuleReconcileLoop interface {
	// AddMulticlusterDestinationRuleReconciler adds a MulticlusterDestinationRuleReconciler to the MulticlusterDestinationRuleReconcileLoop.
	AddMulticlusterDestinationRuleReconciler(ctx context.Context, rec MulticlusterDestinationRuleReconciler, predicates ...predicate.Predicate)
}

type multiclusterDestinationRuleReconcileLoop struct {
	loop multicluster.Loop
}

func (m *multiclusterDestinationRuleReconcileLoop) AddMulticlusterDestinationRuleReconciler(ctx context.Context, rec MulticlusterDestinationRuleReconciler, predicates ...predicate.Predicate) {
	genericReconciler := genericDestinationRuleMulticlusterReconciler{reconciler: rec}

	m.loop.AddReconciler(ctx, genericReconciler, predicates...)
}

func NewMulticlusterDestinationRuleReconcileLoop(name string, cw multicluster.ClusterWatcher) MulticlusterDestinationRuleReconcileLoop {
	return &multiclusterDestinationRuleReconcileLoop{loop: mc_reconcile.NewLoop(name, cw, &networking_istio_io_v1alpha3.DestinationRule{})}
}

type genericDestinationRuleMulticlusterReconciler struct {
	reconciler MulticlusterDestinationRuleReconciler
}

func (g genericDestinationRuleMulticlusterReconciler) ReconcileDeletion(cluster string, req reconcile.Request) {
	if deletionReconciler, ok := g.reconciler.(MulticlusterDestinationRuleDeletionReconciler); ok {
		deletionReconciler.ReconcileDestinationRuleDeletion(cluster, req)
	}
}

func (g genericDestinationRuleMulticlusterReconciler) Reconcile(cluster string, object ezkube.Object) (reconcile.Result, error) {
	obj, ok := object.(*networking_istio_io_v1alpha3.DestinationRule)
	if !ok {
		return reconcile.Result{}, errors.Errorf("internal error: DestinationRule handler received event for %T", object)
	}
	return g.reconciler.ReconcileDestinationRule(cluster, obj)
}

// Reconcile Upsert events for the EnvoyFilter Resource across clusters.
// implemented by the user
type MulticlusterEnvoyFilterReconciler interface {
	ReconcileEnvoyFilter(clusterName string, obj *networking_istio_io_v1alpha3.EnvoyFilter) (reconcile.Result, error)
}

// Reconcile deletion events for the EnvoyFilter Resource across clusters.
// Deletion receives a reconcile.Request as we cannot guarantee the last state of the object
// before being deleted.
// implemented by the user
type MulticlusterEnvoyFilterDeletionReconciler interface {
	ReconcileEnvoyFilterDeletion(clusterName string, req reconcile.Request)
}

type MulticlusterEnvoyFilterReconcilerFuncs struct {
	OnReconcileEnvoyFilter         func(clusterName string, obj *networking_istio_io_v1alpha3.EnvoyFilter) (reconcile.Result, error)
	OnReconcileEnvoyFilterDeletion func(clusterName string, req reconcile.Request)
}

func (f *MulticlusterEnvoyFilterReconcilerFuncs) ReconcileEnvoyFilter(clusterName string, obj *networking_istio_io_v1alpha3.EnvoyFilter) (reconcile.Result, error) {
	if f.OnReconcileEnvoyFilter == nil {
		return reconcile.Result{}, nil
	}
	return f.OnReconcileEnvoyFilter(clusterName, obj)
}

func (f *MulticlusterEnvoyFilterReconcilerFuncs) ReconcileEnvoyFilterDeletion(clusterName string, req reconcile.Request) {
	if f.OnReconcileEnvoyFilterDeletion == nil {
		return
	}
	f.OnReconcileEnvoyFilterDeletion(clusterName, req)
}

type MulticlusterEnvoyFilterReconcileLoop interface {
	// AddMulticlusterEnvoyFilterReconciler adds a MulticlusterEnvoyFilterReconciler to the MulticlusterEnvoyFilterReconcileLoop.
	AddMulticlusterEnvoyFilterReconciler(ctx context.Context, rec MulticlusterEnvoyFilterReconciler, predicates ...predicate.Predicate)
}

type multiclusterEnvoyFilterReconcileLoop struct {
	loop multicluster.Loop
}

func (m *multiclusterEnvoyFilterReconcileLoop) AddMulticlusterEnvoyFilterReconciler(ctx context.Context, rec MulticlusterEnvoyFilterReconciler, predicates ...predicate.Predicate) {
	genericReconciler := genericEnvoyFilterMulticlusterReconciler{reconciler: rec}

	m.loop.AddReconciler(ctx, genericReconciler, predicates...)
}

func NewMulticlusterEnvoyFilterReconcileLoop(name string, cw multicluster.ClusterWatcher) MulticlusterEnvoyFilterReconcileLoop {
	return &multiclusterEnvoyFilterReconcileLoop{loop: mc_reconcile.NewLoop(name, cw, &networking_istio_io_v1alpha3.EnvoyFilter{})}
}

type genericEnvoyFilterMulticlusterReconciler struct {
	reconciler MulticlusterEnvoyFilterReconciler
}

func (g genericEnvoyFilterMulticlusterReconciler) ReconcileDeletion(cluster string, req reconcile.Request) {
	if deletionReconciler, ok := g.reconciler.(MulticlusterEnvoyFilterDeletionReconciler); ok {
		deletionReconciler.ReconcileEnvoyFilterDeletion(cluster, req)
	}
}

func (g genericEnvoyFilterMulticlusterReconciler) Reconcile(cluster string, object ezkube.Object) (reconcile.Result, error) {
	obj, ok := object.(*networking_istio_io_v1alpha3.EnvoyFilter)
	if !ok {
		return reconcile.Result{}, errors.Errorf("internal error: EnvoyFilter handler received event for %T", object)
	}
	return g.reconciler.ReconcileEnvoyFilter(cluster, obj)
}

// Reconcile Upsert events for the Gateway Resource across clusters.
// implemented by the user
type MulticlusterGatewayReconciler interface {
	ReconcileGateway(clusterName string, obj *networking_istio_io_v1alpha3.Gateway) (reconcile.Result, error)
}

// Reconcile deletion events for the Gateway Resource across clusters.
// Deletion receives a reconcile.Request as we cannot guarantee the last state of the object
// before being deleted.
// implemented by the user
type MulticlusterGatewayDeletionReconciler interface {
	ReconcileGatewayDeletion(clusterName string, req reconcile.Request)
}

type MulticlusterGatewayReconcilerFuncs struct {
	OnReconcileGateway         func(clusterName string, obj *networking_istio_io_v1alpha3.Gateway) (reconcile.Result, error)
	OnReconcileGatewayDeletion func(clusterName string, req reconcile.Request)
}

func (f *MulticlusterGatewayReconcilerFuncs) ReconcileGateway(clusterName string, obj *networking_istio_io_v1alpha3.Gateway) (reconcile.Result, error) {
	if f.OnReconcileGateway == nil {
		return reconcile.Result{}, nil
	}
	return f.OnReconcileGateway(clusterName, obj)
}

func (f *MulticlusterGatewayReconcilerFuncs) ReconcileGatewayDeletion(clusterName string, req reconcile.Request) {
	if f.OnReconcileGatewayDeletion == nil {
		return
	}
	f.OnReconcileGatewayDeletion(clusterName, req)
}

type MulticlusterGatewayReconcileLoop interface {
	// AddMulticlusterGatewayReconciler adds a MulticlusterGatewayReconciler to the MulticlusterGatewayReconcileLoop.
	AddMulticlusterGatewayReconciler(ctx context.Context, rec MulticlusterGatewayReconciler, predicates ...predicate.Predicate)
}

type multiclusterGatewayReconcileLoop struct {
	loop multicluster.Loop
}

func (m *multiclusterGatewayReconcileLoop) AddMulticlusterGatewayReconciler(ctx context.Context, rec MulticlusterGatewayReconciler, predicates ...predicate.Predicate) {
	genericReconciler := genericGatewayMulticlusterReconciler{reconciler: rec}

	m.loop.AddReconciler(ctx, genericReconciler, predicates...)
}

func NewMulticlusterGatewayReconcileLoop(name string, cw multicluster.ClusterWatcher) MulticlusterGatewayReconcileLoop {
	return &multiclusterGatewayReconcileLoop{loop: mc_reconcile.NewLoop(name, cw, &networking_istio_io_v1alpha3.Gateway{})}
}

type genericGatewayMulticlusterReconciler struct {
	reconciler MulticlusterGatewayReconciler
}

func (g genericGatewayMulticlusterReconciler) ReconcileDeletion(cluster string, req reconcile.Request) {
	if deletionReconciler, ok := g.reconciler.(MulticlusterGatewayDeletionReconciler); ok {
		deletionReconciler.ReconcileGatewayDeletion(cluster, req)
	}
}

func (g genericGatewayMulticlusterReconciler) Reconcile(cluster string, object ezkube.Object) (reconcile.Result, error) {
	obj, ok := object.(*networking_istio_io_v1alpha3.Gateway)
	if !ok {
		return reconcile.Result{}, errors.Errorf("internal error: Gateway handler received event for %T", object)
	}
	return g.reconciler.ReconcileGateway(cluster, obj)
}

// Reconcile Upsert events for the ServiceEntry Resource across clusters.
// implemented by the user
type MulticlusterServiceEntryReconciler interface {
	ReconcileServiceEntry(clusterName string, obj *networking_istio_io_v1alpha3.ServiceEntry) (reconcile.Result, error)
}

// Reconcile deletion events for the ServiceEntry Resource across clusters.
// Deletion receives a reconcile.Request as we cannot guarantee the last state of the object
// before being deleted.
// implemented by the user
type MulticlusterServiceEntryDeletionReconciler interface {
	ReconcileServiceEntryDeletion(clusterName string, req reconcile.Request)
}

type MulticlusterServiceEntryReconcilerFuncs struct {
	OnReconcileServiceEntry         func(clusterName string, obj *networking_istio_io_v1alpha3.ServiceEntry) (reconcile.Result, error)
	OnReconcileServiceEntryDeletion func(clusterName string, req reconcile.Request)
}

func (f *MulticlusterServiceEntryReconcilerFuncs) ReconcileServiceEntry(clusterName string, obj *networking_istio_io_v1alpha3.ServiceEntry) (reconcile.Result, error) {
	if f.OnReconcileServiceEntry == nil {
		return reconcile.Result{}, nil
	}
	return f.OnReconcileServiceEntry(clusterName, obj)
}

func (f *MulticlusterServiceEntryReconcilerFuncs) ReconcileServiceEntryDeletion(clusterName string, req reconcile.Request) {
	if f.OnReconcileServiceEntryDeletion == nil {
		return
	}
	f.OnReconcileServiceEntryDeletion(clusterName, req)
}

type MulticlusterServiceEntryReconcileLoop interface {
	// AddMulticlusterServiceEntryReconciler adds a MulticlusterServiceEntryReconciler to the MulticlusterServiceEntryReconcileLoop.
	AddMulticlusterServiceEntryReconciler(ctx context.Context, rec MulticlusterServiceEntryReconciler, predicates ...predicate.Predicate)
}

type multiclusterServiceEntryReconcileLoop struct {
	loop multicluster.Loop
}

func (m *multiclusterServiceEntryReconcileLoop) AddMulticlusterServiceEntryReconciler(ctx context.Context, rec MulticlusterServiceEntryReconciler, predicates ...predicate.Predicate) {
	genericReconciler := genericServiceEntryMulticlusterReconciler{reconciler: rec}

	m.loop.AddReconciler(ctx, genericReconciler, predicates...)
}

func NewMulticlusterServiceEntryReconcileLoop(name string, cw multicluster.ClusterWatcher) MulticlusterServiceEntryReconcileLoop {
	return &multiclusterServiceEntryReconcileLoop{loop: mc_reconcile.NewLoop(name, cw, &networking_istio_io_v1alpha3.ServiceEntry{})}
}

type genericServiceEntryMulticlusterReconciler struct {
	reconciler MulticlusterServiceEntryReconciler
}

func (g genericServiceEntryMulticlusterReconciler) ReconcileDeletion(cluster string, req reconcile.Request) {
	if deletionReconciler, ok := g.reconciler.(MulticlusterServiceEntryDeletionReconciler); ok {
		deletionReconciler.ReconcileServiceEntryDeletion(cluster, req)
	}
}

func (g genericServiceEntryMulticlusterReconciler) Reconcile(cluster string, object ezkube.Object) (reconcile.Result, error) {
	obj, ok := object.(*networking_istio_io_v1alpha3.ServiceEntry)
	if !ok {
		return reconcile.Result{}, errors.Errorf("internal error: ServiceEntry handler received event for %T", object)
	}
	return g.reconciler.ReconcileServiceEntry(cluster, obj)
}

// Reconcile Upsert events for the VirtualService Resource across clusters.
// implemented by the user
type MulticlusterVirtualServiceReconciler interface {
	ReconcileVirtualService(clusterName string, obj *networking_istio_io_v1alpha3.VirtualService) (reconcile.Result, error)
}

// Reconcile deletion events for the VirtualService Resource across clusters.
// Deletion receives a reconcile.Request as we cannot guarantee the last state of the object
// before being deleted.
// implemented by the user
type MulticlusterVirtualServiceDeletionReconciler interface {
	ReconcileVirtualServiceDeletion(clusterName string, req reconcile.Request)
}

type MulticlusterVirtualServiceReconcilerFuncs struct {
	OnReconcileVirtualService         func(clusterName string, obj *networking_istio_io_v1alpha3.VirtualService) (reconcile.Result, error)
	OnReconcileVirtualServiceDeletion func(clusterName string, req reconcile.Request)
}

func (f *MulticlusterVirtualServiceReconcilerFuncs) ReconcileVirtualService(clusterName string, obj *networking_istio_io_v1alpha3.VirtualService) (reconcile.Result, error) {
	if f.OnReconcileVirtualService == nil {
		return reconcile.Result{}, nil
	}
	return f.OnReconcileVirtualService(clusterName, obj)
}

func (f *MulticlusterVirtualServiceReconcilerFuncs) ReconcileVirtualServiceDeletion(clusterName string, req reconcile.Request) {
	if f.OnReconcileVirtualServiceDeletion == nil {
		return
	}
	f.OnReconcileVirtualServiceDeletion(clusterName, req)
}

type MulticlusterVirtualServiceReconcileLoop interface {
	// AddMulticlusterVirtualServiceReconciler adds a MulticlusterVirtualServiceReconciler to the MulticlusterVirtualServiceReconcileLoop.
	AddMulticlusterVirtualServiceReconciler(ctx context.Context, rec MulticlusterVirtualServiceReconciler, predicates ...predicate.Predicate)
}

type multiclusterVirtualServiceReconcileLoop struct {
	loop multicluster.Loop
}

func (m *multiclusterVirtualServiceReconcileLoop) AddMulticlusterVirtualServiceReconciler(ctx context.Context, rec MulticlusterVirtualServiceReconciler, predicates ...predicate.Predicate) {
	genericReconciler := genericVirtualServiceMulticlusterReconciler{reconciler: rec}

	m.loop.AddReconciler(ctx, genericReconciler, predicates...)
}

func NewMulticlusterVirtualServiceReconcileLoop(name string, cw multicluster.ClusterWatcher) MulticlusterVirtualServiceReconcileLoop {
	return &multiclusterVirtualServiceReconcileLoop{loop: mc_reconcile.NewLoop(name, cw, &networking_istio_io_v1alpha3.VirtualService{})}
}

type genericVirtualServiceMulticlusterReconciler struct {
	reconciler MulticlusterVirtualServiceReconciler
}

func (g genericVirtualServiceMulticlusterReconciler) ReconcileDeletion(cluster string, req reconcile.Request) {
	if deletionReconciler, ok := g.reconciler.(MulticlusterVirtualServiceDeletionReconciler); ok {
		deletionReconciler.ReconcileVirtualServiceDeletion(cluster, req)
	}
}

func (g genericVirtualServiceMulticlusterReconciler) Reconcile(cluster string, object ezkube.Object) (reconcile.Result, error) {
	obj, ok := object.(*networking_istio_io_v1alpha3.VirtualService)
	if !ok {
		return reconcile.Result{}, errors.Errorf("internal error: VirtualService handler received event for %T", object)
	}
	return g.reconciler.ReconcileVirtualService(cluster, obj)
}

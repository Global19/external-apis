// Code generated by skv2. DO NOT EDIT.

//go:generate mockgen -source ./multicluster_reconcilers.go -destination mocks/multicluster_reconcilers.go

// Definitions for the multicluster Kubernetes Controllers
package controller

import (
	"context"

	admissionregistration_k8s_io_v1 "k8s.io/api/admissionregistration/v1"

	"github.com/pkg/errors"
	"github.com/solo-io/skv2/pkg/ezkube"
	"github.com/solo-io/skv2/pkg/multicluster"
	mc_reconcile "github.com/solo-io/skv2/pkg/multicluster/reconcile"
	"github.com/solo-io/skv2/pkg/reconcile"
	"sigs.k8s.io/controller-runtime/pkg/predicate"
)

// Reconcile Upsert events for the ValidatingWebhookConfiguration Resource across clusters.
// implemented by the user
type MulticlusterValidatingWebhookConfigurationReconciler interface {
	ReconcileValidatingWebhookConfiguration(clusterName string, obj *admissionregistration_k8s_io_v1.ValidatingWebhookConfiguration) (reconcile.Result, error)
}

// Reconcile deletion events for the ValidatingWebhookConfiguration Resource across clusters.
// Deletion receives a reconcile.Request as we cannot guarantee the last state of the object
// before being deleted.
// implemented by the user
type MulticlusterValidatingWebhookConfigurationDeletionReconciler interface {
	ReconcileValidatingWebhookConfigurationDeletion(clusterName string, req reconcile.Request)
}

type MulticlusterValidatingWebhookConfigurationReconcilerFuncs struct {
	OnReconcileValidatingWebhookConfiguration         func(clusterName string, obj *admissionregistration_k8s_io_v1.ValidatingWebhookConfiguration) (reconcile.Result, error)
	OnReconcileValidatingWebhookConfigurationDeletion func(clusterName string, req reconcile.Request)
}

func (f *MulticlusterValidatingWebhookConfigurationReconcilerFuncs) ReconcileValidatingWebhookConfiguration(clusterName string, obj *admissionregistration_k8s_io_v1.ValidatingWebhookConfiguration) (reconcile.Result, error) {
	if f.OnReconcileValidatingWebhookConfiguration == nil {
		return reconcile.Result{}, nil
	}
	return f.OnReconcileValidatingWebhookConfiguration(clusterName, obj)
}

func (f *MulticlusterValidatingWebhookConfigurationReconcilerFuncs) ReconcileValidatingWebhookConfigurationDeletion(clusterName string, req reconcile.Request) {
	if f.OnReconcileValidatingWebhookConfigurationDeletion == nil {
		return
	}
	f.OnReconcileValidatingWebhookConfigurationDeletion(clusterName, req)
}

type MulticlusterValidatingWebhookConfigurationReconcileLoop interface {
	// AddMulticlusterValidatingWebhookConfigurationReconciler adds a MulticlusterValidatingWebhookConfigurationReconciler to the MulticlusterValidatingWebhookConfigurationReconcileLoop.
	AddMulticlusterValidatingWebhookConfigurationReconciler(ctx context.Context, rec MulticlusterValidatingWebhookConfigurationReconciler, predicates ...predicate.Predicate)
}

type multiclusterValidatingWebhookConfigurationReconcileLoop struct {
	loop multicluster.Loop
}

func (m *multiclusterValidatingWebhookConfigurationReconcileLoop) AddMulticlusterValidatingWebhookConfigurationReconciler(ctx context.Context, rec MulticlusterValidatingWebhookConfigurationReconciler, predicates ...predicate.Predicate) {
	genericReconciler := genericValidatingWebhookConfigurationMulticlusterReconciler{reconciler: rec}

	m.loop.AddReconciler(ctx, genericReconciler, predicates...)
}

func NewMulticlusterValidatingWebhookConfigurationReconcileLoop(name string, cw multicluster.ClusterWatcher) MulticlusterValidatingWebhookConfigurationReconcileLoop {
	return &multiclusterValidatingWebhookConfigurationReconcileLoop{loop: mc_reconcile.NewLoop(name, cw, &admissionregistration_k8s_io_v1.ValidatingWebhookConfiguration{})}
}

type genericValidatingWebhookConfigurationMulticlusterReconciler struct {
	reconciler MulticlusterValidatingWebhookConfigurationReconciler
}

func (g genericValidatingWebhookConfigurationMulticlusterReconciler) ReconcileDeletion(cluster string, req reconcile.Request) {
	if deletionReconciler, ok := g.reconciler.(MulticlusterValidatingWebhookConfigurationDeletionReconciler); ok {
		deletionReconciler.ReconcileValidatingWebhookConfigurationDeletion(cluster, req)
	}
}

func (g genericValidatingWebhookConfigurationMulticlusterReconciler) Reconcile(cluster string, object ezkube.Object) (reconcile.Result, error) {
	obj, ok := object.(*admissionregistration_k8s_io_v1.ValidatingWebhookConfiguration)
	if !ok {
		return reconcile.Result{}, errors.Errorf("internal error: ValidatingWebhookConfiguration handler received event for %T", object)
	}
	return g.reconciler.ReconcileValidatingWebhookConfiguration(cluster, obj)
}

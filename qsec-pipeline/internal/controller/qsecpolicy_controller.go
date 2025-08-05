package controller

import (
	"context"

	"k8s.io/apimachinery/pkg/runtime"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/log"

	qsecv1alpha1 "github.com/qsec-pipeline/qsec-pipeline/api/v1alpha1"
)

type QSecPolicyReconciler struct {
	client.Client
	Scheme *runtime.Scheme
}

func (r *QSecPolicyReconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
	log := log.FromContext(ctx)
	var policy qsecv1alpha1.QSecPolicy
	if err := r.Get(ctx, req.NamespacedName, &policy); err != nil {
		return ctrl.Result{}, client.IgnoreNotFound(err)
	}

	// TODO: ensure DaemonSet, update status
	policy.Status.Ready = true
	if err := r.Status().Update(ctx, &policy); err != nil {
		return ctrl.Result{}, err
	}
	log.Info("Policy reconciled", "policy", policy.Name)
	return ctrl.Result{}, nil
}

func (r *QSecPolicyReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&qsecv1alpha1.QSecPolicy{}).
		Complete(r)
}
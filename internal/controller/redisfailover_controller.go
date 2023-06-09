/*
Copyright 2023.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package controller

import (
	"context"
	rfservice "github.com/shshang/redis-operator/operator/redisfailover/service"

	"k8s.io/apimachinery/pkg/runtime"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/log"

	"github.com/shshang/redis-operator/metrics"
	"github.com/shshang/redis-operator/service/k8s"
	databasesv1 "shshang/redis-operator/api/v1"
)

// RedisFailoverReconciler reconciles a RedisFailover object
// RedisFailoverReconciler is very much like RedisFailoverHandler

type RedisFailoverReconciler struct {
	client.Client
	Scheme *runtime.Scheme

	config     Config
	k8sservice k8s.Service
	rfService  rfservice.RedisFailoverClient
	rfChecker  rfservice.RedisFailoverCheck
	rfHealer   rfservice.RedisFailoverHeal
	mClient    metrics.Recorder
	logger     log.Logger
}

//+kubebuilder:rbac:groups=databases.spotahome.com,resources=redisfailovers,verbs=get;list;watch;create;update;patch;delete
//+kubebuilder:rbac:groups=databases.spotahome.com,resources=redisfailovers/status,verbs=get;update;patch
//+kubebuilder:rbac:groups=databases.spotahome.com,resources=redisfailovers/finalizers,verbs=update

// Reconcile is part of the main kubernetes reconciliation loop which aims to
// move the current state of the cluster closer to the desired state.
// TODO(user): Modify the Reconcile function to compare the state specified by
// the RedisFailover object against the actual cluster state, and then
// perform operations to make the cluster state reflect the state specified by
// the user.
//
// For more details, check Reconcile and its Result here:
// - https://pkg.go.dev/sigs.k8s.io/controller-runtime@v0.14.4/pkg/reconcile
func (r *RedisFailoverReconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
	log := log.FromContext(ctx)

	// TODO(user): your logic here
	var redisFailover databasesv1.RedisFailover
	if err := r.Get(ctx, req.NamespacedName, &redisFailover); err != nil {
		log.Error(err, "unable to fetch CronJob")
		// we'll ignore not-found errors, since they can't be fixed by an immediate
		// requeue (we'll need to wait for a new notification), and we can get them
		// on deleted requests.
		return ctrl.Result{}, client.IgnoreNotFound(err)
	}

	// Is this check still valid?
	//if err := rf.Validate(); err != nil {
	//	r.mClient.SetClusterError(rf.Namespace, rf.Name)
	//	return err
	//}

	// Create owner refs so the objects manager by this handler have ownership to the
	// received RF.
	oRefs := r.createOwnerReferences(rf)

	// Create the labels every object derived from this need to have.
	labels := r.getLabels(rf)

	if err := r.Ensure(rf, labels, oRefs, r.mClient); err != nil {
		r.mClient.SetClusterError(rf.Namespace, rf.Name)
		return ctrl.Result{}, err
	}

	if err := r.CheckAndHeal(rf); err != nil {
		r.mClient.SetClusterError(rf.Namespace, rf.Name)
		return ctrl.Result{}, err
	}

	r.mClient.SetClusterOK(rf.Namespace, rf.Name)
	return ctrl.Result{}, nil

	}

	//return ctrl.Result{}, nil
}

// SetupWithManager sets up the controller with the Manager.
func (r *RedisFailoverReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&databasesv1.RedisFailover{}).
		Complete(r)
}

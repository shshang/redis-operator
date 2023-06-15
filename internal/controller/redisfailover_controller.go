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
	"github.com/shshang/redis-operator/log"
	"k8s.io/apimachinery/pkg/runtime"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	//"sigs.k8s.io/controller-runtime/pkg/log"

	databasesv1 "github.com/shshang/redis-operator/api/v1"
)

// RedisFailoverReconciler reconciles a RedisFailover object
// RedisFailoverReconciler is very much like RedisFailoverHandler

type RedisFailoverReconciler struct {
	client.Client
	Scheme *runtime.Scheme
	log    log.Logger
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
	//log := r.log.FromContext(ctx, r.logger)

	r.logger.Errorf("Unable to compile label whitelist regex '%s', ignoring it.", regex)

}

// SetupWithManager sets up the controller with the Manager.
func (r *RedisFailoverReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&databasesv1.RedisFailover{}).
		Complete(r)
}

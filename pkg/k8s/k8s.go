package k8s

import (
	"github.com/shshang/redis-operator/internal/controller/service"
)

// Service is the K8s service entrypoint.
type Services interface {
	ConfigMap
	Secret
	Pod
	PodDisruptionBudget
	service.RedisFailover
	Service
	RBAC
	Deployment
	StatefulSet
}

type services struct {
	ConfigMap
	Secret
	Pod
	PodDisruptionBudget
	service.RedisFailover
	Service
	RBAC
	Deployment
	StatefulSet
}

// New returns a new Kubernetes service.
//func New(kubecli kubernetes.Interface, crdcli redisfailoverclientset.Interface, apiextcli apiextensionscli.Interface, logger log.Logger, metricsRecorder metrics.Recorder) Services {
//	return &services{
//		ConfigMap:           NewConfigMapService(kubecli, logger, metricsRecorder),
//		Secret:              NewSecretService(kubecli, logger, metricsRecorder),
//		Pod:                 NewPodService(kubecli, logger, metricsRecorder),
//		PodDisruptionBudget: NewPodDisruptionBudgetService(kubecli, logger, metricsRecorder),
//		RedisFailover:       service.NewRedisFailoverService(crdcli, logger, metricsRecorder),
//		Service:             NewServiceService(kubecli, logger, metricsRecorder),
//		RBAC:                NewRBACService(kubecli, logger, metricsRecorder),
//		Deployment:          NewDeploymentService(kubecli, logger, metricsRecorder),
//		StatefulSet:         NewStatefulSetService(kubecli, logger, metricsRecorder),
//	}
//}

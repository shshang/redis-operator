package reconciliation

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	redisfailoverv1 "github.com/shshang/redis-operator/api/v1"
	"github.com/shshang/redis-operator/metrics"
)

// Ensure is called to ensure all of the resources associated with a RedisFailover are created
func (r *RedisFailoverHandler) Ensure(rf *redisfailoverv1.RedisFailover, labels map[string]string, or []metav1.OwnerReference, metricsClient metrics.Recorder) error {
	if rf.Spec.Redis.Exporter.Enabled {
		if err := r.rfService.EnsureRedisService(rf, labels, or); err != nil {
			return err
		}
	} else {
		if err := r.rfService.EnsureNotPresentRedisService(rf); err != nil {
			return err
		}
	}

	sentinelsAllowed := rf.SentinelsAllowed()
	if sentinelsAllowed {
		if err := r.rfService.EnsureSentinelService(rf, labels, or); err != nil {
			return err
		}
		if err := r.rfService.EnsureSentinelConfigMap(rf, labels, or); err != nil {
			return err
		}
	}

	if err := r.rfService.EnsureRedisShutdownConfigMap(rf, labels, or); err != nil {
		return err
	}
	if err := r.rfService.EnsureRedisReadinessConfigMap(rf, labels, or); err != nil {
		return err
	}
	if err := r.rfService.EnsureRedisConfigMap(rf, labels, or); err != nil {
		return err
	}
	if err := r.rfService.EnsureRedisStatefulset(rf, labels, or); err != nil {
		return err
	}

	if sentinelsAllowed {
		if err := r.rfService.EnsureSentinelDeployment(rf, labels, or); err != nil {
			return err
		}
	}

	return nil
}

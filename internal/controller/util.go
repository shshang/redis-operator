package controller

import (
	"github.com/shshang/redis-operator/internal/controller/util"
	"regexp"

	redisfailoverv1 "github.com/shshang/redis-operator/api/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

const (
	rfLabelManagedByKey = "app.kubernetes.io/managed-by"
	rfLabelNameKey      = "redisfailovers.databases.spotahome.com/name"
)

var (
	defaultLabels = map[string]string{
		rfLabelManagedByKey: "operatorName",
		//rfLabelManagedByKey: operatorName,
	}
)

// getLabels merges the labels (dynamic and operator static ones).
func (r *RedisFailoverReconciler) getLabels(rf *redisfailoverv1.RedisFailover) map[string]string {
	dynLabels := map[string]string{
		rfLabelNameKey: rf.Name,
	}

	// Filter the labels based on the whitelist
	filteredCustomLabels := make(map[string]string)
	if rf.Spec.LabelWhitelist != nil && len(rf.Spec.LabelWhitelist) != 0 {
		for _, regex := range rf.Spec.LabelWhitelist {
			compiledRegexp, err := regexp.Compile(regex)
			if err != nil {
				r.logger.Errorf("Unable to compile label whitelist regex '%s', ignoring it.", regex)
				continue
			}
			for labelKey, labelValue := range rf.Labels {
				if match := compiledRegexp.MatchString(labelKey); match {
					filteredCustomLabels[labelKey] = labelValue
				}
			}
		}
	} else {
		// If no whitelist is specified then don't filter the labels.
		filteredCustomLabels = rf.Labels
	}
	return util.MergeLabels(defaultLabels, dynLabels, filteredCustomLabels)
}

func (r *RedisFailoverReconciler) createOwnerReferences(rf *redisfailoverv1.RedisFailover) []metav1.OwnerReference {
	rfvk := redisfailoverv1.VersionKind(redisfailoverv1.RFKind)
	return []metav1.OwnerReference{
		*metav1.NewControllerRef(rf, rfvk),
	}
}

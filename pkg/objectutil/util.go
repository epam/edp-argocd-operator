package objectutil

import metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

func ContainsLabel(object metav1.Object, label, labelVal string) bool {
	val, ok := object.GetLabels()[label]

	return ok && val == labelVal
}

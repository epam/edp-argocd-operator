package objectutil

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

func ContainsLabel(object metav1.Object, label, labelVal string) bool {
	val, ok := object.GetLabels()[label]

	return ok && val == labelVal
}

// ObjectMarkDeleted checks if the object is marked for cleaning up by garbage collector.
func ObjectMarkDeleted(object client.Object) bool {
	return len(object.GetFinalizers()) == 0 && object.GetDeletionTimestamp() != nil
}

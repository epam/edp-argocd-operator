package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

//+kubebuilder:object:root=true
//+kubebuilder:subresource:status

// ArgoApplication is the Schema for the argoapplications API.
type ArgoApplication struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty" mapstructure:",squash"`

	Spec   ApplicationSpec   `json:"spec,omitempty"`
	Status ApplicationStatus `json:"status,omitempty"`
}

// ApplicationStatus contains status information for the application.
type ApplicationStatus struct {
	// Status contains the result of reconciliation
	// +optional
	Status ApplicationReconciliationStatus `json:"status,omitempty"`

	// Error show detailed error message if reconciliation fails
	// +optional
	Error string `json:"error,omitempty"`
}

type ApplicationReconciliationStatus string

const (
	ApplicationReconciliationStatusSuccess ApplicationReconciliationStatus = "success"
	ApplicationReconciliationStatusError   ApplicationReconciliationStatus = "error"
)

//+kubebuilder:object:root=true

// ArgoApplicationList contains a list of ArgoApplication.
type ArgoApplicationList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ArgoApplication `json:"items"`
}

func init() {
	SchemeBuilder.Register(&ArgoApplication{}, &ArgoApplicationList{})
}

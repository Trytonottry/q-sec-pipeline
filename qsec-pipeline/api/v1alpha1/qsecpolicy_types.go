package v1alpha1

import metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

// +kubebuilder:object:root=true
type QSecPolicy struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   QSecPolicySpec   `json:"spec"`
	Status QSecPolicyStatus `json:"status,omitempty"`
}

type QSecPolicySpec struct {
	Registry      string `json:"registry"`
	KyberLevel    string `json:"kyberLevel,omitempty"`
	DilithiumLevel string `json:"dilithiumLevel,omitempty"`
	Immutable     bool   `json:"immutable"`
}

type QSecPolicyStatus struct {
	Ready bool   `json:"ready"`
	Hash  string `json:"hash,omitempty"`
}

// +kubebuilder:object:root=true
type QSecPolicyList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []QSecPolicy `json:"items"`
}
package v2

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
type Persona struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   PersonaSpec   `json:"spec,omitempty"`
	Status PersonaStatus `json:"status,omitempty"`
}

// PersonaSpec defines the desired state of Persona
type PersonaSpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "make" to regenerate code after modifying this file
	Name string `json:"name"`
	Age  int    `json:"age"`
}

// PersonaStatus defines the observed state of Persona
type PersonaStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "make" to regenerate code after modifying this file
	State      string `json:"state"`
	Allowed    bool   `json:"allowed"`
	ExpireDate string `json:"expireDate"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// PersonaList contains a list of Persona
type PersonaList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Persona `json:"items"`
}


package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// NetCoreConfigManagementSpec defines the desired state of NetCoreConfigManagement
// +k8s:openapi-gen=true
type NetCoreConfigManagementSpec struct {
	LinkedDeployments []string `json:"linkedDeployments,required"`
	Config            struct {
		ConfigMapName  string `json:"configMapName,required"`
		ConfigMapKey   string `json:"configMapKey,required"`
		ConfigMapValue string `json:"configMapValue,required"`
	} `json:"config"`
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "operator-sdk generate k8s" to regenerate code after modifying this file
	// Add custom validation using kubebuilder tags: https://book-v1.book.kubebuilder.io/beyond_basics/generating_crd.html
}

// NetCoreConfigManagementStatus defines the observed state of NetCoreConfigManagement
// +k8s:openapi-gen=true
type NetCoreConfigManagementStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "operator-sdk generate k8s" to regenerate code after modifying this file
	// Add custom validation using kubebuilder tags: https://book-v1.book.kubebuilder.io/beyond_basics/generating_crd.html
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// NetCoreConfigManagement is the Schema for the netcoreconfigmanagements API
// +k8s:openapi-gen=true
// +kubebuilder:subresource:status
// +kubebuilder:resource:path=netcoreconfigmanagements,scope=Namespaced
type NetCoreConfigManagement struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   NetCoreConfigManagementSpec   `json:"spec,omitempty"`
	Status NetCoreConfigManagementStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// NetCoreConfigManagementList contains a list of NetCoreConfigManagement
type NetCoreConfigManagementList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []NetCoreConfigManagement `json:"items"`
}

func init() {
	SchemeBuilder.Register(&NetCoreConfigManagement{}, &NetCoreConfigManagementList{})
}

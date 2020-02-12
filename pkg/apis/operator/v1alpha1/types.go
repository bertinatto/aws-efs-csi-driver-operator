package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// EFSCSIDriver is a specification for a EFSCSIDriver resource
type EFSCSIDriver struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   EFSCSIDriverSpec   `json:"spec"`
	Status EFSCSIDriverStatus `json:"status"`
}

// EFSCSIDriverSpec is the spec for a EFSCSIDriver resource
type EFSCSIDriverSpec struct {
	DeploymentName string `json:"deploymentName"`
	Replicas       *int32 `json:"replicas"`
}

// EFSCSIDriverStatus is the status for a EFSCSIDriver resource
type EFSCSIDriverStatus struct {
	AvailableReplicas int32 `json:"availableReplicas"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// EFSCSIDriverList is a list of EFSCSIDriver resources
type EFSCSIDriverList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata"`

	Items []EFSCSIDriver `json:"items"`
}

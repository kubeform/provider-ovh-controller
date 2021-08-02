/*
Copyright AppsCode Inc. and Contributors

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Code generated by Kubeform. DO NOT EDIT.

package v1alpha1

import (
	base "kubeform.dev/apimachinery/api/v1alpha1"

	core "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	kmapi "kmodules.xyz/client-go/api/v1"
	"sigs.k8s.io/cli-utils/pkg/kstatus/status"
)

// +genclient
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:printcolumn:name="Phase",type=string,JSONPath=`.status.phase`

type VrackNetwork struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              VrackNetworkSpec   `json:"spec,omitempty"`
	Status            VrackNetworkStatus `json:"status,omitempty"`
}

type VrackNetworkSpec struct {
	State *VrackNetworkSpecResource `json:"state,omitempty" tf:"-"`

	Resource VrackNetworkSpecResource `json:"resource" tf:"resource"`

	UpdatePolicy base.UpdatePolicy `json:"updatePolicy,omitempty" tf:"-"`

	TerminationPolicy base.TerminationPolicy `json:"terminationPolicy,omitempty" tf:"-"`

	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`
}

type VrackNetworkSpecResource struct {
	ID string `json:"id,omitempty" tf:"id,omitempty"`

	// Human readable name for your vrack network
	// +optional
	DisplayName *string `json:"displayName,omitempty" tf:"display_name"`
	// This attribute is there for documentation purpose only and isnt passed to the OVH API as it may conflicts with http/tcp farms `vrack_network_id` attribute
	// +optional
	FarmID []int64 `json:"farmID,omitempty" tf:"farm_id"`
	// An IP block used as a pool of IPs by this Load Balancer to connect to the servers in this private network. The blck must be in the private network and reserved for the Load Balancer
	NatIP *string `json:"natIP" tf:"nat_ip"`
	// The internal name of your IPloadbalancer
	ServiceName *string `json:"serviceName" tf:"service_name"`
	// IP block of the private network in the vRack
	Subnet *string `json:"subnet" tf:"subnet"`
	// VLAN of the private network in the vRack. 0 if the private network is not in a VLAN
	// +optional
	Vlan *int64 `json:"vlan,omitempty" tf:"vlan"`
	// Internal Load Balancer identifier of the vRack private network
	// +optional
	VrackNetworkID *int64 `json:"vrackNetworkID,omitempty" tf:"vrack_network_id"`
}

type VrackNetworkStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`
	// +optional
	Phase status.Status `json:"phase,omitempty"`
	// +optional
	Conditions []kmapi.Condition `json:"conditions,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// VrackNetworkList is a list of VrackNetworks
type VrackNetworkList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of VrackNetwork CRD objects
	Items []VrackNetwork `json:"items,omitempty"`
}
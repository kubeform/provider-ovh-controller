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

type HttpFrontend struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              HttpFrontendSpec   `json:"spec,omitempty"`
	Status            HttpFrontendStatus `json:"status,omitempty"`
}

type HttpFrontendSpec struct {
	State *HttpFrontendSpecResource `json:"state,omitempty" tf:"-"`

	Resource HttpFrontendSpecResource `json:"resource" tf:"resource"`

	UpdatePolicy base.UpdatePolicy `json:"updatePolicy,omitempty" tf:"-"`

	TerminationPolicy base.TerminationPolicy `json:"terminationPolicy,omitempty" tf:"-"`

	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`
}

type HttpFrontendSpecResource struct {
	ID string `json:"id,omitempty" tf:"id,omitempty"`

	// +optional
	AllowedSource []string `json:"allowedSource,omitempty" tf:"allowed_source"`
	// +optional
	DedicatedIpfo []string `json:"dedicatedIpfo,omitempty" tf:"dedicated_ipfo"`
	// +optional
	DefaultFarmID *int64 `json:"defaultFarmID,omitempty" tf:"default_farm_id"`
	// +optional
	DefaultSslID *int64 `json:"defaultSslID,omitempty" tf:"default_ssl_id"`
	// +optional
	Disabled *bool `json:"disabled,omitempty" tf:"disabled"`
	// +optional
	DisplayName *string `json:"displayName,omitempty" tf:"display_name"`
	Port        *string `json:"port" tf:"port"`
	// +optional
	RedirectLocation *string `json:"redirectLocation,omitempty" tf:"redirect_location"`
	ServiceName      *string `json:"serviceName" tf:"service_name"`
	// +optional
	Ssl  *bool   `json:"ssl,omitempty" tf:"ssl"`
	Zone *string `json:"zone" tf:"zone"`
}

type HttpFrontendStatus struct {
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

// HttpFrontendList is a list of HttpFrontends
type HttpFrontendList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of HttpFrontend CRD objects
	Items []HttpFrontend `json:"items,omitempty"`
}
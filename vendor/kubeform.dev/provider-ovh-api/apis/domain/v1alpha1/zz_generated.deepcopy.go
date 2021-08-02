// +build !ignore_autogenerated

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

// Code generated by deepcopy-gen. DO NOT EDIT.

package v1alpha1

import (
	runtime "k8s.io/apimachinery/pkg/runtime"
	v1 "kmodules.xyz/client-go/api/v1"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Zone) DeepCopyInto(out *Zone) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Zone.
func (in *Zone) DeepCopy() *Zone {
	if in == nil {
		return nil
	}
	out := new(Zone)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Zone) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ZoneList) DeepCopyInto(out *ZoneList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Zone, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ZoneList.
func (in *ZoneList) DeepCopy() *ZoneList {
	if in == nil {
		return nil
	}
	out := new(ZoneList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ZoneList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ZoneRecord) DeepCopyInto(out *ZoneRecord) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ZoneRecord.
func (in *ZoneRecord) DeepCopy() *ZoneRecord {
	if in == nil {
		return nil
	}
	out := new(ZoneRecord)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ZoneRecord) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ZoneRecordList) DeepCopyInto(out *ZoneRecordList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]ZoneRecord, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ZoneRecordList.
func (in *ZoneRecordList) DeepCopy() *ZoneRecordList {
	if in == nil {
		return nil
	}
	out := new(ZoneRecordList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ZoneRecordList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ZoneRecordSpec) DeepCopyInto(out *ZoneRecordSpec) {
	*out = *in
	if in.State != nil {
		in, out := &in.State, &out.State
		*out = new(ZoneRecordSpecResource)
		(*in).DeepCopyInto(*out)
	}
	in.Resource.DeepCopyInto(&out.Resource)
	out.ProviderRef = in.ProviderRef
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ZoneRecordSpec.
func (in *ZoneRecordSpec) DeepCopy() *ZoneRecordSpec {
	if in == nil {
		return nil
	}
	out := new(ZoneRecordSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ZoneRecordSpecResource) DeepCopyInto(out *ZoneRecordSpecResource) {
	*out = *in
	if in.Fieldtype != nil {
		in, out := &in.Fieldtype, &out.Fieldtype
		*out = new(string)
		**out = **in
	}
	if in.Subdomain != nil {
		in, out := &in.Subdomain, &out.Subdomain
		*out = new(string)
		**out = **in
	}
	if in.Target != nil {
		in, out := &in.Target, &out.Target
		*out = new(string)
		**out = **in
	}
	if in.Ttl != nil {
		in, out := &in.Ttl, &out.Ttl
		*out = new(int64)
		**out = **in
	}
	if in.Zone != nil {
		in, out := &in.Zone, &out.Zone
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ZoneRecordSpecResource.
func (in *ZoneRecordSpecResource) DeepCopy() *ZoneRecordSpecResource {
	if in == nil {
		return nil
	}
	out := new(ZoneRecordSpecResource)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ZoneRecordStatus) DeepCopyInto(out *ZoneRecordStatus) {
	*out = *in
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]v1.Condition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ZoneRecordStatus.
func (in *ZoneRecordStatus) DeepCopy() *ZoneRecordStatus {
	if in == nil {
		return nil
	}
	out := new(ZoneRecordStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ZoneRedirection) DeepCopyInto(out *ZoneRedirection) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ZoneRedirection.
func (in *ZoneRedirection) DeepCopy() *ZoneRedirection {
	if in == nil {
		return nil
	}
	out := new(ZoneRedirection)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ZoneRedirection) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ZoneRedirectionList) DeepCopyInto(out *ZoneRedirectionList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]ZoneRedirection, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ZoneRedirectionList.
func (in *ZoneRedirectionList) DeepCopy() *ZoneRedirectionList {
	if in == nil {
		return nil
	}
	out := new(ZoneRedirectionList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ZoneRedirectionList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ZoneRedirectionSpec) DeepCopyInto(out *ZoneRedirectionSpec) {
	*out = *in
	if in.State != nil {
		in, out := &in.State, &out.State
		*out = new(ZoneRedirectionSpecResource)
		(*in).DeepCopyInto(*out)
	}
	in.Resource.DeepCopyInto(&out.Resource)
	out.ProviderRef = in.ProviderRef
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ZoneRedirectionSpec.
func (in *ZoneRedirectionSpec) DeepCopy() *ZoneRedirectionSpec {
	if in == nil {
		return nil
	}
	out := new(ZoneRedirectionSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ZoneRedirectionSpecResource) DeepCopyInto(out *ZoneRedirectionSpecResource) {
	*out = *in
	if in.Description != nil {
		in, out := &in.Description, &out.Description
		*out = new(string)
		**out = **in
	}
	if in.Keywords != nil {
		in, out := &in.Keywords, &out.Keywords
		*out = new(string)
		**out = **in
	}
	if in.Subdomain != nil {
		in, out := &in.Subdomain, &out.Subdomain
		*out = new(string)
		**out = **in
	}
	if in.Target != nil {
		in, out := &in.Target, &out.Target
		*out = new(string)
		**out = **in
	}
	if in.Title != nil {
		in, out := &in.Title, &out.Title
		*out = new(string)
		**out = **in
	}
	if in.Type != nil {
		in, out := &in.Type, &out.Type
		*out = new(string)
		**out = **in
	}
	if in.Zone != nil {
		in, out := &in.Zone, &out.Zone
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ZoneRedirectionSpecResource.
func (in *ZoneRedirectionSpecResource) DeepCopy() *ZoneRedirectionSpecResource {
	if in == nil {
		return nil
	}
	out := new(ZoneRedirectionSpecResource)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ZoneRedirectionStatus) DeepCopyInto(out *ZoneRedirectionStatus) {
	*out = *in
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]v1.Condition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ZoneRedirectionStatus.
func (in *ZoneRedirectionStatus) DeepCopy() *ZoneRedirectionStatus {
	if in == nil {
		return nil
	}
	out := new(ZoneRedirectionStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ZoneSpec) DeepCopyInto(out *ZoneSpec) {
	*out = *in
	if in.State != nil {
		in, out := &in.State, &out.State
		*out = new(ZoneSpecResource)
		(*in).DeepCopyInto(*out)
	}
	in.Resource.DeepCopyInto(&out.Resource)
	out.ProviderRef = in.ProviderRef
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ZoneSpec.
func (in *ZoneSpec) DeepCopy() *ZoneSpec {
	if in == nil {
		return nil
	}
	out := new(ZoneSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ZoneSpecOrder) DeepCopyInto(out *ZoneSpecOrder) {
	*out = *in
	if in.Date != nil {
		in, out := &in.Date, &out.Date
		*out = new(string)
		**out = **in
	}
	if in.Details != nil {
		in, out := &in.Details, &out.Details
		*out = make([]ZoneSpecOrderDetails, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.ExpirationDate != nil {
		in, out := &in.ExpirationDate, &out.ExpirationDate
		*out = new(string)
		**out = **in
	}
	if in.OrderID != nil {
		in, out := &in.OrderID, &out.OrderID
		*out = new(int64)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ZoneSpecOrder.
func (in *ZoneSpecOrder) DeepCopy() *ZoneSpecOrder {
	if in == nil {
		return nil
	}
	out := new(ZoneSpecOrder)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ZoneSpecOrderDetails) DeepCopyInto(out *ZoneSpecOrderDetails) {
	*out = *in
	if in.Description != nil {
		in, out := &in.Description, &out.Description
		*out = new(string)
		**out = **in
	}
	if in.Domain != nil {
		in, out := &in.Domain, &out.Domain
		*out = new(string)
		**out = **in
	}
	if in.OrderDetailID != nil {
		in, out := &in.OrderDetailID, &out.OrderDetailID
		*out = new(int64)
		**out = **in
	}
	if in.Quantity != nil {
		in, out := &in.Quantity, &out.Quantity
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ZoneSpecOrderDetails.
func (in *ZoneSpecOrderDetails) DeepCopy() *ZoneSpecOrderDetails {
	if in == nil {
		return nil
	}
	out := new(ZoneSpecOrderDetails)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ZoneSpecPlan) DeepCopyInto(out *ZoneSpecPlan) {
	*out = *in
	if in.CatalogName != nil {
		in, out := &in.CatalogName, &out.CatalogName
		*out = new(string)
		**out = **in
	}
	if in.Configuration != nil {
		in, out := &in.Configuration, &out.Configuration
		*out = make([]ZoneSpecPlanConfiguration, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Duration != nil {
		in, out := &in.Duration, &out.Duration
		*out = new(string)
		**out = **in
	}
	if in.PlanCode != nil {
		in, out := &in.PlanCode, &out.PlanCode
		*out = new(string)
		**out = **in
	}
	if in.PricingMode != nil {
		in, out := &in.PricingMode, &out.PricingMode
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ZoneSpecPlan.
func (in *ZoneSpecPlan) DeepCopy() *ZoneSpecPlan {
	if in == nil {
		return nil
	}
	out := new(ZoneSpecPlan)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ZoneSpecPlanConfiguration) DeepCopyInto(out *ZoneSpecPlanConfiguration) {
	*out = *in
	if in.Label != nil {
		in, out := &in.Label, &out.Label
		*out = new(string)
		**out = **in
	}
	if in.Value != nil {
		in, out := &in.Value, &out.Value
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ZoneSpecPlanConfiguration.
func (in *ZoneSpecPlanConfiguration) DeepCopy() *ZoneSpecPlanConfiguration {
	if in == nil {
		return nil
	}
	out := new(ZoneSpecPlanConfiguration)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ZoneSpecPlanOption) DeepCopyInto(out *ZoneSpecPlanOption) {
	*out = *in
	if in.CatalogName != nil {
		in, out := &in.CatalogName, &out.CatalogName
		*out = new(string)
		**out = **in
	}
	if in.Configuration != nil {
		in, out := &in.Configuration, &out.Configuration
		*out = make([]ZoneSpecPlanOptionConfiguration, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Duration != nil {
		in, out := &in.Duration, &out.Duration
		*out = new(string)
		**out = **in
	}
	if in.PlanCode != nil {
		in, out := &in.PlanCode, &out.PlanCode
		*out = new(string)
		**out = **in
	}
	if in.PricingMode != nil {
		in, out := &in.PricingMode, &out.PricingMode
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ZoneSpecPlanOption.
func (in *ZoneSpecPlanOption) DeepCopy() *ZoneSpecPlanOption {
	if in == nil {
		return nil
	}
	out := new(ZoneSpecPlanOption)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ZoneSpecPlanOptionConfiguration) DeepCopyInto(out *ZoneSpecPlanOptionConfiguration) {
	*out = *in
	if in.Label != nil {
		in, out := &in.Label, &out.Label
		*out = new(string)
		**out = **in
	}
	if in.Value != nil {
		in, out := &in.Value, &out.Value
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ZoneSpecPlanOptionConfiguration.
func (in *ZoneSpecPlanOptionConfiguration) DeepCopy() *ZoneSpecPlanOptionConfiguration {
	if in == nil {
		return nil
	}
	out := new(ZoneSpecPlanOptionConfiguration)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ZoneSpecResource) DeepCopyInto(out *ZoneSpecResource) {
	*out = *in
	if in.DnssecSupported != nil {
		in, out := &in.DnssecSupported, &out.DnssecSupported
		*out = new(bool)
		**out = **in
	}
	if in.HasDNSAnycast != nil {
		in, out := &in.HasDNSAnycast, &out.HasDNSAnycast
		*out = new(bool)
		**out = **in
	}
	if in.LastUpdate != nil {
		in, out := &in.LastUpdate, &out.LastUpdate
		*out = new(string)
		**out = **in
	}
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	if in.NameServers != nil {
		in, out := &in.NameServers, &out.NameServers
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.Order != nil {
		in, out := &in.Order, &out.Order
		*out = make([]ZoneSpecOrder, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.OvhSubsidiary != nil {
		in, out := &in.OvhSubsidiary, &out.OvhSubsidiary
		*out = new(string)
		**out = **in
	}
	if in.PaymentMean != nil {
		in, out := &in.PaymentMean, &out.PaymentMean
		*out = new(string)
		**out = **in
	}
	if in.Plan != nil {
		in, out := &in.Plan, &out.Plan
		*out = new(ZoneSpecPlan)
		(*in).DeepCopyInto(*out)
	}
	if in.PlanOption != nil {
		in, out := &in.PlanOption, &out.PlanOption
		*out = make([]ZoneSpecPlanOption, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ZoneSpecResource.
func (in *ZoneSpecResource) DeepCopy() *ZoneSpecResource {
	if in == nil {
		return nil
	}
	out := new(ZoneSpecResource)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ZoneStatus) DeepCopyInto(out *ZoneStatus) {
	*out = *in
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]v1.Condition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ZoneStatus.
func (in *ZoneStatus) DeepCopy() *ZoneStatus {
	if in == nil {
		return nil
	}
	out := new(ZoneStatus)
	in.DeepCopyInto(out)
	return out
}
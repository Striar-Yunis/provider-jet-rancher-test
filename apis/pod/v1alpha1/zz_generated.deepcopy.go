//go:build !ignore_autogenerated
// +build !ignore_autogenerated

/*
Copyright 2021 The Crossplane Authors.

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

// Code generated by controller-gen. DO NOT EDIT.

package v1alpha1

import (
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AllowedCsiDriverObservation) DeepCopyInto(out *AllowedCsiDriverObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AllowedCsiDriverObservation.
func (in *AllowedCsiDriverObservation) DeepCopy() *AllowedCsiDriverObservation {
	if in == nil {
		return nil
	}
	out := new(AllowedCsiDriverObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AllowedCsiDriverParameters) DeepCopyInto(out *AllowedCsiDriverParameters) {
	*out = *in
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AllowedCsiDriverParameters.
func (in *AllowedCsiDriverParameters) DeepCopy() *AllowedCsiDriverParameters {
	if in == nil {
		return nil
	}
	out := new(AllowedCsiDriverParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AllowedFlexVolumeObservation) DeepCopyInto(out *AllowedFlexVolumeObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AllowedFlexVolumeObservation.
func (in *AllowedFlexVolumeObservation) DeepCopy() *AllowedFlexVolumeObservation {
	if in == nil {
		return nil
	}
	out := new(AllowedFlexVolumeObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AllowedFlexVolumeParameters) DeepCopyInto(out *AllowedFlexVolumeParameters) {
	*out = *in
	if in.Driver != nil {
		in, out := &in.Driver, &out.Driver
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AllowedFlexVolumeParameters.
func (in *AllowedFlexVolumeParameters) DeepCopy() *AllowedFlexVolumeParameters {
	if in == nil {
		return nil
	}
	out := new(AllowedFlexVolumeParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AllowedHostPathObservation) DeepCopyInto(out *AllowedHostPathObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AllowedHostPathObservation.
func (in *AllowedHostPathObservation) DeepCopy() *AllowedHostPathObservation {
	if in == nil {
		return nil
	}
	out := new(AllowedHostPathObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AllowedHostPathParameters) DeepCopyInto(out *AllowedHostPathParameters) {
	*out = *in
	if in.PathPrefix != nil {
		in, out := &in.PathPrefix, &out.PathPrefix
		*out = new(string)
		**out = **in
	}
	if in.ReadOnly != nil {
		in, out := &in.ReadOnly, &out.ReadOnly
		*out = new(bool)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AllowedHostPathParameters.
func (in *AllowedHostPathParameters) DeepCopy() *AllowedHostPathParameters {
	if in == nil {
		return nil
	}
	out := new(AllowedHostPathParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FsGroupObservation) DeepCopyInto(out *FsGroupObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FsGroupObservation.
func (in *FsGroupObservation) DeepCopy() *FsGroupObservation {
	if in == nil {
		return nil
	}
	out := new(FsGroupObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FsGroupParameters) DeepCopyInto(out *FsGroupParameters) {
	*out = *in
	if in.Range != nil {
		in, out := &in.Range, &out.Range
		*out = make([]RangeParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Rule != nil {
		in, out := &in.Rule, &out.Rule
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FsGroupParameters.
func (in *FsGroupParameters) DeepCopy() *FsGroupParameters {
	if in == nil {
		return nil
	}
	out := new(FsGroupParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HostPortObservation) DeepCopyInto(out *HostPortObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HostPortObservation.
func (in *HostPortObservation) DeepCopy() *HostPortObservation {
	if in == nil {
		return nil
	}
	out := new(HostPortObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HostPortParameters) DeepCopyInto(out *HostPortParameters) {
	*out = *in
	if in.Max != nil {
		in, out := &in.Max, &out.Max
		*out = new(float64)
		**out = **in
	}
	if in.Min != nil {
		in, out := &in.Min, &out.Min
		*out = new(float64)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HostPortParameters.
func (in *HostPortParameters) DeepCopy() *HostPortParameters {
	if in == nil {
		return nil
	}
	out := new(HostPortParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RangeObservation) DeepCopyInto(out *RangeObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RangeObservation.
func (in *RangeObservation) DeepCopy() *RangeObservation {
	if in == nil {
		return nil
	}
	out := new(RangeObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RangeParameters) DeepCopyInto(out *RangeParameters) {
	*out = *in
	if in.Max != nil {
		in, out := &in.Max, &out.Max
		*out = new(float64)
		**out = **in
	}
	if in.Min != nil {
		in, out := &in.Min, &out.Min
		*out = new(float64)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RangeParameters.
func (in *RangeParameters) DeepCopy() *RangeParameters {
	if in == nil {
		return nil
	}
	out := new(RangeParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RunAsGroupObservation) DeepCopyInto(out *RunAsGroupObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RunAsGroupObservation.
func (in *RunAsGroupObservation) DeepCopy() *RunAsGroupObservation {
	if in == nil {
		return nil
	}
	out := new(RunAsGroupObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RunAsGroupParameters) DeepCopyInto(out *RunAsGroupParameters) {
	*out = *in
	if in.Range != nil {
		in, out := &in.Range, &out.Range
		*out = make([]RunAsGroupRangeParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Rule != nil {
		in, out := &in.Rule, &out.Rule
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RunAsGroupParameters.
func (in *RunAsGroupParameters) DeepCopy() *RunAsGroupParameters {
	if in == nil {
		return nil
	}
	out := new(RunAsGroupParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RunAsGroupRangeObservation) DeepCopyInto(out *RunAsGroupRangeObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RunAsGroupRangeObservation.
func (in *RunAsGroupRangeObservation) DeepCopy() *RunAsGroupRangeObservation {
	if in == nil {
		return nil
	}
	out := new(RunAsGroupRangeObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RunAsGroupRangeParameters) DeepCopyInto(out *RunAsGroupRangeParameters) {
	*out = *in
	if in.Max != nil {
		in, out := &in.Max, &out.Max
		*out = new(float64)
		**out = **in
	}
	if in.Min != nil {
		in, out := &in.Min, &out.Min
		*out = new(float64)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RunAsGroupRangeParameters.
func (in *RunAsGroupRangeParameters) DeepCopy() *RunAsGroupRangeParameters {
	if in == nil {
		return nil
	}
	out := new(RunAsGroupRangeParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RunAsUserObservation) DeepCopyInto(out *RunAsUserObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RunAsUserObservation.
func (in *RunAsUserObservation) DeepCopy() *RunAsUserObservation {
	if in == nil {
		return nil
	}
	out := new(RunAsUserObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RunAsUserParameters) DeepCopyInto(out *RunAsUserParameters) {
	*out = *in
	if in.Range != nil {
		in, out := &in.Range, &out.Range
		*out = make([]RunAsUserRangeParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Rule != nil {
		in, out := &in.Rule, &out.Rule
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RunAsUserParameters.
func (in *RunAsUserParameters) DeepCopy() *RunAsUserParameters {
	if in == nil {
		return nil
	}
	out := new(RunAsUserParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RunAsUserRangeObservation) DeepCopyInto(out *RunAsUserRangeObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RunAsUserRangeObservation.
func (in *RunAsUserRangeObservation) DeepCopy() *RunAsUserRangeObservation {
	if in == nil {
		return nil
	}
	out := new(RunAsUserRangeObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RunAsUserRangeParameters) DeepCopyInto(out *RunAsUserRangeParameters) {
	*out = *in
	if in.Max != nil {
		in, out := &in.Max, &out.Max
		*out = new(float64)
		**out = **in
	}
	if in.Min != nil {
		in, out := &in.Min, &out.Min
		*out = new(float64)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RunAsUserRangeParameters.
func (in *RunAsUserRangeParameters) DeepCopy() *RunAsUserRangeParameters {
	if in == nil {
		return nil
	}
	out := new(RunAsUserRangeParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RuntimeClassObservation) DeepCopyInto(out *RuntimeClassObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RuntimeClassObservation.
func (in *RuntimeClassObservation) DeepCopy() *RuntimeClassObservation {
	if in == nil {
		return nil
	}
	out := new(RuntimeClassObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RuntimeClassParameters) DeepCopyInto(out *RuntimeClassParameters) {
	*out = *in
	if in.AllowedRuntimeClassNames != nil {
		in, out := &in.AllowedRuntimeClassNames, &out.AllowedRuntimeClassNames
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
	if in.DefaultRuntimeClassName != nil {
		in, out := &in.DefaultRuntimeClassName, &out.DefaultRuntimeClassName
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RuntimeClassParameters.
func (in *RuntimeClassParameters) DeepCopy() *RuntimeClassParameters {
	if in == nil {
		return nil
	}
	out := new(RuntimeClassParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SeLinuxObservation) DeepCopyInto(out *SeLinuxObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SeLinuxObservation.
func (in *SeLinuxObservation) DeepCopy() *SeLinuxObservation {
	if in == nil {
		return nil
	}
	out := new(SeLinuxObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SeLinuxOptionObservation) DeepCopyInto(out *SeLinuxOptionObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SeLinuxOptionObservation.
func (in *SeLinuxOptionObservation) DeepCopy() *SeLinuxOptionObservation {
	if in == nil {
		return nil
	}
	out := new(SeLinuxOptionObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SeLinuxOptionParameters) DeepCopyInto(out *SeLinuxOptionParameters) {
	*out = *in
	if in.Level != nil {
		in, out := &in.Level, &out.Level
		*out = new(string)
		**out = **in
	}
	if in.Role != nil {
		in, out := &in.Role, &out.Role
		*out = new(string)
		**out = **in
	}
	if in.Type != nil {
		in, out := &in.Type, &out.Type
		*out = new(string)
		**out = **in
	}
	if in.User != nil {
		in, out := &in.User, &out.User
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SeLinuxOptionParameters.
func (in *SeLinuxOptionParameters) DeepCopy() *SeLinuxOptionParameters {
	if in == nil {
		return nil
	}
	out := new(SeLinuxOptionParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SeLinuxParameters) DeepCopyInto(out *SeLinuxParameters) {
	*out = *in
	if in.Rule != nil {
		in, out := &in.Rule, &out.Rule
		*out = new(string)
		**out = **in
	}
	if in.SeLinuxOption != nil {
		in, out := &in.SeLinuxOption, &out.SeLinuxOption
		*out = make([]SeLinuxOptionParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SeLinuxParameters.
func (in *SeLinuxParameters) DeepCopy() *SeLinuxParameters {
	if in == nil {
		return nil
	}
	out := new(SeLinuxParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SecurityPolicyTemplate) DeepCopyInto(out *SecurityPolicyTemplate) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SecurityPolicyTemplate.
func (in *SecurityPolicyTemplate) DeepCopy() *SecurityPolicyTemplate {
	if in == nil {
		return nil
	}
	out := new(SecurityPolicyTemplate)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *SecurityPolicyTemplate) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SecurityPolicyTemplateList) DeepCopyInto(out *SecurityPolicyTemplateList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]SecurityPolicyTemplate, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SecurityPolicyTemplateList.
func (in *SecurityPolicyTemplateList) DeepCopy() *SecurityPolicyTemplateList {
	if in == nil {
		return nil
	}
	out := new(SecurityPolicyTemplateList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *SecurityPolicyTemplateList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SecurityPolicyTemplateObservation) DeepCopyInto(out *SecurityPolicyTemplateObservation) {
	*out = *in
	if in.ID != nil {
		in, out := &in.ID, &out.ID
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SecurityPolicyTemplateObservation.
func (in *SecurityPolicyTemplateObservation) DeepCopy() *SecurityPolicyTemplateObservation {
	if in == nil {
		return nil
	}
	out := new(SecurityPolicyTemplateObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SecurityPolicyTemplateParameters) DeepCopyInto(out *SecurityPolicyTemplateParameters) {
	*out = *in
	if in.AllowPrivilegeEscalation != nil {
		in, out := &in.AllowPrivilegeEscalation, &out.AllowPrivilegeEscalation
		*out = new(bool)
		**out = **in
	}
	if in.AllowedCapabilities != nil {
		in, out := &in.AllowedCapabilities, &out.AllowedCapabilities
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
	if in.AllowedCsiDriver != nil {
		in, out := &in.AllowedCsiDriver, &out.AllowedCsiDriver
		*out = make([]AllowedCsiDriverParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.AllowedFlexVolume != nil {
		in, out := &in.AllowedFlexVolume, &out.AllowedFlexVolume
		*out = make([]AllowedFlexVolumeParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.AllowedHostPath != nil {
		in, out := &in.AllowedHostPath, &out.AllowedHostPath
		*out = make([]AllowedHostPathParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.AllowedProcMountTypes != nil {
		in, out := &in.AllowedProcMountTypes, &out.AllowedProcMountTypes
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
	if in.AllowedUnsafeSysctls != nil {
		in, out := &in.AllowedUnsafeSysctls, &out.AllowedUnsafeSysctls
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
	if in.Annotations != nil {
		in, out := &in.Annotations, &out.Annotations
		*out = make(map[string]*string, len(*in))
		for key, val := range *in {
			var outVal *string
			if val == nil {
				(*out)[key] = nil
			} else {
				in, out := &val, &outVal
				*out = new(string)
				**out = **in
			}
			(*out)[key] = outVal
		}
	}
	if in.DefaultAddCapabilities != nil {
		in, out := &in.DefaultAddCapabilities, &out.DefaultAddCapabilities
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
	if in.DefaultAllowPrivilegeEscalation != nil {
		in, out := &in.DefaultAllowPrivilegeEscalation, &out.DefaultAllowPrivilegeEscalation
		*out = new(bool)
		**out = **in
	}
	if in.Description != nil {
		in, out := &in.Description, &out.Description
		*out = new(string)
		**out = **in
	}
	if in.ForbiddenSysctls != nil {
		in, out := &in.ForbiddenSysctls, &out.ForbiddenSysctls
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
	if in.FsGroup != nil {
		in, out := &in.FsGroup, &out.FsGroup
		*out = make([]FsGroupParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.HostIpc != nil {
		in, out := &in.HostIpc, &out.HostIpc
		*out = new(bool)
		**out = **in
	}
	if in.HostNetwork != nil {
		in, out := &in.HostNetwork, &out.HostNetwork
		*out = new(bool)
		**out = **in
	}
	if in.HostPid != nil {
		in, out := &in.HostPid, &out.HostPid
		*out = new(bool)
		**out = **in
	}
	if in.HostPort != nil {
		in, out := &in.HostPort, &out.HostPort
		*out = make([]HostPortParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Labels != nil {
		in, out := &in.Labels, &out.Labels
		*out = make(map[string]*string, len(*in))
		for key, val := range *in {
			var outVal *string
			if val == nil {
				(*out)[key] = nil
			} else {
				in, out := &val, &outVal
				*out = new(string)
				**out = **in
			}
			(*out)[key] = outVal
		}
	}
	if in.Privileged != nil {
		in, out := &in.Privileged, &out.Privileged
		*out = new(bool)
		**out = **in
	}
	if in.ReadOnlyRootFilesystem != nil {
		in, out := &in.ReadOnlyRootFilesystem, &out.ReadOnlyRootFilesystem
		*out = new(bool)
		**out = **in
	}
	if in.RequiredDropCapabilities != nil {
		in, out := &in.RequiredDropCapabilities, &out.RequiredDropCapabilities
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
	if in.RunAsGroup != nil {
		in, out := &in.RunAsGroup, &out.RunAsGroup
		*out = make([]RunAsGroupParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.RunAsUser != nil {
		in, out := &in.RunAsUser, &out.RunAsUser
		*out = make([]RunAsUserParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.RuntimeClass != nil {
		in, out := &in.RuntimeClass, &out.RuntimeClass
		*out = make([]RuntimeClassParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.SeLinux != nil {
		in, out := &in.SeLinux, &out.SeLinux
		*out = make([]SeLinuxParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.SupplementalGroup != nil {
		in, out := &in.SupplementalGroup, &out.SupplementalGroup
		*out = make([]SupplementalGroupParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Volumes != nil {
		in, out := &in.Volumes, &out.Volumes
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SecurityPolicyTemplateParameters.
func (in *SecurityPolicyTemplateParameters) DeepCopy() *SecurityPolicyTemplateParameters {
	if in == nil {
		return nil
	}
	out := new(SecurityPolicyTemplateParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SecurityPolicyTemplateSpec) DeepCopyInto(out *SecurityPolicyTemplateSpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	in.ForProvider.DeepCopyInto(&out.ForProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SecurityPolicyTemplateSpec.
func (in *SecurityPolicyTemplateSpec) DeepCopy() *SecurityPolicyTemplateSpec {
	if in == nil {
		return nil
	}
	out := new(SecurityPolicyTemplateSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SecurityPolicyTemplateStatus) DeepCopyInto(out *SecurityPolicyTemplateStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	in.AtProvider.DeepCopyInto(&out.AtProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SecurityPolicyTemplateStatus.
func (in *SecurityPolicyTemplateStatus) DeepCopy() *SecurityPolicyTemplateStatus {
	if in == nil {
		return nil
	}
	out := new(SecurityPolicyTemplateStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SupplementalGroupObservation) DeepCopyInto(out *SupplementalGroupObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SupplementalGroupObservation.
func (in *SupplementalGroupObservation) DeepCopy() *SupplementalGroupObservation {
	if in == nil {
		return nil
	}
	out := new(SupplementalGroupObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SupplementalGroupParameters) DeepCopyInto(out *SupplementalGroupParameters) {
	*out = *in
	if in.Range != nil {
		in, out := &in.Range, &out.Range
		*out = make([]SupplementalGroupRangeParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Rule != nil {
		in, out := &in.Rule, &out.Rule
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SupplementalGroupParameters.
func (in *SupplementalGroupParameters) DeepCopy() *SupplementalGroupParameters {
	if in == nil {
		return nil
	}
	out := new(SupplementalGroupParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SupplementalGroupRangeObservation) DeepCopyInto(out *SupplementalGroupRangeObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SupplementalGroupRangeObservation.
func (in *SupplementalGroupRangeObservation) DeepCopy() *SupplementalGroupRangeObservation {
	if in == nil {
		return nil
	}
	out := new(SupplementalGroupRangeObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SupplementalGroupRangeParameters) DeepCopyInto(out *SupplementalGroupRangeParameters) {
	*out = *in
	if in.Max != nil {
		in, out := &in.Max, &out.Max
		*out = new(float64)
		**out = **in
	}
	if in.Min != nil {
		in, out := &in.Min, &out.Min
		*out = new(float64)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SupplementalGroupRangeParameters.
func (in *SupplementalGroupRangeParameters) DeepCopy() *SupplementalGroupRangeParameters {
	if in == nil {
		return nil
	}
	out := new(SupplementalGroupRangeParameters)
	in.DeepCopyInto(out)
	return out
}
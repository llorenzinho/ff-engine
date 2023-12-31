//go:build !ignore_autogenerated
// +build !ignore_autogenerated

/*
Copyright Lorenzo.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0
*/
// Code generated by deepcopy-gen. DO NOT EDIT.

package v1alpha1

import (
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FeatureFlag) DeepCopyInto(out *FeatureFlag) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	out.Spec = in.Spec
	out.Status = in.Status
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FeatureFlag.
func (in *FeatureFlag) DeepCopy() *FeatureFlag {
	if in == nil {
		return nil
	}
	out := new(FeatureFlag)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *FeatureFlag) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FeatureFlagList) DeepCopyInto(out *FeatureFlagList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]FeatureFlag, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FeatureFlagList.
func (in *FeatureFlagList) DeepCopy() *FeatureFlagList {
	if in == nil {
		return nil
	}
	out := new(FeatureFlagList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *FeatureFlagList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FeatureFlagSpec) DeepCopyInto(out *FeatureFlagSpec) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FeatureFlagSpec.
func (in *FeatureFlagSpec) DeepCopy() *FeatureFlagSpec {
	if in == nil {
		return nil
	}
	out := new(FeatureFlagSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FeatureFlagStatus) DeepCopyInto(out *FeatureFlagStatus) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FeatureFlagStatus.
func (in *FeatureFlagStatus) DeepCopy() *FeatureFlagStatus {
	if in == nil {
		return nil
	}
	out := new(FeatureFlagStatus)
	in.DeepCopyInto(out)
	return out
}

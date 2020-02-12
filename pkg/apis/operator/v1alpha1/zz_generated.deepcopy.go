// +build !ignore_autogenerated

// Code generated by deepcopy-gen. DO NOT EDIT.

package v1alpha1

import (
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EFSCSIDriver) DeepCopyInto(out *EFSCSIDriver) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	out.Status = in.Status
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EFSCSIDriver.
func (in *EFSCSIDriver) DeepCopy() *EFSCSIDriver {
	if in == nil {
		return nil
	}
	out := new(EFSCSIDriver)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *EFSCSIDriver) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EFSCSIDriverList) DeepCopyInto(out *EFSCSIDriverList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]EFSCSIDriver, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EFSCSIDriverList.
func (in *EFSCSIDriverList) DeepCopy() *EFSCSIDriverList {
	if in == nil {
		return nil
	}
	out := new(EFSCSIDriverList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *EFSCSIDriverList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EFSCSIDriverSpec) DeepCopyInto(out *EFSCSIDriverSpec) {
	*out = *in
	if in.Replicas != nil {
		in, out := &in.Replicas, &out.Replicas
		*out = new(int32)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EFSCSIDriverSpec.
func (in *EFSCSIDriverSpec) DeepCopy() *EFSCSIDriverSpec {
	if in == nil {
		return nil
	}
	out := new(EFSCSIDriverSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EFSCSIDriverStatus) DeepCopyInto(out *EFSCSIDriverStatus) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EFSCSIDriverStatus.
func (in *EFSCSIDriverStatus) DeepCopy() *EFSCSIDriverStatus {
	if in == nil {
		return nil
	}
	out := new(EFSCSIDriverStatus)
	in.DeepCopyInto(out)
	return out
}

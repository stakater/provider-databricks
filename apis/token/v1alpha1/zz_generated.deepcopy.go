//go:build !ignore_autogenerated

/*
Copyright 2022 Upbound Inc.
*/

// Code generated by controller-gen. DO NOT EDIT.

package v1alpha1

import (
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Token) DeepCopyInto(out *Token) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Token.
func (in *Token) DeepCopy() *Token {
	if in == nil {
		return nil
	}
	out := new(Token)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Token) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TokenInitParameters) DeepCopyInto(out *TokenInitParameters) {
	*out = *in
	if in.Comment != nil {
		in, out := &in.Comment, &out.Comment
		*out = new(string)
		**out = **in
	}
	if in.CreationTime != nil {
		in, out := &in.CreationTime, &out.CreationTime
		*out = new(float64)
		**out = **in
	}
	if in.ExpiryTime != nil {
		in, out := &in.ExpiryTime, &out.ExpiryTime
		*out = new(float64)
		**out = **in
	}
	if in.LifetimeSeconds != nil {
		in, out := &in.LifetimeSeconds, &out.LifetimeSeconds
		*out = new(float64)
		**out = **in
	}
	if in.TokenID != nil {
		in, out := &in.TokenID, &out.TokenID
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TokenInitParameters.
func (in *TokenInitParameters) DeepCopy() *TokenInitParameters {
	if in == nil {
		return nil
	}
	out := new(TokenInitParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TokenList) DeepCopyInto(out *TokenList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Token, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TokenList.
func (in *TokenList) DeepCopy() *TokenList {
	if in == nil {
		return nil
	}
	out := new(TokenList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *TokenList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TokenObservation) DeepCopyInto(out *TokenObservation) {
	*out = *in
	if in.Comment != nil {
		in, out := &in.Comment, &out.Comment
		*out = new(string)
		**out = **in
	}
	if in.CreationTime != nil {
		in, out := &in.CreationTime, &out.CreationTime
		*out = new(float64)
		**out = **in
	}
	if in.ExpiryTime != nil {
		in, out := &in.ExpiryTime, &out.ExpiryTime
		*out = new(float64)
		**out = **in
	}
	if in.ID != nil {
		in, out := &in.ID, &out.ID
		*out = new(string)
		**out = **in
	}
	if in.LifetimeSeconds != nil {
		in, out := &in.LifetimeSeconds, &out.LifetimeSeconds
		*out = new(float64)
		**out = **in
	}
	if in.TokenID != nil {
		in, out := &in.TokenID, &out.TokenID
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TokenObservation.
func (in *TokenObservation) DeepCopy() *TokenObservation {
	if in == nil {
		return nil
	}
	out := new(TokenObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TokenParameters) DeepCopyInto(out *TokenParameters) {
	*out = *in
	if in.Comment != nil {
		in, out := &in.Comment, &out.Comment
		*out = new(string)
		**out = **in
	}
	if in.CreationTime != nil {
		in, out := &in.CreationTime, &out.CreationTime
		*out = new(float64)
		**out = **in
	}
	if in.ExpiryTime != nil {
		in, out := &in.ExpiryTime, &out.ExpiryTime
		*out = new(float64)
		**out = **in
	}
	if in.LifetimeSeconds != nil {
		in, out := &in.LifetimeSeconds, &out.LifetimeSeconds
		*out = new(float64)
		**out = **in
	}
	if in.TokenID != nil {
		in, out := &in.TokenID, &out.TokenID
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TokenParameters.
func (in *TokenParameters) DeepCopy() *TokenParameters {
	if in == nil {
		return nil
	}
	out := new(TokenParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TokenSpec) DeepCopyInto(out *TokenSpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	in.ForProvider.DeepCopyInto(&out.ForProvider)
	in.InitProvider.DeepCopyInto(&out.InitProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TokenSpec.
func (in *TokenSpec) DeepCopy() *TokenSpec {
	if in == nil {
		return nil
	}
	out := new(TokenSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TokenStatus) DeepCopyInto(out *TokenStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	in.AtProvider.DeepCopyInto(&out.AtProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TokenStatus.
func (in *TokenStatus) DeepCopy() *TokenStatus {
	if in == nil {
		return nil
	}
	out := new(TokenStatus)
	in.DeepCopyInto(out)
	return out
}

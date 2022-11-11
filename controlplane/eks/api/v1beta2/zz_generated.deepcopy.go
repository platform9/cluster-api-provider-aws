//go:build !ignore_autogenerated
// +build !ignore_autogenerated

/*
Copyright The Kubernetes Authors.

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

package v1beta2

import (
	"k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/runtime"
	apiv1beta2 "sigs.k8s.io/cluster-api-provider-aws/v2/api/v1beta2"
	"sigs.k8s.io/cluster-api/api/v1beta1"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AWSManagedControlPlane) DeepCopyInto(out *AWSManagedControlPlane) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AWSManagedControlPlane.
func (in *AWSManagedControlPlane) DeepCopy() *AWSManagedControlPlane {
	if in == nil {
		return nil
	}
	out := new(AWSManagedControlPlane)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *AWSManagedControlPlane) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AWSManagedControlPlaneList) DeepCopyInto(out *AWSManagedControlPlaneList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]AWSManagedControlPlane, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AWSManagedControlPlaneList.
func (in *AWSManagedControlPlaneList) DeepCopy() *AWSManagedControlPlaneList {
	if in == nil {
		return nil
	}
	out := new(AWSManagedControlPlaneList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *AWSManagedControlPlaneList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AWSManagedControlPlaneSpec) DeepCopyInto(out *AWSManagedControlPlaneSpec) {
	*out = *in
	if in.IdentityRef != nil {
		in, out := &in.IdentityRef, &out.IdentityRef
		*out = new(apiv1beta2.AWSIdentityReference)
		**out = **in
	}
	in.NetworkSpec.DeepCopyInto(&out.NetworkSpec)
	if in.SecondaryCidrBlock != nil {
		in, out := &in.SecondaryCidrBlock, &out.SecondaryCidrBlock
		*out = new(string)
		**out = **in
	}
	if in.SSHKeyName != nil {
		in, out := &in.SSHKeyName, &out.SSHKeyName
		*out = new(string)
		**out = **in
	}
	if in.Version != nil {
		in, out := &in.Version, &out.Version
		*out = new(string)
		**out = **in
	}
	if in.RoleName != nil {
		in, out := &in.RoleName, &out.RoleName
		*out = new(string)
		**out = **in
	}
	if in.RoleAdditionalPolicies != nil {
		in, out := &in.RoleAdditionalPolicies, &out.RoleAdditionalPolicies
		*out = new([]string)
		if **in != nil {
			in, out := *in, *out
			*out = make([]string, len(*in))
			copy(*out, *in)
		}
	}
	if in.Logging != nil {
		in, out := &in.Logging, &out.Logging
		*out = new(ControlPlaneLoggingSpec)
		**out = **in
	}
	if in.EncryptionConfig != nil {
		in, out := &in.EncryptionConfig, &out.EncryptionConfig
		*out = new(EncryptionConfig)
		(*in).DeepCopyInto(*out)
	}
	if in.AdditionalTags != nil {
		in, out := &in.AdditionalTags, &out.AdditionalTags
		*out = make(apiv1beta2.Tags, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.IAMAuthenticatorConfig != nil {
		in, out := &in.IAMAuthenticatorConfig, &out.IAMAuthenticatorConfig
		*out = new(IAMAuthenticatorConfig)
		(*in).DeepCopyInto(*out)
	}
	in.EndpointAccess.DeepCopyInto(&out.EndpointAccess)
	out.ControlPlaneEndpoint = in.ControlPlaneEndpoint
	in.Bastion.DeepCopyInto(&out.Bastion)
	if in.TokenMethod != nil {
		in, out := &in.TokenMethod, &out.TokenMethod
		*out = new(EKSTokenMethod)
		**out = **in
	}
	if in.Addons != nil {
		in, out := &in.Addons, &out.Addons
		*out = new([]Addon)
		if **in != nil {
			in, out := *in, *out
			*out = make([]Addon, len(*in))
			for i := range *in {
				(*in)[i].DeepCopyInto(&(*out)[i])
			}
		}
	}
	if in.OIDCIdentityProviderConfig != nil {
		in, out := &in.OIDCIdentityProviderConfig, &out.OIDCIdentityProviderConfig
		*out = new(OIDCIdentityProviderConfig)
		(*in).DeepCopyInto(*out)
	}
	in.VpcCni.DeepCopyInto(&out.VpcCni)
	out.KubeProxy = in.KubeProxy
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AWSManagedControlPlaneSpec.
func (in *AWSManagedControlPlaneSpec) DeepCopy() *AWSManagedControlPlaneSpec {
	if in == nil {
		return nil
	}
	out := new(AWSManagedControlPlaneSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AWSManagedControlPlaneStatus) DeepCopyInto(out *AWSManagedControlPlaneStatus) {
	*out = *in
	in.Network.DeepCopyInto(&out.Network)
	if in.FailureDomains != nil {
		in, out := &in.FailureDomains, &out.FailureDomains
		*out = make(v1beta1.FailureDomains, len(*in))
		for key, val := range *in {
			(*out)[key] = *val.DeepCopy()
		}
	}
	if in.Bastion != nil {
		in, out := &in.Bastion, &out.Bastion
		*out = new(apiv1beta2.Instance)
		(*in).DeepCopyInto(*out)
	}
	out.OIDCProvider = in.OIDCProvider
	if in.ExternalManagedControlPlane != nil {
		in, out := &in.ExternalManagedControlPlane, &out.ExternalManagedControlPlane
		*out = new(bool)
		**out = **in
	}
	if in.FailureMessage != nil {
		in, out := &in.FailureMessage, &out.FailureMessage
		*out = new(string)
		**out = **in
	}
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make(v1beta1.Conditions, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Addons != nil {
		in, out := &in.Addons, &out.Addons
		*out = make([]AddonState, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	out.IdentityProviderStatus = in.IdentityProviderStatus
	if in.Version != nil {
		in, out := &in.Version, &out.Version
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AWSManagedControlPlaneStatus.
func (in *AWSManagedControlPlaneStatus) DeepCopy() *AWSManagedControlPlaneStatus {
	if in == nil {
		return nil
	}
	out := new(AWSManagedControlPlaneStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Addon) DeepCopyInto(out *Addon) {
	*out = *in
	if in.ConflictResolution != nil {
		in, out := &in.ConflictResolution, &out.ConflictResolution
		*out = new(AddonResolution)
		**out = **in
	}
	if in.ServiceAccountRoleArn != nil {
		in, out := &in.ServiceAccountRoleArn, &out.ServiceAccountRoleArn
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Addon.
func (in *Addon) DeepCopy() *Addon {
	if in == nil {
		return nil
	}
	out := new(Addon)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AddonIssue) DeepCopyInto(out *AddonIssue) {
	*out = *in
	if in.Code != nil {
		in, out := &in.Code, &out.Code
		*out = new(string)
		**out = **in
	}
	if in.Message != nil {
		in, out := &in.Message, &out.Message
		*out = new(string)
		**out = **in
	}
	if in.ResourceIDs != nil {
		in, out := &in.ResourceIDs, &out.ResourceIDs
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AddonIssue.
func (in *AddonIssue) DeepCopy() *AddonIssue {
	if in == nil {
		return nil
	}
	out := new(AddonIssue)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AddonState) DeepCopyInto(out *AddonState) {
	*out = *in
	if in.ServiceAccountRoleArn != nil {
		in, out := &in.ServiceAccountRoleArn, &out.ServiceAccountRoleArn
		*out = new(string)
		**out = **in
	}
	in.CreatedAt.DeepCopyInto(&out.CreatedAt)
	in.ModifiedAt.DeepCopyInto(&out.ModifiedAt)
	if in.Status != nil {
		in, out := &in.Status, &out.Status
		*out = new(string)
		**out = **in
	}
	if in.Issues != nil {
		in, out := &in.Issues, &out.Issues
		*out = make([]AddonIssue, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AddonState.
func (in *AddonState) DeepCopy() *AddonState {
	if in == nil {
		return nil
	}
	out := new(AddonState)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ControlPlaneLoggingSpec) DeepCopyInto(out *ControlPlaneLoggingSpec) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ControlPlaneLoggingSpec.
func (in *ControlPlaneLoggingSpec) DeepCopy() *ControlPlaneLoggingSpec {
	if in == nil {
		return nil
	}
	out := new(ControlPlaneLoggingSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EncryptionConfig) DeepCopyInto(out *EncryptionConfig) {
	*out = *in
	if in.Provider != nil {
		in, out := &in.Provider, &out.Provider
		*out = new(string)
		**out = **in
	}
	if in.Resources != nil {
		in, out := &in.Resources, &out.Resources
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

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EncryptionConfig.
func (in *EncryptionConfig) DeepCopy() *EncryptionConfig {
	if in == nil {
		return nil
	}
	out := new(EncryptionConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EndpointAccess) DeepCopyInto(out *EndpointAccess) {
	*out = *in
	if in.Public != nil {
		in, out := &in.Public, &out.Public
		*out = new(bool)
		**out = **in
	}
	if in.PublicCIDRs != nil {
		in, out := &in.PublicCIDRs, &out.PublicCIDRs
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
	if in.Private != nil {
		in, out := &in.Private, &out.Private
		*out = new(bool)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EndpointAccess.
func (in *EndpointAccess) DeepCopy() *EndpointAccess {
	if in == nil {
		return nil
	}
	out := new(EndpointAccess)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IAMAuthenticatorConfig) DeepCopyInto(out *IAMAuthenticatorConfig) {
	*out = *in
	if in.RoleMappings != nil {
		in, out := &in.RoleMappings, &out.RoleMappings
		*out = make([]RoleMapping, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.UserMappings != nil {
		in, out := &in.UserMappings, &out.UserMappings
		*out = make([]UserMapping, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IAMAuthenticatorConfig.
func (in *IAMAuthenticatorConfig) DeepCopy() *IAMAuthenticatorConfig {
	if in == nil {
		return nil
	}
	out := new(IAMAuthenticatorConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IdentityProviderStatus) DeepCopyInto(out *IdentityProviderStatus) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IdentityProviderStatus.
func (in *IdentityProviderStatus) DeepCopy() *IdentityProviderStatus {
	if in == nil {
		return nil
	}
	out := new(IdentityProviderStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KubeProxy) DeepCopyInto(out *KubeProxy) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KubeProxy.
func (in *KubeProxy) DeepCopy() *KubeProxy {
	if in == nil {
		return nil
	}
	out := new(KubeProxy)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KubernetesMapping) DeepCopyInto(out *KubernetesMapping) {
	*out = *in
	if in.Groups != nil {
		in, out := &in.Groups, &out.Groups
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KubernetesMapping.
func (in *KubernetesMapping) DeepCopy() *KubernetesMapping {
	if in == nil {
		return nil
	}
	out := new(KubernetesMapping)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OIDCIdentityProviderConfig) DeepCopyInto(out *OIDCIdentityProviderConfig) {
	*out = *in
	if in.GroupsClaim != nil {
		in, out := &in.GroupsClaim, &out.GroupsClaim
		*out = new(string)
		**out = **in
	}
	if in.GroupsPrefix != nil {
		in, out := &in.GroupsPrefix, &out.GroupsPrefix
		*out = new(string)
		**out = **in
	}
	if in.RequiredClaims != nil {
		in, out := &in.RequiredClaims, &out.RequiredClaims
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.UsernameClaim != nil {
		in, out := &in.UsernameClaim, &out.UsernameClaim
		*out = new(string)
		**out = **in
	}
	if in.UsernamePrefix != nil {
		in, out := &in.UsernamePrefix, &out.UsernamePrefix
		*out = new(string)
		**out = **in
	}
	if in.Tags != nil {
		in, out := &in.Tags, &out.Tags
		*out = make(apiv1beta2.Tags, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OIDCIdentityProviderConfig.
func (in *OIDCIdentityProviderConfig) DeepCopy() *OIDCIdentityProviderConfig {
	if in == nil {
		return nil
	}
	out := new(OIDCIdentityProviderConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OIDCProviderStatus) DeepCopyInto(out *OIDCProviderStatus) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OIDCProviderStatus.
func (in *OIDCProviderStatus) DeepCopy() *OIDCProviderStatus {
	if in == nil {
		return nil
	}
	out := new(OIDCProviderStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RoleMapping) DeepCopyInto(out *RoleMapping) {
	*out = *in
	in.KubernetesMapping.DeepCopyInto(&out.KubernetesMapping)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RoleMapping.
func (in *RoleMapping) DeepCopy() *RoleMapping {
	if in == nil {
		return nil
	}
	out := new(RoleMapping)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *UserMapping) DeepCopyInto(out *UserMapping) {
	*out = *in
	in.KubernetesMapping.DeepCopyInto(&out.KubernetesMapping)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new UserMapping.
func (in *UserMapping) DeepCopy() *UserMapping {
	if in == nil {
		return nil
	}
	out := new(UserMapping)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VpcCni) DeepCopyInto(out *VpcCni) {
	*out = *in
	if in.Env != nil {
		in, out := &in.Env, &out.Env
		*out = make([]v1.EnvVar, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VpcCni.
func (in *VpcCni) DeepCopy() *VpcCni {
	if in == nil {
		return nil
	}
	out := new(VpcCni)
	in.DeepCopyInto(out)
	return out
}

/*
Copyright 2021 The Kubernetes Authors.

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

package v1beta1

import (
	apiconversion "k8s.io/apimachinery/pkg/conversion"
	infrav1beta1 "sigs.k8s.io/cluster-api-provider-aws/api/v1beta1"
	infrav1beta2 "sigs.k8s.io/cluster-api-provider-aws/api/v1beta2"
	ekscontrolplanev1 "sigs.k8s.io/cluster-api-provider-aws/controlplane/eks/api/v1beta2"
	"sigs.k8s.io/controller-runtime/pkg/conversion"
)

// ConvertTo converts the v1beta1 AWSManagedControlPlane receiver to a v1beta2 AWSManagedControlPlane.
func (r *AWSManagedControlPlane) ConvertTo(dstRaw conversion.Hub) error {
	dst := dstRaw.(*ekscontrolplanev1.AWSManagedControlPlane)

	if err := Convert_v1beta1_AWSManagedControlPlane_To_v1beta2_AWSManagedControlPlane(r, dst, nil); err != nil {
		return err
	}

	return nil
}

// ConvertFrom converts the v1beta2 AWSManagedControlPlane receiver to a v1beta1 AWSManagedControlPlane.
func (r *AWSManagedControlPlane) ConvertFrom(srcRaw conversion.Hub) error {
	src := srcRaw.(*ekscontrolplanev1.AWSManagedControlPlane)

	if err := Convert_v1beta2_AWSManagedControlPlane_To_v1beta1_AWSManagedControlPlane(src, r, nil); err != nil {
		return err
	}

	return nil
}

// ConvertTo converts the v1beta1 AWSManagedControlPlaneList receiver to a v1beta2 AWSManagedControlPlaneList.
func (r *AWSManagedControlPlaneList) ConvertTo(dstRaw conversion.Hub) error {
	dst := dstRaw.(*ekscontrolplanev1.AWSManagedControlPlaneList)

	return Convert_v1beta1_AWSManagedControlPlaneList_To_v1beta2_AWSManagedControlPlaneList(r, dst, nil)
}

// ConvertFrom converts the v1beta2 AWSManagedControlPlaneList receiver to a v1beta1 AWSManagedControlPlaneList.
func (r *AWSManagedControlPlaneList) ConvertFrom(srcRaw conversion.Hub) error {
	src := srcRaw.(*ekscontrolplanev1.AWSManagedControlPlaneList)

	return Convert_v1beta2_AWSManagedControlPlaneList_To_v1beta1_AWSManagedControlPlaneList(src, r, nil)
}

// Convert_v1beta1_NetworkSpec_To_v1beta2_NetworkSpec is a conversion function.
func Convert_v1beta1_NetworkSpec_To_v1beta2_NetworkSpec(in *infrav1beta1.NetworkSpec, out *infrav1beta2.NetworkSpec, s apiconversion.Scope) error {
	return infrav1beta1.Convert_v1beta1_NetworkSpec_To_v1beta2_NetworkSpec(in, out, s)
}

// Convert_v1beta2_NetworkSpec_To_v1beta1_NetworkSpec is a generated conversion function.
func Convert_v1beta2_NetworkSpec_To_v1beta1_NetworkSpec(in *infrav1beta2.NetworkSpec, out *infrav1beta1.NetworkSpec, s apiconversion.Scope) error {
	return infrav1beta1.Convert_v1beta2_NetworkSpec_To_v1beta1_NetworkSpec(in, out, s)
}

// Convert_v1beta1_NetworkStatus_To_v1beta2_NetworkStatus is a conversion function.
func Convert_v1beta1_NetworkStatus_To_v1beta2_NetworkStatus(in *infrav1beta1.NetworkStatus, out *infrav1beta2.NetworkStatus, s apiconversion.Scope) error {
	return infrav1beta1.Convert_v1beta1_NetworkStatus_To_v1beta2_NetworkStatus(in, out, s)
}

// Convert_v1beta2_NetworkStatus_To_v1beta1_NetworkStatus is a conversion function.
func Convert_v1beta2_NetworkStatus_To_v1beta1_NetworkStatus(in *infrav1beta2.NetworkStatus, out *infrav1beta1.NetworkStatus, s apiconversion.Scope) error {
	return infrav1beta1.Convert_v1beta2_NetworkStatus_To_v1beta1_NetworkStatus(in, out, s)
}

// Convert_v1beta1_Bastion_To_v1beta2_Bastion is a generated conversion function.
func Convert_v1beta1_Bastion_To_v1beta2_Bastion(in *infrav1beta1.Bastion, out *infrav1beta2.Bastion, s apiconversion.Scope) error {
	return infrav1beta1.Convert_v1beta1_Bastion_To_v1beta2_Bastion(in, out, s)
}

// Convert_v1beta2_Bastion_To_v1beta1_Bastion is a generated conversion function.
func Convert_v1beta2_Bastion_To_v1beta1_Bastion(in *infrav1beta2.Bastion, out *infrav1beta1.Bastion, s apiconversion.Scope) error {
	return infrav1beta1.Convert_v1beta2_Bastion_To_v1beta1_Bastion(in, out, s)
}

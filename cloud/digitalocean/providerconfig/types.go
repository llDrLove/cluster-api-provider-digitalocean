// Copyright © 2018 The Kubernetes Authors.
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package providerconfig

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// DigitalOceanMachineProviderConfig contains Config for DigitalOcean machines.
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
type DigitalOceanMachineProviderConfig struct {
	metav1.TypeMeta `json:",inline"`
}

// DigitalOceanClusterProviderConfig contains Config for DigitalOcean clusters.
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
type DigitalOceanClusterProviderConfig struct {
	metav1.TypeMeta `json:",inline"`
}

// DigitalOceanMachineProviderStatus is status of a DigitalOcean Machine.
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
type DigitalOceanMachineProviderStatus struct {
	metav1.TypeMeta `json:",inline"`
}

// DigitalOceanClusterProviderStatus is a status of a DigitalOcean cluster.
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
type DigitalOceanClusterProviderStatus struct {
	metav1.TypeMeta `json:",inline"`
}
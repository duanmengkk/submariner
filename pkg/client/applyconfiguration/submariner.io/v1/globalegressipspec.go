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

// Code generated by applyconfiguration-gen. DO NOT EDIT.

package v1

import (
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// GlobalEgressIPSpecApplyConfiguration represents an declarative configuration of the GlobalEgressIPSpec type for use
// with apply.
type GlobalEgressIPSpecApplyConfiguration struct {
	NumberOfIPs *int              `json:"numberOfIPs,omitempty"`
	PodSelector *v1.LabelSelector `json:"podSelector,omitempty"`
}

// GlobalEgressIPSpecApplyConfiguration constructs an declarative configuration of the GlobalEgressIPSpec type for use with
// apply.
func GlobalEgressIPSpec() *GlobalEgressIPSpecApplyConfiguration {
	return &GlobalEgressIPSpecApplyConfiguration{}
}

// WithNumberOfIPs sets the NumberOfIPs field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the NumberOfIPs field is set to the value of the last call.
func (b *GlobalEgressIPSpecApplyConfiguration) WithNumberOfIPs(value int) *GlobalEgressIPSpecApplyConfiguration {
	b.NumberOfIPs = &value
	return b
}

// WithPodSelector sets the PodSelector field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the PodSelector field is set to the value of the last call.
func (b *GlobalEgressIPSpecApplyConfiguration) WithPodSelector(value v1.LabelSelector) *GlobalEgressIPSpecApplyConfiguration {
	b.PodSelector = &value
	return b
}

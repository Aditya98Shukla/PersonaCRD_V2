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

package v2

// PersonaStatusApplyConfiguration represents an declarative configuration of the PersonaStatus type for use
// with apply.
type PersonaStatusApplyConfiguration struct {
	State      *string `json:"state,omitempty"`
	Allowed    *bool   `json:"allowed,omitempty"`
	ExpireDate *string `json:"expireDate,omitempty"`
}

// PersonaStatusApplyConfiguration constructs an declarative configuration of the PersonaStatus type for use with
// apply.
func PersonaStatus() *PersonaStatusApplyConfiguration {
	return &PersonaStatusApplyConfiguration{}
}

// WithState sets the State field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the State field is set to the value of the last call.
func (b *PersonaStatusApplyConfiguration) WithState(value string) *PersonaStatusApplyConfiguration {
	b.State = &value
	return b
}

// WithAllowed sets the Allowed field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Allowed field is set to the value of the last call.
func (b *PersonaStatusApplyConfiguration) WithAllowed(value bool) *PersonaStatusApplyConfiguration {
	b.Allowed = &value
	return b
}

// WithExpireDate sets the ExpireDate field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the ExpireDate field is set to the value of the last call.
func (b *PersonaStatusApplyConfiguration) WithExpireDate(value string) *PersonaStatusApplyConfiguration {
	b.ExpireDate = &value
	return b
}

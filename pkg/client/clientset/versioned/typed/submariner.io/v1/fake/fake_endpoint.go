/*
SPDX-License-Identifier: Apache-2.0

Copyright Contributors to the Submariner project.

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

// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	v1 "github.com/submariner-io/submariner/pkg/apis/submariner.io/v1"
	submarineriov1 "github.com/submariner-io/submariner/pkg/client/applyconfiguration/submariner.io/v1"
	typedsubmarineriov1 "github.com/submariner-io/submariner/pkg/client/clientset/versioned/typed/submariner.io/v1"
	gentype "k8s.io/client-go/gentype"
)

// fakeEndpoints implements EndpointInterface
type fakeEndpoints struct {
	*gentype.FakeClientWithListAndApply[*v1.Endpoint, *v1.EndpointList, *submarineriov1.EndpointApplyConfiguration]
	Fake *FakeSubmarinerV1
}

func newFakeEndpoints(fake *FakeSubmarinerV1, namespace string) typedsubmarineriov1.EndpointInterface {
	return &fakeEndpoints{
		gentype.NewFakeClientWithListAndApply[*v1.Endpoint, *v1.EndpointList, *submarineriov1.EndpointApplyConfiguration](
			fake.Fake,
			namespace,
			v1.SchemeGroupVersion.WithResource("endpoints"),
			v1.SchemeGroupVersion.WithKind("Endpoint"),
			func() *v1.Endpoint { return &v1.Endpoint{} },
			func() *v1.EndpointList { return &v1.EndpointList{} },
			func(dst, src *v1.EndpointList) { dst.ListMeta = src.ListMeta },
			func(list *v1.EndpointList) []*v1.Endpoint { return gentype.ToPointerSlice(list.Items) },
			func(list *v1.EndpointList, items []*v1.Endpoint) { list.Items = gentype.FromPointerSlice(items) },
		),
		fake,
	}
}

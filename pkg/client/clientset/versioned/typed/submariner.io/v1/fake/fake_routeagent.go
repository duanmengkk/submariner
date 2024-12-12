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

// fakeRouteAgents implements RouteAgentInterface
type fakeRouteAgents struct {
	*gentype.FakeClientWithListAndApply[*v1.RouteAgent, *v1.RouteAgentList, *submarineriov1.RouteAgentApplyConfiguration]
	Fake *FakeSubmarinerV1
}

func newFakeRouteAgents(fake *FakeSubmarinerV1, namespace string) typedsubmarineriov1.RouteAgentInterface {
	return &fakeRouteAgents{
		gentype.NewFakeClientWithListAndApply[*v1.RouteAgent, *v1.RouteAgentList, *submarineriov1.RouteAgentApplyConfiguration](
			fake.Fake,
			namespace,
			v1.SchemeGroupVersion.WithResource("routeagents"),
			v1.SchemeGroupVersion.WithKind("RouteAgent"),
			func() *v1.RouteAgent { return &v1.RouteAgent{} },
			func() *v1.RouteAgentList { return &v1.RouteAgentList{} },
			func(dst, src *v1.RouteAgentList) { dst.ListMeta = src.ListMeta },
			func(list *v1.RouteAgentList) []*v1.RouteAgent { return gentype.ToPointerSlice(list.Items) },
			func(list *v1.RouteAgentList, items []*v1.RouteAgent) { list.Items = gentype.FromPointerSlice(items) },
		),
		fake,
	}
}

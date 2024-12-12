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

// Code generated by lister-gen. DO NOT EDIT.

package v1

import (
	submarineriov1 "github.com/submariner-io/submariner/pkg/apis/submariner.io/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	listers "k8s.io/client-go/listers"
	cache "k8s.io/client-go/tools/cache"
)

// GlobalIngressIPLister helps list GlobalIngressIPs.
// All objects returned here must be treated as read-only.
type GlobalIngressIPLister interface {
	// List lists all GlobalIngressIPs in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*submarineriov1.GlobalIngressIP, err error)
	// GlobalIngressIPs returns an object that can list and get GlobalIngressIPs.
	GlobalIngressIPs(namespace string) GlobalIngressIPNamespaceLister
	GlobalIngressIPListerExpansion
}

// globalIngressIPLister implements the GlobalIngressIPLister interface.
type globalIngressIPLister struct {
	listers.ResourceIndexer[*submarineriov1.GlobalIngressIP]
}

// NewGlobalIngressIPLister returns a new GlobalIngressIPLister.
func NewGlobalIngressIPLister(indexer cache.Indexer) GlobalIngressIPLister {
	return &globalIngressIPLister{listers.New[*submarineriov1.GlobalIngressIP](indexer, submarineriov1.Resource("globalingressip"))}
}

// GlobalIngressIPs returns an object that can list and get GlobalIngressIPs.
func (s *globalIngressIPLister) GlobalIngressIPs(namespace string) GlobalIngressIPNamespaceLister {
	return globalIngressIPNamespaceLister{listers.NewNamespaced[*submarineriov1.GlobalIngressIP](s.ResourceIndexer, namespace)}
}

// GlobalIngressIPNamespaceLister helps list and get GlobalIngressIPs.
// All objects returned here must be treated as read-only.
type GlobalIngressIPNamespaceLister interface {
	// List lists all GlobalIngressIPs in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*submarineriov1.GlobalIngressIP, err error)
	// Get retrieves the GlobalIngressIP from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*submarineriov1.GlobalIngressIP, error)
	GlobalIngressIPNamespaceListerExpansion
}

// globalIngressIPNamespaceLister implements the GlobalIngressIPNamespaceLister
// interface.
type globalIngressIPNamespaceLister struct {
	listers.ResourceIndexer[*submarineriov1.GlobalIngressIP]
}

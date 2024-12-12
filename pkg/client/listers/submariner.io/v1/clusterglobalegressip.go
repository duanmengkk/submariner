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

// ClusterGlobalEgressIPLister helps list ClusterGlobalEgressIPs.
// All objects returned here must be treated as read-only.
type ClusterGlobalEgressIPLister interface {
	// List lists all ClusterGlobalEgressIPs in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*submarineriov1.ClusterGlobalEgressIP, err error)
	// ClusterGlobalEgressIPs returns an object that can list and get ClusterGlobalEgressIPs.
	ClusterGlobalEgressIPs(namespace string) ClusterGlobalEgressIPNamespaceLister
	ClusterGlobalEgressIPListerExpansion
}

// clusterGlobalEgressIPLister implements the ClusterGlobalEgressIPLister interface.
type clusterGlobalEgressIPLister struct {
	listers.ResourceIndexer[*submarineriov1.ClusterGlobalEgressIP]
}

// NewClusterGlobalEgressIPLister returns a new ClusterGlobalEgressIPLister.
func NewClusterGlobalEgressIPLister(indexer cache.Indexer) ClusterGlobalEgressIPLister {
	return &clusterGlobalEgressIPLister{listers.New[*submarineriov1.ClusterGlobalEgressIP](indexer, submarineriov1.Resource("clusterglobalegressip"))}
}

// ClusterGlobalEgressIPs returns an object that can list and get ClusterGlobalEgressIPs.
func (s *clusterGlobalEgressIPLister) ClusterGlobalEgressIPs(namespace string) ClusterGlobalEgressIPNamespaceLister {
	return clusterGlobalEgressIPNamespaceLister{listers.NewNamespaced[*submarineriov1.ClusterGlobalEgressIP](s.ResourceIndexer, namespace)}
}

// ClusterGlobalEgressIPNamespaceLister helps list and get ClusterGlobalEgressIPs.
// All objects returned here must be treated as read-only.
type ClusterGlobalEgressIPNamespaceLister interface {
	// List lists all ClusterGlobalEgressIPs in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*submarineriov1.ClusterGlobalEgressIP, err error)
	// Get retrieves the ClusterGlobalEgressIP from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*submarineriov1.ClusterGlobalEgressIP, error)
	ClusterGlobalEgressIPNamespaceListerExpansion
}

// clusterGlobalEgressIPNamespaceLister implements the ClusterGlobalEgressIPNamespaceLister
// interface.
type clusterGlobalEgressIPNamespaceLister struct {
	listers.ResourceIndexer[*submarineriov1.ClusterGlobalEgressIP]
}

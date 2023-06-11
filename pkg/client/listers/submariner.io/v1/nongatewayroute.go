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

// Code generated by lister-gen. DO NOT EDIT.

package v1

import (
	v1 "github.com/submariner-io/submariner/pkg/apis/submariner.io/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// NonGatewayRouteLister helps list NonGatewayRoutes.
// All objects returned here must be treated as read-only.
type NonGatewayRouteLister interface {
	// List lists all NonGatewayRoutes in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1.NonGatewayRoute, err error)
	// NonGatewayRoutes returns an object that can list and get NonGatewayRoutes.
	NonGatewayRoutes(namespace string) NonGatewayRouteNamespaceLister
	NonGatewayRouteListerExpansion
}

// nonGatewayRouteLister implements the NonGatewayRouteLister interface.
type nonGatewayRouteLister struct {
	indexer cache.Indexer
}

// NewNonGatewayRouteLister returns a new NonGatewayRouteLister.
func NewNonGatewayRouteLister(indexer cache.Indexer) NonGatewayRouteLister {
	return &nonGatewayRouteLister{indexer: indexer}
}

// List lists all NonGatewayRoutes in the indexer.
func (s *nonGatewayRouteLister) List(selector labels.Selector) (ret []*v1.NonGatewayRoute, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1.NonGatewayRoute))
	})
	return ret, err
}

// NonGatewayRoutes returns an object that can list and get NonGatewayRoutes.
func (s *nonGatewayRouteLister) NonGatewayRoutes(namespace string) NonGatewayRouteNamespaceLister {
	return nonGatewayRouteNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// NonGatewayRouteNamespaceLister helps list and get NonGatewayRoutes.
// All objects returned here must be treated as read-only.
type NonGatewayRouteNamespaceLister interface {
	// List lists all NonGatewayRoutes in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1.NonGatewayRoute, err error)
	// Get retrieves the NonGatewayRoute from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1.NonGatewayRoute, error)
	NonGatewayRouteNamespaceListerExpansion
}

// nonGatewayRouteNamespaceLister implements the NonGatewayRouteNamespaceLister
// interface.
type nonGatewayRouteNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all NonGatewayRoutes in the indexer for a given namespace.
func (s nonGatewayRouteNamespaceLister) List(selector labels.Selector) (ret []*v1.NonGatewayRoute, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1.NonGatewayRoute))
	})
	return ret, err
}

// Get retrieves the NonGatewayRoute from the indexer for a given namespace and name.
func (s nonGatewayRouteNamespaceLister) Get(name string) (*v1.NonGatewayRoute, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1.Resource("nongatewayroute"), name)
	}
	return obj.(*v1.NonGatewayRoute), nil
}
/*
The MIT License (MIT)

Copyright (c) 2016-2020 Containous SAS; 2020-2023 Traefik Labs

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in
all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
THE SOFTWARE.
*/

// Code generated by lister-gen. DO NOT EDIT.

package v1alpha1

import (
	v1alpha1 "github.com/traefik/traefik/v2/pkg/provider/kubernetes/crd/traefik/v1alpha1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// TraefikServiceLister helps list TraefikServices.
// All objects returned here must be treated as read-only.
type TraefikServiceLister interface {
	// List lists all TraefikServices in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.TraefikService, err error)
	// TraefikServices returns an object that can list and get TraefikServices.
	TraefikServices(namespace string) TraefikServiceNamespaceLister
	TraefikServiceListerExpansion
}

// traefikServiceLister implements the TraefikServiceLister interface.
type traefikServiceLister struct {
	indexer cache.Indexer
}

// NewTraefikServiceLister returns a new TraefikServiceLister.
func NewTraefikServiceLister(indexer cache.Indexer) TraefikServiceLister {
	return &traefikServiceLister{indexer: indexer}
}

// List lists all TraefikServices in the indexer.
func (s *traefikServiceLister) List(selector labels.Selector) (ret []*v1alpha1.TraefikService, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.TraefikService))
	})
	return ret, err
}

// TraefikServices returns an object that can list and get TraefikServices.
func (s *traefikServiceLister) TraefikServices(namespace string) TraefikServiceNamespaceLister {
	return traefikServiceNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// TraefikServiceNamespaceLister helps list and get TraefikServices.
// All objects returned here must be treated as read-only.
type TraefikServiceNamespaceLister interface {
	// List lists all TraefikServices in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.TraefikService, err error)
	// Get retrieves the TraefikService from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1alpha1.TraefikService, error)
	TraefikServiceNamespaceListerExpansion
}

// traefikServiceNamespaceLister implements the TraefikServiceNamespaceLister
// interface.
type traefikServiceNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all TraefikServices in the indexer for a given namespace.
func (s traefikServiceNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.TraefikService, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.TraefikService))
	})
	return ret, err
}

// Get retrieves the TraefikService from the indexer for a given namespace and name.
func (s traefikServiceNamespaceLister) Get(name string) (*v1alpha1.TraefikService, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("traefikservice"), name)
	}
	return obj.(*v1alpha1.TraefikService), nil
}
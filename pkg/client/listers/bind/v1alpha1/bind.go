/*
Copyright 2018 Google, Inc. All rights reserved.

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

// This file was automatically generated by lister-gen

package v1alpha1

import (
	v1alpha1 "github.com/elafros/eventing/pkg/apis/bind/v1alpha1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// BindLister helps list Binds.
type BindLister interface {
	// List lists all Binds in the indexer.
	List(selector labels.Selector) (ret []*v1alpha1.Bind, err error)
	// Binds returns an object that can list and get Binds.
	Binds(namespace string) BindNamespaceLister
	BindListerExpansion
}

// bindLister implements the BindLister interface.
type bindLister struct {
	indexer cache.Indexer
}

// NewBindLister returns a new BindLister.
func NewBindLister(indexer cache.Indexer) BindLister {
	return &bindLister{indexer: indexer}
}

// List lists all Binds in the indexer.
func (s *bindLister) List(selector labels.Selector) (ret []*v1alpha1.Bind, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.Bind))
	})
	return ret, err
}

// Binds returns an object that can list and get Binds.
func (s *bindLister) Binds(namespace string) BindNamespaceLister {
	return bindNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// BindNamespaceLister helps list and get Binds.
type BindNamespaceLister interface {
	// List lists all Binds in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v1alpha1.Bind, err error)
	// Get retrieves the Bind from the indexer for a given namespace and name.
	Get(name string) (*v1alpha1.Bind, error)
	BindNamespaceListerExpansion
}

// bindNamespaceLister implements the BindNamespaceLister
// interface.
type bindNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all Binds in the indexer for a given namespace.
func (s bindNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.Bind, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.Bind))
	})
	return ret, err
}

// Get retrieves the Bind from the indexer for a given namespace and name.
func (s bindNamespaceLister) Get(name string) (*v1alpha1.Bind, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("bind"), name)
	}
	return obj.(*v1alpha1.Bind), nil
}

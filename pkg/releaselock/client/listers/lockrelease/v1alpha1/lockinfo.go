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

package v1alpha1

import (
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
	v1alpha1 "sigs.k8s.io/gcp-filestore-csi-driver/pkg/apis/lockrelease/v1alpha1"
)

// LockInfoLister helps list LockInfos.
// All objects returned here must be treated as read-only.
type LockInfoLister interface {
	// List lists all LockInfos in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.LockInfo, err error)
	// LockInfos returns an object that can list and get LockInfos.
	LockInfos(namespace string) LockInfoNamespaceLister
	LockInfoListerExpansion
}

// lockInfoLister implements the LockInfoLister interface.
type lockInfoLister struct {
	indexer cache.Indexer
}

// NewLockInfoLister returns a new LockInfoLister.
func NewLockInfoLister(indexer cache.Indexer) LockInfoLister {
	return &lockInfoLister{indexer: indexer}
}

// List lists all LockInfos in the indexer.
func (s *lockInfoLister) List(selector labels.Selector) (ret []*v1alpha1.LockInfo, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.LockInfo))
	})
	return ret, err
}

// LockInfos returns an object that can list and get LockInfos.
func (s *lockInfoLister) LockInfos(namespace string) LockInfoNamespaceLister {
	return lockInfoNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// LockInfoNamespaceLister helps list and get LockInfos.
// All objects returned here must be treated as read-only.
type LockInfoNamespaceLister interface {
	// List lists all LockInfos in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.LockInfo, err error)
	// Get retrieves the LockInfo from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1alpha1.LockInfo, error)
	LockInfoNamespaceListerExpansion
}

// lockInfoNamespaceLister implements the LockInfoNamespaceLister
// interface.
type lockInfoNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all LockInfos in the indexer for a given namespace.
func (s lockInfoNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.LockInfo, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.LockInfo))
	})
	return ret, err
}

// Get retrieves the LockInfo from the indexer for a given namespace and name.
func (s lockInfoNamespaceLister) Get(name string) (*v1alpha1.LockInfo, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("lockinfo"), name)
	}
	return obj.(*v1alpha1.LockInfo), nil
}

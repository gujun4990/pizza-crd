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
	v1alpha1 "github.com/programming-kubernetes/pizza-crd/pkg/apis/restaurant/v1alpha1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// PizzaLister helps list Pizzas.
type PizzaLister interface {
	// List lists all Pizzas in the indexer.
	List(selector labels.Selector) (ret []*v1alpha1.Pizza, err error)
	// Pizzas returns an object that can list and get Pizzas.
	Pizzas(namespace string) PizzaNamespaceLister
	PizzaListerExpansion
}

// pizzaLister implements the PizzaLister interface.
type pizzaLister struct {
	indexer cache.Indexer
}

// NewPizzaLister returns a new PizzaLister.
func NewPizzaLister(indexer cache.Indexer) PizzaLister {
	return &pizzaLister{indexer: indexer}
}

// List lists all Pizzas in the indexer.
func (s *pizzaLister) List(selector labels.Selector) (ret []*v1alpha1.Pizza, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.Pizza))
	})
	return ret, err
}

// Pizzas returns an object that can list and get Pizzas.
func (s *pizzaLister) Pizzas(namespace string) PizzaNamespaceLister {
	return pizzaNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// PizzaNamespaceLister helps list and get Pizzas.
type PizzaNamespaceLister interface {
	// List lists all Pizzas in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v1alpha1.Pizza, err error)
	// Get retrieves the Pizza from the indexer for a given namespace and name.
	Get(name string) (*v1alpha1.Pizza, error)
	PizzaNamespaceListerExpansion
}

// pizzaNamespaceLister implements the PizzaNamespaceLister
// interface.
type pizzaNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all Pizzas in the indexer for a given namespace.
func (s pizzaNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.Pizza, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.Pizza))
	})
	return ret, err
}

// Get retrieves the Pizza from the indexer for a given namespace and name.
func (s pizzaNamespaceLister) Get(name string) (*v1alpha1.Pizza, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("pizza"), name)
	}
	return obj.(*v1alpha1.Pizza), nil
}
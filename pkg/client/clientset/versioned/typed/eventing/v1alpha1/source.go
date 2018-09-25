/*
Copyright 2018 The Knative Authors

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

package v1alpha1

import (
	v1alpha1 "github.com/knative/eventing/pkg/apis/eventing/v1alpha1"
	scheme "github.com/knative/eventing/pkg/client/clientset/versioned/scheme"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// SourcesGetter has a method to return a SourceInterface.
// A group's client should implement this interface.
type SourcesGetter interface {
	Sources(namespace string) SourceInterface
}

// SourceInterface has methods to work with Source resources.
type SourceInterface interface {
	Create(*v1alpha1.Source) (*v1alpha1.Source, error)
	Update(*v1alpha1.Source) (*v1alpha1.Source, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.Source, error)
	List(opts v1.ListOptions) (*v1alpha1.SourceList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.Source, err error)
	SourceExpansion
}

// sources implements SourceInterface
type sources struct {
	client rest.Interface
	ns     string
}

// newSources returns a Sources
func newSources(c *EventingV1alpha1Client, namespace string) *sources {
	return &sources{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the source, and returns the corresponding source object, and an error if there is any.
func (c *sources) Get(name string, options v1.GetOptions) (result *v1alpha1.Source, err error) {
	result = &v1alpha1.Source{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("sources").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of Sources that match those selectors.
func (c *sources) List(opts v1.ListOptions) (result *v1alpha1.SourceList, err error) {
	result = &v1alpha1.SourceList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("sources").
		VersionedParams(&opts, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested sources.
func (c *sources) Watch(opts v1.ListOptions) (watch.Interface, error) {
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("sources").
		VersionedParams(&opts, scheme.ParameterCodec).
		Watch()
}

// Create takes the representation of a source and creates it.  Returns the server's representation of the source, and an error, if there is any.
func (c *sources) Create(source *v1alpha1.Source) (result *v1alpha1.Source, err error) {
	result = &v1alpha1.Source{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("sources").
		Body(source).
		Do().
		Into(result)
	return
}

// Update takes the representation of a source and updates it. Returns the server's representation of the source, and an error, if there is any.
func (c *sources) Update(source *v1alpha1.Source) (result *v1alpha1.Source, err error) {
	result = &v1alpha1.Source{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("sources").
		Name(source.Name).
		Body(source).
		Do().
		Into(result)
	return
}

// Delete takes name of the source and deletes it. Returns an error if one occurs.
func (c *sources) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("sources").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *sources) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("sources").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched source.
func (c *sources) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.Source, err error) {
	result = &v1alpha1.Source{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("sources").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}

/*
Copyright 2019 The Knative Authors

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

// Code generated by informer-gen. DO NOT EDIT.

package v1alpha1

import (
	time "time"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
	flowsv1alpha1 "knative.dev/eventing/pkg/apis/flows/v1alpha1"
	versioned "knative.dev/eventing/pkg/client/clientset/versioned"
	internalinterfaces "knative.dev/eventing/pkg/client/informers/externalversions/internalinterfaces"
	v1alpha1 "knative.dev/eventing/pkg/client/listers/flows/v1alpha1"
)

// SequenceInformer provides access to a shared informer and lister for
// Sequences.
type SequenceInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1alpha1.SequenceLister
}

type sequenceInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
	namespace        string
}

// NewSequenceInformer constructs a new informer for Sequence type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewSequenceInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredSequenceInformer(client, namespace, resyncPeriod, indexers, nil)
}

// NewFilteredSequenceInformer constructs a new informer for Sequence type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredSequenceInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options v1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.FlowsV1alpha1().Sequences(namespace).List(options)
			},
			WatchFunc: func(options v1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.FlowsV1alpha1().Sequences(namespace).Watch(options)
			},
		},
		&flowsv1alpha1.Sequence{},
		resyncPeriod,
		indexers,
	)
}

func (f *sequenceInformer) defaultInformer(client versioned.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredSequenceInformer(client, f.namespace, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *sequenceInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&flowsv1alpha1.Sequence{}, f.defaultInformer)
}

func (f *sequenceInformer) Lister() v1alpha1.SequenceLister {
	return v1alpha1.NewSequenceLister(f.Informer().GetIndexer())
}

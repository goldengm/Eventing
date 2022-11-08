/*
Copyright 2021 The Knative Authors

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

// Code generated by injection-gen. DO NOT EDIT.

package filtered

import (
	context "context"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	cache "k8s.io/client-go/tools/cache"
	apismessagingv1 "knative.dev/eventing/pkg/apis/messaging/v1"
	versioned "knative.dev/eventing/pkg/client/clientset/versioned"
	v1 "knative.dev/eventing/pkg/client/informers/externalversions/messaging/v1"
	client "knative.dev/eventing/pkg/client/injection/client"
	filtered "knative.dev/eventing/pkg/client/injection/informers/factory/filtered"
	messagingv1 "knative.dev/eventing/pkg/client/listers/messaging/v1"
	controller "knative.dev/pkg/controller"
	injection "knative.dev/pkg/injection"
	logging "knative.dev/pkg/logging"
)

func init() {
	injection.Default.RegisterFilteredInformers(withInformer)
	injection.Dynamic.RegisterDynamicInformer(withDynamicInformer)
}

// Key is used for associating the Informer inside the context.Context.
type Key struct {
	Selector string
}

func withInformer(ctx context.Context) (context.Context, []controller.Informer) {
	untyped := ctx.Value(filtered.LabelKey{})
	if untyped == nil {
		logging.FromContext(ctx).Panic(
			"Unable to fetch labelkey from context.")
	}
	labelSelectors := untyped.([]string)
	infs := []controller.Informer{}
	for _, selector := range labelSelectors {
		f := filtered.Get(ctx, selector)
		inf := f.Messaging().V1().Channels()
		ctx = context.WithValue(ctx, Key{Selector: selector}, inf)
		infs = append(infs, inf.Informer())
	}
	return ctx, infs
}

func withDynamicInformer(ctx context.Context) context.Context {
	untyped := ctx.Value(filtered.LabelKey{})
	if untyped == nil {
		logging.FromContext(ctx).Panic(
			"Unable to fetch labelkey from context.")
	}
	labelSelectors := untyped.([]string)
	for _, selector := range labelSelectors {
		inf := &wrapper{client: client.Get(ctx), selector: selector}
		ctx = context.WithValue(ctx, Key{Selector: selector}, inf)
	}
	return ctx
}

// Get extracts the typed informer from the context.
func Get(ctx context.Context, selector string) v1.ChannelInformer {
	untyped := ctx.Value(Key{Selector: selector})
	if untyped == nil {
		logging.FromContext(ctx).Panicf(
			"Unable to fetch knative.dev/eventing/pkg/client/informers/externalversions/messaging/v1.ChannelInformer with selector %s from context.", selector)
	}
	return untyped.(v1.ChannelInformer)
}

type wrapper struct {
	client versioned.Interface

	namespace string

	selector string
}

var _ v1.ChannelInformer = (*wrapper)(nil)
var _ messagingv1.ChannelLister = (*wrapper)(nil)

func (w *wrapper) Informer() cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(nil, &apismessagingv1.Channel{}, 0, nil)
}

func (w *wrapper) Lister() messagingv1.ChannelLister {
	return w
}

func (w *wrapper) Channels(namespace string) messagingv1.ChannelNamespaceLister {
	return &wrapper{client: w.client, namespace: namespace, selector: w.selector}
}

func (w *wrapper) List(selector labels.Selector) (ret []*apismessagingv1.Channel, err error) {
	reqs, err := labels.ParseToRequirements(w.selector)
	if err != nil {
		return nil, err
	}
	selector = selector.Add(reqs...)
	lo, err := w.client.MessagingV1().Channels(w.namespace).List(context.TODO(), metav1.ListOptions{
		LabelSelector: selector.String(),
		// TODO(mattmoor): Incorporate resourceVersion bounds based on staleness criteria.
	})
	if err != nil {
		return nil, err
	}
	for idx := range lo.Items {
		ret = append(ret, &lo.Items[idx])
	}
	return ret, nil
}

func (w *wrapper) Get(name string) (*apismessagingv1.Channel, error) {
	// TODO(mattmoor): Check that the fetched object matches the selector.
	return w.client.MessagingV1().Channels(w.namespace).Get(context.TODO(), name, metav1.GetOptions{
		// TODO(mattmoor): Incorporate resourceVersion bounds based on staleness criteria.
	})
}

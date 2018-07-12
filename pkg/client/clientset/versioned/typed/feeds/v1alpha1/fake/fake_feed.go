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

package fake

import (
	v1alpha1 "github.com/knative/eventing/pkg/apis/feeds/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeFeeds implements FeedInterface
type FakeFeeds struct {
	Fake *FakeFeedsV1alpha1
	ns   string
}

var feedsResource = schema.GroupVersionResource{Group: "feeds.knative.dev", Version: "v1alpha1", Resource: "feeds"}

var feedsKind = schema.GroupVersionKind{Group: "feeds.knative.dev", Version: "v1alpha1", Kind: "Feed"}

// Get takes name of the feed, and returns the corresponding feed object, and an error if there is any.
func (c *FakeFeeds) Get(name string, options v1.GetOptions) (result *v1alpha1.Feed, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(feedsResource, c.ns, name), &v1alpha1.Feed{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Feed), err
}

// List takes label and field selectors, and returns the list of Feeds that match those selectors.
func (c *FakeFeeds) List(opts v1.ListOptions) (result *v1alpha1.FeedList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(feedsResource, feedsKind, c.ns, opts), &v1alpha1.FeedList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.FeedList{}
	for _, item := range obj.(*v1alpha1.FeedList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested feeds.
func (c *FakeFeeds) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(feedsResource, c.ns, opts))

}

// Create takes the representation of a feed and creates it.  Returns the server's representation of the feed, and an error, if there is any.
func (c *FakeFeeds) Create(feed *v1alpha1.Feed) (result *v1alpha1.Feed, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(feedsResource, c.ns, feed), &v1alpha1.Feed{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Feed), err
}

// Update takes the representation of a feed and updates it. Returns the server's representation of the feed, and an error, if there is any.
func (c *FakeFeeds) Update(feed *v1alpha1.Feed) (result *v1alpha1.Feed, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(feedsResource, c.ns, feed), &v1alpha1.Feed{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Feed), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeFeeds) UpdateStatus(feed *v1alpha1.Feed) (*v1alpha1.Feed, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(feedsResource, "status", c.ns, feed), &v1alpha1.Feed{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Feed), err
}

// Delete takes name of the feed and deletes it. Returns an error if one occurs.
func (c *FakeFeeds) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(feedsResource, c.ns, name), &v1alpha1.Feed{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeFeeds) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(feedsResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.FeedList{})
	return err
}

// Patch applies the patch and returns the patched feed.
func (c *FakeFeeds) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.Feed, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(feedsResource, c.ns, name, data, subresources...), &v1alpha1.Feed{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Feed), err
}

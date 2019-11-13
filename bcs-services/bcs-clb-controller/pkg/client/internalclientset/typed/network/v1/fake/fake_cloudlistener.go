/*
 * Tencent is pleased to support the open source community by making Blueking Container Service available.
 * Copyright (C) 2019 THL A29 Limited, a Tencent company. All rights reserved.
 * Licensed under the MIT License (the "License"); you may not use this file except
 * in compliance with the License. You may obtain a copy of the License at
 * http://opensource.org/licenses/MIT
 * Unless required by applicable law or agreed to in writing, software distributed under
 * the License is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND,
 * either express or implied. See the License for the specific language governing permissions and
 * limitations under the License.
 *
 */
// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	network_v1 "bk-bcs/bcs-services/bcs-clb-controller/pkg/apis/network/v1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeCloudListeners implements CloudListenerInterface
type FakeCloudListeners struct {
	Fake *FakeNetworkV1
	ns   string
}

var cloudlistenersResource = schema.GroupVersionResource{Group: "network.bmsf.tencent.com", Version: "v1", Resource: "cloudlisteners"}

var cloudlistenersKind = schema.GroupVersionKind{Group: "network.bmsf.tencent.com", Version: "v1", Kind: "CloudListener"}

// Get takes name of the cloudListener, and returns the corresponding cloudListener object, and an error if there is any.
func (c *FakeCloudListeners) Get(name string, options v1.GetOptions) (result *network_v1.CloudListener, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(cloudlistenersResource, c.ns, name), &network_v1.CloudListener{})

	if obj == nil {
		return nil, err
	}
	return obj.(*network_v1.CloudListener), err
}

// List takes label and field selectors, and returns the list of CloudListeners that match those selectors.
func (c *FakeCloudListeners) List(opts v1.ListOptions) (result *network_v1.CloudListenerList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(cloudlistenersResource, cloudlistenersKind, c.ns, opts), &network_v1.CloudListenerList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &network_v1.CloudListenerList{ListMeta: obj.(*network_v1.CloudListenerList).ListMeta}
	for _, item := range obj.(*network_v1.CloudListenerList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested cloudListeners.
func (c *FakeCloudListeners) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(cloudlistenersResource, c.ns, opts))

}

// Create takes the representation of a cloudListener and creates it.  Returns the server's representation of the cloudListener, and an error, if there is any.
func (c *FakeCloudListeners) Create(cloudListener *network_v1.CloudListener) (result *network_v1.CloudListener, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(cloudlistenersResource, c.ns, cloudListener), &network_v1.CloudListener{})

	if obj == nil {
		return nil, err
	}
	return obj.(*network_v1.CloudListener), err
}

// Update takes the representation of a cloudListener and updates it. Returns the server's representation of the cloudListener, and an error, if there is any.
func (c *FakeCloudListeners) Update(cloudListener *network_v1.CloudListener) (result *network_v1.CloudListener, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(cloudlistenersResource, c.ns, cloudListener), &network_v1.CloudListener{})

	if obj == nil {
		return nil, err
	}
	return obj.(*network_v1.CloudListener), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeCloudListeners) UpdateStatus(cloudListener *network_v1.CloudListener) (*network_v1.CloudListener, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(cloudlistenersResource, "status", c.ns, cloudListener), &network_v1.CloudListener{})

	if obj == nil {
		return nil, err
	}
	return obj.(*network_v1.CloudListener), err
}

// Delete takes name of the cloudListener and deletes it. Returns an error if one occurs.
func (c *FakeCloudListeners) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(cloudlistenersResource, c.ns, name), &network_v1.CloudListener{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeCloudListeners) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(cloudlistenersResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &network_v1.CloudListenerList{})
	return err
}

// Patch applies the patch and returns the patched cloudListener.
func (c *FakeCloudListeners) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *network_v1.CloudListener, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(cloudlistenersResource, c.ns, name, data, subresources...), &network_v1.CloudListener{})

	if obj == nil {
		return nil, err
	}
	return obj.(*network_v1.CloudListener), err
}

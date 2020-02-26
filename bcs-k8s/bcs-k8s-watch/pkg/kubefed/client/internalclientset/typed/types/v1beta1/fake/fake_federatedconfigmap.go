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

// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	v1beta1 "bk-bcs/bcs-k8s/bcs-k8s-watch/pkg/kubefed/apis/types/v1beta1"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeFederatedConfigMaps implements FederatedConfigMapInterface
type FakeFederatedConfigMaps struct {
	Fake *FakeTypesV1beta1
	ns   string
}

var federatedconfigmapsResource = schema.GroupVersionResource{Group: "types.kubefed.io", Version: "v1beta1", Resource: "federatedconfigmaps"}

var federatedconfigmapsKind = schema.GroupVersionKind{Group: "types.kubefed.io", Version: "v1beta1", Kind: "FederatedConfigMap"}

// Get takes name of the federatedConfigMap, and returns the corresponding federatedConfigMap object, and an error if there is any.
func (c *FakeFederatedConfigMaps) Get(name string, options v1.GetOptions) (result *v1beta1.FederatedConfigMap, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(federatedconfigmapsResource, c.ns, name), &v1beta1.FederatedConfigMap{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.FederatedConfigMap), err
}

// List takes label and field selectors, and returns the list of FederatedConfigMaps that match those selectors.
func (c *FakeFederatedConfigMaps) List(opts v1.ListOptions) (result *v1beta1.FederatedConfigMapList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(federatedconfigmapsResource, federatedconfigmapsKind, c.ns, opts), &v1beta1.FederatedConfigMapList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1beta1.FederatedConfigMapList{ListMeta: obj.(*v1beta1.FederatedConfigMapList).ListMeta}
	for _, item := range obj.(*v1beta1.FederatedConfigMapList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested federatedConfigMaps.
func (c *FakeFederatedConfigMaps) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(federatedconfigmapsResource, c.ns, opts))

}

// Create takes the representation of a federatedConfigMap and creates it.  Returns the server's representation of the federatedConfigMap, and an error, if there is any.
func (c *FakeFederatedConfigMaps) Create(federatedConfigMap *v1beta1.FederatedConfigMap) (result *v1beta1.FederatedConfigMap, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(federatedconfigmapsResource, c.ns, federatedConfigMap), &v1beta1.FederatedConfigMap{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.FederatedConfigMap), err
}

// Update takes the representation of a federatedConfigMap and updates it. Returns the server's representation of the federatedConfigMap, and an error, if there is any.
func (c *FakeFederatedConfigMaps) Update(federatedConfigMap *v1beta1.FederatedConfigMap) (result *v1beta1.FederatedConfigMap, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(federatedconfigmapsResource, c.ns, federatedConfigMap), &v1beta1.FederatedConfigMap{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.FederatedConfigMap), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeFederatedConfigMaps) UpdateStatus(federatedConfigMap *v1beta1.FederatedConfigMap) (*v1beta1.FederatedConfigMap, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(federatedconfigmapsResource, "status", c.ns, federatedConfigMap), &v1beta1.FederatedConfigMap{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.FederatedConfigMap), err
}

// Delete takes name of the federatedConfigMap and deletes it. Returns an error if one occurs.
func (c *FakeFederatedConfigMaps) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(federatedconfigmapsResource, c.ns, name), &v1beta1.FederatedConfigMap{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeFederatedConfigMaps) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(federatedconfigmapsResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &v1beta1.FederatedConfigMapList{})
	return err
}

// Patch applies the patch and returns the patched federatedConfigMap.
func (c *FakeFederatedConfigMaps) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1beta1.FederatedConfigMap, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(federatedconfigmapsResource, c.ns, name, data, subresources...), &v1beta1.FederatedConfigMap{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.FederatedConfigMap), err
}

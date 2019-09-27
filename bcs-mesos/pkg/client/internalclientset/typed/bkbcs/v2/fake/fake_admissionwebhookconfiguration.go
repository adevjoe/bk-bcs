// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	v2 "bk-bcs/bcs-mesos/pkg/apis/bkbcs/v2"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeAdmissionWebhookConfigurations implements AdmissionWebhookConfigurationInterface
type FakeAdmissionWebhookConfigurations struct {
	Fake *FakeBkbcsV2
	ns   string
}

var admissionwebhookconfigurationsResource = schema.GroupVersionResource{Group: "bkbcs.bkbcs.tencent.com", Version: "v2", Resource: "admissionwebhookconfigurations"}

var admissionwebhookconfigurationsKind = schema.GroupVersionKind{Group: "bkbcs.bkbcs.tencent.com", Version: "v2", Kind: "AdmissionWebhookConfiguration"}

// Get takes name of the admissionWebhookConfiguration, and returns the corresponding admissionWebhookConfiguration object, and an error if there is any.
func (c *FakeAdmissionWebhookConfigurations) Get(name string, options v1.GetOptions) (result *v2.AdmissionWebhookConfiguration, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(admissionwebhookconfigurationsResource, c.ns, name), &v2.AdmissionWebhookConfiguration{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v2.AdmissionWebhookConfiguration), err
}

// List takes label and field selectors, and returns the list of AdmissionWebhookConfigurations that match those selectors.
func (c *FakeAdmissionWebhookConfigurations) List(opts v1.ListOptions) (result *v2.AdmissionWebhookConfigurationList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(admissionwebhookconfigurationsResource, admissionwebhookconfigurationsKind, c.ns, opts), &v2.AdmissionWebhookConfigurationList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v2.AdmissionWebhookConfigurationList{ListMeta: obj.(*v2.AdmissionWebhookConfigurationList).ListMeta}
	for _, item := range obj.(*v2.AdmissionWebhookConfigurationList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested admissionWebhookConfigurations.
func (c *FakeAdmissionWebhookConfigurations) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(admissionwebhookconfigurationsResource, c.ns, opts))

}

// Create takes the representation of a admissionWebhookConfiguration and creates it.  Returns the server's representation of the admissionWebhookConfiguration, and an error, if there is any.
func (c *FakeAdmissionWebhookConfigurations) Create(admissionWebhookConfiguration *v2.AdmissionWebhookConfiguration) (result *v2.AdmissionWebhookConfiguration, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(admissionwebhookconfigurationsResource, c.ns, admissionWebhookConfiguration), &v2.AdmissionWebhookConfiguration{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v2.AdmissionWebhookConfiguration), err
}

// Update takes the representation of a admissionWebhookConfiguration and updates it. Returns the server's representation of the admissionWebhookConfiguration, and an error, if there is any.
func (c *FakeAdmissionWebhookConfigurations) Update(admissionWebhookConfiguration *v2.AdmissionWebhookConfiguration) (result *v2.AdmissionWebhookConfiguration, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(admissionwebhookconfigurationsResource, c.ns, admissionWebhookConfiguration), &v2.AdmissionWebhookConfiguration{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v2.AdmissionWebhookConfiguration), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeAdmissionWebhookConfigurations) UpdateStatus(admissionWebhookConfiguration *v2.AdmissionWebhookConfiguration) (*v2.AdmissionWebhookConfiguration, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(admissionwebhookconfigurationsResource, "status", c.ns, admissionWebhookConfiguration), &v2.AdmissionWebhookConfiguration{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v2.AdmissionWebhookConfiguration), err
}

// Delete takes name of the admissionWebhookConfiguration and deletes it. Returns an error if one occurs.
func (c *FakeAdmissionWebhookConfigurations) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(admissionwebhookconfigurationsResource, c.ns, name), &v2.AdmissionWebhookConfiguration{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeAdmissionWebhookConfigurations) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(admissionwebhookconfigurationsResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &v2.AdmissionWebhookConfigurationList{})
	return err
}

// Patch applies the patch and returns the patched admissionWebhookConfiguration.
func (c *FakeAdmissionWebhookConfigurations) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v2.AdmissionWebhookConfiguration, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(admissionwebhookconfigurationsResource, c.ns, name, data, subresources...), &v2.AdmissionWebhookConfiguration{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v2.AdmissionWebhookConfiguration), err
}

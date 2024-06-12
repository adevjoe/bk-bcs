/*
Copyright 2023.

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
	"context"

	v1 "github.com/Tencent/bk-bcs/bcs-scenarios/bcs-gitops-workflow/pkg/apis/gitopsworkflow/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeWorkflows implements WorkflowInterface
type FakeWorkflows struct {
	Fake *FakeGitopsworkflowV1
	ns   string
}

var workflowsResource = v1.SchemeGroupVersion.WithResource("workflows")

var workflowsKind = v1.SchemeGroupVersion.WithKind("Workflow")

// Get takes name of the workflow, and returns the corresponding workflow object, and an error if there is any.
func (c *FakeWorkflows) Get(ctx context.Context, name string, options metav1.GetOptions) (result *v1.Workflow, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(workflowsResource, c.ns, name), &v1.Workflow{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1.Workflow), err
}

// List takes label and field selectors, and returns the list of Workflows that match those selectors.
func (c *FakeWorkflows) List(ctx context.Context, opts metav1.ListOptions) (result *v1.WorkflowList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(workflowsResource, workflowsKind, c.ns, opts), &v1.WorkflowList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1.WorkflowList{ListMeta: obj.(*v1.WorkflowList).ListMeta}
	for _, item := range obj.(*v1.WorkflowList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested workflows.
func (c *FakeWorkflows) Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(workflowsResource, c.ns, opts))

}

// Create takes the representation of a workflow and creates it.  Returns the server's representation of the workflow, and an error, if there is any.
func (c *FakeWorkflows) Create(ctx context.Context, workflow *v1.Workflow, opts metav1.CreateOptions) (result *v1.Workflow, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(workflowsResource, c.ns, workflow), &v1.Workflow{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1.Workflow), err
}

// Update takes the representation of a workflow and updates it. Returns the server's representation of the workflow, and an error, if there is any.
func (c *FakeWorkflows) Update(ctx context.Context, workflow *v1.Workflow, opts metav1.UpdateOptions) (result *v1.Workflow, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(workflowsResource, c.ns, workflow), &v1.Workflow{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1.Workflow), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeWorkflows) UpdateStatus(ctx context.Context, workflow *v1.Workflow, opts metav1.UpdateOptions) (*v1.Workflow, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(workflowsResource, "status", c.ns, workflow), &v1.Workflow{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1.Workflow), err
}

// Delete takes name of the workflow and deletes it. Returns an error if one occurs.
func (c *FakeWorkflows) Delete(ctx context.Context, name string, opts metav1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteActionWithOptions(workflowsResource, c.ns, name, opts), &v1.Workflow{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeWorkflows) DeleteCollection(ctx context.Context, opts metav1.DeleteOptions, listOpts metav1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(workflowsResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1.WorkflowList{})
	return err
}

// Patch applies the patch and returns the patched workflow.
func (c *FakeWorkflows) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (result *v1.Workflow, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(workflowsResource, c.ns, name, pt, data, subresources...), &v1.Workflow{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1.Workflow), err
}

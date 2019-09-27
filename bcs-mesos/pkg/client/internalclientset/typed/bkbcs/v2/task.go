// Code generated by client-gen. DO NOT EDIT.

package v2

import (
	v2 "bk-bcs/bcs-mesos/pkg/apis/bkbcs/v2"
	scheme "bk-bcs/bcs-mesos/pkg/client/internalclientset/scheme"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// TasksGetter has a method to return a TaskInterface.
// A group's client should implement this interface.
type TasksGetter interface {
	Tasks(namespace string) TaskInterface
}

// TaskInterface has methods to work with Task resources.
type TaskInterface interface {
	Create(*v2.Task) (*v2.Task, error)
	Update(*v2.Task) (*v2.Task, error)
	UpdateStatus(*v2.Task) (*v2.Task, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v2.Task, error)
	List(opts v1.ListOptions) (*v2.TaskList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v2.Task, err error)
	TaskExpansion
}

// tasks implements TaskInterface
type tasks struct {
	client rest.Interface
	ns     string
}

// newTasks returns a Tasks
func newTasks(c *BkbcsV2Client, namespace string) *tasks {
	return &tasks{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the task, and returns the corresponding task object, and an error if there is any.
func (c *tasks) Get(name string, options v1.GetOptions) (result *v2.Task, err error) {
	result = &v2.Task{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("tasks").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of Tasks that match those selectors.
func (c *tasks) List(opts v1.ListOptions) (result *v2.TaskList, err error) {
	result = &v2.TaskList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("tasks").
		VersionedParams(&opts, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested tasks.
func (c *tasks) Watch(opts v1.ListOptions) (watch.Interface, error) {
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("tasks").
		VersionedParams(&opts, scheme.ParameterCodec).
		Watch()
}

// Create takes the representation of a task and creates it.  Returns the server's representation of the task, and an error, if there is any.
func (c *tasks) Create(task *v2.Task) (result *v2.Task, err error) {
	result = &v2.Task{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("tasks").
		Body(task).
		Do().
		Into(result)
	return
}

// Update takes the representation of a task and updates it. Returns the server's representation of the task, and an error, if there is any.
func (c *tasks) Update(task *v2.Task) (result *v2.Task, err error) {
	result = &v2.Task{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("tasks").
		Name(task.Name).
		Body(task).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *tasks) UpdateStatus(task *v2.Task) (result *v2.Task, err error) {
	result = &v2.Task{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("tasks").
		Name(task.Name).
		SubResource("status").
		Body(task).
		Do().
		Into(result)
	return
}

// Delete takes name of the task and deletes it. Returns an error if one occurs.
func (c *tasks) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("tasks").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *tasks) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("tasks").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched task.
func (c *tasks) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v2.Task, err error) {
	result = &v2.Task{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("tasks").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}

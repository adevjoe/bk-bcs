// Code generated by lister-gen. DO NOT EDIT.

package v2

import (
	v2 "bk-bcs/bcs-mesos/pkg/apis/bkbcs/v2"

	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// TaskLister helps list Tasks.
type TaskLister interface {
	// List lists all Tasks in the indexer.
	List(selector labels.Selector) (ret []*v2.Task, err error)
	// Tasks returns an object that can list and get Tasks.
	Tasks(namespace string) TaskNamespaceLister
	TaskListerExpansion
}

// taskLister implements the TaskLister interface.
type taskLister struct {
	indexer cache.Indexer
}

// NewTaskLister returns a new TaskLister.
func NewTaskLister(indexer cache.Indexer) TaskLister {
	return &taskLister{indexer: indexer}
}

// List lists all Tasks in the indexer.
func (s *taskLister) List(selector labels.Selector) (ret []*v2.Task, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v2.Task))
	})
	return ret, err
}

// Tasks returns an object that can list and get Tasks.
func (s *taskLister) Tasks(namespace string) TaskNamespaceLister {
	return taskNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// TaskNamespaceLister helps list and get Tasks.
type TaskNamespaceLister interface {
	// List lists all Tasks in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v2.Task, err error)
	// Get retrieves the Task from the indexer for a given namespace and name.
	Get(name string) (*v2.Task, error)
	TaskNamespaceListerExpansion
}

// taskNamespaceLister implements the TaskNamespaceLister
// interface.
type taskNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all Tasks in the indexer for a given namespace.
func (s taskNamespaceLister) List(selector labels.Selector) (ret []*v2.Task, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v2.Task))
	})
	return ret, err
}

// Get retrieves the Task from the indexer for a given namespace and name.
func (s taskNamespaceLister) Get(name string) (*v2.Task, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v2.Resource("task"), name)
	}
	return obj.(*v2.Task), nil
}

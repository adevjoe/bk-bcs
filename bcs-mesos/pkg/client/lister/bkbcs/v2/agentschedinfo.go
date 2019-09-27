// Code generated by lister-gen. DO NOT EDIT.

package v2

import (
	v2 "bk-bcs/bcs-mesos/pkg/apis/bkbcs/v2"

	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// AgentSchedInfoLister helps list AgentSchedInfos.
type AgentSchedInfoLister interface {
	// List lists all AgentSchedInfos in the indexer.
	List(selector labels.Selector) (ret []*v2.AgentSchedInfo, err error)
	// AgentSchedInfos returns an object that can list and get AgentSchedInfos.
	AgentSchedInfos(namespace string) AgentSchedInfoNamespaceLister
	AgentSchedInfoListerExpansion
}

// agentSchedInfoLister implements the AgentSchedInfoLister interface.
type agentSchedInfoLister struct {
	indexer cache.Indexer
}

// NewAgentSchedInfoLister returns a new AgentSchedInfoLister.
func NewAgentSchedInfoLister(indexer cache.Indexer) AgentSchedInfoLister {
	return &agentSchedInfoLister{indexer: indexer}
}

// List lists all AgentSchedInfos in the indexer.
func (s *agentSchedInfoLister) List(selector labels.Selector) (ret []*v2.AgentSchedInfo, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v2.AgentSchedInfo))
	})
	return ret, err
}

// AgentSchedInfos returns an object that can list and get AgentSchedInfos.
func (s *agentSchedInfoLister) AgentSchedInfos(namespace string) AgentSchedInfoNamespaceLister {
	return agentSchedInfoNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// AgentSchedInfoNamespaceLister helps list and get AgentSchedInfos.
type AgentSchedInfoNamespaceLister interface {
	// List lists all AgentSchedInfos in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v2.AgentSchedInfo, err error)
	// Get retrieves the AgentSchedInfo from the indexer for a given namespace and name.
	Get(name string) (*v2.AgentSchedInfo, error)
	AgentSchedInfoNamespaceListerExpansion
}

// agentSchedInfoNamespaceLister implements the AgentSchedInfoNamespaceLister
// interface.
type agentSchedInfoNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all AgentSchedInfos in the indexer for a given namespace.
func (s agentSchedInfoNamespaceLister) List(selector labels.Selector) (ret []*v2.AgentSchedInfo, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v2.AgentSchedInfo))
	})
	return ret, err
}

// Get retrieves the AgentSchedInfo from the indexer for a given namespace and name.
func (s agentSchedInfoNamespaceLister) Get(name string) (*v2.AgentSchedInfo, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v2.Resource("agentschedinfo"), name)
	}
	return obj.(*v2.AgentSchedInfo), nil
}

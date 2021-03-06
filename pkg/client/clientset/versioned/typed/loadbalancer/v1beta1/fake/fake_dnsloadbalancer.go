/*
SPDX-FileCopyrightText: 2019 SAP SE or an SAP affiliate company and Gardener contributors

SPDX-License-Identifier: Apache-2.0
*/
// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	v1beta1 "github.com/gardener/dnslb-controller-manager/pkg/apis/loadbalancer/v1beta1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeDNSLoadBalancers implements DNSLoadBalancerInterface
type FakeDNSLoadBalancers struct {
	Fake *FakeLoadbalancerV1beta1
	ns   string
}

var dnsloadbalancersResource = schema.GroupVersionResource{Group: "loadbalancer.gardener.cloud", Version: "v1beta1", Resource: "dnsloadbalancers"}

var dnsloadbalancersKind = schema.GroupVersionKind{Group: "loadbalancer.gardener.cloud", Version: "v1beta1", Kind: "DNSLoadBalancer"}

// Get takes name of the dNSLoadBalancer, and returns the corresponding dNSLoadBalancer object, and an error if there is any.
func (c *FakeDNSLoadBalancers) Get(name string, options v1.GetOptions) (result *v1beta1.DNSLoadBalancer, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(dnsloadbalancersResource, c.ns, name), &v1beta1.DNSLoadBalancer{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.DNSLoadBalancer), err
}

// List takes label and field selectors, and returns the list of DNSLoadBalancers that match those selectors.
func (c *FakeDNSLoadBalancers) List(opts v1.ListOptions) (result *v1beta1.DNSLoadBalancerList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(dnsloadbalancersResource, dnsloadbalancersKind, c.ns, opts), &v1beta1.DNSLoadBalancerList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1beta1.DNSLoadBalancerList{ListMeta: obj.(*v1beta1.DNSLoadBalancerList).ListMeta}
	for _, item := range obj.(*v1beta1.DNSLoadBalancerList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested dNSLoadBalancers.
func (c *FakeDNSLoadBalancers) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(dnsloadbalancersResource, c.ns, opts))

}

// Create takes the representation of a dNSLoadBalancer and creates it.  Returns the server's representation of the dNSLoadBalancer, and an error, if there is any.
func (c *FakeDNSLoadBalancers) Create(dNSLoadBalancer *v1beta1.DNSLoadBalancer) (result *v1beta1.DNSLoadBalancer, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(dnsloadbalancersResource, c.ns, dNSLoadBalancer), &v1beta1.DNSLoadBalancer{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.DNSLoadBalancer), err
}

// Update takes the representation of a dNSLoadBalancer and updates it. Returns the server's representation of the dNSLoadBalancer, and an error, if there is any.
func (c *FakeDNSLoadBalancers) Update(dNSLoadBalancer *v1beta1.DNSLoadBalancer) (result *v1beta1.DNSLoadBalancer, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(dnsloadbalancersResource, c.ns, dNSLoadBalancer), &v1beta1.DNSLoadBalancer{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.DNSLoadBalancer), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeDNSLoadBalancers) UpdateStatus(dNSLoadBalancer *v1beta1.DNSLoadBalancer) (*v1beta1.DNSLoadBalancer, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(dnsloadbalancersResource, "status", c.ns, dNSLoadBalancer), &v1beta1.DNSLoadBalancer{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.DNSLoadBalancer), err
}

// Delete takes name of the dNSLoadBalancer and deletes it. Returns an error if one occurs.
func (c *FakeDNSLoadBalancers) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(dnsloadbalancersResource, c.ns, name), &v1beta1.DNSLoadBalancer{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeDNSLoadBalancers) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(dnsloadbalancersResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &v1beta1.DNSLoadBalancerList{})
	return err
}

// Patch applies the patch and returns the patched dNSLoadBalancer.
func (c *FakeDNSLoadBalancers) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1beta1.DNSLoadBalancer, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(dnsloadbalancersResource, c.ns, name, data, subresources...), &v1beta1.DNSLoadBalancer{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.DNSLoadBalancer), err
}

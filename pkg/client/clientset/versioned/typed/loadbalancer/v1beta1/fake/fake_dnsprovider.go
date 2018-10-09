/*
Copyright (c) 2018 SAP SE or an SAP affiliate company. All rights reserved. This file is licensed under the Apache Software License, v. 2 except as noted otherwise in the LICENSE file

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
	v1beta1 "github.com/gardener/dnslb-controller-manager/pkg/apis/loadbalancer/v1beta1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeDNSProviders implements DNSProviderInterface
type FakeDNSProviders struct {
	Fake *FakeLoadbalancerV1beta1
	ns   string
}

var dnsprovidersResource = schema.GroupVersionResource{Group: "loadbalancer.gardener.cloud", Version: "v1beta1", Resource: "dnsproviders"}

var dnsprovidersKind = schema.GroupVersionKind{Group: "loadbalancer.gardener.cloud", Version: "v1beta1", Kind: "DNSProvider"}

// Get takes name of the dNSProvider, and returns the corresponding dNSProvider object, and an error if there is any.
func (c *FakeDNSProviders) Get(name string, options v1.GetOptions) (result *v1beta1.DNSProvider, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(dnsprovidersResource, c.ns, name), &v1beta1.DNSProvider{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.DNSProvider), err
}

// List takes label and field selectors, and returns the list of DNSProviders that match those selectors.
func (c *FakeDNSProviders) List(opts v1.ListOptions) (result *v1beta1.DNSProviderList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(dnsprovidersResource, dnsprovidersKind, c.ns, opts), &v1beta1.DNSProviderList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1beta1.DNSProviderList{ListMeta: obj.(*v1beta1.DNSProviderList).ListMeta}
	for _, item := range obj.(*v1beta1.DNSProviderList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested dNSProviders.
func (c *FakeDNSProviders) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(dnsprovidersResource, c.ns, opts))

}

// Create takes the representation of a dNSProvider and creates it.  Returns the server's representation of the dNSProvider, and an error, if there is any.
func (c *FakeDNSProviders) Create(dNSProvider *v1beta1.DNSProvider) (result *v1beta1.DNSProvider, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(dnsprovidersResource, c.ns, dNSProvider), &v1beta1.DNSProvider{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.DNSProvider), err
}

// Update takes the representation of a dNSProvider and updates it. Returns the server's representation of the dNSProvider, and an error, if there is any.
func (c *FakeDNSProviders) Update(dNSProvider *v1beta1.DNSProvider) (result *v1beta1.DNSProvider, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(dnsprovidersResource, c.ns, dNSProvider), &v1beta1.DNSProvider{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.DNSProvider), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeDNSProviders) UpdateStatus(dNSProvider *v1beta1.DNSProvider) (*v1beta1.DNSProvider, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(dnsprovidersResource, "status", c.ns, dNSProvider), &v1beta1.DNSProvider{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.DNSProvider), err
}

// Delete takes name of the dNSProvider and deletes it. Returns an error if one occurs.
func (c *FakeDNSProviders) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(dnsprovidersResource, c.ns, name), &v1beta1.DNSProvider{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeDNSProviders) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(dnsprovidersResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &v1beta1.DNSProviderList{})
	return err
}

// Patch applies the patch and returns the patched dNSProvider.
func (c *FakeDNSProviders) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1beta1.DNSProvider, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(dnsprovidersResource, c.ns, name, data, subresources...), &v1beta1.DNSProvider{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.DNSProvider), err
}

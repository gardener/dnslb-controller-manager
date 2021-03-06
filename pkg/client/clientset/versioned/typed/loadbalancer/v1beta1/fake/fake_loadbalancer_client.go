/*
SPDX-FileCopyrightText: 2019 SAP SE or an SAP affiliate company and Gardener contributors

SPDX-License-Identifier: Apache-2.0
*/
// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	v1beta1 "github.com/gardener/dnslb-controller-manager/pkg/client/clientset/versioned/typed/loadbalancer/v1beta1"
	rest "k8s.io/client-go/rest"
	testing "k8s.io/client-go/testing"
)

type FakeLoadbalancerV1beta1 struct {
	*testing.Fake
}

func (c *FakeLoadbalancerV1beta1) DNSLoadBalancers(namespace string) v1beta1.DNSLoadBalancerInterface {
	return &FakeDNSLoadBalancers{c, namespace}
}

func (c *FakeLoadbalancerV1beta1) DNSLoadBalancerEndpoints(namespace string) v1beta1.DNSLoadBalancerEndpointInterface {
	return &FakeDNSLoadBalancerEndpoints{c, namespace}
}

// RESTClient returns a RESTClient that is used to communicate
// with API server by this client implementation.
func (c *FakeLoadbalancerV1beta1) RESTClient() rest.Interface {
	var ret *rest.RESTClient
	return ret
}

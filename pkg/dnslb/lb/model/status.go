// Copyright 2018 SAP SE or an SAP affiliate company. All rights reserved. This file is licensed under the Apache Software License, v. 2 except as noted otherwise in the LICENSE file
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package model

import (
	"fmt"
	"reflect"
	"sort"

	api "github.com/gardener/dnslb-controller-manager/pkg/apis/loadbalancer/v1beta1"
	lbutils "github.com/gardener/dnslb-controller-manager/pkg/dnslb/utils"

	"github.com/gardener/dnslb-controller-manager/pkg/server/metrics"
	corev1 "k8s.io/api/core/v1"
)

type DNSDone struct {
	dnslb     *lbutils.DNSLoadBalancerObject
	model     *Model
	done      bool
	message   string
	hcount    int
	ishealthy bool
	invalid   bool
	active    map[string]*lbutils.DNSLoadBalancerEndpointObject
	healthy   map[string]*lbutils.DNSLoadBalancerEndpointObject
	unhealthy map[string]*lbutils.DNSLoadBalancerEndpointObject
}

func NewStatusUpdate(m *Model, w *Watch) *DNSDone {
	return &DNSDone{
		model:     m,
		dnslb:     w.DNSLB,
		active:    map[string]*lbutils.DNSLoadBalancerEndpointObject{},
		healthy:   map[string]*lbutils.DNSLoadBalancerEndpointObject{},
		unhealthy: map[string]*lbutils.DNSLoadBalancerEndpointObject{},
	}
}

func (this *DNSDone) SetInvalid() {
	this.invalid = true
}

func (this *DNSDone) IsInvalid() bool {
	return this.invalid
}

func (this *DNSDone) IsHealthy() bool {
	return this.ishealthy
}

func (this *DNSDone) SetHealthy(a bool) *DNSDone {
	this.ishealthy = a
	return this
}

func (this *DNSDone) SetMessage(msg string) *DNSDone {
	this.message = msg
	return this
}

func (this *DNSDone) AddHealthyTarget(target *Target) {
	this.hcount++
	if target.DNSEP != nil {
		this.healthy[target.DNSEP.GetName()] = target.DNSEP
	}
}

func (this *DNSDone) AddActiveTarget(target *Target) {
	this.hcount++
	if target.DNSEP != nil {
		this.active[target.DNSEP.GetName()] = target.DNSEP
	}
}

func (this *DNSDone) AddUnhealthyTarget(target *Target) {
	if target.DNSEP != nil {
		this.unhealthy[target.DNSEP.GetName()] = target.DNSEP
	}
}

func (this *DNSDone) HasHealthy() bool {
	return this.hcount != 0
}

func (this *DNSDone) Eventf(ty, reason string, msgfmt string, args ...interface{}) {
	if this.dnslb != nil {
		this.dnslb.Eventf(ty, reason, msgfmt, args...)
	}
}

func (this *DNSDone) Event(ty, reason string, msg string) {
	if this.dnslb != nil {
		this.dnslb.Event(ty, reason, msg)
	}
}

func (this *DNSDone) updateStatus() {
	if this.dnslb != nil {
		this._updateLoadBalancerStatus(true, "")
		for _, t := range this.healthy {
			this._updateEndpointStatus(t, true, this.active[t.GetName()] != nil)
		}
		for _, t := range this.unhealthy {
			this._updateEndpointStatus(t, false, false)
		}
	}
}

func (this *DNSDone) _updateLoadBalancerStatus(activeupd bool, fail string) {
	dnslb := this.dnslb.Copy()
	status:=dnslb.Status()
	if this.ishealthy {
		status.State = "healthy"
	} else {
		status.State = "unreachable"
	}
	status.Message = fail
	if activeupd {
		if len(this.active) > 0 {
			status.Active = []api.DNSLoadBalancerActive{}
			keys := []string{}
			for _, t := range this.active {
				keys = append(keys, t.GetName())
			}
			sort.Strings(keys)
			for _, k := range keys {
				t := this.active[k]
				status.Active = append(status.Active,
					api.DNSLoadBalancerActive{
						Name:      t.GetName(),
						IPAddress: t.Spec().IPAddress,
						CName:     t.Spec().CName,
					})
			}
		} else {
			status.Active = nil
		}
	}
	if this.invalid {
		status.Active = nil
	}
	if !reflect.DeepEqual(this.dnslb.Status(), status) {
		this.model.Infof("old: %#v", this.dnslb.Status())
		this.model.Infof("new: %#v", status)
		this.model.Infof("updating status for dns load balancer %s/%s", dnslb.GetNamespace(), dnslb.GetName())
		err := dnslb.Update()
		if err != nil {
			this.model.Errorf("cannot update dns load balancer status for %s/%s: %s", dnslb.GetNamespace(), dnslb.GetName(), err)
		}
	}
}

func (this *DNSDone) _updateEndpointStatus(ep *lbutils.DNSLoadBalancerEndpointObject, healthy, active bool) {
	if ep.Status().Healthy != healthy || ep.Status().Active != active {
		dnsep := ep.Copy()
		dnsep.Status().Healthy = healthy
		dnsep.Status().Active = active
		this.model.Infof("updating status for endpoint %s/%s: healthy %t, active %t", dnsep.GetNamespace(), dnsep.GetName(), healthy, active)
		err := ep.Update()
		if err != nil {
			this.model.Errorf("cannot update dns endpoint status for %s/%s: %s", dnsep.GetNamespace(), dnsep.GetName(), err)
		}
	}

	metrics.ReportActiveEndpoint(this.dnslb.ObjectName(), ep.ObjectName(), active)
}

///////////////////////////////////////

func (this *DNSDone) Failed(err error) {
	this.Error(false, err)
}

func (this *DNSDone) Error(activeupd bool, err error) {
	if !this.done {
		this.done = true
		msg := ""
		if this.message != "" {
			msg = fmt.Sprintf("%s: %s", this.message, err)
		} else {
			msg = err.Error()
		}
		this.Event(corev1.EventTypeWarning, "sync", msg)
		this._updateLoadBalancerStatus(activeupd, msg)
	}
}

func (this *DNSDone) Succeeded() {
	if !this.done {
		this.done = true
		if this.message != "" {
			this.Eventf(corev1.EventTypeNormal, "sync", "%s", this.message)
		}
		if this.IsHealthy() {
			this.Eventf(corev1.EventTypeNormal, "sync", "healthy again")
		}
		this.updateStatus()
	}
}

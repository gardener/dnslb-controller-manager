// Code generated by informer-gen. DO NOT EDIT.

package v1alpha1

import (
	internalinterfaces "github.com/gardener/controller-manager-library/pkg/client/gardenextensions/informers/externalversions/internalinterfaces"
)

// Interface provides access to all the informers in this group version.
type Interface interface {
	// ControlPlanes returns a ControlPlaneInformer.
	ControlPlanes() ControlPlaneInformer
	// DNSs returns a DNSInformer.
	DNSs() DNSInformer
	// Infrastructures returns a InfrastructureInformer.
	Infrastructures() InfrastructureInformer
}

type version struct {
	factory          internalinterfaces.SharedInformerFactory
	namespace        string
	tweakListOptions internalinterfaces.TweakListOptionsFunc
}

// New returns a new Interface.
func New(f internalinterfaces.SharedInformerFactory, namespace string, tweakListOptions internalinterfaces.TweakListOptionsFunc) Interface {
	return &version{factory: f, namespace: namespace, tweakListOptions: tweakListOptions}
}

// ControlPlanes returns a ControlPlaneInformer.
func (v *version) ControlPlanes() ControlPlaneInformer {
	return &controlPlaneInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// DNSs returns a DNSInformer.
func (v *version) DNSs() DNSInformer {
	return &dNSInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// Infrastructures returns a InfrastructureInformer.
func (v *version) Infrastructures() InfrastructureInformer {
	return &infrastructureInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

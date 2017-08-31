// This file was automatically generated by informer-gen

package v1

import (
	internalinterfaces "github.com/sapcc/kubernikus/pkg/generated/informers/externalversions/internalinterfaces"
)

// Interface provides access to all the informers in this group version.
type Interface interface {
	// Klusters returns a KlusterInformer.
	Klusters() KlusterInformer
}

type version struct {
	internalinterfaces.SharedInformerFactory
}

// New returns a new Interface.
func New(f internalinterfaces.SharedInformerFactory) Interface {
	return &version{f}
}

// Klusters returns a KlusterInformer.
func (v *version) Klusters() KlusterInformer {
	return &klusterInformer{factory: v.SharedInformerFactory}
}

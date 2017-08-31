// This file was automatically generated by informer-gen

package kubernikus

import (
	internalinterfaces "github.com/sapcc/kubernikus/pkg/generated/informers/externalversions/internalinterfaces"
	v1 "github.com/sapcc/kubernikus/pkg/generated/informers/externalversions/kubernikus/v1"
)

// Interface provides access to each of this group's versions.
type Interface interface {
	// V1 provides access to shared informers for resources in V1.
	V1() v1.Interface
}

type group struct {
	internalinterfaces.SharedInformerFactory
}

// New returns a new Interface.
func New(f internalinterfaces.SharedInformerFactory) Interface {
	return &group{f}
}

// V1 returns a new v1.Interface.
func (g *group) V1() v1.Interface {
	return v1.New(g.SharedInformerFactory)
}

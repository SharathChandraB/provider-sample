/*
Copyright 2021 Upbound Inc.
*/

package bucket

import (
	ujconfig "github.com/upbound/upjet/pkg/config"
)

// Configure configures the null group
func Configure(p *ujconfig.Provider) {
	p.AddResourceConfigurator("oci_objectstorage_bucket", func(r *ujconfig.Resource) {
		r.Kind = "Resource"
		r.ShortGroup = "ObjectStorage"
		// And other overrides.
	})
}

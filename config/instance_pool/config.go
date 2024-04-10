package instance_pool

import "github.com/crossplane/upjet/pkg/config"

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("databricks_instance_pool", func(r *config.Resource) {
		r.ShortGroup = "instance_pool"
	})
}

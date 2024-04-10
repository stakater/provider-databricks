package secret_scope

import "github.com/crossplane/upjet/pkg/config"

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("databricks_secret_scope", func(r *config.Resource) {
		r.ShortGroup = "secret_scope"
	})
}

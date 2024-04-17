// SPDX-FileCopyrightText: 2023 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/pkg/controller"

	cluster "github.com/stakater/provider-databricks/internal/controller/databricks/cluster"
	clusterpolicy "github.com/stakater/provider-databricks/internal/controller/databricks/clusterpolicy"
	instancepool "github.com/stakater/provider-databricks/internal/controller/databricks/instancepool"
	job "github.com/stakater/provider-databricks/internal/controller/databricks/job"
	notebook "github.com/stakater/provider-databricks/internal/controller/databricks/notebook"
	secret "github.com/stakater/provider-databricks/internal/controller/databricks/secret"
	secretscope "github.com/stakater/provider-databricks/internal/controller/databricks/secretscope"
	token "github.com/stakater/provider-databricks/internal/controller/databricks/token"
	providerconfig "github.com/stakater/provider-databricks/internal/controller/providerconfig"
)

// Setup creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		cluster.Setup,
		clusterpolicy.Setup,
		instancepool.Setup,
		job.Setup,
		notebook.Setup,
		secret.Setup,
		secretscope.Setup,
		token.Setup,
		providerconfig.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}
